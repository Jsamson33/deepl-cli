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

func Load() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	configFilePath := filepath.Join(homeDir, ".deepl", ".deepl.toml")
	if _, err := os.Stat(configFilePath); os.IsNotExist(err) {
		return nil, errors.New("config file not found: " + configFilePath)
	}

	var cfg Config
	if _, err := toml.DecodeFile(configFilePath, &cfg); err != nil {
		return nil, err
	}

	if cfg.APIKey == "" {
		return nil, errors.New("API key not found in config file")
	}

	if cfg.TargetLang == "" {
		return nil, errors.New("target language not found in config file")
	}

	if cfg.SourceLang == "" {
		cfg.SourceLang = "en" // Default source language if not specified
	}

	return &cfg, nil
}
