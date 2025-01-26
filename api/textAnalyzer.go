package handler

import (
	"net/http"
	"text-analyzer/handlers"
)

func HandlerTextAnalyzer(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/textAnalyzer", handlers.TextAnalyzer)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}
