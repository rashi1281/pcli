package internal

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

func GetLogs(logGroup string, follow bool, since time.Duration) error {
	// Build base command
	cmdArgs := []string{"logs", "tail", logGroup}

	// Add flags dynamically
	if follow {
		cmdArgs = append(cmdArgs, "--follow")
	}
	if since > 0 {
		// Format duration for AWS CLI (e.g. "10m", "1h")
		cmdArgs = append(cmdArgs, "--since", formatSinceForAWS(since))
	}

	// Create the command
	awsCmd := exec.Command("aws", cmdArgs...)

	// Stream output directly to terminal
	awsCmd.Stdout = os.Stdout
	awsCmd.Stderr = os.Stderr
	awsCmd.Stdin = os.Stdin // allows Ctrl+C interrupt to propagate

	// Print what we're running (for debug)
	fmt.Printf("Running: aws %s\n\n", strings.Join(cmdArgs, " "))

	// Execute
	if err := awsCmd.Run(); err != nil {
		return fmt.Errorf("failed to run aws logs tail: %w", err)
	}

	return nil
}

// formatSinceForAWS converts a duration into the compact format expected by
// `aws logs tail --since`, e.g. 90m, 2h, 1d. It rounds toward the largest
// sensible unit and avoids verbose strings like "1h0m0s".
func formatSinceForAWS(d time.Duration) string {
	if d <= 0 {
		return "0s"
	}
	// Prefer days, hours, then minutes
	if d%(24*time.Hour) == 0 {
		return fmt.Sprintf("%dd", int(d/(24*time.Hour)))
	}
	if d%time.Hour == 0 {
		return fmt.Sprintf("%dh", int(d/time.Hour))
	}
	if d%time.Minute == 0 {
		return fmt.Sprintf("%dm", int(d/time.Minute))
	}
	// Fallback to seconds for odd durations
	return fmt.Sprintf("%ds", int(d/time.Second))
}
