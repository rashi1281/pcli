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

// tailCmd represents the tail command
var tailCmd = &cobra.Command{
	Use:               "tail [log-group]",
	Short:             "View / stream service logs",
	Long:              `Fetch application logs from a specific log group, with options to follow live output or query recent history.`,
	Args:              cobra.ExactArgs(1),
	ValidArgsFunction: internal.AutoCompleteLogGroups,
	Run: func(cmd *cobra.Command, args []string) {
		logGroup := args[0]
		if logGroup == "" {
			cmd.Help()
		}

		err := internal.GetLogs(logGroup, follow, since)
		if err != nil {
			fmt.Println("Error:", err)
		}

	},
}

func init() {

	LogsCmd.AddCommand(tailCmd)

	// Flags
	// --follow, -f
	tailCmd.Flags().BoolVarP(&follow, "follow", "f", false, "Stream logs in real time (like tail -f)")

	// --since, -s
	tailCmd.Flags().DurationVarP(&since, "since", "s", 0*time.Hour, "How far back to fetch logs (e.g. 10m, 1h, 24h). Ignored with --follow")
}
