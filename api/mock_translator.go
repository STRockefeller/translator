package api

type MockTranslator struct {
	TranslatedText string
	Err            error
}

func (m *MockTranslator) Translate(text, sourceLang, targetLang string) (string, error) {
	return m.TranslatedText, m.Err
}

func (m *MockTranslator) CanRetry(err error) bool {
	return true
}
