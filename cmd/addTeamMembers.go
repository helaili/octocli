// Copyright © 2018 Alain Hélaïli <helaili@github.com>


package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var role string

// listCmd represents the list command
var addTeamMembersCmd = &cobra.Command{
	Use:   "add",
	Args: cobra.MinimumNArgs(1),
	Short: "Add members to a team",
	Run: func(cmd *cobra.Command, members []string) {
		if role == "member" || role == "maintainer" {
			api.AddTeamMembers(server, viper.GetString("token"), org, team, members, role)
		} else {
			fmt.Printf("%s is not a valid role. Accepted values are 'member' and 'maintainer'\n", role)
		}
	},
}

func init() {
	teamMembersCmd.AddCommand(addTeamMembersCmd)
	addTeamMembersCmd.Flags().StringVarP(&team, "team", "t", "", "The team's slug")
	addTeamMembersCmd.MarkFlagRequired("team")
	addTeamMembersCmd.Flags().StringVarP(&org, "org", "o", "", "The parent organization's name")
	addTeamMembersCmd.MarkFlagRequired("org")
	addTeamMembersCmd.Flags().StringVarP(&role, "role", "r", "member", "User role, member (default) or maintainer")
}
