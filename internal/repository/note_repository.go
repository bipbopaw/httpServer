// Package repository provides storage interfaces and implementations for notes
package repository

import (
	"context"
	"errors"
	"httpServer/internal/model"
	"sync"
)

var ErrNotFound = errors.New("note no found")

type NoteRepository interface {
	Create(ctx context.Context, note model.Note) error
	Get(ctx context.Context, id string) (model.Note, error)
	Update(ctx context.Context, note model.Note) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]model.Note, error)
}

type InMemoryNoteRepository struct {
	mu    sync.RWMutex
	notes map[string]model.Note
}

func NewInMemoryNoteRepository() *InMemoryNoteRepository {
	return &InMemoryNoteRepository{
		notes: make(map[string]model.Note),
	}
}

func (r *InMemoryNoteRepository) Create(ctx context.Context, note model.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.notes[note.ID] = note
	return nil
}

func (r *InMemoryNoteRepository) Update(ctx context.Context, note model.Note) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.notes[note.ID]; !ok {
		return ErrNotFound
	}

	r.notes[note.ID] = note
	return nil
}

func (r *InMemoryNoteRepository) Delete(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.notes[id]; !ok {
		return ErrNotFound
	}

	delete(r.notes, id)
	return nil
}

func (r *InMemoryNoteRepository) Get(ctx context.Context, id string) (model.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	note, ok := r.notes[id]
	if !ok {
		return model.Note{}, ErrNotFound
	}

	return note, nil
}

func (r *InMemoryNoteRepository) List(ctx context.Context) ([]model.Note, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result []model.Note

	for _, note := range r.notes {
		result = append(result, note)
	}

	return result, nil
}
