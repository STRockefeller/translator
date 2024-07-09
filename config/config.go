package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	ApiKeys        APIKeys
	RetryDelay     int
	MaxRetries     int
	SourceLang     string
	TargetLang     string
	InputFilePath  string
	OutputFilePath string
}

type APIKeys struct {
	NiuTrans []string
	Deepl    []string
}

func LoadConfig(filepath string) (*Config, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	if err := decoder.Decode(config); err != nil {
		return nil, err
	}

	return config, nil
}
