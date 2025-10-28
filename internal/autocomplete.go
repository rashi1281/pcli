package internal

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Struct for unmarshalling JSON from `aws logs describe-log-groups`
type logGroupsResponse struct {
	LogGroups []struct {
		LogGroupName string `json:"logGroupName"`
	} `json:"logGroups"`
}

// AutoCompleteLogGroups dynamically fetches CloudWatch log groups for completion.
func AutoCompleteLogGroups(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {

	logGroup := viper.GetStringSlice("cache.log_groups")
	if len(logGroup) <= 0 {
		out, err := exec.Command("aws", "logs", "describe-log-groups", "--output", "json").Output()
		if err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		var resp logGroupsResponse
		if err := json.Unmarshal(out, &resp); err != nil {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}

		toComplete = strings.ToLower(toComplete)
		logGroup = make([]string, len(resp.LogGroups))
		for i, lg := range resp.LogGroups {
			logGroup[i] = lg.LogGroupName
		}
		viper.Set("cache.log_groups", logGroup)
		viper.WriteConfig()
	}

	var suggestions []string
	for _, lg := range logGroup {
		name := lg
		// case-insensitive substring match
		if strings.Contains(strings.ToLower(name), toComplete) {
			suggestions = append(suggestions, name)
		}
	}

	return suggestions, cobra.ShellCompDirectiveNoFileComp
}
