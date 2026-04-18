package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
