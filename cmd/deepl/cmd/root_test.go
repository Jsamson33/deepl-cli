package cmd

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTranslateText(t *testing.T) {
	// Mock DeepL API server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/v2/translate" {
			t.Errorf("Expected to request '/v2/translate', got: %s", r.URL.Path)
		}
		if r.Header.Get("Authorization") != "DeepL-Auth-Key test_api_key" {
			t.Errorf("Expected Authorization header 'DeepL-Auth-Key test_api_key', got: %s", r.Header.Get("Authorization"))
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type 'application/json', got: %s", r.Header.Get("Content-Type"))
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("Failed to read request body: %v", err)
		}

		expectedRequestBody := `{"text":["Hello"],"source_lang":"en","target_lang":"fr"}`
		if string(body) != expectedRequestBody {
			t.Errorf("Expected request body '%s', got '%s'", expectedRequestBody, string(body))
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"translations":[{"detected_source_language":"en","text":"Bonjour"}]}`))
	}))
	defer server.Close()

	// Override apiURL for testing
	oldAPIURL := apiURL
	apiURL = server.URL + "/v2/translate"
	defer func() { apiURL = oldAPIURL }()

	// Test case: successful translation
	translatedText, err := translateText("Hello", "en", "fr", "test_api_key")
	if err != nil {
		t.Errorf("translateText failed: %v", err)
	}
	if translatedText != "Bonjour" {
		t.Errorf("Expected translated text 'Bonjour', got '%s'", translatedText)
	}

	// Test case: API error
	serverError := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"DeepL API error"}`))
	}))
	defer serverError.Close()

	apiURL = serverError.URL + "/v2/translate"
	_, err = translateText("Hello", "en", "fr", "test_api_key")
	if err == nil || err.Error() != "DeepL API error: {\"message\":\"DeepL API error\"} (Status: 500)" {
		t.Errorf("Expected API error, got: %v", err)
	}

	// Test case: No translation found
	serverNoTranslation := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"translations":[]}`))
	}))
	defer serverNoTranslation.Close()

	apiURL = serverNoTranslation.URL + "/v2/translate"
	translatedText, err = translateText("Hello", "en", "fr", "test_api_key")
	if err != nil {
		t.Errorf("translateText failed for no translation: %v", err)
	}
	if translatedText != "No translation found." {
		t.Errorf("Expected 'No translation found.', got '%s'", translatedText)
	}
}
