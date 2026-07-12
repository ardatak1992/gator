package config

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	configFileName = ".gatorconfig.json"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(newUser string) error {
	cfg.CurrentUserName = newUser
	err := write(*cfg)
	if err != nil {
		return fmt.Errorf("error setting username: %v", err)
	}
	return nil
}

func Read() (Config, error) {
	configFileDir, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("Error getting home dir: %v", err)
	}

	fileBytes, err := os.ReadFile(configFileDir)
	if err != nil {
		return Config{}, fmt.Errorf("Error getting home dir: %v", err)
	}

	var cfg Config
	err = json.Unmarshal(fileBytes, &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("Error unmarshaling config file: %v", err)
	}

	return cfg, nil
}
