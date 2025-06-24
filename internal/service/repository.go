package service

import (
	"errors"
	"projectIO/dto"
	"time"
)

var tasks = make(map[int]*dto.Task)

type TaskRepo struct{}

func NewRepository() TaskRepositoryInterface {
	return &TaskRepo{}
}

func (t *TaskRepo) CreateTask(taskID int, initialStatus string, createdAt time.Time, expiry time.Time) error {
	if _, ok := tasks[taskID]; ok {
		return errors.New("task already exists")
	}

	tasks[taskID] = &dto.Task{
		ID:        taskID,
		Status:    initialStatus,
		CreatedAt: createdAt,
		Expiry:    expiry,
	}
	return nil
}

func (t *TaskRepo) GetTaskByID(taskID int) (*dto.Task, error) {
	task, ok := tasks[taskID]
	if !ok {
		return nil, errors.New("task not found")
	}
	return task, nil
}

func (t *TaskRepo) UpdateTaskStatus(taskID int, newStatus string) error {
	task, ok := tasks[taskID]
	if !ok {
		return errors.New("task not found")
	}
	task.Status = newStatus
	return nil
}

func (t *TaskRepo) DeleteTask(taskID int) error {
	if _, ok := tasks[taskID]; !ok {
		return errors.New("task not found")
	}
	delete(tasks, taskID)
	return nil
}
