package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/actionforge/internal/config"
	"github.com/actionforge/pkg/amazonq"
	"github.com/actionforge/pkg/utils"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new GitHub Action workflow",
	Long: `Create a new GitHub Action workflow with intelligent suggestions.
You can specify the type of workflow you want to create and ActionForge
will guide you through the process.`,
	Run: func(cmd *cobra.Command, args []string) {
		workflowType, _ := cmd.Flags().GetString("type")
		outputDir, _ := cmd.Flags().GetString("output")
		filename, _ := cmd.Flags().GetString("filename")
		
		utils.PrintInfo(fmt.Sprintf("Creating a new %s workflow...", workflowType))
		
		// Load configuration
		cfg, err := config.LoadConfig("")
		if err != nil {
			utils.PrintError(fmt.Sprintf("Failed to load configuration: %v", err))
			return
		}
		
		// Detect the primary language of the codebase
		language, err := amazonq.DetectLanguage()
		if err != nil {
			utils.PrintWarning(fmt.Sprintf("Could not detect language: %v", err))
			language = "unknown"
		}
		utils.PrintInfo(fmt.Sprintf("Detected primary language: %s", language))
		
		// Generate workflow using Amazon Q
		utils.PrintInfo("Generating workflow with Amazon Q...")
		yamlContent, err := amazonq.GenerateWorkflow(workflowType, language, cfg)
		if err != nil {
			utils.PrintError(fmt.Sprintf("Failed to generate workflow: %v", err))
			return
		}
		
		// Ensure the output directory exists
		if err := utils.EnsureDirectoryExists(outputDir); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to create output directory: %v", err))
			return
		}
		
		// Write the workflow file
		outputPath := filepath.Join(outputDir, filename)
		if err := os.WriteFile(outputPath, []byte(yamlContent), 0644); err != nil {
			utils.PrintError(fmt.Sprintf("Failed to write workflow file: %v", err))
			return
		}
		
		utils.PrintSuccess(fmt.Sprintf("Workflow created successfully: %s", outputPath))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Add local flags for the create command
	createCmd.Flags().StringP("type", "t", "ci", "Type of workflow to create (ci, cd, etc.)")
	createCmd.Flags().StringP("output", "o", ".github/workflows/", "Output directory for the workflow file")
	createCmd.Flags().StringP("filename", "f", "github-actions.yml", "Name of the workflow file")
}
