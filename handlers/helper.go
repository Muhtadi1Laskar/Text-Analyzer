package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type CommonRequest struct {
	Message string `json:"message" validate:"required"`
}

type CommonResponse struct {
	Message string `json:"message" validate:"required"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func readRequestBody(r *http.Request, target interface{}) error {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(reqBody, target); err != nil {
		return fmt.Errorf("unable to read request body: %v", err)
	}

	validate := validator.New()
	err = validate.Struct(target)
	if err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	return nil
}

func writeJSONResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func writeError(w http.ResponseWriter, statusCode int, err string) {
	writeJSONResponse(w, statusCode, ErrorResponse{ Error: err })
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
