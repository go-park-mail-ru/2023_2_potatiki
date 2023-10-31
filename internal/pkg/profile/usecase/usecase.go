package usecase

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"slices"
	"strings"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/hasher"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type ProfileUsecase struct {
	log  *slog.Logger
	repo profile.ProfileRepo
}

func NewProfileUsecase(log *slog.Logger, repo profile.ProfileRepo) *ProfileUsecase {
	return &ProfileUsecase{
		log:  log,
		repo: repo,
	}
}

var (
	acceptingFileTypes      = []string{"image/webp", "image/png", "image/jpeg"}
	ErrorForbiddenExtension = errors.New("this file extension is not allowed")
	ErrPassMismatch         = errors.New("password does not match")
	ErrInvalidSignUpPayload = errors.New("SignUpPayload is not valid")
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

func (uc *ProfileUsecase) UpdatePhoto(ctx context.Context, Id uuid.UUID, filePhotoByte []byte, fileType string) error {
	if !slices.Contains(acceptingFileTypes, fileType) {
		return ErrorForbiddenExtension
	}
	fileExtension := strings.TrimPrefix(fileType, "image/")

	photoName := uuid.New().String() + "." + fileExtension

	file, err := os.Create(photoName)
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

func (uc *ProfileUsecase) UpdateInfo(ctx context.Context, Id uuid.UUID, profileInfo *models.ProfileInfo) error {
	if err := validator.New().Struct(profileInfo.SignUpPayload); err != nil {
		return ErrInvalidSignUpPayload
	}

	if profileInfo.Description == profileInfo.NewDescription {
		err := errors.New("profile description and new description are the same")

		return err
	}

	if profileInfo.Password == profileInfo.NewPassword {
		err := errors.New("profile password and new password are the same")

		return err
	}

	profileId, err := uc.repo.GetProfileIdByLogin(ctx, profileInfo.SignUpPayload.Login)
	if err != nil {
		err = fmt.Errorf("error happened in repo.GetProfileIdByUser: %w", err)

		return err
	}

	profile, err := uc.repo.ReadProfile(ctx, profileId)
	if err != nil {
		err = fmt.Errorf("error happened in repo.GetProfileIdByUser: %w", err)

		return err
	}

	if profile.Login != profileInfo.Login {
		err = errors.New("profile login in db and login from client does not match")

		return err
	}

	profile.Description = profileInfo.NewDescription
	profile.PasswordHash = hasher.HashPass(profileInfo.NewPassword)

	err = uc.repo.UpdateProfile(ctx, profile)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdateProfile: %w", err)

		return err
	}

	return nil
}
