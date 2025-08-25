package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {

	filePath, _ := getConfigFilePath()

	fmt.Println(filePath)

	jsonFile, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	defer jsonFile.Close()

	data, err := io.ReadAll(jsonFile)
	if err != nil {
		return Config{}, err
	}

	var conf Config

	if err := json.Unmarshal(data, &conf); err != nil {
		return Config{}, err
	}

	return conf, nil
}

func SetUser(username string) error {

}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	filePath := filepath.Join(homeDir, configFileName)
	return filePath, nil
}

func write() error {

}
