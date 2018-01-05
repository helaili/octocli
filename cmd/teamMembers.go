package cmd

import (
	"github.com/spf13/cobra"
)

// membersCmd represents the members command
var teamMembersCmd = &cobra.Command{
	Use:   "members",
	Short: "Base for various team members related commands",
}

func init() {
	teamCmd.AddCommand(teamMembersCmd)
}
