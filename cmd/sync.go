package cmd

import (
	"log"
	"os"
	"strings"

	"renatoaraujo/gh-insights/pkg/github"
	"renatoaraujo/gh-insights/pkg/infrastructure"
	"renatoaraujo/gh-insights/pkg/sync"

	"github.com/spf13/cobra"
)

var repository string

func init() {
	syncCmd.Flags().StringVarP(&repository, "repository", "r", "", "Repository using the format owner/repository")
	rootCmd.AddCommand(syncCmd)
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync GitHub Repository",
	Long:  `Sync all the issues and pull requests from a given repository`,
	Run: func(cmd *cobra.Command, args []string) {
		split := strings.Split(repository, "/")
		ghClient := github.NewAuthenticatedClient(split[0], split[1], os.Getenv("GITHUB_TOKEN"))

		db, err := infrastructure.NewDatabase(cmd.Context(), os.Getenv("DATABASE_DSN"))
		if err != nil {
			log.Fatal(err)
		}

		sync.Sync(cmd.Context(), ghClient, db)
		db.Close()
	},
}
