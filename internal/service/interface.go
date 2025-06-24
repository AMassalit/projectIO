package service

import (
	"projectIO/dto"
	"time"
)

type TaskRepositoryInterface interface {
	CreateTask(taskID int, initialStatus string, createdAt time.Time, expiry time.Time) error
	GetTaskByID(taskID int) (*dto.Task, error)
	UpdateTaskStatus(taskID int, newStatus string) error
	DeleteTask(taskID int) error
}

type TaskServiceInterface interface {
	CreateTask(task int) error
	DeleteTaskByID(task int) error
	GetTaskInfoByID(task int) (dto.TaskResponse, error)
}
