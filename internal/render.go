package internal

import (
	"context"
	"github.com/marcbran/gensonnet/pkg/gensonnet/config"
	"os"
	"path"
)

func Render(ctx context.Context, config config.RenderConfig) error {
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
