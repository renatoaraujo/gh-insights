package cmd

import (
	"context"
	"log"
	"renatoaraujo/gh-insights/pkg/infrastructure"
	"strings"

	"renatoaraujo/gh-insights/pkg/github"
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

		ctx := context.Background()
		split := strings.Split(repository, "/")
		ghClient := github.NewPublic(split[0], split[1])

		db, err := infrastructure.NewDatabase("postgresql://postgres:example@localhost/postgres?sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		sync.Sync(ctx, ghClient, db)
	},
}
