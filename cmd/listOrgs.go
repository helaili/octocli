package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

// listOrgsCmd represents the list command
var listOrgsCmd = &cobra.Command{
	Use:   "list",
	Short: "List organizations",
	Long: `List all organizations when nothing is specified, or list the Organizations
which the specified user belongs to`,
  Run: func(cmd *cobra.Command, args []string) {
		if server == "github.com" && user == "" {
			fmt.Println("Browsing all the organizations on github.com is not allowed.")
		} else if user != "" {
			api.PrintUserOrgs(user)
		} else {
			api.PrintAllOrgs()
		}
	},
}

func init() {
	orgCmd.AddCommand(listOrgsCmd)
	listOrgsCmd.Flags().StringVarP(&user, "user", "u", "", "Only retrieves organizations which this user belongs to")
}
