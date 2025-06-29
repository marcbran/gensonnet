package internal

import (
	"embed"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/google/go-jsonnet"
	"github.com/marcbran/gensonnet/internal/fun"
	"github.com/marcbran/jsonnet-kit/pkg/jsonnext"
)

//go:embed lib
var lib embed.FS

type Lib struct {
	manifestDir string
	jpath       []string
}

type LibConfig struct {
	ManifestDir string   `json:"manifestDir"`
	Jpath       []string `json:"jpath"`
}

func NewLib(
	config LibConfig,
) *Lib {
	return &Lib{
		manifestDir: config.ManifestDir,
		jpath:       config.Jpath,
	}
}

func (l Lib) vm() *jsonnet.VM {
	vm := jsonnet.MakeVM()
	var paths []string
	for _, p := range l.jpath {
		paths = append(paths, filepath.Join(l.manifestDir, p))
	}
	vm.Importer(jsonnext.CompoundImporter{
		Importers: []jsonnet.Importer{
			&jsonnext.FSImporter{Fs: lib},
			&jsonnet.FileImporter{JPaths: paths},
		},
	})
	vm.NativeFunction(fun.FormatJsonnet())
	vm.NativeFunction(fun.ParseMarkdown())
	vm.TLACode("manifest", fmt.Sprintf("import '%s/manifest.jsonnet'", l.manifestDir))
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

func (l Lib) renderPath(path string, config ServeConfig, watch bool) (string, error) {
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
