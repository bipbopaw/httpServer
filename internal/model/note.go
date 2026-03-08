// Package model defines domain entities used by the calendar service
package model

import "time"

type Note struct {
	ID           string
	Title        string
	Description  string
	EventTime    time.Time
	NotifyBefore time.Duration
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
