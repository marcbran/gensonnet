package gensonnet

import (
	"context"
	"github.com/marcbran/gensonnet/internal"
)

func RenderDir(ctx context.Context, dirname string) error {
	config, err := internal.NewConfig(dirname)
	if err != nil {
		return err
	}
	err = internal.Render(ctx, config.Render)
	if err != nil {
		return err
	}
	return nil
}

func RenderWithConfig(ctx context.Context, config internal.Config) error {
	err := internal.Render(ctx, config.Render)
	if err != nil {
		return err
	}
	return nil
}
