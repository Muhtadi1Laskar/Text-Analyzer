package handlers

import (
	"net/http"
	"text-analyzer/core"
)

type PlagrismRequest struct {
	TextOne     string `json:"textOne" validate:"required"`
	TextTwo     string `json:"textTwo" validate:"required"`
	CheckerType string `json:"checkerType" validate:"required"`
}

type PlagrismResponse struct {
	Similarity         float64  `json:"similarity" validate:"required"`
	KnownText          Features `json:"knownText" validate:"required"`
	UnknownText        Features `json:"unknownText" validate:"required"`
}

type Features struct {
	AvgWordLength      float64 `json:"avgWordLength"`
	AvgSentenceLength  float64 `json:"avgSentenceLength"`
	StopWordFrequency  float64 `json:"stopWordFrequency"`
	VocabularyRichness float64 `json:"vocabularyRichness"`
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

	textOne, textTwo := fileOne["message"], fileTwo["message"]

	var similarity float64
	switch fileOne["checkerType"] {
	case "cosine-similarity":
		similarity = core.CheckPlagrism(textOne, textTwo)
	case "minhash":
		similarity = core.MinHash(textOne, textTwo)
	case "robin-karp":
		similarity = core.RabinKarp(textOne, textTwo)
	default:
		writeError(w, http.StatusInternalServerError, "Invalid Checker Type")
		return
	}

	textOneFeatures := core.ExtractFeature(textOne)
	textTwoFeatures := core.ExtractFeature(textTwo)

	response := PlagrismResponse{
		Similarity: similarity,
		KnownText: Features{
			AvgWordLength:      textOneFeatures.AvgWordLength,
			AvgSentenceLength:  textOneFeatures.AvgSentenceLength,
			StopWordFrequency:  textOneFeatures.StopWordFrequency,
			VocabularyRichness: textOneFeatures.VocabularyRichness,
		},
		UnknownText: Features{
			AvgWordLength:      textTwoFeatures.AvgWordLength,
			AvgSentenceLength:  textTwoFeatures.AvgSentenceLength,
			StopWordFrequency:  textTwoFeatures.StopWordFrequency,
			VocabularyRichness: textTwoFeatures.VocabularyRichness,
		},
	}

	writeJSONResponse(w, http.StatusOK, response)
}
