package config

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

var Modules = fx.Options(
	fx.Provide(loadConfig),
)

type DatabaseConfig struct {
	DBType string `yaml:"db_type"`
	DSN    string `yaml:"dsn"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

const (
	configFilePath = "./config.yaml"
)

func loadConfig() *Config {
	file, err := os.Open(configFilePath)
	if err != nil {
		panic(fmt.Errorf("could not open config file: %v", err))
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("could not read config file: %v", err))
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		panic(fmt.Errorf("could not parse config file: %v", err))
	}
	return &config
}
