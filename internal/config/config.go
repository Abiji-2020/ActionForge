package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Config represents the application configuration
type Config struct {
	DefaultOutputDir string            `yaml:"defaultOutputDir"`
	Templates        map[string]string `yaml:"templates"`
	AmazonQ          AmazonQConfig     `yaml:"amazonQ"`
}

// AmazonQConfig represents the Amazon Q configuration
type AmazonQConfig struct {
	Enabled bool   `yaml:"enabled"`
	Region  string `yaml:"region"`
}

// LoadConfig loads the configuration from the specified file
func LoadConfig(configPath string) (*Config, error) {
	if configPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return nil, err
		}
		configPath = filepath.Join(homeDir, ".actionforge.yaml")
	}

	// Check if config file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// Return default config if file doesn't exist
		return getDefaultConfig(), nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// SaveConfig saves the configuration to the specified file
func SaveConfig(config *Config, configPath string) error {
	if configPath == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return err
		}
		configPath = filepath.Join(homeDir, ".actionforge.yaml")
	}

	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// getDefaultConfig returns the default configuration
func getDefaultConfig() *Config {
	return &Config{
		DefaultOutputDir: ".github/workflows",
		Templates: map[string]string{
			"ci": "templates/ci.yml",
			"cd": "templates/cd.yml",
		},
		AmazonQ: AmazonQConfig{
			Enabled: true,
			Region:  "us-east-1",
		},
	}
}
