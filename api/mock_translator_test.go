package api

import (
	"errors"
	"testing"
)

func TestMockTranslator_Translate(t *testing.T) {
	translator := &MockTranslator{
		TranslatedText: "你好",
		Err:            nil,
	}

	text := "Hello"
	sourceLang := "en"
	targetLang := "zh"

	translatedText, err := translator.Translate(text, sourceLang, targetLang)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if translatedText != "你好" {
		t.Fatalf("Expected translated text to be '你好', got %s", translatedText)
	}
}

func TestMockTranslator_CanRetry(t *testing.T) {
	translator := &MockTranslator{}

	err := errors.New("QPS limit exceeded")
	if !translator.CanRetry(err) {
		t.Fatalf("Expected true for retryable error, got false")
	}

	err = errors.New("Some other error")
	if !translator.CanRetry(err) {
		t.Fatalf("Expected true for non-retryable error, got false")
	}
}
