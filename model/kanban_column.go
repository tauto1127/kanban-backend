package entities

import "time"

type KanbanColumn struct {
	Id                int
	Name              string
	KanbanColumnItems []KanbanColumnItem
	CreatedAt         time.Time
}
