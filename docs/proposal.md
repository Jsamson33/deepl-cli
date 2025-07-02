# Feature Proposals for DeepL CLI

This document outlines potential feature enhancements for the `deepl-cli` project, aiming to improve its functionality, user experience, and leverage more capabilities of the DeepL API.

## 1. File Translation

*   **Description:** Allow users to translate entire text files (e.g., `.txt`, `.md`) and output the translated content to a new file. This would be highly useful for translating documents without copying and pasting.
*   **Example Usage:** `deepl --file input.txt --output translated_fr.txt -t fr`

## 2. Interactive Translation Mode

*   **Description:** Implement a mode where the user can continuously type text into the terminal, and the tool translates it line by line or paragraph by paragraph until the user exits (e.g., by pressing Ctrl+D). This would provide a more fluid translation experience for longer texts.
*   **Example Usage:** `deepl --interactive -t es`

## 3. Glossary Support

*   **Description:** Integrate DeepL's glossary feature, allowing users to specify a glossary ID to ensure specific terms are translated consistently according to their predefined rules. This is crucial for technical or specialized translations.
*   **Example Usage:** `deepl --glossary my_tech_glossary_id "My technical term"`

## 4. Formality Control

*   **Description:** For languages that support it (e.g., German, French), allow users to specify the desired formality of the translation (e.g., "more formal", "less formal").
*   **Example Usage:** `deepl -t de --formality formal "How are you?"`

## 5. List Supported Languages

*   **Description:** Add a command to query the DeepL API and display a comprehensive list of all supported source and target languages, along with their respective codes. This would help users discover available languages without consulting external documentation.
*   **Example Usage:** `deepl --list-languages`

## 6. Auto-Update Mechanism

*   **Description:** Implement the proposed auto-update system (as detailed in `docs/auto-update-proposal.html`) to allow the CLI tool to update itself to the latest version automatically or on user command. This significantly improves maintainability and ensures users always have the latest features and bug fixes.
