package entities

import "time"

type Kanban struct {
	Id            int
	Name          string
	KanbanColumns []KanbanColumn
	CreatedAt     time.Time
}
