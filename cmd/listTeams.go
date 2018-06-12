package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var listTeamsCmdOrg string

// listCmd represents the list command
var listTeamsCmd = &cobra.Command{
	Use:   "list",
	Short: "List teams within an organization",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeams(listTeamsCmdOrg)
	},
}

func init() {
	teamCmd.AddCommand(listTeamsCmd)
	listTeamsCmd.Flags().StringVarP(&listTeamsCmdOrg, "org", "o", "", "The parent organization's name")
	listTeamsCmd.MarkFlagRequired("org")
}
