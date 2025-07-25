package internal

import (
	"context"
	"errors"
	"fmt"
	"github.com/marcbran/gensonnet/pkg/gensonnet/config"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
	"golang.org/x/sync/errgroup"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func Serve(ctx context.Context, config config.ServeConfig) error {
	reloadCh := make(chan struct{}, 1)
	broadcaster := Broadcaster{}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		watchCtx, watchCancel := context.WithCancel(gCtx)
		defer watchCancel()

		err := watchFiles(watchCtx, config.Lib.ManifestDir, reloadCh)
		if err != nil && !errors.Is(err, context.Canceled) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		publishCtx, publishCancel := context.WithCancel(gCtx)
		defer publishCancel()

		err := broadcaster.Publish(publishCtx, reloadCh)
		if err != nil && !errors.Is(err, context.Canceled) {
			return err
		}
		return nil
	})

	g.Go(func() error {
		runCtx, runCancel := context.WithCancel(gCtx)
		defer runCancel()

		err := runServer(runCtx, config, &broadcaster)
		if err != nil && !errors.Is(err, context.Canceled) {
			return err
		}
		return nil
	})

	return g.Wait()
}

type Broadcaster struct {
	subscribers sync.Map
}

func (b *Broadcaster) Publish(ctx context.Context, ch <-chan struct{}) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case value, ok := <-ch:
			if !ok {
				return nil
			}
			b.subscribers.Range(func(k, v interface{}) bool {
				select {
				case k.(chan struct{}) <- value:
				default:
					// Drop update if channel is full
				}
				return true
			})
		}
	}
}

func (b *Broadcaster) Subscribe() (<-chan struct{}, func()) {
	ch := make(chan struct{}, 100)

	b.subscribers.Store(ch, struct{}{})

	unsubscribe := func() {
		b.subscribers.Delete(ch)
		close(ch)
	}

	return ch, unsubscribe
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func runServer(ctx context.Context, config config.ServeConfig, reloadBroadcaster *Broadcaster) error {
	mux := http.NewServeMux()

	for endpoint, filePath := range config.Server.StaticFiles {
		if !strings.HasPrefix(endpoint, "/") {
			endpoint = "/" + endpoint
		}
		filePath = filepath.Join(config.Server.StaticBaseDir, filePath)

		stat, err := os.Stat(filePath)
		if err != nil {
			log.WithError(err).
				WithField("file", filePath).
				Error("failed to stat static file/directory")
			continue
		}

		if stat.IsDir() {
			mux.Handle(endpoint+"/", http.StripPrefix(endpoint, http.FileServer(http.Dir(filePath))))
			log.WithField("endpoint", endpoint).
				WithField("directory", filePath).
				Info("registered static directory handler")
		} else {
			mux.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
				http.ServeFile(w, r, filePath)
			})
			log.WithField("endpoint", endpoint).
				WithField("file", filePath).
				Info("registered static file handler")
		}
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.WithField("path", r.URL.Path).
			WithField("method", r.Method).
			Info("request received")

		p := strings.TrimPrefix(r.URL.Path, "/")

		if config.Server.DirectoryIndex != "" {
			if p == "" {
				p = config.Server.DirectoryIndex
			} else if !strings.Contains(filepath.Base(p), ".") {
				p = filepath.Join(p, config.Server.DirectoryIndex)
			}
		}

		lib := NewLib(config.Lib)
		file, err := lib.renderPath(p, config, true)
		if err != nil {
			log.WithError(err).
				WithField("path", p).
				Error("failed to handle view")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(file))
		if err != nil {
			log.WithError(err).
				Error("failed to write response")
			return
		}
	})

	mux.HandleFunc("/_reload", func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Query().Get("path")
		if p == "" {
			log.Error("path is required")
			http.Error(w, "path is required", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.WithError(err).
				Error("failed to upgrade connection to websocket")
			return
		}
		defer func() {
			err := conn.Close()
			if err != nil {
				log.WithError(err).
					Error("failed to close websocket connection")
			}
		}()

		reloads, unsubscribe := reloadBroadcaster.Subscribe()
		defer unsubscribe()

		for {
			select {
			case <-r.Context().Done():
				return
			case _, ok := <-reloads:
				if !ok {
					return
				}

				lib := NewLib(config.Lib)
				_, err := lib.renderPath(p, config, true)
				if err != nil {
					log.WithError(err).
						WithField("path", p).
						Warn("failed to render path")
					continue
				}

				err = conn.WriteMessage(websocket.TextMessage, []byte(("reload")))
				if err != nil {
					log.WithError(err).
						Error("failed to write websocket message")
					return
				}
			}
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", config.Server.Port),
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		err := server.Shutdown(context.Background())
		if err != nil {
			log.WithError(err).
				Error("failed to shutdown server")
		}
	}()

	log.WithField("port", config.Server.Port).
		Info("starting server")
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func watchFiles(ctx context.Context, dir string, reloadCh chan<- struct{}) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer func() {
		err := watcher.Close()
		if err != nil {
			log.WithError(err).
				Error("failed to close watcher")
		}
	}()

	err = watcher.Add(dir)
	if err != nil {
		return err
	}

	log.WithField("directory", dir).
		Info("watching directory for changes")

	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return watcher.Add(path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event, ok := <-watcher.Events:
			if !ok {
				return nil
			}
			if event.Op&(fsnotify.Write|fsnotify.Create|fsnotify.Remove) != 0 {
				log.WithField("file", event.Name).
					Info("file changed, triggering reload")
				select {
				case reloadCh <- struct{}{}:
				default:
				}
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return nil
			}
			log.WithError(err).
				Error("file watcher error")
		}
	}
}
