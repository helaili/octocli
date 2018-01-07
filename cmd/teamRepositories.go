package cmd

import (
	"github.com/spf13/cobra"
)

// teamRepositoriesCmd represents the members command
var teamRepositoriesCmd = &cobra.Command{
	Use:   "repositories",
	Short: "Base for various team repositories related commands",
}

func init() {
	teamCmd.AddCommand(teamRepositoriesCmd)
}
