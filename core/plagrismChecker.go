package core

import (
	"math"
	"strings"
)

func preprocessText(text string) []string {
	cleanText := RemovePunctuation(text)
	cleanText = RemoveStopWords(cleanText)

	return strings.Fields(cleanText)
}

func createFrequencyMap(words []string) map[string]int {
	freqMap := make(map[string]int)

	for _, word := range words {
		freqMap[word]++
	}

	return freqMap
}

func getAllWords(freqMap1, freqMap2 map[string]int) []string {
	allWords := make(map[string]bool)

	for word := range freqMap1 {
		allWords[word] = true
	}
	for word := range freqMap2 {
		allWords[word] = true
	}

	uniqueWords := make([]string, 0, len(allWords))
	for word := range allWords {
		uniqueWords = append(uniqueWords, word)
	}
	return uniqueWords
}


func createVector(allWords []string, freqMap map[string]int) []float64 {
	vector := make([]float64, len(allWords))
	for i, word := range allWords {
		vector[i] = float64(freqMap[word])
	}
	return vector
}

func cosineSimilarity(vector1, vector2 []float64) float64 {
	dotProduct := 0.0
	magnitude1 := 0.0
	magnitude2 := 0.0

	for i := 0; i < len(vector1); i++ {
		dotProduct += vector1[i] * vector2[i]
		magnitude1 += vector1[i] * vector1[i]
		magnitude2 += vector2[i] * vector2[i]
	}

	magnitude1 = math.Sqrt(magnitude1)
	magnitude2 = math.Sqrt(magnitude2)

	if magnitude1 == 0 || magnitude2 == 0 {
		return 0.0
	}

	return dotProduct / (magnitude1 * magnitude2)
}

func CheckPlagrism(textOne, textTwo string) float64 {
	wordsOne := preprocessText(textOne)
	wordsTwo := preprocessText(textTwo)

	freqMapOne := createFrequencyMap(wordsOne)
	freqMapTwo := createFrequencyMap(wordsTwo)

	allWords := getAllWords(freqMapOne, freqMapTwo)

	vectorOne := createVector(allWords, freqMapOne)
	vectorTwo := createVector(allWords, freqMapTwo)

	similarity := cosineSimilarity(vectorOne, vectorTwo)

	return similarity * 100
}
