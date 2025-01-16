package handlers

import (
	"net/http"
	"strings"
)

func TextAnalyzer(w http.ResponseWriter, r *http.Request) {
	var requestBody CommonRequest

	if err := readRequestBody(r, &requestBody); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	response := CommonResponse {
		Message: strings.ToUpper(requestBody.Message),
	}

	writeJSONResponse(w, http.StatusOK, response)
}