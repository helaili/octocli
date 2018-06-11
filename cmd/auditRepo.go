package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/helaili/octocli/api"
)

// auditRepoCmd represents the audit command for a repository
var auditRepoCmd = &cobra.Command{
	Use:   "audit",
	TraverseChildren: true,
	Short: "Audit access to a repository",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintRepoAudit(viper.GetString("user"), viper.GetString("name"), viper.GetBool("isGitHubber"))
	},
}

func init() {
	repoCmd.AddCommand(auditRepoCmd)
	auditRepoCmd.Flags().StringP("name", "n", "", "The repository's name")
	auditRepoCmd.MarkFlagRequired("name")

	auditRepoCmd.Flags().StringP("user", "u", "", "The repository's owner")
	auditRepoCmd.MarkFlagRequired("user")

	auditRepoCmd.Flags().BoolP("gitHubber", "g", false, "Is this user a GitHub employee")
}
