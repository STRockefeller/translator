package translator

import (
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
		translatedText, err := t.translateLine(currentLine)
		if err != nil {
			return err
		}

		t.progress.CompletedLines = append(t.progress.CompletedLines, translatedText)
		t.progress.RemainingLines = t.progress.RemainingLines[1:]

		if err := utils.SaveProgress(t.progressPath, t.progress); err != nil {
			return err
		}
	}

	return nil
}

func (t *Translator) translateLine(text string) (string, error) {
	var translatedText string
	var err error

	for retries := 0; retries < t.config.MaxRetries; retries++ {
		translatedText, err = t.apis[t.apiIndex].Translate(text, t.config.SourceLang, t.config.TargetLang)
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
