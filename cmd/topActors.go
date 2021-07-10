package cmd

import (
	"fmt"

	"github.com/e5pecial/github_analytics/internal"
	"github.com/spf13/cobra"
)

var topActorsCmd = &cobra.Command{
	Use:   "topActors",
	Short: "Top active users",
	Long:  `Top 10 active users sorted by amount of PRs created and commits pushed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		internal.GetAutorsByCommits(10)
	},
}

func init() {
	rootCmd.AddCommand(topActorsCmd)
}
