package cmd

import (
	"github.com/spf13/cobra"
)

// repoCmd represents the repo command
var repoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Base for various GitHub repository related commands",
}

func init() {
	rootCmd.AddCommand(repoCmd)
}
