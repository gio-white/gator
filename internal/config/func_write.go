package config

import (
	"encoding/json"
	"os"
	"fmt"
)

func Write(cfg Config) error {
    data, err := json.MarshalIndent(cfg, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal config: %w", err)
    }

    path, err := GetConfigFilePath()
    if err != nil {
        return err
    }

    err = os.WriteFile(path, data, 0644)
    if err != nil {
        return fmt.Errorf("failed to write config file: %w", err)
    }

    return nil
}