package task

import (
	"errors"
	"net/http"

	requestDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/request"
	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httputil "github.com/kialkuz/task-manager/internal/delivery/http/httputil"
	"github.com/kialkuz/task-manager/internal/delivery/http/mapper"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
)

type EditTask struct {
	ID string `json:"id"`
}

func (h *Handler) EditTaskHandler(w http.ResponseWriter, r *http.Request) {
	req, err := httputil.DecodeJSON[requestDto.TaskRequest](r)
	if err != nil {
		h.logger.Error(err)
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if err := h.validator.ValidateStructDTO(req); err != nil {
		h.logger.WithFields(h.formatter.PrepareForLogs(err)).Error(errors.New("validation error"))
		httpService.WriteJsonBadResponse(w, responseDto.NewDetailsError("validation error", h.formatter.ViewFormat(err)))
		return
	}

	task := mapper.TaskRequestToDomain(*req)

	err = h.service.Update(*task)
	if err != nil {
		h.logger.Error(err)
		httpService.WriteJsonInternalServerError(w, responseDto.ErrorResponse{ErrorText: "ошибка редактирования задачи"})
		return
	}

	httpService.WriteJsonOKResponse(w, AddedTask{
		ID: task.ID,
	})
}
