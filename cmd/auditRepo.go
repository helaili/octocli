package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

var isGitHubber bool

// auditRepoCmd represents the audit command for a repository
var auditRepoCmd = &cobra.Command{
	Use:   "audit",
	TraverseChildren: true,
	Short: "Audit access to a repository",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintRepoAudit(server, viper.GetString("token"), user, name, isGitHubber)
	},
}

func init() {
	repoCmd.AddCommand(auditRepoCmd)
	auditRepoCmd.Flags().StringVarP(&name, "name", "n", "", "The repository's name")
	auditRepoCmd.MarkFlagRequired("name")

	auditRepoCmd.Flags().StringVarP(&user, "user", "u", "", "The repository's owner")
	auditRepoCmd.MarkFlagRequired("user")

	auditRepoCmd.Flags().BoolVarP(&isGitHubber, "gitHubber", "g", false, "Is this user a GitHub employee")
}
