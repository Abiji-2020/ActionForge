package amazonq

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/actionforge/internal/config"
)

// DetectLanguage determines the primary language of the codebase
func DetectLanguage() (string, error) {
	// Check for common language indicators
	languages := map[string][]string{
		"go":         {".go", "go.mod", "go.sum"},
		"javascript": {".js", "package.json", "node_modules"},
		"typescript": {".ts", "tsconfig.json"},
		"python":     {".py", "requirements.txt", "setup.py", "Pipfile"},
		"java":       {".java", "pom.xml", "build.gradle"},
		"ruby":       {".rb", "Gemfile"},
		"php":        {".php", "composer.json"},
		"csharp":     {".cs", ".sln", ".csproj"},
		"rust":       {".rs", "Cargo.toml"},
	}

	files, err := os.ReadDir(".")
	if err != nil {
		return "", err
	}

	counts := make(map[string]int)
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		
		name := file.Name()
		
		for lang, indicators := range languages {
			for _, indicator := range indicators {
				if strings.HasSuffix(name, indicator) {
					counts[lang]++
				}
			}
		}
	}

	// Find the language with the highest count
	maxCount := 0
	primaryLang := "unknown"
	for lang, count := range counts {
		if count > maxCount {
			maxCount = count
			primaryLang = lang
		}
	}

	return primaryLang, nil
}

// GenerateWorkflow uses Amazon Q to generate a GitHub Actions workflow
func GenerateWorkflow(workflowType string, language string, cfg *config.Config) (string, error) {
	if !cfg.AmazonQ.Enabled {
		return "", fmt.Errorf("Amazon Q integration is disabled in configuration")
	}

	// Prepare the prompt for Amazon Q
	prompt := fmt.Sprintf("Generate a GitHub Actions workflow for a %s project with %s workflow. Include best practices and optimizations. Return only the YAML content without any explanations.", language, workflowType)

	// Execute the Amazon Q CLI command
	cmd := exec.Command("q", "chat", "--no-interactive", "--trust-all-tools", prompt)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to execute Amazon Q: %v\n%s", err, stderr.String())
	}

	// Extract the YAML content from the response
	response := stdout.String()
	
	// Clean up the response to extract just the YAML
	yamlContent := extractYAML(response)
	if yamlContent == "" {
		return "", fmt.Errorf("failed to extract YAML content from Amazon Q response")
	}

	return yamlContent, nil
}

// extractYAML extracts YAML content from the Amazon Q response
func extractYAML(response string) string {
	// Look for content between triple backticks
	parts := strings.Split(response, "```yaml")
	if len(parts) < 2 {
		// Try without the language specifier
		parts = strings.Split(response, "```")
		if len(parts) < 2 {
			return response // Return the whole response if no code blocks found
		}
	}
	
	// Get the content after the opening backticks
	content := parts[1]
	
	// Remove everything after the closing backticks if they exist
	if idx := strings.Index(content, "```"); idx != -1 {
		content = content[:idx]
	}
	
	return strings.TrimSpace(content)
}
