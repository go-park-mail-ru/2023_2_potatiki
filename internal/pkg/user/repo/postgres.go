package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	profileExistsByID    = "SELECT login, description, imgsrc FROM profile WHERE id=$1;"
	profileExistsByLogin = "SELECT login, description, imgsrc, passwordhash FROM profile WHERE login=$1;"
	addProfile           = "INSERT INTO profile(id, login, description, imgsrc, passwordhash) VALUES($1, $2, $3, $4, $5);"
	updateProfileInfo    = "UPDATE profile SET description=$1, passwordhash=$2 WHERE id=$3;"
	updateProfilePhoto   = "UPDATE profile SET imgsrc=$1 WHERE id=$2;"
)

var (
	ErrInvalidPass = errors.New("invalid pass")
)

type UserRepo struct {
	db pgxtype.Querier
}

func NewUserRepo(db pgxtype.Querier) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user models.User) (models.Profile, error) {
	profileID := uuid.New()
	_, err := r.db.Exec(ctx, addProfile,
		profileID, user.Login, "", "default.png", user.PasswordHash)
	if err != nil { // !errcheck.Is(err, sql.ErrNoRows) будут проверять на рк
		err = fmt.Errorf("error happened in rows.Scan: %w", err)

		return models.Profile{}, err
	}

	profile := models.Profile{
		Id:          profileID,
		Login:       user.Login,
		Description: "",
		ImgSrc:      "default.png",
	}

	return profile, nil
}

func (r *UserRepo) CheckUser(ctx context.Context, user models.User) (models.Profile, error) {
	row := r.db.QueryRow(ctx, profileExistsByLogin, user.Login)
	pr := models.Profile{
		Login: user.Login,
	}
	var userPasswordHash string
	if err := row.Scan(&pr.Id, &pr.Description, &pr.ImgSrc, &userPasswordHash); err != nil {
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Profile{}, err
	}

	if userPasswordHash == user.PasswordHash {
		return pr, nil
	}

	return models.Profile{}, ErrInvalidPass
}

func (r *UserRepo) ReadProfile(ctx context.Context, userID uuid.UUID) (models.Profile, error) {
	row := r.db.QueryRow(ctx, profileExistsByID, userID)
	pr := models.Profile{
		Id: userID,
	}
	if err := row.Scan(&pr.Login, &pr.Description, &pr.ImgSrc); err != nil {
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return models.Profile{}, err
	}

	return pr, nil
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
