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
		cmdArgs = append(cmdArgs, "--since", since.String())
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
