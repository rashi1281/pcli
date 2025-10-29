package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

// DeleteConfigKey removes a top-level key (like "cache") from the file
// that Viper is currently using, and persists the change.
// Works for YAML (.yml/.yaml) and JSON (.json).
func DeleteConfigKey(key string) error {
	// 1. Find which file viper is actually using
	cfgPath := viper.ConfigFileUsed()
	if cfgPath == "" {
		return fmt.Errorf("no config file bound to viper (ConfigFileUsed() is empty)")
	}

	// 2. Read raw file bytes
	raw, err := os.ReadFile(cfgPath)
	if err != nil {
		return fmt.Errorf("read config: %w", err)
	}

	// 3. Figure out format from extension
	ext := strings.ToLower(filepath.Ext(cfgPath))

	// We'll load into a generic map
	data := map[string]any{}

	switch ext {
	case ".yml", ".yaml":
		if err := yaml.Unmarshal(raw, &data); err != nil {
			return fmt.Errorf("unmarshal yaml: %w", err)
		}
	case ".json":
		if err := json.Unmarshal(raw, &data); err != nil {
			return fmt.Errorf("unmarshal json: %w", err)
		}
	default:
		return fmt.Errorf("unsupported config format: %s", ext)
	}

	// 4. Delete the key from the top level
	delete(data, key)

	// 5. Marshal back to original format
	var updated []byte
	switch ext {
	case ".yml", ".yaml":
		updated, err = yaml.Marshal(data)
		if err != nil {
			return fmt.Errorf("marshal yaml: %w", err)
		}
	case ".json":
		updated, err = json.MarshalIndent(data, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal json: %w", err)
		}
		updated = append(updated, '\n') // pretty end newline
	}

	// 6. Write back to same file
	if err := os.WriteFile(cfgPath, updated, 0o644); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	// 7. Reload into viper so in-memory matches disk
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("reload viper: %w", err)
	}

	return nil
}
