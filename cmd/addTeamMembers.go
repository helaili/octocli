// Copyright © 2018 Alain Hélaïli <helaili@github.com>


package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var addTeamMembersCmdTeam, addTeamMembersCmdOrg, addTeamMembersCmdRole string

// listCmd represents the list command
var addTeamMembersCmd = &cobra.Command{
	Use:   "add [space separated list of members]",
	Args: cobra.MinimumNArgs(1),
	Short: "Add one or more members to a team",
	Run: func(cmd *cobra.Command, members []string) {
		if addTeamMembersCmdRole == "member" || addTeamMembersCmdRole == "maintainer" {
			api.AddTeamMembers(addTeamMembersCmdOrg, addTeamMembersCmdTeam, members, addTeamMembersCmdRole)
		} else {
			fmt.Printf("%s is not a valid role. Accepted values are 'member' and 'maintainer'\n", addTeamMembersCmdRole)
			return
		}
	},
}

func init() {
	teamMembersCmd.AddCommand(addTeamMembersCmd)
	addTeamMembersCmd.Flags().StringVarP(&addTeamMembersCmdTeam, "team", "t", "", "The team's slug")
	addTeamMembersCmd.MarkFlagRequired("team")
	addTeamMembersCmd.Flags().StringVarP(&addTeamMembersCmdOrg, "org", "o", "", "The parent organization's name")
	addTeamMembersCmd.MarkFlagRequired("org")
	addTeamMembersCmd.Flags().StringVarP(&addTeamMembersCmdRole, "role", "r", "member", "User role, member or maintainer")
}
