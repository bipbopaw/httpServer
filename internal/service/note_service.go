// Package service implements business logic for working with calendar notes
package service

import (
	"calendar-notes-api/internal/model"
	"calendar-notes-api/internal/repository"
	"context"
	"time"

	"github.com/google/uuid"
)

type NoteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) CreateNote(ctx context.Context, title string, description string, eventTime time.Time, notifyBefore time.Duration) (model.Note, error) {

	now := time.Now()

	note := model.Note{
		ID:           uuid.NewString(),
		Title:        title,
		Description:  description,
		EventTime:    eventTime,
		NotifyBefore: notifyBefore,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	err := s.repo.Create(ctx, note)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (s *NoteService) GetNote(ctx context.Context, id string) (model.Note, error) {
	return s.repo.Get(ctx, id)
}

func (s *NoteService) ListNotes(ctx context.Context) ([]model.Note, error) {
	return s.repo.List(ctx)
}

func (s *NoteService) UpdateNote(ctx context.Context, note model.Note) (model.Note, error) {
	note.UpdatedAt = time.Now()

	err := s.repo.Update(ctx, note)
	if err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (s *NoteService) DeleteNote(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
