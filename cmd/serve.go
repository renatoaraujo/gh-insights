package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"renatoaraujo/gh-insights/cmd/app"
	"renatoaraujo/gh-insights/pkg/infrastructure"
	"renatoaraujo/gh-insights/pkg/server"
)

var port string

func init() {
	serveCmd.Flags().StringVarP(&port, "port", "p", "8080", "HTTP port to serve the charts. Default 8080")
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start a HTTP app",
	Long:  `Start a HTTP app to visualise the generated charts`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := infrastructure.NewDatabase(cmd.Context(), os.Getenv("DATABASE_DSN"))
		if err != nil {
			log.Fatal(err)
		}

		app := app.NewApp(cmd.Context(), db)
		server.Serve(port, app.Router)
	},
}
