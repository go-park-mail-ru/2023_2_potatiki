package repo

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	updateProfileInfo = "UPDATE profile(Description, PasswordHash) SET Description=$1, PasswordHash=$2 WHERE Id=$3;"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) UpdatePhoto(ctx context.Context, userID uuid.UUID) error {
	return nil
}

func (r *UserRepo) UpdateInfo(ctx context.Context, userID uuid.UUID, userInfo models.UserInfo) error {
	_, err := r.db.Exec(ctx, updateProfileInfo, userInfo.NewDescription, userInfo.NewPasswordHash, userID)
	if err != nil { // !errcheck.Is(err, sql.ErrNoRows) будут проверять на рк
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return err
	}

	return nil
}
