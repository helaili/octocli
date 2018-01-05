package cmd

import (
	"github.com/spf13/cobra"
)

// teamCmd represents the team command
var teamCmd = &cobra.Command{
	Use:   "team",
	Short: "Base for various team related commands",
}

func init() {
	rootCmd.AddCommand(teamCmd)
}
