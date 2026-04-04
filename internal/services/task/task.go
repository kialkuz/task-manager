package task

import (
	"time"

	"github.com/kialkuz/task-manager/internal/domain"
	"github.com/kialkuz/task-manager/internal/infrastructure/repository"
)

type TaskService struct {
	repository repository.TaskRepository
}

func NewService(repository repository.TaskRepository) TaskService {
	return TaskService{repository: repository}
}

func (t *TaskService) Add(task domain.Task) (int64, error) {
	err := task.PrepareDateByRules()
	if err != nil {
		return 0, err
	}

	id, err := t.repository.Add(task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (t *TaskService) Update(task domain.Task) error {
	err := task.PrepareDateByRules()
	if err != nil {
		return err
	}

	err = t.repository.Update(task)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskService) Delete(id int) error {
	err := t.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

func (t *TaskService) Get(id int) (*domain.Task, error) {
	task, err := t.repository.Get(id)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskService) GetList(limit int) ([]*domain.Task, error) {
	tasks, err := t.repository.GetList(limit)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) SearchByText(data string, limit int) ([]*domain.Task, error) {
	tasks, err := t.repository.SearchByText(data, limit)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (t *TaskService) SearchByDate(searchDate time.Time, limit int) ([]*domain.Task, error) {
	tasks, err := t.repository.SearchByDate(searchDate, limit)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}
