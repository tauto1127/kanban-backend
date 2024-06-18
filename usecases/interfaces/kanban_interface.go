package interfaces
import "github.com/tauto1127/kanban-backend/model"

type KanbanInterface interface {
	CreateKanban(name string) (entities.Kanban, error)
	GetKanban(id int) (entities.Kanban, error)
	DeleteKanban(id int) error
}