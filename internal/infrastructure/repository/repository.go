package repository

import (
	"time"

	"github.com/kialkuz/task-manager/internal/domain"
)

type TaskRepository interface {
	Add(task domain.Task) (int64, error)
	Update(task domain.Task) error
	Delete(id int) error
	Get(id int) (*domain.Task, error)
	GetList(limit int) ([]*domain.Task, error)
	SearchByText(data string, limit int) ([]*domain.Task, error)
	SearchByDate(searchDate time.Time, limit int) ([]*domain.Task, error)
}
