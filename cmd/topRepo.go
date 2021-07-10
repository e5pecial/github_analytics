package cmd

import (
	"fmt"

	"github.com/e5pecial/github_analytics/internal"
	"github.com/spf13/cobra"
)

var topRepoCmd = &cobra.Command{
	Use:   "topRepo",
	Short: "Top repositories",
	Long: `Top repositories sorted by amount of commits pushed (use "commits" key)
	or Top repositories sorted by amount of watch events (use "watch" key)
	Example: topRepo -t watch`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Short)
		flag, _ := cmd.Flags().GetString("type")
		topN, _ := cmd.Flags().GetInt("topN")

		if topN < 0 {
			fmt.Println("Negative numbers unsupported. Set n=10")
			topN = 10
		}
		switch flag {
		case "watch":
			internal.GetRepositoriesByWatchEvents(topN)
		case "commits":
			internal.GetRepositoriesByCommits(topN)
		default:
			fmt.Println("Undefined type key: Use key -t and params `watch` or `commits`")
		}
	},
}

func init() {
	rootCmd.AddCommand(topRepoCmd)
	topRepoCmd.Flags().StringP("type", "t", "", "Get top repos by commits pushed or watch event")
}
