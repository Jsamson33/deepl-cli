![Release Workflow](https://github.com/Jsamson33/deepl-cli/workflows/Release/badge.svg) ![Go Version](https://img.shields.io/badge/Go-1.22-blue.svg) ![GitHub Pages](https://github.com/Jsamson33/deepl-cli/workflows/pages%20build%20and%20deployment/badge.svg) ![Tests](https://img.shields.io/badge/tests-passing-brightgreen.svg) ![Linter](https://img.shields.io/badge/linter-passing-brightgreen.svg)

# 🌍 deepl-cli: Your DeepL Command-Line Translator

`deepl-cli` is a lightweight and powerful command-line tool written in Go, designed to give you access to DeepL's exceptional translation quality directly from your terminal. Quickly translate text without leaving your working environment!

## ✨ Features

*   Fast text translation via the DeepL API.
*   Automatic source language detection.
*   Easy configuration of your API key and default languages via a `~/.deepl/.deepl.toml` file.
*   Support for source and target languages directly from the command line.

### Architecture Overview

![DeepL CLI Architecture](docs/architecture.png)

## 🚀 Installation

### Prerequisites

*   A DeepL API key. You can obtain one by signing up for the [DeepL API Free plan](https://www.deepl.com/pro/developer).
*   Go 1.16+ installed on your machine.

### From Pre-compiled Binaries (Recommended)

Pre-compiled binaries for Linux, macOS, and Windows are available in the [Releases section of this GitHub repository](https://github.com/votre-utilisateur/deepl-cli/releases).
1.  Download the archive corresponding to your operating system and architecture.
2.  Unzip the archive.
3.  Place the `deepl` executable (or `deepl.exe` on Windows) in a directory present in your `PATH` (e.g., `/usr/local/bin` on Linux/macOS, or a custom folder added to your `PATH` on Windows).

### From Source

```bash
git clone [https://github.com/votre-utilisateur/deepl-cli.git](https://github.com/votre-utilisateur/deepl-cli.git)
cd deepl-cli
go build -o deepl ./cmd/deepl
# Move the executable to a directory in your PATH
# mv deepl /usr/local/bin/deepl
```

## 💡 Philosophy

DeepL CLI embraces the core tenets of the Unix philosophy:

*   **Do one thing and do it well:** Its primary function is text translation via the DeepL API.
*   **Work together:** Designed to be easily integrated with other command-line tools through pipes (e.g., `cat file.txt | deepl | less`).
*   **Text is the universal interface:** Input and output are plain text, making it highly interoperable.
*   **Build a prototype as soon as you can:** Focus on core functionality first, then iterate.

## ⚙️ Configuration

Before using deepl-cli, you need to configure your DeepL API key and optionally your default languages.

Create a `.deepl` directory in your home directory:

```bash
mkdir -p ~/.deepl
```

Create a file named `.deepl.toml` inside this directory (`~/.deepl/.deepl.toml`) with the following content, replacing `YOUR_DEEPL_API_KEY` with your actual key:

```toml
api_key = "YOUR_DEEPL_API_KEY"
target_lang = "en" # Default target language (e.g., "en" for English, "fr" for French)
# source_lang = "fr" # Optional: default source language (can be auto-detected)
```

**Note on languages:** Use ISO 639-1 codes (e.g., `en`, `fr`, `de`). For a complete list, consult the DeepL API documentation.

## 💻 Usage

### Translate text (default languages)

```bash
deepl "Hello world, how are you?"
```

The text will be translated to the `target_lang` configured in your `.deepl.toml` file. The source language will be automatically detected.

### Specify source and target languages

Use the `-s` (source) and `-t` (target) flags:

```bash
deepl -s en -t fr "Hello world, how are you?"
deepl -s fr -t en "Bonjour le monde, comment allez-vous ?"
```

### Help

To see all available options:

```bash
deepl --help
```