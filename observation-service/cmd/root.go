package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/caraml-dev/timber/common/log"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "observation-service",
	Short: "Observation Service CMux Server",
	Long:  `Observation Service CMux Server for logging ML ground-truth observations`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
