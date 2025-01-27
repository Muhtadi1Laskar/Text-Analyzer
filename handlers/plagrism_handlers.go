package handlers

import (
	"net/http"
	"text-analyzer/core"
)

type PlagrismRequest struct {
	TextOne string `json:"textOne" validate:"required"`
	TextTwo string `json:"textTwo" validate:"required"`
	CheckerType string `json:"checkerType" validate:"required"`
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

	var similarity float64
	switch fileOne["checkerType"] {
	case "cosine-similarity":
		similarity = core.CheckPlagrism(fileOne["message"], fileTwo["message"])
	case "minhash":
		similarity = core.MinHash(fileOne["message"], fileTwo["message"])
	default:
		writeError(w, http.StatusInternalServerError, "Invalid Checker Type")
		return
	}


	response := PlagrismResponse{
		Similarity: similarity,
	}

	writeJSONResponse(w, http.StatusOK, response)
}