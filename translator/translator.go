package translator

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"translator/api"
	"translator/config"
	"translator/utils"
)

type Translator struct {
	apis         []api.TranslatorAPI
	config       *config.Config
	apiIndex     int
	progress     *utils.Progress
	progressPath string
}

func NewTranslator(apis []api.TranslatorAPI, config *config.Config, progressPath string) (*Translator, error) {
	progress, err := utils.LoadProgress(progressPath)
	if err != nil {
		return nil, err
	}

	return &Translator{
		apis:         apis,
		config:       config,
		progress:     progress,
		progressPath: progressPath,
	}, nil
}

func (t *Translator) TranslateText() error {
	for len(t.progress.RemainingLines) > 0 {
		currentLine := t.progress.RemainingLines[0]
		var translatedText string
		if len(t.config.Translations) > 0 {
			var err error
			translatedText, err = t.translateLine(currentLine)
			if err != nil {
				return err
			}
		} else {
			var err error
			translatedText, err = t.translateFragment(currentLine)
			if err != nil {
				return err
			}
		}

		t.progress.CompletedLines = append(t.progress.CompletedLines, translatedText)
		t.progress.RemainingLines = t.progress.RemainingLines[1:]

		if err := utils.SaveProgress(t.progressPath, t.progress); err != nil {
			return err
		}
		// show progress
		if l := len(t.progress.RemainingLines); l%100 == 0 {
			fmt.Println("remaining lines:", l)
		}
	}

	// delete progress file if there is no error
	os.Remove(t.progressPath)
	return nil
}

var regexPattern *regexp.Regexp

func (t *Translator) translateLine(text string) (string, error) {
	translatedText := text

	for _, translationConfig := range t.config.Translations {
		if regexPattern == nil {
			var err error
			regexPattern, err = regexp.Compile(translationConfig.Pattern)
			if err != nil {
				return "", err
			}
		}

		matches := regexPattern.FindAllStringSubmatchIndex(translatedText, -1)
		if matches == nil {
			continue
		}

		for _, match := range matches {
			original := translatedText[match[0]:match[1]]
			translatedParts := make([]string, len(match)/2)

			for i := 0; i < len(match)/2; i++ {
				translatedParts[i] = translatedText[match[2*i]:match[2*i+1]]
			}

			for _, groupIndex := range translationConfig.TranslateGroups {
				if groupIndex < len(translatedParts) {
					translatedFragment, err := t.translateFragment(translatedParts[groupIndex])
					if err != nil {
						return "", err
					}
					translatedParts[groupIndex] = translatedFragment
				}
			}

			translatedPair := original
			for i, part := range translatedParts {
				translatedPair = strings.Replace(translatedPair, translatedText[match[2*i]:match[2*i+1]], part, 1)
			}

			translatedText = strings.Replace(translatedText, original, translatedPair, 1)
		}
	}

	return translatedText, nil
}

func (t *Translator) translateFragment(fragment string) (string, error) {
	var translatedText string
	var err error

	for retries := 0; retries < t.config.MaxRetries; retries++ {
		translatedText, err = t.apis[t.apiIndex].Translate(fragment, t.config.SourceLang, t.config.TargetLang)
		if err == nil {
			return translatedText, nil
		}
		if !t.apis[t.apiIndex].CanRetry(err) {
			return "", err
		}

		utils.HandleError(err)
		t.apiIndex = (t.apiIndex + 1) % len(t.apis)
		time.Sleep(time.Duration(t.config.RetryDelay) * time.Second)
	}

	return "", err
}

func (t *Translator) GetProgress() *utils.Progress {
	return t.progress
}
