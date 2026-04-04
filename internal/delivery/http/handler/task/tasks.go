package task

import (
	"log"
	"net/http"
	"time"

	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
	"github.com/kialkuz/task-manager/internal/domain"
)

type TasksResp struct {
	Tasks []*domain.Task `json:"tasks"`
}

const tasksLimit = 50

func (h *Handler) TasksHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []*domain.Task
	var err error

	q := r.URL.Query()

	search := q.Get("search")
	if search == "" {
		tasks, err = h.service.GetList(50)
	} else {
		var searchDate time.Time
		searchDate, err = time.Parse("02.01.2006", search)

		if err != nil {
			tasks, err = h.service.SearchByText(search, tasksLimit)
		} else {
			tasks, err = h.service.SearchByDate(searchDate, tasksLimit)
		}
	}

	if err != nil {
		log.Println(err)
		httpService.WriteJsonInternalServerError(w, responseDto.ErrorResponse{ErrorText: "ошибка получения"})
		return
	}

	httpService.WriteJsonOKResponse(w, TasksResp{
		Tasks: tasks,
	})
}
