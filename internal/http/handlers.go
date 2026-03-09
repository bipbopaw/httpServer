// Package http contains HTTP handlers, middleware and routing
package http

import (
	"calendar-notes-api/internal/model"
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

type updateNoteRequest struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	EventTime    string `json:"event_time"`
	NotifyBefore string `json:"notify_before"`
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

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.service.ListNotes(r.Context())
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Tyoe", "application/json")

	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) GetNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	note, err := h.service.GetNote(r.Context(), id)
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Tyoe", "application/json")

	if err := json.NewEncoder(w).Encode(note); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func (h *Handler) DeleteNote(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	err := h.service.DeleteNote(r.Context(), id)
	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *Handler) UpdateNote(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	var req updateNoteRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
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

	note := model.Note{
		ID:           id,
		Title:        req.Title,
		Description:  req.Description,
		EventTime:    eventTime,
		NotifyBefore: notifyBefore,
	}

	updatedNote, err := h.service.UpdateNote(r.Context(), note)

	if err != nil {
		http.Error(w, "note not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(updatedNote); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
		return
	}
}
