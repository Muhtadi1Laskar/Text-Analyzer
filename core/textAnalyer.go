package core

import (
	"strings"
)

type Analysis struct {
	WordCount int
}

func countWord (text string) int {
	text = strings.TrimSpace(text)
	return len(strings.Fields(text))
}

func MainFunc(text string) Analysis {
	wordCount := countWord(text)

	return Analysis{
		WordCount: wordCount,
	}
}

