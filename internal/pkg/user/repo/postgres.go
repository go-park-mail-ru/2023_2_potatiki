package repo

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepo struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) UpdatePhoto(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (r *UserRepo) UpdatePassword(ctx context.Context, userID uuid.UUID, password string) error {
	return nil
}

func (r *UserRepo) UpdateDescription(ctx context.Context, userID uuid.UUID, description string) error {
	return nil
}
