package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// EnsureDirectoryExists checks if a directory exists and creates it if it doesn't
func EnsureDirectoryExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.MkdirAll(path, 0755)
	}
	return nil
}

// IsGitHubRepo checks if the current directory is a GitHub repository
func IsGitHubRepo() bool {
	_, err := os.Stat(".git")
	return !os.IsNotExist(err)
}

// GetWorkflowsDir returns the path to the GitHub workflows directory
func GetWorkflowsDir() string {
	return filepath.Join(".github", "workflows")
}

// PrintSuccess prints a success message
func PrintSuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

// PrintError prints an error message
func PrintError(message string) {
	fmt.Printf("❌ %s\n", message)
}

// PrintWarning prints a warning message
func PrintWarning(message string) {
	fmt.Printf("⚠️ %s\n", message)
}

// PrintInfo prints an info message
func PrintInfo(message string) {
	fmt.Printf("ℹ️ %s\n", message)
}
