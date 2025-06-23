package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gensonnet",
	Short: "Tool to create manifest pages with Jsonnet",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(renderCmd)
	rootCmd.AddCommand(serveCmd)
	rootCmd.PersistentFlags().StringArrayP("jpath", "J", []string{}, "Specify an additional library search dir (right-most wins)")
}
