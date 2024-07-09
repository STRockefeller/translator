package utils

import (
	"regexp"
)

func MatchPattern(text, pattern string) ([]string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	matches := re.FindAllString(text, -1)
	return matches, nil
}

func ReplacePattern(text, pattern, replacement string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	result := re.ReplaceAllString(text, replacement)
	return result, nil
}
