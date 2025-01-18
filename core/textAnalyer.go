package core

import (
	"strconv"
	"strings"
)

type Analysis struct {
	WordCount string
}

func countWord (text string) int {
	text = strings.TrimSpace(text)
	return len(strings.Fields(text))
}

func MainFunc(text string) Analysis {
	wordCount := strconv.Itoa(countWord(text))

	return Analysis{
		WordCount: wordCount,
	}
}

