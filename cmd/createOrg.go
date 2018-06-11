package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// createCmd represents the create organization command
var createOrgCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new organization in GitHub Enterprise",
	Long: `Creates a new organization with the manadatory name and owner.
This command only works on GitHub Enterprise`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("server") == "github.com" {
			fmt.Println("Creating an organization on github.com is not allowed.")
		} else {
			api.CreateOrg(viper.GetString("name"), viper.GetString("owner"), viper.GetString("display"))
		}
	},
}

func init() {
	orgCmd.AddCommand(createOrgCmd)
	createOrgCmd.Flags().StringP("owner", "o", "", "The organization owner's handle")
	createOrgCmd.Flags().StringP("name", "n", "", "The organization's name")
	createOrgCmd.Flags().StringP("display", "d", "", "The organization's display name")

	createOrgCmd.MarkFlagRequired("owner")
	createOrgCmd.MarkFlagRequired("name")
}
