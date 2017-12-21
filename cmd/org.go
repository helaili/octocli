package cmd

import (
	"github.com/spf13/cobra"
)

// orgCmd represents the org command
var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Base for various GitHub Organizations related commands",
}

func init() {
	rootCmd.AddCommand(orgCmd)
}
