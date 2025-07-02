# DeepL CLI Project Technical Architecture

## 1. Project Overview

The DeepL CLI project is a command-line utility developed in Go (Golang) that allows users to translate text by leveraging the DeepL API. The main objective is to provide a simple, fast, and configurable tool for text translations directly from the terminal.

## 2. General Architecture

The project's architecture is designed to be modular and easy to maintain, separating concerns between command-line management, configuration, and interaction with the external API.

```mermaid
graph TD
    A[User] -->|Executes deepl command| B(CLI - cmd/deepl/main.go)
    B -->|Parses arguments (cobra)| C{Root Command - cmd/deepl/cmd/root.go}
    C -->|Loads configuration| D[Configuration - internal/config/config.go]
    C -->|Prioritizes arguments over config|
    C -->|Calls DeepL API| E(DeepL API)
    E -->|Returns translation| C
    C -->|Displays result| A
    D -->|.deepl.toml (config file)|
```

## 3. Key Components

### 3.1. `cmd/deepl/main.go`

Main entry point of the application. It initializes and executes the root command defined by Cobra.

### 3.2. `cmd/deepl/cmd/root.go`

Contains the definition of the root command (`deepl`) and handles the parsing of command-line arguments (`-s`, `-t`, text to translate). It also orchestrates the loading of the configuration and the call to the translation function.

### 3.3. `internal/config/config.go`

Manages the loading and validation of the `.deepl.toml` configuration file. This file is located in the user's `~/.deepl/` directory and contains the DeepL API key, as well as the default source and target languages.

### 3.4. Translation Logic (`translateText` function in `root.go`)

This function is responsible for constructing the HTTP request to the DeepL API, sending the request, handling responses (success or errors), and extracting the translated text.

## 4. Data Flow

1.  The user executes the `deepl` command with the text to translate and optional flags (source language, target language).
2.  The `main.go` module launches the root command.
3.  The root command (`root.go`) uses Cobra to parse CLI arguments.
4.  It loads the configuration from `.deepl.toml` via `config.go`.
5.  CLI arguments take precedence over default values in the configuration file.
6.  The `translateText` function is called with the text, languages (source and target), and the API key.
7.  `translateText` constructs an HTTP POST request to the DeepL API, including the text, languages, and API key (via the `Authorization` header).
8.  The DeepL API processes the request and returns the translation.
9.  The API response is read, deserialized, and the translated text is extracted.
10. The translated text is displayed on the user's console.

## 5. Configuration

The `.deepl.toml` configuration file must be manually created by the user in `~/.deepl/`. It must contain at least the DeepL API key and the default target language. Example:

```toml
api_key = "YOUR_DEEPL_API_KEY"
target_lang = "en"
source_lang = "fr" # Optional
```

## 6. DeepL API Usage

The tool uses the DeepL REST API. Requests are made to `https://api-free.deepl.com/v2/translate` using the POST method with a JSON request body. Authentication is done via the `Authorization` header with the format `DeepL-Auth-Key YOUR_API_KEY`.