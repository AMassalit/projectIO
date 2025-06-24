package service

import (
	"math/rand"
	"projectIO/infrastructure/logger"
	"projectIO/utils"
	"time"

	"projectIO/dto"
)

type TaskService struct {
	repo   TaskRepositoryInterface
	logger *logger.Logger
}

func NewService(repo TaskRepositoryInterface, logger *logger.Logger) TaskServiceInterface {
	return &TaskService{repo: repo, logger: logger}
}

func (t *TaskService) CreateTask(taskID int) error {

	rand.Seed(time.Now().UnixNano())
	duration := time.Duration(3+rand.Intn(3)) * time.Minute

	return t.repo.CreateTask(taskID, "pending", time.Now(), time.Now().Add(duration))
}
func (t *TaskService) GetTaskInfoByID(taskID int) (dto.TaskResponse, error) {
	task, err := t.repo.GetTaskByID(taskID)
	if err != nil {
		return dto.TaskResponse{}, err
	}

	if time.Now().After(task.Expiry) {
		if task.Status != "completed" {
			task.Status = "completed"

			t.repo.UpdateTaskStatus(taskID, "completed")
		}
	}

	durationFormatted := utils.FormatDuration(time.Since(task.CreatedAt))
	createdAtFormatted := task.CreatedAt.Format("02.01.2006 15:04:05")

	return dto.TaskResponse{
		ID:        task.ID,
		Status:    task.Status,
		CreatedAt: createdAtFormatted,
		Duration:  durationFormatted,
	}, nil
}

func (t *TaskService) DeleteTaskByID(TaskID int) error {
	return t.repo.DeleteTask(TaskID)
}
