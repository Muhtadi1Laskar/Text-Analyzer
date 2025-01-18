package handlers

import (
	"net/http"
	"text-analyzer/core"
)

func TextAnalyzer(w http.ResponseWriter, r *http.Request) {
	var requestBody CommonRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := core.MainFunc(requestBody.Message)

	response := CommonResponse {
		Message: string(data.WordCount),
	}

	writeJSONResponse(w, http.StatusOK, response)
}