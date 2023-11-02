package usecase

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
)

type ProfileUsecase struct {
	log       *slog.Logger
	repo      profile.ProfileRepo
	photoPath string
}

func NewProfileUsecase(log *slog.Logger, repo profile.ProfileRepo, cfg profile.ProfileConfig) *ProfileUsecase {
	return &ProfileUsecase{
		log:       log,
		repo:      repo,
		photoPath: cfg.GetPhotosFilePath(),
	}
}

var (
	acceptingFileTypes = []string{"image/webp", "image/png", "image/jpeg"}
)

// TODO: nginx path from env to save imgs

func (uc *ProfileUsecase) GetProfile(ctx context.Context, Id uuid.UUID) (*models.Profile, error) {
	profile, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return &models.Profile{}, err
	}
	profile.HidePass()

	return profile, nil
}

func (uc *ProfileUsecase) UpdateData(ctx context.Context, Id uuid.UUID, payload *models.UpdateProfileDataPayload) error {
	if err := validator.New().Struct(payload); err != nil {
		return err
	}

	profil, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		err = fmt.Errorf("error happened in repo.ReadProfile: %w", err)

		return err
	}

	if profil.Phone == payload.Phone &&
		hasher.CheckPass(profil.PasswordHash, payload.Password) {
		return profile.ErrDoubleData
	}

	profil.Phone = payload.Phone
	profil.PasswordHash = hasher.HashPass(payload.Password)

	err = uc.repo.UpdateProfile(ctx, profil)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdateProfile: %w", err)

		return err
	}

	return nil
}

func (uc *ProfileUsecase) UpdatePhoto(ctx context.Context, Id uuid.UUID, filePhotoByte []byte, fileType string) error {
	if !slices.Contains(acceptingFileTypes, fileType) {
		return profile.ErrorForbiddenExtension
	}

	_, err := os.Stat(uc.photoPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("photos upload file path [%s] is not exist, error: %w", uc.photoPath, err)

			return err
		}

		return fmt.Errorf("bad photos file path [%s], error: %w", uc.photoPath, err)
	}

	fileExtension := strings.TrimPrefix(fileType, "image/")

	photoName := uuid.NewV4().String() + "." + fileExtension

	file, err := os.Create(uc.photoPath + photoName)
	if err != nil {
		err = fmt.Errorf("error happened in create file: %w", err)

		return err
	}
	defer file.Close()

	_, err = file.Write(filePhotoByte)
	if err != nil {
		err = fmt.Errorf("error happened in write to file: %w", err)

		return err
	}

	err = uc.repo.UpdatePhoto(ctx, Id, photoName)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdatePhoto: %w", err)

		return err
	}

	return nil
}
