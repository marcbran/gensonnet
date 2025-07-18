package cmd

import (
	"github.com/marcbran/gensonnet/internal"
	"github.com/marcbran/gensonnet/pkg/gensonnet/config"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve [flags] directory",
	Short: "Serves the provided directory in an HTTP server according to its manifest",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		cmd.SilenceErrors = true
		dirname := "."
		if len(args) > 0 {
			dirname = args[0]
		}
		config, err := config.New(dirname)
		if err != nil {
			return err
		}
		err = internal.Serve(cmd.Context(), config.Serve)
		if err != nil {
			return err
		}
		return nil
	},
}
