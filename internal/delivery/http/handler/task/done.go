package task

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
	"github.com/kialkuz/task-manager/internal/services/task/nextDate"
)

func (h *Handler) DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		h.logger.Error(errors.New("Не передан ID " + id))
		httpService.WriteJsonBadResponse(w, responseDto.ErrorResponse{ErrorText: "Не передан ID"})
		return
	}

	numericId, err := strconv.Atoi(id)
	if err != nil {
		h.logger.Error(errors.New("Передан некорректный ID " + id))
		httpService.WriteJsonBadResponse(w, responseDto.ErrorResponse{ErrorText: "Передан некорректный ID"})
		return
	}

	task, err := h.service.Get(numericId)
	if err != nil {
		h.logger.Error(err)
		httpService.WriteJson(w, responseDto.ErrorResponse{ErrorText: "Задача не найдена"}, http.StatusNotFound)
		return
	}

	if task.Repeat == "" {
		err = h.service.Delete(numericId)
		if err != nil {
			h.logger.Error(err)
			httpService.WriteJsonInternalServerError(w, responseDto.ErrorResponse{ErrorText: "Ошибка удаления задачи"})
			return
		}
	} else {
		nextDate, err := nextDate.NextDate(time.Now(), task.Date, task.Repeat)
		if err != nil {
			h.logger.Error(err)
			httpService.WriteJsonInternalServerError(
				w,
				responseDto.ErrorResponse{ErrorText: "Ошибка отметки задачи выполненной"},
			)
			return
		}

		task.Date = nextDate
		err = h.service.Update(*task)
		if err != nil {
			h.logger.Error(err)
			httpService.WriteJsonInternalServerError(
				w,
				responseDto.ErrorResponse{ErrorText: "Ошибка отметки задачи выполненной"},
			)
			return
		}
	}

	httpService.WriteJsonOKResponse(w, responseDto.EmptyResponse{})
}
