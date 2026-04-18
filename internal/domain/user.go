package domain

import (
	"time"

	"github.com/google/uuid"
)

// User
type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
)

type UserStatus string

const (
	UserStatusOnline  UserStatus = "online"
	UserStatusOffline UserStatus = "offline"
)

type User struct {
	ID           uuid.UUID
	Name         string
	Login        string
	PasswordHash string
	Status       UserStatus
	Role         UserRole
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
