package cmd

import (
	"github.com/spf13/cobra"

	"renatoaraujo/gh-insights/pkg/server"
)

var port string

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "HTTP port to serve the charts. Default 8080")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a HTTP application",
	Long:  `Start a HTTP application to visualise the generated charts`,
	Run: func(cmd *cobra.Command, args []string) {
		server := server.NewServer(port)
		server.Serve()
	},
}
