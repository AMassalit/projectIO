package httpserver

import (
	"context"
	"time"

	"github.com/labstack/echo"
)

type Server struct {
	server          *echo.Echo
	address         string
	notify          chan error
	shutdownTimeout time.Duration
}

func New(e *echo.Echo, address string) *Server {
	s := &Server{
		server:  e,
		address: address,
		notify:  make(chan error, 1),
	}

	s.start()

	return s
}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.Start(s.address)
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
