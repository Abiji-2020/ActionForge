package actions

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

// WorkflowFile represents a GitHub Actions workflow file
type WorkflowFile struct {
	Name     string                 `yaml:"name"`
	On       interface{}            `yaml:"on"`
	Jobs     map[string]Job         `yaml:"jobs"`
	Env      map[string]string      `yaml:"env,omitempty"`
	Defaults map[string]interface{} `yaml:"defaults,omitempty"`
}

// Job represents a job in a GitHub Actions workflow
type Job struct {
	Name        string                 `yaml:"name,omitempty"`
	RunsOn      interface{}            `yaml:"runs-on"`
	Needs       interface{}            `yaml:"needs,omitempty"`
	If          string                 `yaml:"if,omitempty"`
	Steps       []Step                 `yaml:"steps"`
	Environment interface{}            `yaml:"environment,omitempty"`
	Env         map[string]string      `yaml:"env,omitempty"`
	Defaults    map[string]interface{} `yaml:"defaults,omitempty"`
	Strategy    map[string]interface{} `yaml:"strategy,omitempty"`
	Container   interface{}            `yaml:"container,omitempty"`
	Services    map[string]interface{} `yaml:"services,omitempty"`
	Outputs     map[string]string      `yaml:"outputs,omitempty"`
	Timeout     string                 `yaml:"timeout-minutes,omitempty"`
}

// Step represents a step in a GitHub Actions job
type Step struct {
	Name      string            `yaml:"name,omitempty"`
	ID        string            `yaml:"id,omitempty"`
	If        string            `yaml:"if,omitempty"`
	Uses      string            `yaml:"uses,omitempty"`
	Run       string            `yaml:"run,omitempty"`
	With      map[string]string `yaml:"with,omitempty"`
	Env       map[string]string `yaml:"env,omitempty"`
	ContinueOnError bool        `yaml:"continue-on-error,omitempty"`
	TimeoutMinutes  int         `yaml:"timeout-minutes,omitempty"`
}

// ParseWorkflow parses a GitHub Actions workflow file
func ParseWorkflow(filePath string) (*WorkflowFile, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var workflow WorkflowFile
	err = yaml.Unmarshal(data, &workflow)
	if err != nil {
		return nil, err
	}

	return &workflow, nil
}

// SaveWorkflow saves a GitHub Actions workflow to a file
func SaveWorkflow(workflow *WorkflowFile, filePath string) error {
	data, err := yaml.Marshal(workflow)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(filePath, data, 0644)
}
