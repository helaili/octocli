package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var listTeamMembersCmdTeam, listTeamMembersCmdOrg string

// listCmd represents the list command
var listTeamMembersCmd = &cobra.Command{
	Use:   "list",
	TraverseChildren: true,
	Short: "List members of a team",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeamMembers(listTeamMembersCmdOrg, listTeamMembersCmdTeam)
	},
}

func init() {
	teamMembersCmd.AddCommand(listTeamMembersCmd)
	listTeamMembersCmd.Flags().StringVarP(&listTeamMembersCmdTeam, "team", "t", "", "The team's name")
	listTeamMembersCmd.MarkFlagRequired("team")
	listTeamMembersCmd.Flags().StringVarP(&listTeamMembersCmdOrg, "org", "o", "", "The parent organization's name")
	listTeamMembersCmd.MarkFlagRequired("org")
}
