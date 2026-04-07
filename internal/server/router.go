package server

import (
	"net/http"

	"github.com/mickey-mickser/mini-bank/internal/handler"
)

func newRoutes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthHandler)
	return mux
}
