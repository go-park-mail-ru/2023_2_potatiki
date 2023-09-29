package repo

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
)

const (
	profileExists = "SELECT Id, Description, ImgSrc, PasswordHash FROM public.profiles WHERE login=$1;"
	addProfile    = "INSERT INTO public.profiles(Id, Login, Description, ImgSrc, PasswordHash) VALUES($1, $2, $3, $4, $5);"
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
		profileId, user.Login, "", "default.png", user.PasswordHash) //sql иньекции + константу
	if err != nil {
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
	row := r.db.QueryRowContext(ctx, profileExists, user.Login)
	var (
		userId          uuid.UUID
		userDescription string
		userImgSrc      string
	)
	var userPasswordHash string
	if err := row.Scan(&userId, &userDescription, &userImgSrc, &userPasswordHash); err != nil {
		return models.Profile{}, err
	}

	if userPasswordHash == user.PasswordHash {
		return models.Profile{
			Id:          userId,
			Login:       user.Login,
			Description: userDescription,
			ImgSrc:      userImgSrc,
		}, nil
	}
	return models.Profile{}, ErrInvalidPass
}
func (r *AuthRepo) ReadProfile(context.Context, uuid.UUID) (models.Profile, error) {
	panic("unimplemented")
}
