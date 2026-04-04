package task

import (
	"errors"
	"net/http"
	"strconv"

	requestDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/request"
	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httputil "github.com/kialkuz/task-manager/internal/delivery/http/httputil"
	"github.com/kialkuz/task-manager/internal/delivery/http/mapper"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
)

type AddedTask struct {
	ID string `json:"id"`
}

func (h *Handler) AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	req, err := httputil.DecodeJSON[requestDto.TaskRequest](r)
	if err != nil {
		h.logger.Error(err)
		httpService.WriteJsonBadResponse(w, responseDto.ErrorResponse{ErrorText: "invalid body"})
		return
	}

	if err := h.validator.ValidateStructDTO(req); err != nil {
		h.logger.WithFields(h.formatter.PrepareForLogs(err)).Error(errors.New("validation error"))
		httpService.WriteJsonBadResponse(w, responseDto.NewDetailsError("validation error", h.formatter.ViewFormat(err)))
		return
	}

	task := mapper.TaskRequestToDomain(*req)

	id, err := h.service.Add(*task)
	if err != nil {
		h.logger.Error(err)
		httpService.WriteJsonInternalServerError(w, responseDto.ErrorResponse{ErrorText: "error prepare task time"})
		return
	}

	httpService.WriteJson(w, AddedTask{
		ID: strconv.FormatInt(id, 10),
	}, http.StatusCreated)
}
