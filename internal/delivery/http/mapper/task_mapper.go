package mapper

import (
	dto "github.com/kialkuz/task-manager/internal/delivery/http/dto/request"
	"github.com/kialkuz/task-manager/internal/domain"
)

func TaskRequestToDomain(req dto.TaskRequest) *domain.Task {
	return &domain.Task{
		ID:      req.ID,
		Date:    req.Date,
		Title:   req.Title,
		Comment: req.Comment,
		Repeat:  req.Repeat,
	}
}
