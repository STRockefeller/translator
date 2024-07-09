package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type DeepLTranslator struct {
	apiKey string
}

func NewDeepLTranslator(apiKey string) *DeepLTranslator {
	return &DeepLTranslator{apiKey: apiKey}
}

func (d *DeepLTranslator) Translate(text, sourceLang, targetLang string) (string, error) {
	url := fmt.Sprintf("https://api.deepl.com/v2/translate?auth_key=%s&text=%s&source_lang=%s&target_lang=%s", d.apiKey, text, sourceLang, targetLang)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get translation from DeepL")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	translations, ok := result["translations"].([]interface{})
	if !ok || len(translations) == 0 {
		return "", errors.New("invalid response from DeepL")
	}

	translation, ok := translations[0].(map[string]interface{})
	if !ok {
		return "", errors.New("invalid response from DeepL")
	}

	translatedText, ok := translation["text"].(string)
	if !ok {
		return "", errors.New("invalid response from DeepL")
	}

	return translatedText, nil
}

func (d *DeepLTranslator) CanRetry(err error) bool {
	// 根據實際錯誤信息來判斷是否可以重試
	return true
}
