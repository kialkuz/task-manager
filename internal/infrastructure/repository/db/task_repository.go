package db

import (
	"fmt"
	"time"

	"database/sql"

	"github.com/kialkuz/task-manager/internal/domain"
	"github.com/kialkuz/task-manager/internal/infrastructure/repository"
	datePkg "github.com/kialkuz/task-manager/pkg/date"

	_ "github.com/lib/pq"
)

type TaskRepository struct {
	db *sql.DB
}

func NewTaskRepository(db *sql.DB) repository.TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) Add(task domain.Task) (int64, error) {
	var id int64
	query := `INSERT INTO scheduler (date, title, comment, repeat) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRow(query, task.Date, task.Title, task.Comment, task.Repeat).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *TaskRepository) Update(task domain.Task) error {
	query := `UPDATE scheduler SET date = $1, title = $2, comment = $3, repeat = $4 WHERE id = $5`
	res, err := r.db.Exec(query,
		task.Date,
		task.Title,
		task.Comment,
		task.Repeat,
		task.ID,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf(`incorrect id for updating task`)
	}
	return nil
}

func (r *TaskRepository) Delete(id int) error {
	err := r.db.QueryRow("DELETE FROM scheduler WHERE id = $1", id).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) Get(id int) (*domain.Task, error) {
	task := domain.Task{}
	query := "SELECT id, date, title, comment, repeat FROM scheduler WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepository) GetList(limit int) ([]*domain.Task, error) {
	rows, err := r.db.Query("SELECT id, date, title, comment, repeat FROM scheduler ORDER BY date DESC LIMIT $1",
		limit,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	return makeList(rows)
}

func (r *TaskRepository) SearchByText(data string, limit int) ([]*domain.Task, error) {
	var rows *sql.Rows

	search := "%" + data + "%"

	rows, err := r.db.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE title LIKE $1
		OR comment LIKE $1
		ORDER BY date LIMIT $2`,
		search,
		limit,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return makeList(rows)
}

func (r *TaskRepository) SearchByDate(searchDate time.Time, limit int) ([]*domain.Task, error) {
	var rows *sql.Rows

	rows, err := r.db.Query(`
		SELECT id, date, title, comment, repeat
		FROM scheduler
		WHERE date = $1 LIMIT $2`,
		searchDate.Format(datePkg.DateFormat),
		limit,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	return makeList(rows)
}

func makeList(rows *sql.Rows) ([]*domain.Task, error) {
	tasks := make([]*domain.Task, 0)

	for rows.Next() {
		task := domain.Task{}

		err := rows.Scan(&task.ID, &task.Date, &task.Title, &task.Comment, &task.Repeat)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, &task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
