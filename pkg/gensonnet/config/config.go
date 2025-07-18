package config

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/marcbran/devsonnet/pkg/jsonnext"

	"github.com/google/go-jsonnet"
)

//go:embed lib
var lib embed.FS

type Config struct {
	Render RenderConfig `json:"render"`
	Serve  ServeConfig  `json:"serve"`
}

type RenderConfig struct {
	TargetDir string    `json:"targetDir"`
	Lib       LibConfig `json:"lib"`
}

type ServeConfig struct {
	Server ServerConfig `json:"server"`
	Lib    LibConfig    `json:"lib"`
}

type ServerConfig struct {
	Port           int               `json:"port"`
	DirectoryIndex string            `json:"directoryIndex"`
	StaticBaseDir  string            `json:"staticBaseDir"`
	StaticFiles    map[string]string `json:"staticFiles"`
}

type LibConfig struct {
	ManifestDir  string   `json:"manifestDir"`
	ManifestCode string   `json:"manifestStr"`
	Jpath        []string `json:"jpath"`
	Filesystems  []embed.FS
	Imports      map[string]string `json:"imports"`
}

func New(
	manifestDir string,
) (Config, error) {
	configLib := Lib{
		manifestDir: manifestDir,
	}
	config, err := configLib.readConfig()
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

type Lib struct {
	manifestDir string
}

func (l Lib) vm() *jsonnet.VM {
	vm := jsonnet.MakeVM()
	vm.Importer(jsonnext.CompoundImporter{
		Importers: []jsonnet.Importer{
			&jsonnext.FSImporter{Fs: lib},
			&jsonnet.FileImporter{},
		},
	})
	vm.TLACode("manifest", fmt.Sprintf("import '%s/manifest.jsonnet'", l.manifestDir))
	return vm
}

func (l Lib) readConfig() (Config, error) {
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
