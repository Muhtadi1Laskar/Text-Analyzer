package handlers

import (
	"fmt"
	"io"
	"net/http"
	"text-analyzer/core"
)

type AnalyzerResponse struct {
	WordCount      int `json:"wordCount"`
	CharacterCount int `json:"characterCount"`
	LetterCount    int `json:"letterCount"`
	SentenceCount  int `json:"sentenceCount"`
	AverageWordCount float32 `json:"averageWordCount"`
}

func TextAnalyzer(w http.ResponseWriter, r *http.Request) {
	fileText, err := UploadFile(r)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := core.MainFunc(fileText["message"])

	response := AnalyzerResponse{
		WordCount:      data.WordCount,
		CharacterCount: data.CharacterCount,
		LetterCount:    data.LetterCount,
		SentenceCount:  data.SentenceCount,
		AverageWordCount: data.AverageWordCount,
	}

	writeJSONResponse(w, http.StatusOK, response)
}

func UploadFile(r *http.Request) (map[string]string, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return nil, fmt.Errorf("unable to read request body: %v", err)
	}

	file, _, err := r.FormFile("myFile")
	if err != nil {
		return nil, fmt.Errorf("unable to read request body: %v", err)
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("unable to read file bytes: %v", err)
	}

	temp := map[string]string{
		"message": string(fileBytes),
	}

	for key, data := range r.MultipartForm.Value {
		temp[key] = data[0]
	}

	return temp, nil
}
