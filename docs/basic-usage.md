# Basic Usage of DeepL CLI

This guide covers the fundamental ways to use the `deepl-cli` tool for text translation.

## Translating Text

The primary function of `deepl-cli` is to translate text. You can provide the text directly as an argument to the `deepl` command.

### Default Translation

If you have configured your `~/.deepl/.deepl.toml` file with a `target_lang`, you can simply provide the text, and the tool will translate it to your default target language. The source language will be automatically detected.

```bash
deepl "Hello world, how are you?"
# Example output: Bonjour le monde, comment allez-vous ?
```

### Specifying Source and Target Languages

You can explicitly define the source and target languages using the `-s` (source) and `-t` (target) flags. This overrides any default languages set in your configuration file.

```bash
deepl -s en -t fr "Hello world, how are you?"
# Example output: Bonjour le monde, comment allez-vous ?

deepl -s fr -t en "Bonjour le monde, comment allez-vous ?"
# Example output: Hello world, how are you?
```

**Note on languages:** Use ISO 639-1 codes (e.g., `en`, `fr`, `de`). For a complete list of supported languages, please refer to the official DeepL API documentation.

## Getting Help

To see all available command-line options and flags, you can use the `--help` flag:

```bash
deepl --help
```

This will display a comprehensive list of commands and their descriptions.