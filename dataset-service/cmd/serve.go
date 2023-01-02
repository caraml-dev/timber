package cmd

import (
	"github.com/spf13/cobra"

	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/dataset-service/server"
)

var cfgFile []string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts Dataset Service HTTP server with configured API",
	Run: func(cmd *cobra.Command, args []string) {

		server, err := server.NewServer(cfgFile)
		if err != nil {
			log.Panic(err)
		}
		server.Start()
	},
}

func init() {
	serveCmd.Flags().StringArrayVar(&cfgFile, "config", []string{},
		`Path to one or more configuration files. The flag can be set multiple times
	and the later values will take precedence.`)
	RootCmd.AddCommand(serveCmd)
}
