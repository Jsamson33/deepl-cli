package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// Create a temporary config file for testing
	tempDir := t.TempDir()

	// Set HOME environment variable to the temporary directory
	oldHome := os.Getenv("HOME")
	defer os.Setenv("HOME", oldHome) // Restore HOME after test
	os.Setenv("HOME", tempDir)

	configDir := filepath.Join(tempDir, ".deepl")
	err := os.MkdirAll(configDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp config directory: %v", err)
	}
	configPath := filepath.Join(configDir, ".deepl.toml")

	// Test case 1: Valid config file
	validConfigContent := `api_key = "test_api_key"
target_lang = "en"
source_lang = "fr"`
	err = os.WriteFile(configPath, []byte(validConfigContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write valid config file: %v", err)
	}

	cfg, err := Load()
	if err != nil {
		t.Errorf("Load() failed for valid config: %v", err)
	}
	if cfg.APIKey != "test_api_key" {
		t.Errorf("Expected APIKey 'test_api_key', got '%s'", cfg.APIKey)
	}
	if cfg.TargetLang != "en" {
		t.Errorf("Expected TargetLang 'en', got '%s'", cfg.TargetLang)
	}
	if cfg.SourceLang != "fr" {
		t.Errorf("Expected SourceLang 'fr', got '%s'", cfg.SourceLang)
	}

	// Test case 2: Missing config file
	os.Remove(configPath)
	_, err = Load()
	if err == nil || err.Error() != "config file not found: "+configPath {
		t.Errorf("Load() did not return expected error for missing config file. Got: %v", err)
	}

	// Test case 3: Missing API Key
	missingAPIKeyContent := `target_lang = "en"`
	err = os.WriteFile(configPath, []byte(missingAPIKeyContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write config file with missing API key: %v", err)
	}
	_, err = Load()
	if err == nil || err.Error() != "API key not found in config file" {
		t.Errorf("Load() did not return expected error for missing API key. Got: %v", err)
	}

	// Test case 4: Missing Target Language
	missingTargetLangContent := `api_key = "test_api_key"`
	err = os.WriteFile(configPath, []byte(missingTargetLangContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write config file with missing target lang: %v", err)
	}
	_, err = Load()
	if err == nil || err.Error() != "target language not found in config file" {
		t.Errorf("Load() did not return expected error for missing target lang. Got: %v", err)
	}
}
