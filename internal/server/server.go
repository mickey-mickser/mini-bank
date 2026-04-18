package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/mickey-mickser/mini-bank/internal/handler"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(port string, handlers handler.Handlers) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:              ":" + port,
			ReadTimeout:       5 * time.Second,
			WriteTimeout:      10 * time.Second,
			ReadHeaderTimeout: 2 * time.Second,
			IdleTimeout:       time.Minute,
			MaxHeaderBytes:    1 << 20,
			Handler:           newRoutes(handlers), //TODO change chi
		},
	}
}
func (s *Server) Run() error {
	if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		if closeErr := s.httpServer.Close(); closeErr != nil {
			return errors.Join(err, closeErr)
			//return closeErr
		}
		return err
	}
	return nil
}
