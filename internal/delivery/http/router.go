package http

import (
	"net/http"
	"projectIO/infrastructure/logger"

	"projectIO/internal/service"

	"github.com/labstack/echo"
)

type processRouters struct {
	actions service.TaskServiceInterface
	log     logger.Logger
}

func NewRouter(e *echo.Echo, g service.TaskServiceInterface, log logger.Logger) {
	e.GET("/healthz/l", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "ok",
			"Status":  true,
		})
	})

	newTaskRouters(e, g, log)
}
