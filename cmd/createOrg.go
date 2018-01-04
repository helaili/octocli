package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var owner string
var name string
var profileName string

// createCmd represents the create command
var createOrgCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new organization in GitHub Enterprise",
	Long: `Creates a new organization with the manadatory name and owner.
This command only works on GitHub Enterprise`,
	Run: func(cmd *cobra.Command, args []string) {
		if server == "github.com" && user == "" {
			fmt.Println("Creating an organization on github.com is not allowed.")
		} else {
			api.CreateOrg(server, viper.GetString("token"), name, owner, profileName)
		}
	},
}

func init() {
	orgCmd.AddCommand(createOrgCmd)
	createOrgCmd.Flags().StringVarP(&owner, "owner", "o", "", "The organization owner's handle")
	createOrgCmd.Flags().StringVarP(&name, "name", "n", "", "The organization's name")
	createOrgCmd.Flags().StringVarP(&profileName, "display", "d", "", "The organization's display name")

	createOrgCmd.MarkFlagRequired("owner")
	createOrgCmd.MarkFlagRequired("name")
}
