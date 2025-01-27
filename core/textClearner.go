package core

import (
	"regexp"
	"strings"
	"unicode"
)

var stopWordsSet = map[string]struct{}{
	"a": {}, "an": {}, "this": {}, "the": {}, "is": {}, "are": {}, "was": {}, "were": {}, "will": {}, "be": {},
	"in": {}, "on": {}, "at": {}, "of": {}, "for": {}, "to": {}, "from": {}, "with": {},
	"and": {}, "or": {}, "but": {}, "not": {}, "if": {}, "then": {}, "else": {},
	"i": {}, "you": {}, "he": {}, "she": {}, "it": {}, "we": {}, "they": {}, "my": {}, "your": {}, "his": {}, "her": {}, "its": {}, "our": {}, "their": {},
	"couldve": {}, "couldnt": {}, "wouldnt": {}, "shouldnt": {}, "wasnt": {}, "wont": {}, "shallnt": {}, "didnt": {}, "weev": {}, "im": {}, "as": {}, "would": {}, "have": {}, "had": {},
}

func RemovePunctuation(text string) string {
	text = strings.ToLower(text)
	reg := regexp.MustCompile(`[^\w\s]`)
	text = reg.ReplaceAllString(text, "")

	return text
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

func RemoveStopWords(texts string) string {
	texts = RemovePunctuation(texts)
	text := strings.Fields(texts)
	var filteredStr []string

	for _, word := range text {
		if _, exists := stopWordsSet[word]; !exists {
			filteredStr = append(filteredStr, word)
		}
	}
	return strings.Join(filteredStr, " ")
}
