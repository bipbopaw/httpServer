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

func (s *NoteService) CreateNote(
	ctx context.Context,
	title string,
	description string,
	eventTime time.Time,
	notifyBefore time.Duration,
) (model.Note, error) {

	note := model.Note{
		ID:           uuid.NewString(),
		Title:        title,
		Description:  description,
		EventTime:    eventTime,
		NotifyBefore: notifyBefore,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	err := s.repo.Create(ctx, note)
	if err != nil {
		return model.Note{}, nil
	}

	return note, nil
}
