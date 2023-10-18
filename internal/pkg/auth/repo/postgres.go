package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

const (
	profileExistsByLogin = "SELECT Id, Description, ImgSrc, PasswordHash FROM profiles WHERE login=$1;"
	profileExistsById    = "SELECT Login, Description, ImgSrc FROM profiles WHERE Id=$1;"
	addProfile           = "INSERT INTO profiles(Id, Login, Description, ImgSrc, PasswordHash) VALUES($1, $2, $3, $4, $5);"
)

var (
	ErrInvalidPass = errors.New("invalid pass")
)

type AuthRepo struct {
	db *sql.DB
}

func NewAuthRepo(db *sql.DB) *AuthRepo {
	return &AuthRepo{
		db: db,
	}
}

func (r *AuthRepo) CreateUser(ctx context.Context, user models.User) (models.Profile, error) {
	profileId := uuid.New()
	_, err := r.db.ExecContext(ctx, addProfile,
		profileId, user.Login, "", "default.png", user.PasswordHash)
	if err != nil && !errors.Is(err, sql.ErrNoRows) { // !errors.Is(err, sql.ErrNoRows) будут проверять на рк
		err = fmt.Errorf("error happened in rows.Scan: %w", err)
		return models.Profile{}, err
	}

	profile := models.Profile{
		Id:          profileId,
		Login:       user.Login,
		Description: "",
		ImgSrc:      "default.png",
	}

	return profile, nil
}

func (r *AuthRepo) CheckUser(ctx context.Context, user models.User) (models.Profile, error) {
	row := r.db.QueryRowContext(ctx, profileExistsByLogin, user.Login)
	pr := models.Profile{
		Login: user.Login,
	}
	var userPasswordHash string
	if err := row.Scan(&pr.Id, &pr.Description, &pr.ImgSrc, &userPasswordHash); err != nil {
		return models.Profile{}, err
	}

	if userPasswordHash == user.PasswordHash {
		return pr, nil
	}

	return models.Profile{}, ErrInvalidPass
}
func (r *AuthRepo) ReadProfile(ctx context.Context, userId uuid.UUID) (models.Profile, error) {
	row := r.db.QueryRowContext(ctx, profileExistsById, userId)
	pr := models.Profile{
		Id: userId,
	}
	if err := row.Scan(&pr.Login, &pr.Description, &pr.ImgSrc); err != nil {
		err = fmt.Errorf("error happened in row.Scan: %w", err)
		return models.Profile{}, err
	}

	return pr, nil
}
