package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var listTeamRepositoriesCmdOrg, listTeamRepositoriesCmdTeam string

// listCmd represents the list command
var listTeamRepositoriesCmd = &cobra.Command{
	Use:   "list",
	TraverseChildren: true,
	Short: "List a team's repositories",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeamRepositories(listTeamRepositoriesCmdOrg, listTeamRepositoriesCmdTeam)
	},
}

func init() {
	teamRepositoriesCmd.AddCommand(listTeamRepositoriesCmd)
	listTeamRepositoriesCmd.Flags().StringVarP(&listTeamRepositoriesCmdTeam, "team", "t", "", "The team's name")
	listTeamRepositoriesCmd.MarkFlagRequired("team")
	listTeamRepositoriesCmd.Flags().StringVarP(&listTeamRepositoriesCmdOrg, "org", "o", "", "The parent organization's name")
	listTeamRepositoriesCmd.MarkFlagRequired("org")
}
