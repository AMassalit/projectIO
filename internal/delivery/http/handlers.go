package http

import (
	"fmt"
	"net/http"
	"projectIO/infrastructure/logger"
	"projectIO/internal/service"
	"strconv"

	"github.com/labstack/echo"
)

func newTaskRouters(e *echo.Echo, g service.TaskServiceInterface, log logger.Logger) {
	r := &processRouters{
		actions: g,
		log:     log,
	}

	TGroup := e.Group("/task")
	TGroup.GET("/create", r.CreateTask)
	TGroup.GET("/delete", r.DeleteTaskByID)
	TGroup.GET("/getinfo", r.GetTaskInfoByID)

}

func (r *processRouters) CreateTask(c echo.Context) error {
	taskID := c.QueryParam("task_id")
	if taskID == "" {
		r.log.Error("task_id is required")
		return echo.NewHTTPError(http.StatusBadRequest, "task_id is required")
	}

	taskIDint, err := strconv.Atoi(taskID)
	if err != nil {
		msg := map[string]interface{}{
			"event":  "projectIO",
			"method": "createTask",
			"error":  err.Error(),
		}
		r.log.WithFields(r.log.MapFields(msg)).Error()
		return echo.NewHTTPError(http.StatusBadRequest, "task_id must be a valid integer")
	}

	if err := r.actions.CreateTask(taskIDint); err != nil {
		msg := map[string]interface{}{
			"event":  "projectIO",
			"method": "createTask",
			"error":  err.Error(),
		}
		r.log.WithFields(r.log.MapFields(msg)).Error(fmt.Errorf("failed to create task:  %w", err).Error())
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("failed to create task:  %w", err).Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "task created successfully",
		"task_id": taskIDint,
	})
}

func (r *processRouters) DeleteTaskByID(c echo.Context) error {
	taskID := c.QueryParam("task_id")
	if taskID == "" {
		r.log.Error("task_id is required")
		return echo.NewHTTPError(http.StatusBadRequest, "task_id is required")
	}

	taskIDint, err := strconv.Atoi(taskID)
	if err != nil {
		msg := map[string]interface{}{
			"event":  "projectIO",
			"method": "deleteTaskByID",
			"error":  err.Error(),
		}
		r.log.WithFields(r.log.MapFields(msg)).Error()
		return echo.NewHTTPError(http.StatusBadRequest, "task_id must be a valid integer")
	}

	if err := r.actions.DeleteTaskByID(taskIDint); err != nil {
		msg := map[string]interface{}{
			"event":  "projectIO",
			"method": "deleteTaskByID",
			"error":  err.Error(),
		}
		r.log.WithFields(r.log.MapFields(msg)).Error(fmt.Sprintf("Error while deleting task: %v", err))
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error while deleting task: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "task deleted successfully",
		"task_id": taskIDint,
	})
}

func (r *processRouters) GetTaskInfoByID(c echo.Context) error {
	taskID := c.QueryParam("task_id")
	if taskID == "" {
		r.log.Error("task_id is required")
		return echo.NewHTTPError(http.StatusBadRequest, "task_id is required")
	}

	taskIDint, err := strconv.Atoi(taskID)
	if err != nil {
		msg := map[string]interface{}{
			"event":  "projectIO",
			"method": "getTaskInfoByID",
			"error":  err.Error(),
		}
		r.log.WithFields(r.log.MapFields(msg)).Error()
		return echo.NewHTTPError(http.StatusBadRequest, "task_id must be a valid integer")
	}

	tast, err := r.actions.GetTaskInfoByID(taskIDint)
	if err != nil {
		msg := map[string]interface{}{
			"event":  "projectIO",
			"method": "getTaskInfoByID",
			"error":  err.Error(),
		}
		r.log.WithFields(r.log.MapFields(msg)).Error(fmt.Sprintf("Error while getting task: %v", err))
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error while getting task: %v", err))
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "task info",
		"task":    tast,
	})
}
