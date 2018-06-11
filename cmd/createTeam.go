package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var description, privacy string

// listCmd represents the list command
var createTeamCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new team",
	Run: func(cmd *cobra.Command, args []string) {
		if privacy != "secret" && privacy != "closed" {
			fmt.Printf("%s is not a valid privacy setting. Accepted values are 'secret' and 'closed'\n", privacy)
			return
		}
		api.CreateTeam(viper.GetString("org"), viper.GetString("name"), viper.GetString("description"), viper.GetString("privacy"))
	},
}

func init() {
	teamCmd.AddCommand(createTeamCmd)
	createTeamCmd.Flags().StringP("org", "o", "", "The parent organization's name")
	createTeamCmd.MarkFlagRequired("org")
	createTeamCmd.Flags().StringP("name", "n", "", "The team's name")
	createTeamCmd.MarkFlagRequired("name")
	createTeamCmd.Flags().StringP("description", "d", "", "The team's description")
	createTeamCmd.Flags().StringP("privacy", "p", "secret", "The level of privacy this team should have: secret or closed")
}
