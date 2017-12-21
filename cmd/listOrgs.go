package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var user string

// listOrgsCmd represents the list command
var listOrgsCmd = &cobra.Command{
	Use:   "list",
	Short: "List organizations",
	Long: `List all organizations when nothing is specified, or list the Organizations
which the specified user belongs to`,
  Run: func(cmd *cobra.Command, args []string) {
		if server == "github.com" && user == "" {
			log.Fatal("Browsing all the organizations on github.com is not allowed.")
		} else if user != "" {
			api.GetUserOrgs(server, viper.GetString("token"), user, "")
		} else {
			//TODO: implement retrival of all orgs on GHE
		}
	},
}

func init() {
	orgCmd.AddCommand(listOrgsCmd)
	listOrgsCmd.Flags().StringVarP(&user, "user", "u", "", "Only retrieves organizations which this user belongs to")
}
