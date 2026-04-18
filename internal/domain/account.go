package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

// Account
type AccountStatus string

const (
	AccountStatusActive   AccountStatus = "active"
	AccountStatusInactive AccountStatus = "inactive"
)

type Account struct {
	ID        uuid.UUID
	UserID    string
	Currency  string
	Balance   decimal.Decimal
	Status    AccountStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}
