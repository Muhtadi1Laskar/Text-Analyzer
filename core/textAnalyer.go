package core

import (
	"regexp"
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

type FeatureVector struct {
	AvgWordLength      float64
	AvgSentenceLength  float64
	StopWordFrequency  float64
	VocabularyRichness float64
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
	tokenizeText := Tokenize(text)

	var totalStopWords int = 0
	for _, elem := range tokenizeText {
		if _, exists := stopWordsSet[elem]; exists {
			totalStopWords++
		}
	}
	return totalStopWords
}

func Preprocess(text string) []string {
	text = RemovePunctuation(text)
	tokens := Tokenize(text)

	return tokens
}

func ExtractFeature(text string) FeatureVector {
	tokens := Preprocess(text)
	sentences := splitSentences(text)

	if len(tokens) == 0 || len(sentences) == 0 {
		return FeatureVector{}
	}

	totalWordLength := 0
	for _, word := range tokens {
		totalWordLength += len(word)
	}
	avgWordLength := float64(totalWordLength) / float64(len(tokens))
	avgSentenceLength := float64(len(tokens)) / float64(len(sentences))
	totalStopWords := countStopWords(text)
	stopwordFreq := float64(totalStopWords) / float64(len(tokens))

	uniqueWords := make(map[string]bool)
	for _, word := range tokens {
		uniqueWords[word] = true
	}
	vocabRichness := float64(len(uniqueWords)) / float64(len(tokens))

	return FeatureVector{
		AvgWordLength: avgWordLength,
		AvgSentenceLength: avgSentenceLength,
		StopWordFrequency: stopwordFreq,
		VocabularyRichness: vocabRichness,
	}
}

func averageWordCount(text string) float32 {
	return float32(characterCount(text)) / float32(countWord(text))
}

func isWhiteSpace(char rune) bool {
	return char == ' ' || char == '\n' || char == '\t'
}

func splitSentences(text string) []string {
	reg := regexp.MustCompile(`[.!?]+`)
	sentences := reg.Split(text, -1)

	var result []string
	for _, sentence := range sentences {
		if strings.TrimSpace(sentence) != "" {
			result = append(result, sentence)
		}
	}
	return result
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
