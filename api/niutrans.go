package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type NiuTransTranslator struct {
	apiKey string
}

func NewNiuTransTranslator(apiKey string) *NiuTransTranslator {
	return &NiuTransTranslator{apiKey: apiKey}
}

func (n *NiuTransTranslator) Translate(text, sourceLang, targetLang string) (string, error) {
	apiUrl := "https://api.niutrans.com/NiuTransServer/translation"

	// Prepare the request parameters
	params := url.Values{}
	params.Add("from", sourceLang)
	params.Add("to", targetLang)
	params.Add("apikey", n.apiKey)
	params.Add("src_text", url.QueryEscape(text))

	// Determine whether to use GET or POST based on text length
	var resp *http.Response
	var err error
	if len(text) > 1500 {
		resp, err = http.PostForm(apiUrl, params)
	} else {
		apiUrl += "?" + params.Encode()
		resp, err = http.Get(apiUrl)
	}

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get translation from NiuTrans")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	// Check for errors in the response
	if errorCode, exists := result["error_code"].(string); exists && errorCode != "0" {
		errorMsg := result["error_msg"].(string)
		return "", fmt.Errorf("NiuTrans error %s: %s", errorCode, errorMsg)
	}

	translatedText, ok := result["tgt_text"].(string)
	if !ok {
		return "", errors.New("invalid response from NiuTrans")
	}

	return translatedText, nil
}

func (n *NiuTransTranslator) CanRetry(err error) bool {
	if strings.Contains(err.Error(), "QPS") ||
		strings.Contains(err.Error(), "timeout") ||
		strings.Contains(err.Error(), "running out") {
		return true
	}
	return false
}
