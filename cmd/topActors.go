package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// topActorsCmd represents the topActors command
var topActorsCmd = &cobra.Command{
	Use:   "topActors",
	Short: "Top 10 active users",
	Long:  `Top 10 active users sorted by amount of PRs created and commits pushed`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("topActors called")
	},
}

func init() {
	rootCmd.AddCommand(topActorsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// topActorsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// topActorsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
