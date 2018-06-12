package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var listOrgsCmdLogin string

// listOrgsCmd represents the list command
var listOrgsCmd = &cobra.Command{
	Use:   "list",
	Short: "List organizations",
	Long: `List all organizations when nothing is specified, or list the Organizations
which the specified user belongs to`,
  Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("server") == "github.com" && listOrgsCmdLogin == "" {
			fmt.Println("Browsing all the organizations on github.com is not allowed.")
		} else if listOrgsCmdLogin != "" {
			api.PrintUserOrgs(listOrgsCmdLogin)
		} else {
			api.PrintAllOrgs()
		}
	},
}

func init() {
	orgCmd.AddCommand(listOrgsCmd)
	listOrgsCmd.Flags().StringVarP(&listOrgsCmdLogin, "login", "l", "", "Only retrieves organizations which this user belongs to")
}
