package handler

import (
	"net/http"
	"text-analyzer/handlers"
)

func HandlePlagrism(w http.ResponseWriter, r *http.Request) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/plagrismChecker", handlers.PlagrismChecker)
	handlers.EnableCORS(mux).ServeHTTP(w, r)
}