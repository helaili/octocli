package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var listTeamMembersCmd = &cobra.Command{
	Use:   "list",
	TraverseChildren: true,
	Short: "List members of a team",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeamMembers(viper.GetString("org"), viper.GetString("team"))
	},
}

func init() {
	teamMembersCmd.AddCommand(listTeamMembersCmd)
	listTeamMembersCmd.Flags().StringP("team", "t", "", "The team's name")
	listTeamMembersCmd.MarkFlagRequired("team")
	listTeamMembersCmd.Flags().StringP("org", "o", "", "The parent organization's name")
	listTeamMembersCmd.MarkFlagRequired("org")
}
