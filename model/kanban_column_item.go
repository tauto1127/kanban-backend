package entities

import "time"

type KanbanColumnItem struct {
	Index       int
	Name        string
	Description string
	DoneAt      time.Time
	CreatedAt   time.Time
}
