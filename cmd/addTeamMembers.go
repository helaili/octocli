// Copyright © 2018 Alain Hélaïli <helaili@github.com>


package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var addTeamMembersCmd = &cobra.Command{
	Use:   "add",
	Args: cobra.MinimumNArgs(1),
	Short: "Add members to a team",
	Run: func(cmd *cobra.Command, members []string) {
		api.AddTeamMembers(server, viper.GetString("token"), org, team, members)
	},
}

func init() {
	teamMembersCmd.AddCommand(addTeamMembersCmd)
	addTeamMembersCmd.Flags().StringVarP(&team, "team", "t", "", "The team's slug")
	addTeamMembersCmd.MarkFlagRequired("team")
	addTeamMembersCmd.Flags().StringVarP(&org, "org", "o", "", "The parent organization's name")
	addTeamMembersCmd.MarkFlagRequired("org")
}
