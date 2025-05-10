package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "actionforge",
	Short: "ActionForge - Craft your GitHub Actions seamlessly",
	Long: `ActionForge is a CLI tool that helps you create, manage, and optimize 
GitHub Actions workflows with ease. Powered by Amazon Q, it provides
intelligent suggestions and validations for your workflows.`,
	Run: func(cmd *cobra.Command, args []string) {
		// This is the default command when no subcommand is provided
		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings
	rootCmd.PersistentFlags().StringP("config", "c", "", "config file (default is $HOME/.actionforge.yaml)")
	
	// Add more global flags here
}
