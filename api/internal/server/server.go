// Package server wires the HTTP routes: the JSON API under /api/ and the
// embedded frontend as catch-all.
package server

import (
	"encoding/json"
	"net/http"

	"github.com/LuisMDeveloper/fullstack-template/api/internal/webui"
)

func New() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"status": "ok"})
	})
	mux.HandleFunc("GET /api/hello", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]string{"message": "Hello, World!"})
	})
	// Static payload to exercise the frontend.
	mux.HandleFunc("GET /api/info", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, map[string]any{
			"name":     "fullstack-template",
			"version":  "0.1.0",
			"language": "Go",
			"features": []string{"embedded frontend", "json api"},
		})
	})
	// Anything else under /api/ must be JSON, never the frontend's index.html.
	mux.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
	})
	mux.Handle("/", webui.Handler())
	return mux
}

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(v)
}
