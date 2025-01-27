package handlers

import (
	"net/http"
	"strings"
	"text-analyzer/core"
)

type Request struct {
	Text      string `json:"text" validate:"required"`
	Operation string `json:"operation" validate:"required"`
}

type Response struct {
	Data string `json:"data"`
}

func TextCleaner(w http.ResponseWriter, r *http.Request) {
	var requestBody Request

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var result string
	switch strings.ToLower(requestBody.Operation) {
	case "stopwords":
		result = core.RemoveStopWords(requestBody.Text)
	case "lower-case":
		result = strings.ToLower(requestBody.Text)
	case "upper-case":
		result = strings.ToUpper(requestBody.Text)
	case "remove-punctuation":
		result = core.RemovePunctuation(requestBody.Text)
	default:
		writeError(w, http.StatusInternalServerError, "invalid operation. The following operations are valid: stopwords, lower-case, upper-case, remove-punctuation")
		return
	}

	responseBody := Response{
		Data: result,
	}

	writeJSONResponse(w, http.StatusOK, responseBody)
}
