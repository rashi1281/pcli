/*
Copyright © 2025 Rashi M
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rashi1281/pcli/cmd/cache"
	"github.com/rashi1281/pcli/cmd/logs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "pcli",
	Short: "🚀 Internal CLI - Developer productivity tool",
	Long: `🚀 pcli - Internal CLI Tool

A comprehensive developer productivity tool for managing logs, deployments, 
and other internal services. Built to streamline common development tasks 
and improve team efficiency.

Features:
  📋 Log Management    - View and stream application logs
  💾 Cache Management  - Manage CLI cache and data
  🔧 AWS Integration  - Seamless AWS service integration
  ⚡ Auto-completion  - Smart command completion

Examples:
  pcli logs tail my-service --follow
  pcli cache refresh
  pcli cache list
  pcli --help

For more information about a specific command, use:
  pcli <command> --help`,
	Run: func(cmd *cobra.Command, args []string) {
		// Display welcome message and available commands
		fmt.Println("🚀 Welcome to pcli - Internal CLI Tool")
		fmt.Println()
		fmt.Println("Available commands:")
		fmt.Println("  logs    📋 View and stream application logs")
		fmt.Println("  cache   💾 Manage CLI cache and data")
		fmt.Println()
		fmt.Println("Use 'pcli <command> --help' for more information about a command.")
		fmt.Println("Use 'pcli --help' to see all available options.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(logs.LogsCmd)
	rootCmd.AddCommand(cache.CacheCmd)

	// Global configuration flags available to all commands
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "",
		"📁 Config file path (default: $HOME/.pcli.json)")

	// Add version flag
	rootCmd.Flags().BoolP("version", "v", false, "📋 Show version information")

	// Add verbose flag for detailed output
	rootCmd.PersistentFlags().BoolP("verbose", "V", false, "🔍 Enable verbose output")

	// Add quiet flag for minimal output
	rootCmd.PersistentFlags().BoolP("quiet", "q", false, "🔇 Suppress non-essential output")

	// Bind flags to viper so initConfig can respect --quiet/--verbose
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))

	// Handle --version: print version and exit before running commands
	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			// Version could be set at build time via -ldflags "-X main.version=..."
			// Fallback to 'dev' if not set
			version := os.Getenv("PCLI_VERSION")
			if version == "" {
				version = "dev"
			}
			fmt.Printf("pcli version %s\n", version)
			os.Exit(0)
		}
	}
}

// initConfig reads in config file and ENV variables if set.
// It handles configuration initialization with proper error handling and user feedback.
func initConfig() {
	// Set configuration file path
	if cfgFile != "" {
		// Use config file from the command line flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory for default config location
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Printf("⚠️  Warning: Could not find home directory: %v\n", err)
			return
		}

		// Search for config file in home directory with name ".pcli.json"
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".pcli")
	}

	// Enable automatic environment variable reading
	// Environment variables will override config file values
	viper.AutomaticEnv()

	// Try to read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found, create a default one
			if !viper.GetBool("quiet") {
				fmt.Println("📁 Config file not found, creating default configuration...")
			}

			// Set some default values
			viper.SetDefault("version", "v0.1.0")

			// Write the default config file
			if err := viper.SafeWriteConfig(); err != nil {
				fmt.Printf("⚠️  Warning: Could not create config file: %v\n", err)
			} else if !viper.GetBool("quiet") {
				fmt.Println("✅ Default configuration created successfully")
			}
		} else {
			// Other configuration errors
			fmt.Printf("⚠️  Warning: Error reading config file: %v\n", err)
		}
	} else if viper.GetBool("verbose") {
		// Config file loaded successfully
		fmt.Printf("📁 Using config file: %s\n", viper.ConfigFileUsed())
	}
}
