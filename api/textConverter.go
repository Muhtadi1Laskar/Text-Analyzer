package handler

import (
	"net/http"
	"text-analyzer/handlers"
)

func HandleTextConverter(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/textConverter", handlers.TextCleaner)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}