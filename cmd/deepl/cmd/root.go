package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"deepl-cli/internal/config"
	"github.com/spf13/cobra"
)

var (
	srcLang string
	tgtLang string
	apiURL = "https://api-free.deepl.com/v2/"
	showLanguages bool
)

var RootCmd = &cobra.Command{
	Use:   "deepl [text to translate]",
	Short: "A simple CLI for translating text using the DeepL API",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Printf("Error loading configuration: %v\n", err)
			os.Exit(1)
		}

		if showLanguages {
			type Language struct {
				Language string `json:"language"`
				Name     string `json:"name"`
			}

			req, err := http.NewRequest("GET", apiURL + "languages", nil)
			if err != nil {
				fmt.Printf("Error creating request: %v\n", err)
				os.Exit(1)
			}
			req.Header.Set("Authorization", "DeepL-Auth-Key "+cfg.APIKey)

			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("Error fetching languages: %v\n", err)
				os.Exit(1)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("Error reading response body: %v\n", err)
				os.Exit(1)
			}

			if resp.StatusCode != http.StatusOK {
				fmt.Printf("DeepL API error: %s (Status: %d)\n", string(body), resp.StatusCode)
				os.Exit(1)
			}

			var languages []Language
			if err := json.Unmarshal(body, &languages); err != nil {
				fmt.Printf("Error unmarshalling languages: %v\n", err)
				os.Exit(1)
			}

			fmt.Println("Supported DeepL Languages:")
			fmt.Println("--------------------------")
			for _, lang := range languages {
				fmt.Printf("%s: %s\n", lang.Language, lang.Name)
			}
			return
		}

		if len(args) == 0 {
			cmd.Help()
			return
		}

		if srcLang == "" {
			srcLang = cfg.SourceLang
		}
		if tgtLang == "" {
			tgtLang = cfg.TargetLang
		}

		if tgtLang == "" {
			fmt.Println("Error: Target language must be specified either via --target flag or in the config file.")
			os.Exit(1)
		}

		textToTranslate := args[0]

		translatedText, err := translateText(textToTranslate, srcLang, tgtLang, cfg.APIKey)
		if err != nil {
			fmt.Printf("Error translating text: %v\n", err)
			os.Exit(1)
		}

		fmt.Println(translatedText)
	},
}

func init() {
	RootCmd.Flags().StringVarP(&srcLang, "source", "s", "", "Source language")
	RootCmd.Flags().StringVarP(&tgtLang, "target", "t", "", "Target language")
	RootCmd.Flags().BoolVarP(&showLanguages, "languages", "l", false, "List all supported DeepL languages")
}



func translateText(text, srcLangParam, tgtLangParam, apiKey string) (string, error) {

	type DeepLRequest struct {
		Text       []string `json:"text"`
		SourceLang string   `json:"source_lang,omitempty"`
		TargetLang string   `json:"target_lang"`
	}

	type DeepLResponse struct {
		Translations []struct {
			DetectedSourceLanguage string `json:"detected_source_language"`
			Text                   string `json:"text"`
		} `json:"translations"`
	}

	requestBody := DeepLRequest{
		Text:       []string{text},
		SourceLang: srcLangParam,
		TargetLang: tgtLangParam,
	}

	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("error marshalling request: %w", err)
	}

	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "DeepL-Auth-Key "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making API request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("DeepL API error: %s (Status: %d)", string(body), resp.StatusCode)
	}

	var deeplResponse DeepLResponse
	if err := json.Unmarshal(body, &deeplResponse); err != nil {
		return "", fmt.Errorf("error unmarshalling response: %w", err)
	}

	if len(deeplResponse.Translations) > 0 {
		return deeplResponse.Translations[0].Text, nil
	} else {
		return "No translation found.", nil
	}
}


