package api

type TranslatorAPI interface {
	Translate(text, sourceLang, targetLang string) (string, error)
	CanRetry(err error) bool
}
