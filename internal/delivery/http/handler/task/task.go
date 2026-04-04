package task

import (
	"log"
	"net/http"
	"strconv"

	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
)

func (h *Handler) TaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("Не передан ID " + id)
		httpService.WriteJsonBadResponse(w, responseDto.ErrorResponse{ErrorText: "Не передан ID"})
		return
	}

	numericId, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Передан некорректный ID " + id)
		httpService.WriteJsonBadResponse(w, responseDto.ErrorResponse{ErrorText: "Передан некорректный ID"})
		return
	}

	task, err := h.service.Get(numericId)
	if err != nil {
		log.Println(err)
		httpService.WriteJson(w, responseDto.ErrorResponse{ErrorText: "Задача не найдена"}, http.StatusNotFound)
		return
	}

	httpService.WriteJsonOKResponse(w, task)
}
