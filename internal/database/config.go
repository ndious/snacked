package database

import (
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/ndious/snacked/internal"
)

type DbConfigSchema struct {
	User     string
	Password string
	Database string
	Url      string
	Port     int32
}

type ConfigSchema struct {
	Db DbConfigSchema `toml:"db"`
}

func GetConfig() (ConfigSchema, error) {
	configFilePath := filepath.Join(internal.GetDir("config"), "config.toml")

	content, err := os.ReadFile(configFilePath)

	if err != nil {
		return ConfigSchema{}, err
	}

	var config ConfigSchema

	if err := toml.Unmarshal(content, &config); err != nil {
		return ConfigSchema{}, err
	}

	return config, nil
}
