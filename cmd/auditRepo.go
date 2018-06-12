package cmd

import (
	"github.com/spf13/cobra"
	"github.com/helaili/octocli/api"
)

var auditRepoCmdRepo, auditRepoCmdLogin string
var auditRepoCmdGitHubber bool

// auditRepoCmd represents the audit command for a repository
var auditRepoCmd = &cobra.Command{
	Use:   "audit",
	TraverseChildren: true,
	Short: "Audit access to a repository",
	Run: func(cmd *cobra.Command, args []string) {
		api.PrintRepoAudit(auditRepoCmdLogin, auditRepoCmdRepo, auditRepoCmdGitHubber)
	},
}

func init() {
	repoCmd.AddCommand(auditRepoCmd)
	auditRepoCmd.Flags().StringVarP(&auditRepoCmdRepo, "repo", "r", "", "The repository's name")
	auditRepoCmd.MarkFlagRequired("repo")

	auditRepoCmd.Flags().StringVarP(&auditRepoCmdLogin, "login", "l", "", "The repository's owner")
	auditRepoCmd.MarkFlagRequired("login")

	auditRepoCmd.Flags().BoolVarP(&auditRepoCmdGitHubber, "gitHubber", "g", false, "Is this user a GitHub employee")
	auditRepoCmd.Flags().MarkHidden("gitHubber")
}
