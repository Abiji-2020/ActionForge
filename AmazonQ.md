# Amazon Q Integration

ActionForge leverages Amazon Q to provide intelligent suggestions and validations for your GitHub Actions workflows.

## Features

- **Smart Workflow Creation**: Get intelligent suggestions for workflow steps based on your project type and requirements.
- **Workflow Validation**: Validate your workflows against best practices and identify potential issues.
- **Optimization Recommendations**: Receive recommendations to optimize your workflows for better performance and cost efficiency.

## Configuration

Amazon Q integration can be configured in your `.actionforge.yaml` file:

```yaml
amazonQ:
  enabled: true
  region: us-east-1
```

## Requirements

To use Amazon Q integration, you need:

1. AWS CLI configured with appropriate credentials
2. Permissions to access Amazon Q services
3. Internet connectivity to AWS services

## Usage

Amazon Q integration is automatically used when creating, validating, or optimizing workflows. You can disable it by setting `amazonQ.enabled` to `false` in your configuration file.
