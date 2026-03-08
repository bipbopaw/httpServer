// Package http contains HTTP handlers, middleware and routing
package http

import "net/http"

func NewRouter(handler *Handler) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /notes", handler.CreateNote)
	return mux
}
