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
	Comparision float64 `json:"comparision" validate:"required"`
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

	comparisionRatio := core.CheckPlagrism(fileOne["message"], fileTwo["message"])

	response := PlagrismResponse{
		Comparision: comparisionRatio,
	}

	writeJSONResponse(w, http.StatusOK, response)
}