package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/caraml-dev/timber/common/log"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dataset-service",
	Short: "Dataset Service HTTP Server",
	Long:  `Dataset Service HTTP Server for deployment of logging infrastructure and generation of ML datasets.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
