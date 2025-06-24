package app

import (
	"os"
	"os/signal"
	"projectIO/config"
	"projectIO/infrastructure/logger"
	http "projectIO/internal/delivery/http"
	"projectIO/internal/service"
	"projectIO/pkg/httpserver"
	"syscall"

	"github.com/labstack/echo"
)

func Run(cfg *config.Config) {

	newLogger := logger.NewLogger(cfg.Log.Level)
	taskRepo := service.NewRepository()
	taskService := service.NewService(taskRepo, newLogger)
	newLogger.WithField("address", cfg.HTTP.Address()).Info("Start web service")

	httpRouter := echo.New()
	http.NewRouter(httpRouter, taskService, *newLogger)
	httpServer := httpserver.New(httpRouter, cfg.HTTP.Address())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		newLogger.WithFields(map[string]interface{}{"server": "interrupt process"}).Info()
	case err := <-httpServer.Notify():
		newLogger.WithFields(map[string]interface{}{"server": "shutdown signal"}).Info()
		if err != nil {
			newLogger.WithFields(map[string]interface{}{"shutdown": err.Error()}).Error()
		}
	}
	err := httpServer.Shutdown()
	if err != nil {
		newLogger.WithFields(map[string]interface{}{"server": "shutdown signal error: " + err.Error()}).Fatal()
	}

}
