package internal

import (
	"bytes"
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

// Struct for unmarshalling JSON from `aws logs describe-log-groups`
type logGroupsResponse struct {
	LogGroups []struct {
		LogGroupName string `json:"logGroupName"`
	} `json:"logGroups"`
}

// AutoCompleteLogGroups dynamically fetches CloudWatch log groups for completion.
func AutoCompleteLogGroups(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	// Call AWS CLI to list log groups
	out, err := exec.Command("aws", "logs", "describe-log-groups", "--output", "json").Output()
	if err != nil {
		// Return empty but prevent file completion
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// Parse JSON response
	var resp logGroupsResponse
	dec := json.NewDecoder(bytes.NewReader(out))
	if err := dec.Decode(&resp); err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	// Collect matching log groups
	var suggestions []string
	for _, lg := range resp.LogGroups {
		name := lg.LogGroupName
		if toComplete == "" || strings.HasPrefix(name, toComplete) {
			suggestions = append(suggestions, name)
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}
