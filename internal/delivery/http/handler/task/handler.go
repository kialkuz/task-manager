package task

import (
	"net/http"

	"github.com/kialkuz/task-manager/internal/middleware"
	"github.com/kialkuz/task-manager/internal/services/task"
	"github.com/kialkuz/task-manager/internal/validator"
	"github.com/kialkuz/task-manager/pkg/logger"
)

type Handler struct {
	validator *validator.Validator
	formatter *validator.ErrorFormatter
	logger    logger.LoggerInterface
	service   task.TaskService
}

func NewHandler(
	validator *validator.Validator,
	formatter *validator.ErrorFormatter,
	logger logger.LoggerInterface,
	service task.TaskService,
) *Handler {
	return &Handler{
		validator: validator,
		formatter: formatter,
		logger:    logger,
		service:   service,
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/nextdate", middleware.Auth(h.NextDayHandler))
	mux.HandleFunc("GET /api/task", middleware.Auth(h.TaskHandler))
	mux.HandleFunc("POST /api/task", middleware.Auth(h.AddTaskHandler))
	mux.HandleFunc("PUT /api/task", middleware.Auth(h.EditTaskHandler))
	mux.HandleFunc("GET /api/tasks", h.TasksHandler)
	mux.HandleFunc("DELETE /api/task", middleware.Auth(h.DeleteTaskHandler))
	mux.HandleFunc("POST /api/task/done", middleware.Auth(h.DoneTaskHandler))
}
