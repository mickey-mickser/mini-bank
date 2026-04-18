package service

import (
	"time"

	"github.com/google/uuid"
)

type CreateUserOutput struct {
	ID        uuid.UUID
	Name      string
	CreatedAt time.Time
}
