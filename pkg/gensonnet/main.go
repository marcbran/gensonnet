package gensonnet

import (
	"context"
	"github.com/marcbran/gensonnet/internal"
	"github.com/marcbran/gensonnet/pkg/gensonnet/config"
)

func RenderDir(ctx context.Context, dirname string) error {
	config, err := config.New(dirname)
	if err != nil {
		return err
	}
	err = internal.Render(ctx, config.Render)
	if err != nil {
		return err
	}
	return nil
}

func RenderWithConfig(ctx context.Context, config config.Config) error {
	err := internal.Render(ctx, config.Render)
	if err != nil {
		return err
	}
	return nil
}
