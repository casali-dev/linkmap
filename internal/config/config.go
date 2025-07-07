package config

import (
	"fmt"
	"os"
	"strings"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	BaseURL string
	DBPath  string
	Port    string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		BaseURL: os.Getenv("BASE_URL"),
		DBPath:  os.Getenv("DB_PATH"),
		Port:    os.Getenv("PORT"),
	}

	var missing []string

	if cfg.BaseURL == "" {
		missing = append(missing, "BASE_URL")
	}
	if cfg.DBPath == "" {
		missing = append(missing, "DB_PATH")
	}
	if cfg.Port == "" {
		missing = append(missing, "PORT")
	}

	if len(missing) > 0 {
		return nil, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return cfg, nil
}
