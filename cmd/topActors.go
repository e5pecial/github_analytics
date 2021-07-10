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
		topN, _ := cmd.Flags().GetInt("topN")
		if topN < 0 {
			fmt.Println("Negative numbers unsupported. Set n=10")
			topN = 10
		}
		internal.GetAutorsByCommits(topN)
	},
}

func init() {
	rootCmd.AddCommand(topActorsCmd)
}
