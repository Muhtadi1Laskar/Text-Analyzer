package core

import (
	"strings"
	"unicode"
)

func RemovePunctuation(text string) string {
	var builder strings.Builder

	for _, word := range text {
		if word != '-' || word != '\'' || !unicode.IsPunct(word) {
			builder.WriteRune(word)
		}
	}
	return builder.String()
}

func Tokenize(text string) []string {
	var tokens []string
	var wordBuilder strings.Builder

	for _, char := range text {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			wordBuilder.WriteRune(unicode.ToLower(char))
		} else if wordBuilder.Len() > 0 {
			tokens = append(tokens, wordBuilder.String())
			wordBuilder.Reset()
		}
	}

	if wordBuilder.Len() > 0 {
		tokens = append(tokens, wordBuilder.String())
	}

	return tokens
}