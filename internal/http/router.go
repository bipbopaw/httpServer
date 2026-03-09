// Package http contains HTTP handlers, middleware and routing
package http

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /notes", handler.CreateNote)
	mux.HandleFunc("GET /notes", handler.ListNotes)
	mux.HandleFunc("GET /notes/{id}", handler.GetNote)
	mux.HandleFunc("DELETE /notes/{id}", handler.DeleteNote)
	mux.HandleFunc("PATCH /notes/{id}", handler.UpdateNote)
	return mux
}
