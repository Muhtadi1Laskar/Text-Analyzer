package handlers

import (
	"net/http"
	"text-analyzer/core"
)

type PlagrismRequest struct {
	TextOne string `json:"textOne" validate:"required"`
	TextTwo string `json:"textTwo" validate:"required"`
}

type PlagrismResponse struct {
	Similarity float64 `json:"similarity" validate:"required"`
}

func PlagrismChecker(w http.ResponseWriter, r *http.Request) {
	fileOne, err := UploadFile(r, "fileOne")
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	
	fileTwo, err := UploadFile(r, "fileTwo")
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	similarity := core.CheckPlagrism(fileOne["message"], fileTwo["message"])

	response := PlagrismResponse{
		Similarity: similarity,
	}

	writeJSONResponse(w, http.StatusOK, response)
}