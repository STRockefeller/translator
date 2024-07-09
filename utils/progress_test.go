package utils

import (
	"os"
	"testing"
)

func TestProgress_SaveAndLoad(t *testing.T) {
	progress := &Progress{
		CompletedLines: []string{"Hello"},
		RemainingLines: []string{"world"},
	}

	filepath := "progress_test.json"
	err := SaveProgress(filepath, progress)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	loadedProgress, err := LoadProgress(filepath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(loadedProgress.CompletedLines) != len(progress.CompletedLines) {
		t.Fatalf("Expected %d completed lines, got %d", len(progress.CompletedLines), len(loadedProgress.CompletedLines))
	}

	if len(loadedProgress.RemainingLines) != len(progress.RemainingLines) {
		t.Fatalf("Expected %d remaining lines, got %d", len(progress.RemainingLines), len(loadedProgress.RemainingLines))
	}

	os.Remove(filepath)
}
