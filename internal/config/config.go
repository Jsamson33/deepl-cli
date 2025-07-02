package config

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type Config struct {
	APIKey     string `toml:"api_key"`
	SourceLang string `toml:"source_lang"`
	TargetLang string `toml:"target_lang"`
}

func LoadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configPath := filepath.Join(home, ".deepl", ".deepl.toml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, errors.New("config file not found: " + configPath)
	}

	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}

	if config.APIKey == "" {
		return nil, errors.New("api_key not found in config file")
	}

	if config.SourceLang == "" {
		config.SourceLang = "fr"
	}

	if config.TargetLang == "" {
		config.TargetLang = "en"
	}

	return &config, nil
}
