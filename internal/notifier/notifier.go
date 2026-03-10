package notifier

import (
	"calendar-notes-api/internal/model"
	"log"
)

type Notifier interface {
	Notify(note model.Note) error
}

type ConsoleNotifier struct{}

func NewConsoleNotifier() *ConsoleNotifier {
	return &ConsoleNotifier{}
}

func (c *ConsoleNotifier) Notify(note model.Note) error {
	log.Printf("notification: event '%s' happening at %s (notify before %v)\n", note.Title, note.EventTime.Format("2006-01-02 15:04:05"), note.NotifyBefore)
	return nil
}
