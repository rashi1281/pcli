/*
Copyright © 2025 Rashi M

pcli - Internal CLI Tool

A comprehensive developer productivity tool for managing logs, deployments,
and other internal services. Built to streamline common development tasks
and improve team efficiency.

Features:

	📋 Log Management    - View and stream application logs from AWS CloudWatch
	💾 Cache Management  - Manage CLI cache and cached data
	🔧 AWS Integration  - Seamless AWS service integration
	⚡ Auto-completion  - Smart command completion for better UX

Usage:

	pcli [command] [flags]

Examples:

	pcli logs tail my-service --follow
	pcli cache refresh
	pcli --help

For more information, visit: https://github.com/rashi1281/pcli
*/
package main

import "github.com/rashi1281/pcli/cmd"

// main is the entry point of the pcli application.
// It delegates execution to the cobra command structure defined in the cmd package.
func main() {
	// Execute the root command and all its subcommands
	// This will handle command parsing, flag processing, and command execution
	cmd.Execute()
}
