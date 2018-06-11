package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var listTeamMembersCmd = &cobra.Command{
	Use:   "list",
	TraverseChildren: true,
	Short: "List members of a team",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeamMembers(org, team)
	},
}

func init() {
	teamMembersCmd.AddCommand(listTeamMembersCmd)
	listTeamMembersCmd.Flags().StringVarP(&team, "team", "t", "", "The team's name")
	listTeamMembersCmd.MarkFlagRequired("team")
	listTeamMembersCmd.Flags().StringVarP(&org, "org", "o", "", "The parent organization's name")
	listTeamMembersCmd.MarkFlagRequired("org")
}
