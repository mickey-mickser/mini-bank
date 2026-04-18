package server

import (
	"net/http"

	"github.com/mickey-mickser/mini-bank/internal/handler"
)

func newRoutes(h handler.Handlers) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/createUser", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h.User.CreateUser(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	return mux
}
