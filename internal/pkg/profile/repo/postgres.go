package repo

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v4"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile"
	"github.com/jackc/pgtype/pgxtype"
	uuid "github.com/satori/go.uuid"
)

const (
	profileById        = "SELECT login, description, imgsrc, phone, passwordhash FROM profile WHERE id=$1;"
	profileIdByLogin   = "SELECT id FROM profile WHERE login=$1;"
	addProfile         = "INSERT INTO profile(id, login, description, imgsrc, phone, passwordhash) VALUES($1, $2, $3, $4, $5, $6);"
	updateProfileInfo  = "UPDATE profile SET phone=$1, passwordhash=$2 WHERE id=$3;"
	updateProfilePhoto = "UPDATE profile SET imgsrc=$1 WHERE id=$2;"
)

type ProfileRepo struct {
	db pgxtype.Querier
}

func NewProfileRepo(db pgxtype.Querier) *ProfileRepo {
	return &ProfileRepo{
		db: db,
	}
}

func (r *ProfileRepo) CreateProfile(ctx context.Context, p *models.Profile) error {
	_, err := r.db.Exec(ctx, addProfile,
		p.Id, p.Login, p.Description, p.ImgSrc, p.Phone, p.PasswordHash)

	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *ProfileRepo) ReadProfile(ctx context.Context, Id uuid.UUID) (*models.Profile, error) {
	p := &models.Profile{Id: Id}
	if err := r.db.QueryRow(ctx, profileById, Id).
		Scan(&p.Login, &p.Description, &p.ImgSrc, &p.Phone, &p.PasswordHash); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return &models.Profile{}, profile.ErrProfileNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return &models.Profile{}, err
	}
	return p, nil
}

func (r *ProfileRepo) GetProfileIdByLogin(ctx context.Context, login string) (uuid.UUID, error) {
	row := r.db.QueryRow(ctx, profileIdByLogin, login)
	var Id uuid.UUID
	if err := row.Scan(&Id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return uuid.UUID{}, profile.ErrProfileNotFound
		}
		err = fmt.Errorf("error happened in row.Scan: %w", err)

		return uuid.UUID{}, err
	}

	return Id, nil
}

func (r *ProfileRepo) UpdateProfile(ctx context.Context, p *models.Profile) error {
	_, err := r.db.Exec(ctx, updateProfileInfo, p.Phone, p.PasswordHash, p.Id)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}

func (r *ProfileRepo) UpdatePhoto(ctx context.Context, userID uuid.UUID, photoName string) error {
	_, err := r.db.Exec(ctx, updateProfilePhoto, photoName, userID)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)

		return err
	}

	return nil
}
