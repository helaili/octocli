// Copyright © 2018 Alain Hélaïli <helaili@github.com>


package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var addTeamMembersCmd = &cobra.Command{
	Use:   "add [space separated list of members]",
	Args: cobra.MinimumNArgs(1),
	Short: "Add one or more members to a team",
	Run: func(cmd *cobra.Command, members []string) {
		role := viper.GetString("role")
		if role == "member" || role == "maintainer" {
			api.AddTeamMembers(viper.GetString("org"), viper.GetString("team"), members, role)
		} else {
			fmt.Printf("%s is not a valid role. Accepted values are 'member' and 'maintainer'\n", role)
			return
		}
	},
}

func init() {
	teamMembersCmd.AddCommand(addTeamMembersCmd)
	addTeamMembersCmd.Flags().StringP("team", "t", "", "The team's slug")
	addTeamMembersCmd.MarkFlagRequired("team")
	addTeamMembersCmd.Flags().StringP("org", "o", "", "The parent organization's name")
	addTeamMembersCmd.MarkFlagRequired("org")
	addTeamMembersCmd.Flags().StringP("role", "r", "member", "User role, member or maintainer")
}
