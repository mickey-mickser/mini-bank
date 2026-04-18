package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mickey-mickser/mini-bank/internal/handler/dto"
	"github.com/mickey-mickser/mini-bank/internal/service"
)

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var input dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json body")
		return
	}
	if input.Login == "" {
		writeJSONError(w, http.StatusBadRequest, "Login is Empty")
		return
	}
	if input.Password == "" {
		writeJSONError(w, http.StatusBadRequest, "Password is Empty")
		return
	}
	user, err := h.userService.CreateUser(r.Context(), service.CreateUserInput{
		Name:     input.Name,
		Login:    input.Login,
		Password: input.Password,
	})
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, dto.CreateUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	})
}
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(data)
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{
		"error": message,
	})
}
