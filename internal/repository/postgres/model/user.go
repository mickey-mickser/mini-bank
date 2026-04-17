package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Role struct {
	value string
}

var (
	UserRole  = Role{value: "user"}
	AdminRole = Role{value: "admin"}
)

func NewRole(v string) (Role, error) {
	switch role {
	case "user":
		return UserRole, nil
	case "admin":
		return AdminRole, nil
	default:
		return Role{}, fmt.Errorf("invalid role: %q", v)
	}
}
func (v Role) String() string {
	return v.value
}

type UserStatus string
type AccountStatus string

type Users struct {
	ID           uuid.UUID ``
	Name         string
	Login        string
	PasswordHash string
	Status       string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Accounts struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	Currency  string
	Balance   int64
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
