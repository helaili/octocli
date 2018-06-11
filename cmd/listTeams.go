package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var listTeamsCmd = &cobra.Command{
	Use:   "list",
	Short: "List teams within an organization",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeams(viper.GetString("org"))
	},
}

func init() {
	teamCmd.AddCommand(listTeamsCmd)
	listTeamsCmd.Flags().StringP("org", "o", "", "The parent organization's name")
	listTeamsCmd.MarkFlagRequired("org")
}
