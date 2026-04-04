package task

import (
	"log"
	"net/http"
	"time"

	responseDto "github.com/kialkuz/task-manager/internal/delivery/http/dto/response"
	httpService "github.com/kialkuz/task-manager/internal/delivery/http/services"
	"github.com/kialkuz/task-manager/internal/services/task/nextDate"
	datePkg "github.com/kialkuz/task-manager/pkg/date"
)

func (h *Handler) NextDayHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	now, err := time.Parse(datePkg.DateFormat, q.Get("now"))
	if err != nil {
		log.Println(err)
		httpService.WriteJsonBadResponse(w, responseDto.ErrorResponse{ErrorText: "ошибка получения новой даты"})
		return
	}

	nextDate, err := nextDate.NextDate(now, q.Get("date"), q.Get("repeat"))
	if err != nil {
		log.Println(err)
		httpService.WriteJsonInternalServerError(w, responseDto.ErrorResponse{ErrorText: "ошибка получения новой даты"})
		return
	}

	httpService.WriteJsonWithoutSerialize(w, []byte(nextDate), http.StatusOK)
}
