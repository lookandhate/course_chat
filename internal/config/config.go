package config

import (
	"flag"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

const DefaultConfig = "config/local.config.yml"

// AppConfig config for application.
type AppConfig struct {
	Env  string         `yaml:"env" env-default:"local"`
	GPRC GRPCConfig     `yaml:"gprc"`
	DB   PostgresConfig `yaml:"db"`
}

// MustLoad creates AppConfig and loads it from yaml file.
func MustLoad() *AppConfig {
	configPath := fetchConfigPath()
	if configPath == "" {
		panic("config path is empty")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}

	var cfg AppConfig

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("config path is empty: " + err.Error())
	}

	return &cfg
}

// fetchConfigPath fetches config path from command line flag or environment variable.
// Priority: flag > env > default.
// Default value is empty string.
func fetchConfigPath() string {
	var res string

	flag.StringVar(&res, "config", "", "path to config file")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	if res == "" {
		res = DefaultConfig
	}

	return res
}
