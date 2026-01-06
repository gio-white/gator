package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func GetConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir: %w", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}


func Read() (Config, error) {
	configPath, err := GetConfigFilePath()
	if err != nil {
        return Config{}, err
    }

	file, err := os.ReadFile(configPath)
    if err != nil {
        return Config{}, fmt.Errorf("could not read config file at %s: %w", configPath, err)
    }

    var cfg Config
    err = json.Unmarshal(file, &cfg)
    if err != nil {
        return Config{}, fmt.Errorf("failed to decode json: %w", err)
    }

    return cfg, nil
}