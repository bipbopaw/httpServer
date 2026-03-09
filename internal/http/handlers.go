// Package http contains HTTP handlers, middleware and routing
package http

import (
	"calendar-notes-api/internal/service"
	"encoding/json"
	"net/http"
	"time"
)

type Handler struct {
	service *service.NoteService
}

func NewHandler(service *service.NoteService) *Handler {
	return &Handler{service: service}
}

type createNoteRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	EventTime    string `json:"event_time"`
	NotifyBefore string `json:"notify_before"`
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req createNoteRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	eventTime, err := time.Parse(time.RFC3339, req.EventTime)
	if err != nil {
		http.Error(w, "invalid event_time", http.StatusBadRequest)
		return
	}

	notifyBefore, err := time.ParseDuration(req.NotifyBefore)
	if err != nil {
		http.Error(w, "invalid notify_before", http.StatusBadRequest)
		return
	}

	note, err := h.service.CreateNote(
		r.Context(),
		req.Title,
		req.Description,
		eventTime,
		notifyBefore,
	)

	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(note)
	if err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
