package core

import (
	"strings"
	"unicode"
)

type Analysis struct {
	WordCount        int
	CharacterCount   int
	LetterCount      int
	SentenceCount    int
	AverageWordCount float32
	TotalStopWords   int
}

func countWord(text string) int {
	return len(strings.Fields(text))
}

func characterCount(text string) int {
	var count int = 0

	for _, char := range text {
		if !isWhiteSpace(char) {
			count++
		}
	}
	return count
}

func letterCount(text string) int {
	var count int = 0

	for _, char := range text {
		if unicode.IsLetter(char) && !isWhiteSpace(char) {
			count++
		}
	}
	return count
}

func countSentence(text string) int {
	var count int = 0
	insideSentence := false

	for _, char := range text {
		if char == '.' || char == '?' || char == '!' {
			if insideSentence {
				count++
				insideSentence = false
			}
		} else if isWhiteSpace(char) {
			insideSentence = true
		}
	}

	if insideSentence {
		count++
	}

	return count
}

func countStopWords(text string) int {
	text = RemovePunctuation(text)
	tokenizeText := strings.Fields(text)

	var totalStopWords int = 0
	for _, elem := range tokenizeText {
		if _, exists := stopWordsSet[elem]; !exists {
			totalStopWords++
		}
	}
	return totalStopWords
}

func averageWordCount(text string) float32 {
	return float32(characterCount(text)) / float32(countWord(text))
}

func isWhiteSpace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t'
}

func MainFunc(text string) Analysis {
	text = strings.TrimSpace(text)

	wordCount := countWord(text)
	characterCount := characterCount(text)
	letterCount := letterCount(text)
	sentenceCount := countSentence(text)
	averageWordCount := averageWordCount(text)
	totalStopWords := countStopWords(text)

	return Analysis{
		WordCount:        wordCount,
		CharacterCount:   characterCount,
		LetterCount:      letterCount,
		SentenceCount:    sentenceCount,
		AverageWordCount: averageWordCount,
		TotalStopWords:   totalStopWords,
	}
}
