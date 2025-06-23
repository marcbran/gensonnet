package internal

import (
	"encoding/json"
	"fmt"
	"github.com/marcbran/jsonnet-kit/pkg/jsonnext"

	"github.com/google/go-jsonnet"
	"github.com/marcbran/gensonnet/internal/fun"
)

type ConfigLib struct {
	manifestDir string
}

type Config struct {
	Render RenderConfig `json:"render"`
	Serve  ServeConfig  `json:"serve"`
}

func NewConfig(
	manifestDir string,
) (Config, error) {
	configLib := ConfigLib{
		manifestDir: manifestDir,
	}
	config, err := configLib.readConfig()
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func (l ConfigLib) vm() *jsonnet.VM {
	vm := jsonnet.MakeVM()
	vm.Importer(jsonnext.CompoundImporter{
		Importers: []jsonnet.Importer{
			&jsonnext.FSImporter{Fs: lib},
			&jsonnet.FileImporter{},
		},
	})
	vm.NativeFunction(fun.FormatJsonnet())
	vm.TLACode("manifest", fmt.Sprintf("import '%s/manifest.jsonnet'", l.manifestDir))
	return vm
}

func (l ConfigLib) readConfig() (Config, error) {
	vm := l.vm()

	vm.TLAVar("manifestDir", l.manifestDir)
	rawConfig, err := vm.EvaluateFile("./lib/read_config.libsonnet")
	if err != nil {
		return Config{}, err
	}
	var config Config
	err = json.Unmarshal([]byte(rawConfig), &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
