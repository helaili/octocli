package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var createOrgCmdLogin, createOrgCmdOrg, createOrgCmdDisplay string

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
			api.CreateOrg(createOrgCmdOrg, createOrgCmdLogin, createOrgCmdDisplay)
		}
	},
}

func init() {
	orgCmd.AddCommand(createOrgCmd)
	createOrgCmd.Flags().StringVarP(&createOrgCmdLogin, "login", "l", "", "The organization owner's handle")
	createOrgCmd.MarkFlagRequired("login")
	createOrgCmd.Flags().StringVarP(&createOrgCmdOrg, "org", "o", "", "The organization's name")
	createOrgCmd.MarkFlagRequired("org")
	createOrgCmd.Flags().StringVarP(&createOrgCmdDisplay, "display", "d", "", "The organization's display name")
}
