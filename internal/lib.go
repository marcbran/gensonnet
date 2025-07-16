package internal

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/marcbran/gensonnet/internal/fun"
	"github.com/marcbran/gensonnet/pkg/gensonnet/config"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/marcbran/jsonnet-kit/pkg/jsonnext"
)

//go:embed lib
var lib embed.FS

type Lib struct {
	manifestDir string
	manifestStr string
	jpath       []string
	filesystems []embed.FS
}

func NewLib(
	config config.LibConfig,
) *Lib {
	return &Lib{
		manifestDir: config.ManifestDir,
		manifestStr: config.ManifestStr,
		jpath:       config.Jpath,
		filesystems: config.Filesystems,
	}
}

func (l Lib) vm() *jsonnet.VM {
	vm := jsonnet.MakeVM()
	var paths []string
	for _, p := range l.jpath {
		if l.manifestDir != "" {
			paths = append(paths, filepath.Join(l.manifestDir, p))
		} else {
			paths = append(paths, p)
		}
	}
	importers := []jsonnet.Importer{
		&jsonnext.FSImporter{Fs: lib},
		&jsonnet.FileImporter{JPaths: paths},
	}
	for _, fs := range l.filesystems {
		importers = append(importers, &jsonnext.FSImporter{Fs: fs})
	}
	vm.Importer(jsonnext.CompoundImporter{
		Importers: importers,
	})
	var manifestCode string
	if l.manifestDir != "" {
		manifestCode = fmt.Sprintf("import '%s/manifest.jsonnet'", l.manifestDir)
	} else {
		manifestCode = l.manifestStr
	}
	vm.TLACode("manifest", manifestCode)
	vm.NativeFunction(fun.FormatJsonnet())
	vm.NativeFunction(fun.ManifestJsonnet())
	vm.NativeFunction(fun.ParseJsonnet())
	vm.NativeFunction(fun.ManifestMarkdown())
	vm.NativeFunction(fun.ParseMarkdown())
	return vm
}

func (l Lib) render() (map[string]string, error) {
	vm := l.vm()
	vm.StringOutput = true

	files, err := vm.EvaluateFileMulti("./lib/render.libsonnet")
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (l Lib) renderPath(path string, config config.ServeConfig, watch bool) (string, error) {
	vm := l.vm()
	vm.TLAVar("path", path)
	jsonConfig, err := json.Marshal(config)
	if err != nil {
		return "", err
	}
	vm.TLACode("config", string(jsonConfig))
	vm.TLACode("watch", fmt.Sprintf("%t", watch))
	vm.StringOutput = true

	file, err := vm.EvaluateFile("./lib/render_path.libsonnet")
	if err != nil {
		return "", err
	}
	return file, nil
}
