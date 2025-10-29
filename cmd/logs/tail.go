package logs

import (
	"fmt"
	"time"

	"github.com/rashi1281/pcli/internal"
	"github.com/spf13/cobra"
)

var (
	follow bool
	since  time.Duration
)

// tailCmd represents the tail command for streaming logs
var tailCmd = &cobra.Command{
	Use:   "tail [log-group]",
	Short: "📊 Stream logs from a specific log group",
	Long: `📊 Stream Logs

Stream application logs from a specific AWS CloudWatch Log Group. This command 
provides real-time log streaming similar to 'tail -f' but for AWS CloudWatch Logs.

Features:
  🔄 Real-time streaming    - Follow logs as they're written (--follow)
  📅 Time-based filtering   - View logs from specific time ranges (--since)
  ⚡ Auto-completion       - Tab completion for log group names
  📊 Clean formatting      - Formatted, readable log output
  🎯 Smart filtering       - Built-in log filtering and highlighting

The command will automatically detect the log group and stream logs accordingly.
Use Ctrl+C to stop streaming when using --follow mode.

Examples:
  pcli logs tail my-service                    # View recent logs
  pcli logs tail my-service --follow          # Stream logs in real-time
  pcli logs tail my-service --since 1h        # View logs from last hour
  pcli logs tail my-service -f -s 30m         # Stream logs from last 30 minutes
  pcli logs tail /aws/lambda/my-function      # Stream Lambda function logs`,
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: internal.AutoCompleteLogGroups,
	Run: func(cmd *cobra.Command, args []string) {
		logGroup := args[0]

		// Validate log group name
		if logGroup == "" {
			fmt.Println("❌ Error: Log group name is required")
			fmt.Println("Usage: pcli logs tail <log-group>")
			return
		}

		// Display operation info
		if follow {
			fmt.Printf("🔄 Streaming logs from '%s' (Press Ctrl+C to stop)...\n", logGroup)
		} else {
			fmt.Printf("📊 Fetching logs from '%s'", logGroup)
			if since > 0 {
				fmt.Printf(" (since %v)", since)
			}
			fmt.Println("...")
		}
		fmt.Println()

		// Fetch and display logs
		err := internal.GetLogs(logGroup, follow, since)
		if err != nil {
			fmt.Printf("❌ Error fetching logs: %v\n", err)
			fmt.Println()
			fmt.Println("Troubleshooting:")
			fmt.Println("  • Check if the log group exists")
			fmt.Println("  • Verify AWS credentials and permissions")
			fmt.Println("  • Use 'pcli cache refresh' to update log groups")
			return
		}

		if !follow {
			fmt.Println()
			fmt.Println("✅ Log fetch completed")
			fmt.Println("Use --follow to stream logs in real-time")
		}
	},
}

func init() {
	// Add tail command to logs command
	LogsCmd.AddCommand(tailCmd)

	// Define command flags with improved descriptions
	tailCmd.Flags().BoolVarP(&follow, "follow", "f", false,
		"🔄 Stream logs in real-time (like tail -f). Use Ctrl+C to stop.")

	tailCmd.Flags().DurationVarP(&since, "since", "s", 0*time.Hour,
		"📅 How far back to fetch logs (e.g. 10m, 1h, 24h). Ignored with --follow")
}
