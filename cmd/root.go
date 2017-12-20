package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile, token, server string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "octocli",
	Short: "A Command Line Interface for GitHub",
	Long: `octocli is a command line interface which goal is to make life easier
for GitHub Organizations and GitHub Enterprise administrators

Developement takes place at https://github.com/helaili/octocli.
Pull Requests are welcomed!`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.octocli.yaml)")
	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "personal authentication token to use. Required when environement variable GITHUB_AUTH_TOKEN is not set")
	rootCmd.PersistentFlags().StringVarP(&server, "server", "s", "github.com", "Hostname of the GitHub Enterprise server. Using github.com if omitted")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".octocli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".octocli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
