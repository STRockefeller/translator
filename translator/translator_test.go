package translator

import (
	"os"
	"testing"
	"translator/api"
	"translator/config"
)

func TestTranslator_TranslateText(t *testing.T) {
	// using the mock one
	mockAPI := &api.MockTranslator{
		TranslatedText: "你好",
		Err:            nil,
	}

	apis := []api.TranslatorAPI{mockAPI}
	cfg := &config.Config{
		RetryDelay: 5,
		MaxRetries: 3,
		SourceLang: "en",
		TargetLang: "zh",
	}

	progressPath := "progress_test.json"
	translator, err := NewTranslator(apis, cfg, progressPath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	translator.GetProgress().RemainingLines = []string{"Hello", "world"}

	err = translator.TranslateText()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(translator.GetProgress().CompletedLines) != 2 {
		t.Fatalf("Expected 2 completed lines, got %d", len(translator.GetProgress().CompletedLines))
	}

	// remove progress_test.json
	if err := os.Remove(progressPath); err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}
