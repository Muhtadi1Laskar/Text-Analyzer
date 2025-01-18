package handlers

import (
	"net/http"
	"text-analyzer/core"
)

type AnalyzerResponse struct {
	WordCount      int `json:"wordCount"`
	CharacterCount int `json:"characterCount"`
	LetterCount    int `json:"letterCount"`
	SentenceCount  int `json:"sentenceCount"`
	AverageWordCount float64 `json:"averageWordCount"`
}

func TextAnalyzer(w http.ResponseWriter, r *http.Request) {
	var requestBody CommonRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := core.MainFunc(requestBody.Message)

	response := AnalyzerResponse{
		WordCount:      data.WordCount,
		CharacterCount: data.CharacterCount,
		LetterCount:    data.LetterCount,
		SentenceCount:  data.SentenceCount,
		AverageWordCount: data.AverageWordCount,
	}

	writeJSONResponse(w, http.StatusOK, response)
}
