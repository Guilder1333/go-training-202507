package userapp

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

type UserAppConfig struct {
	MySQL MySQLConfig `json:"mysql"`
}

type MySQLConfig struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Address  string `json:"address"`
	DBName   string `json:"dbName"`
}

func loadConfig() (*UserAppConfig, error) {
	path := os.Getenv("CONFIG_FILE_NAME")

	log.Info().Msg(fmt.Sprintf("Loading config file '%s'", path))
	configJson, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %w", err)
	}

	var config UserAppConfig
	err = json.Unmarshal(configJson, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	return &config, nil
}
