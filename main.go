package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"translator/api"
	"translator/config"
	"translator/input"
	"translator/output"
	"translator/translator"
)

func main() {
	configPath := flag.String("config", "config.json", "path to the configuration file")
	sourceLang := flag.String("source", "", "source language")
	targetLang := flag.String("target", "", "target language")
	inputFilePath := flag.String("input", "", "path to the input file")
	outputFilePath := flag.String("output", "", "path to the output file")
	progressFilePath := flag.String("progress", "progress.json", "path to the progress file")
	showHelp := flag.Bool("h", false, "show help")

	flag.Parse()

	if *showHelp {
		flag.Usage()
		os.Exit(0)
	}

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		os.Exit(1)
	}

	// Override config with flags if provided
	if *sourceLang != "" {
		cfg.SourceLang = *sourceLang
	}
	if *targetLang != "" {
		cfg.TargetLang = *targetLang
	}
	if *inputFilePath != "" {
		cfg.InputFilePath = *inputFilePath
	}
	if *outputFilePath != "" {
		cfg.OutputFilePath = *outputFilePath
	}

	// Initialize translator
	var apis []api.TranslatorAPI
	for _, key := range cfg.ApiKeys.Deepl {
		apis = append(apis, api.NewDeepLTranslator(key))
	}
	for _, key := range cfg.ApiKeys.NiuTrans {
		apis = append(apis, api.NewNiuTransTranslator(key))
	}
	translator, err := translator.NewTranslator(apis, cfg, *progressFilePath)
	if err != nil {
		fmt.Println("Error initializing translator:", err)
		os.Exit(1)
	}

	// Read input
	text, err := input.ReadFromFile(cfg.InputFilePath)
	if err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}

	// Split text into lines
	lines := strings.Split(text, "\n")
	if len(translator.GetProgress().RemainingLines) == 0 {
		translator.GetProgress().RemainingLines = lines
	}

	// Perform translation
	if err := translator.TranslateText(); err != nil {
		fmt.Println("Error translating text:", err)
		os.Exit(1)
	}

	// Write output
	translatedText := strings.Join(translator.GetProgress().CompletedLines, "\n")
	if err := output.WriteToFile(cfg.OutputFilePath, translatedText); err != nil {
		fmt.Println("Error writing output file:", err)
		os.Exit(1)
	}

	fmt.Println("Translation completed successfully!")
}
