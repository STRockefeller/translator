package utils

import (
	"encoding/json"
	"os"
)

type Progress struct {
	CompletedLines []string `json:"completed_lines"`
	RemainingLines []string `json:"remaining_lines"`
}

func LoadProgress(filepath string) (*Progress, error) {
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return &Progress{}, nil
	}

	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var progress Progress
	if err := json.Unmarshal(data, &progress); err != nil {
		return nil, err
	}

	return &progress, nil
}

func SaveProgress(filepath string, progress *Progress) error {
	data, err := json.Marshal(progress)
	if err != nil {
		return err
	}

	return os.WriteFile(filepath, data, 0644)
}
