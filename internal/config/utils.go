package config

import (
	"encoding/json"
	"fmt"
	"os"
)

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homeDir + "/" + configFileName, nil
}

func write(cfg Config) error {
	configFileDir, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("error getting home dir: %v", err)
	}

	data, err := json.Marshal(cfg)

	err = os.WriteFile(configFileDir, data, 0666)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}

	return err
}
