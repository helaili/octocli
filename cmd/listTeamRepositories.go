package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// listCmd represents the list command
var listTeamRepositoriesCmd = &cobra.Command{
	Use:   "list",
	TraverseChildren: true,
	Short: "List a team's repositories",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintTeamRepositories(viper.GetString("org"), viper.GetString("team"))
	},
}

func init() {
	teamRepositoriesCmd.AddCommand(listTeamRepositoriesCmd)
	listTeamRepositoriesCmd.Flags().StringP("team", "t", "", "The team's name")
	listTeamRepositoriesCmd.MarkFlagRequired("team")
	listTeamRepositoriesCmd.Flags().StringP("org", "o", "", "The parent organization's name")
	listTeamRepositoriesCmd.MarkFlagRequired("org")
}
