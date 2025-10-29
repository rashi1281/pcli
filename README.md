# 🚀 pcli - Internal CLI Tool

[![Go Version](https://img.shields.io/badge/go-1.24.2-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Version](https://img.shields.io/badge/version-dev-orange.svg)]()

A comprehensive developer productivity tool for managing logs, deployments, and other internal services. Built to streamline common development tasks and improve team efficiency with a beautiful, user-friendly interface.

## ✨ Features

- 📋 **Log Management** - View and stream application logs from AWS CloudWatch
- 💾 **Cache Management** - Manage CLI cache and cached data with smart persistence
- 🔧 **AWS Integration** - Seamless AWS service integration with auto-completion
- ⚡ **Auto-completion** - Smart command completion for log group names
- 🎨 **Beautiful UI** - Rich, emoji-enhanced interface with clear feedback
- 🔄 **Real-time Streaming** - Live log streaming with `tail -f` functionality
- 📅 **Time-based Filtering** - View logs from specific time ranges

## 🚀 Quick Start

### Installation

```bash
# Install the latest version
go install github.com/rashi1281/pcli@latest

# Or build from source
git clone https://github.com/rashi1281/pcli.git
cd pcli
go build -o pcli .
```

### Shell Completion

```bash
# Zsh
pcli completion zsh > "${fpath[1]}/_pcli"

# Bash
pcli completion bash > /etc/bash_completion.d/pcli

# Fish
pcli completion fish > ~/.config/fish/completions/pcli.fish
```

## 📖 Usage

### Basic Commands

```bash
# Show help and available commands
pcli --help

# Show version information
pcli --version

# View logs from a specific service
pcli logs tail my-service

# Stream logs in real-time
pcli logs tail my-service --follow

# View logs from the last hour
pcli logs tail my-service --since 1h

# Manage cache
pcli cache list
pcli cache refresh
pcli cache clear
```

### Global Flags

| Flag | Short | Description |
|------|-------|-------------|
| `--config` | | 📁 Config file path (default: $HOME/.pcli.json) |
| `--verbose` | `-V` | 🔍 Enable verbose output |
| `--quiet` | `-q` | 🔇 Suppress non-essential output |
| `--version` | `-v` | 📋 Show version information |

## 📋 Log Management

### Viewing Logs

```bash
# Basic log viewing
pcli logs tail my-service

# Real-time streaming (like tail -f)
pcli logs tail my-service --follow

# Time-based filtering
pcli logs tail my-service --since 30m
pcli logs tail my-service --since 2h
pcli logs tail my-service --since 1d

# Combined flags
pcli logs tail my-service -f -s 1h
```

### Supported Time Formats

- `30m` - 30 minutes
- `2h` - 2 hours
- `1d` - 1 day
- `90s` - 90 seconds

### Auto-completion

The tool provides intelligent auto-completion for log group names:

```bash
# Press Tab to see available log groups
pcli logs tail <TAB>
```

## 💾 Cache Management

### Cache Operations

```bash
# List all cached entries with details
pcli cache list

# Clear all cached data
pcli cache clear

# Refresh cache by fetching latest data
pcli cache refresh

# Get specific cached entry
pcli cache get log-groups
```

### Cache Features

- **Automatic Management** - Cache is automatically managed and refreshed when needed
- **Smart Persistence** - Cached data persists between sessions
- **Detailed Information** - View cache entries with type and size information
- **Error Handling** - Robust error handling with helpful suggestions

## 🔧 AWS Integration

### Prerequisites

- AWS CLI installed and configured
- Appropriate AWS permissions for CloudWatch Logs
- Valid AWS credentials (via AWS CLI, environment variables, or IAM roles)

### Required Permissions

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "logs:DescribeLogGroups",
        "logs:FilterLogEvents",
        "logs:GetLogEvents"
      ],
      "Resource": "*"
    }
  ]
}
```

## 🎨 User Interface

### Beautiful Output

The tool features a rich, emoji-enhanced interface that provides:

- 🚀 **Welcome Messages** - Friendly welcome screens
- 📊 **Status Indicators** - Clear success/error/warning indicators
- 🔄 **Progress Feedback** - Real-time operation status
- 📋 **Formatted Tables** - Clean, readable data presentation
- ❌ **Error Messages** - Helpful error messages with troubleshooting tips

### Example Output

```
🚀 Welcome to pcli - Internal CLI Tool

Available commands:
  logs    📋 View and stream application logs
  cache   💾 Manage CLI cache and data

Use 'pcli <command> --help' for more information about a command.
```

### Project Structure

```
pcli/
├── cmd/                    # Command implementations
│   ├── root.go            # Root command and configuration
│   ├── cache/             # Cache management commands
│   └── logs/              # Log management commands
├── internal/              # Internal packages
│   ├── aws.go            # AWS integration
│   ├── autocomplete.go   # Auto-completion logic
│   └── cache.go          # Cache utilities
├── main.go               # Application entry point
├── go.mod               # Go module definition
└── README.md            # This file
```

## 🤝 Contributing

We welcome contributions!

### Development Setup

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Troubleshooting

### Common Issues

**Q: "Error fetching logs: failed to run aws logs tail"**
- Ensure AWS CLI is installed and configured
- Check your AWS credentials and permissions
- Verify the log group exists

**Q: "Cache entry not found"**
- Run `pcli cache refresh` to update the cache
- Check if the log group name is correct

**Q: "Config file not found"**
- The tool will create a default configuration automatically
- Check file permissions in your home directory

### Getting Help

- Check the help: `pcli --help`
- Command-specific help: `pcli <command> --help`
- Enable verbose output: `pcli -V <command>`
- Check logs: `pcli logs tail <log-group> --follow`

## 🚀 Roadmap

- [ ] Support for multiple AWS profiles
- [ ] Log filtering and search capabilities
- [ ] Export logs to files
- [ ] Integration with other cloud providers
- [ ] Plugin system for custom commands
- [ ] Web dashboard for log monitoring

---