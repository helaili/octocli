package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var createTeamCmdOrg, createTeamCmdTeam, createTeamCmdDescription, createTeamCmdPrivacy string

// listCmd represents the list command
var createTeamCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new team",
	Run: func(cmd *cobra.Command, args []string) {
		if createTeamCmdPrivacy != "secret" && createTeamCmdPrivacy != "closed" {
			fmt.Printf("%s is not a valid privacy setting. Accepted values are 'secret' and 'closed'\n", createTeamCmdPrivacy)
			return
		}
		api.CreateTeam(createTeamCmdOrg, createTeamCmdTeam, createTeamCmdDescription, createTeamCmdPrivacy)
	},
}

func init() {
	teamCmd.AddCommand(createTeamCmd)
	createTeamCmd.Flags().StringVarP(&createTeamCmdOrg, "org", "o", "", "The parent organization's name")
	createTeamCmd.MarkFlagRequired("org")
	createTeamCmd.Flags().StringVarP(&createTeamCmdTeam, "team", "t", "", "The team's name")
	createTeamCmd.MarkFlagRequired("team")
	createTeamCmd.Flags().StringVarP(&createTeamCmdDescription, "description", "d", "", "The team's description")
	createTeamCmd.Flags().StringVarP(&createTeamCmdPrivacy, "privacy", "p", "secret", "The level of privacy this team should have: secret or closed")
}
