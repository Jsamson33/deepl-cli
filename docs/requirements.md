# Requirements for the DeepL CLI Translation Tool

This document describes the functional and technical requirements for the development of a command-line translation tool named `deepl`, using the DeepL API.

## 1. Project Description

The goal is to create a simple and efficient command-line utility that allows users to translate text using the DeepL API. The tool will be configurable and easy to use for quick translations from the terminal.

## 2. Features

### 2.1 Core Features

*   **Text Translation:** Translate a piece of text provided as an argument.
*   **Automatic Source Language Detection:** If the source language is not specified, the tool will attempt to detect it automatically.
*   **Flexible Configuration:** The DeepL API key, as well as default source and target languages, must be configurable via a dedicated configuration file.
*   **Result Display:** The translated text will be displayed directly in the terminal.

### 2.2 Optional Features (for future versions)

*   **File Translation:** Ability to translate the content of a file and write the translation to a new file.
*   **Glossary Support:** Use DeepL glossaries for specific translations.
*   **Interactive Mode:** A mode where the user can type text line by line for translation.
*   **Display Supported Languages:** A command to list all languages supported by DeepL.

## 3. Technical Specifications

### 3.1 Development Language

*   The project will be entirely coded in **Go (Golang)**.

### 3.2 Command-Line Usage

*   The final executable must be usable via the `deepl` command.
*   Usage examples:
    *   `deepl "Hello world"` (translates to default target language)
    *   `deepl -s fr -t en "Bonjour le monde"` (translates from French to English)

### 3.3 Configuration File

*   A configuration file named `.deepl.toml` (or another format like YAML/JSON if preferred, but TOML is common for Go) will be created in the `~/.deepl/` directory (or `$HOME/.deepl/` on Linux/macOS, `%USERPROFILE%\.deepl\` on Windows).
*   This file will contain at minimum:
    *   `api_key`: The DeepL API key.
    *   `source_lang`: The default source language (optional, can be auto-detected).
    *   `target_lang`: The default target language (mandatory).

### 3.4 DeepL API

*   Use of the DeepL REST API. A Go library for the DeepL API could be used if available and maintained, otherwise direct HTTP requests will need to be implemented.
*   The DeepL API Free plan will be used for development and testing.

## 4. Implementation Steps

### 4.1 Phase 1: Initialization and Configuration

1.  **Go Project Initialization:** Create a new Go module.
2.  **Directory Structure:** Set up the basic project structure (see section 7).
3.  **Configuration Management:**
    *   Implement logic to read the `~/.deepl/.deepl.toml` file.
    *   Handle errors if the file or API key is missing.
    *   Use a configuration parsing library (e.g., `viper`, `gopkg.in/toml.v2`).

### 4.2 Phase 2: Translation Logic

1.  **CLI Argument Parsing:**
    *   Use a library like `cobra` or `flag` to manage command-line arguments (`-s`, `-t`, text to translate).
    *   Prioritize CLI arguments over default values from the configuration file.
2.  **DeepL API Call:**
    *   Implement logic to construct the HTTP request to the DeepL API.
    *   Handle API responses (success, errors, HTTP codes).
    *   Display the translation or error messages clearly.

### 4.3 Phase 3: Build and Distribution

1.  **Compilation Script:** Prepare a script (or Makefile target) to compile the executable for different platforms (Linux, macOS, Windows).
2.  **GitHub Actions for Releases:**
    *   Set up a GitHub Actions workflow to compile the tool for different architectures and operating systems.
    *   This workflow should create GitHub *releases* with attached binaries upon a tag push.

### 4.4 Phase 4: Documentation and Tests

1.  **Unit and Integration Tests:** Write tests for key components (configuration parsing, API calls, etc.).
2.  **Documentation:**
    *   Create a detailed `README.md` file.
    *   Create more in-depth documentation in the `doc/` directory.