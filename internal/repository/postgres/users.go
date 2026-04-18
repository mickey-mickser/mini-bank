package postgres

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mickey-mickser/mini-bank/internal/domain"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *domain.User) error
}
type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) CreateUser(ctx context.Context, user *domain.User) error {
	query := `INSERT INTO users (
                   name,
                   login,
                   password_hash
                   ) 
	VALUES ($1, $2, $3) 
	RETURNING id, name, created_at`

	err := r.db.QueryRow(ctx, query,
		user.Name,
		user.Login,
		user.PasswordHash,
	).Scan(
		&user.ID,
		&user.Name,
		&user.CreatedAt,
	)
	if err != nil {
		if strings.Contains(err.Error(), "uq_users_login_lower") {
			return errors.New("login already exist")
		}
		return err
	}
	return nil
}
