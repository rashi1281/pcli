package cache

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"

	"github.com/rashi1281/pcli/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cacheRefreshFunc = []func() error{
	internal.CacheLogGroups,
}

// CacheCmd represents the cache management command
var CacheCmd = &cobra.Command{
	Use:   "cache [command]",
	Short: "💾 Manage CLI cache and cached data",
	Long: `💾 Cache Management

Manage cached data for the CLI tool. The cache stores frequently accessed 
information like AWS log groups, service configurations, and other data 
to improve performance and reduce API calls.

Available Commands:
  clear    🗑️  Clear all cached data
  list     📋 List all cached entries with details
  get      🔍 Get specific cached entry by key
  refresh  🔄 Refresh cache by fetching latest data

Examples:
  pcli cache list                    # Show all cached entries
  pcli cache clear                   # Clear all cache
  pcli cache get log-groups          # Get cached log groups
  pcli cache refresh                 # Refresh all cache data

The cache is automatically managed and will be refreshed when needed.
Use 'pcli cache --help' for more information about specific commands.`,
	Args:      cobra.RangeArgs(1, 2),
	ValidArgs: []string{"clear", "list", "get", "refresh"},
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]

		switch command {
		case "clear":
			handleCacheClear()
		case "list":
			handleCacheList()
		case "get":
			if len(args) < 2 {
				fmt.Println("❌ Error: Key is required for 'get' command")
				fmt.Println("Usage: pcli cache get <key>")
				return
			}
			handleCacheGet(args[1])
		case "refresh":
			handleCacheRefresh()
		default:
			fmt.Printf("❌ Error: Unknown command '%s'\n", command)
			fmt.Println("Available commands: clear, list, get, refresh")
			cmd.Help()
		}
	},
}

// handleCacheClear clears all cached data
func handleCacheClear() {
	fmt.Println("🗑️  Clearing cache...")
	err := internal.DeleteConfigKey("cache")
	if err != nil {
		fmt.Printf("❌ Error clearing cache: %v\n", err)
		return
	}
	fmt.Println("✅ Cache cleared successfully")
}

// handleCacheList displays all cached entries in a formatted table
func handleCacheList() {
	cache := viper.GetStringMap("cache")

	if len(cache) == 0 {
		fmt.Println("📋 Cache is empty")
		return
	}

	fmt.Println("📋 Cached entries:")
	fmt.Println()

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"#", "Key", "Type", "Size"})

	i := 1
	for key, val := range cache {
		valType := fmt.Sprintf("%T", val)
		size := "N/A"

		// Try to get size information
		if valStr, ok := val.(string); ok {
			size = fmt.Sprintf("%d chars", len(valStr))
		} else if valMap, ok := val.(map[string]interface{}); ok {
			size = fmt.Sprintf("%d items", len(valMap))
		} else if valSlice, ok := val.([]interface{}); ok {
			size = fmt.Sprintf("%d items", len(valSlice))
		}

		table.Append([]string{
			fmt.Sprintf("%d", i),
			key,
			valType,
			size,
		})
		i++
	}

	table.Render()
	fmt.Printf("\n📊 Total entries: %d\n", len(cache))
}

// handleCacheGet retrieves and displays a specific cached entry
func handleCacheGet(key string) {
	cacheKey := fmt.Sprintf("cache.%s", key)
	cache := viper.Get(cacheKey)

	if cache == nil {
		fmt.Printf("❌ Cache entry '%s' not found\n", key)
		fmt.Println("Use 'pcli cache list' to see available entries")
		return
	}

	fmt.Printf("🔍 Cache entry '%s':\n", key)
	fmt.Println()

	jsonString, err := json.MarshalIndent(cache, "", "  ")
	if err != nil {
		fmt.Printf("❌ Error formatting cache data: %v\n", err)
		return
	}

	fmt.Println(string(jsonString))
}

// handleCacheRefresh refreshes all cached data
func handleCacheRefresh() {
	fmt.Println("🔄 Refreshing cache...")

	successCount := 0
	totalCount := len(cacheRefreshFunc)

	for i, refreshFunc := range cacheRefreshFunc {
		if err := refreshFunc(); err != nil {
			fmt.Printf("⚠️  Warning: Error refreshing cache item %d: %v\n", i+1, err)
		} else {
			successCount++
		}
	}

	if successCount == totalCount {
		fmt.Println("✅ Cache refreshed successfully")
	} else {
		fmt.Printf("⚠️  Cache refresh completed with %d/%d successful\n", successCount, totalCount)
	}
}

func init() {
}
