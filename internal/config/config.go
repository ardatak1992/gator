package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	write(*c)
	return nil
}

func Read() (Config, error) {

	filePath, _ := getConfigFilePath()

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)

	var config Config

	if err := decoder.Decode(&config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(homeDir, configFileName)
	return filePath, nil
}

func write(cfg Config) error {
	filepath, _ := getConfigFilePath()
	f, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("unable to open file %v", err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}
