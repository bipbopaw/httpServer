package scheduler

import (
	"calendar-notes-api/internal/notifier"
	"calendar-notes-api/internal/repository"
	"context"
	"log"
	"time"
)

type Scheduler struct {
	repo     repository.NoteRepository
	notifier notifier.Notifier
}

func NewScheduler(repo repository.NoteRepository, notifier notifier.Notifier) *Scheduler {
	return &Scheduler{
		repo:     repo,
		notifier: notifier,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			s.checkAndNotify(ctx)
		}
	}
}

func (s *Scheduler) checkAndNotify(ctx context.Context) {
	notes, err := s.repo.List(ctx)
	if err != nil {
		log.Printf("scheduler error: %v\n", err)
		return
	}

	now := time.Now()
	for _, note := range notes {
		timeUntilEvent := note.EventTime.Sub(now)
		if timeUntilEvent <= note.NotifyBefore && timeUntilEvent > 0 {
			if err := s.notifier.Notify(note); err != nil {
				log.Printf("notification error: %v\n", err)
			}
		}
	}
}
