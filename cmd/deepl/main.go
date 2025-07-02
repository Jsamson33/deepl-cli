package main

import (
	"fmt"
	"log"

	"deepl-cli/internal/config"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	fmt.Printf("Source Lang: %s\n", cfg.SourceLang)
	fmt.Printf("Target Lang: %s\n", cfg.TargetLang)
}