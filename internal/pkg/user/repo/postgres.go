package repo

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	updateProfileInfo  = "UPDATE profile SET Description=$1, PasswordHash=$2 WHERE Id=$3;"
	updateProfilePhoto = "UPDATE profile SET imgsrc=$1 WHERE Id=$2;"
)

type UserRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) UpdatePhoto(ctx context.Context, userID uuid.UUID, photoName string) error {
	_, err := r.db.Exec(ctx, updateProfilePhoto, photoName, userID)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *UserRepo) UpdateInfo(ctx context.Context, userID uuid.UUID, userInfo models.UserInfo) error {
	_, err := r.db.Exec(ctx, updateProfileInfo, userInfo.NewDescription, userInfo.NewPasswordHash, userID)
	if err != nil { // !errcheck.Is(err, sql.ErrNoRows) будут проверять на рк
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}
