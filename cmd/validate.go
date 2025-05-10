package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate GitHub Action workflows",
	Long: `Validate your GitHub Action workflows for syntax errors, 
best practices, and potential optimizations.`,
	Run: func(cmd *cobra.Command, args []string) {
		filePath, _ := cmd.Flags().GetString("file")
		fmt.Printf("Validating workflow file: %s\n", filePath)
		// Implement the validation logic here
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)

	// Add local flags for the validate command
	validateCmd.Flags().StringP("file", "f", "", "Path to the workflow file to validate")
	validateCmd.MarkFlagRequired("file")
}
