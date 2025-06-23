package internal

import (
	"embed"
	"fmt"

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
	vm.Importer(jsonnext.CompoundImporter{
		Importers: []jsonnet.Importer{
			&jsonnext.FSImporter{Fs: lib},
			&jsonnet.FileImporter{JPaths: l.jpath},
		},
	})
	vm.NativeFunction(fun.FormatJsonnet())
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

func (l Lib) renderPath(path string) (string, error) {
	vm := l.vm()
	vm.TLAVar("path", path)
	vm.StringOutput = true

	file, err := vm.EvaluateFile("./lib/render_path.libsonnet")
	if err != nil {
		return "", err
	}
	return file, nil
}
