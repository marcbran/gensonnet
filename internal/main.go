package internal

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

type RenderConfig struct {
	TargetDir string    `json:"targetDir"`
	Lib       LibConfig `json:"lib"`
}

func Render(ctx context.Context, config RenderConfig) error {
	lib := NewLib(config.Lib)
	files, err := lib.render()
	if err != nil {
		return err
	}
	for name, content := range files {
		filename := path.Join(config.TargetDir, name)
		err := os.MkdirAll(path.Dir(filename), 0755)
		if err != nil {
			return err
		}
		err = os.WriteFile(filename, []byte(content), 0666)
		if err != nil {
			return err
		}
	}
	return nil
}

type ServeConfig struct {
	Server ServerConfig `json:"server"`
	Lib    LibConfig    `json:"lib"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

func Serve(ctx context.Context, config ServeConfig) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.WithField("path", r.URL.Path).
			WithField("method", r.Method).
			Info("request received")

		p := strings.TrimPrefix(r.URL.Path, "/")
		lib := NewLib(config.Lib)
		file, err := lib.renderPath(p)
		if err != nil {
			log.WithError(err).
				WithField("path", p).
				Error("failed to handle view")
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(file))
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
