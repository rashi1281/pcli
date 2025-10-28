package internal

import (
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
	out, err := exec.Command("aws", "logs", "describe-log-groups", "--output", "json").Output()
	if err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	var resp logGroupsResponse
	if err := json.Unmarshal(out, &resp); err != nil {
		return nil, cobra.ShellCompDirectiveNoFileComp
	}

	toComplete = strings.ToLower(toComplete)
	var suggestions []string

	for _, lg := range resp.LogGroups {
		name := lg.LogGroupName
		// case-insensitive substring match
		if strings.Contains(strings.ToLower(name), toComplete) {
			suggestions = append(suggestions, name)
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}
