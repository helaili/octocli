package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var listTeamRepositoriesCmd = &cobra.Command{
	Use:   "list",
	TraverseChildren: true,
	Short: "List a team's repositories",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeamRepositories(org, team)
	},
}

func init() {
	teamRepositoriesCmd.AddCommand(listTeamRepositoriesCmd)
	listTeamRepositoriesCmd.Flags().StringVarP(&team, "team", "t", "", "The team's name")
	listTeamRepositoriesCmd.MarkFlagRequired("team")
	listTeamRepositoriesCmd.Flags().StringVarP(&org, "org", "o", "", "The parent organization's name")
	listTeamRepositoriesCmd.MarkFlagRequired("org")
}
