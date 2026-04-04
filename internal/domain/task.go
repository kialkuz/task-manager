package domain

import (
	"time"

	"github.com/kialkuz/task-manager/internal/services/task/nextDate"
	datePkg "github.com/kialkuz/task-manager/pkg/date"
)

type Task struct {
	ID      string `json:"id"`
	Date    string `json:"date"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Repeat  string `json:"repeat"`
}

func (t *Task) PrepareDateByRules() error {
	now := time.Now()

	if t.Date != "" {
		parsedTime, _ := time.Parse(datePkg.DateFormat, t.Date)
		t.Date = parsedTime.Format(datePkg.DateFormat)

		if datePkg.IsDateAfter(now, parsedTime) {
			if len(t.Repeat) == 0 {
				// если правила повторения нет, то берём сегодняшнее число
				t.Date = now.Format(datePkg.DateFormat)
			} else {
				// в противном случае, берём вычисленную ранее следующую дату
				next, err := nextDate.NextDate(now, t.Date, t.Repeat)
				if err != nil {
					return err
				}
				t.Date = next
			}
		}
	} else {
		t.Date = now.Format(datePkg.DateFormat)
	}

	return nil
}
