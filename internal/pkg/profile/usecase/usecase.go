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

	return profile, nil
}

func (uc *ProfileUsecase) UpdateData(ctx context.Context, Id uuid.UUID, payload *models.UpdateProfileDataPayload) (*models.Profile, error) { //nolint:cyclop
	if err := validator.New().Struct(payload); err != nil {
		return nil, err
	}

	if payload.Phone != "" {
		if (payload.Passwords.OldPass == "") !=
			(payload.Passwords.NewPass == "") {
			return nil, profile.ErrBadUpdateData
		}
	} else {
		if payload.Passwords.OldPass == "" ||
			payload.Passwords.NewPass == "" {
			return nil, profile.ErrBadUpdateData
		}
	}

	profileInfo, err := uc.repo.ReadProfile(ctx, Id)
	if err != nil {
		return nil, fmt.Errorf("error happened in repo.ReadProfile: %w", err)
	}

	if payload.Passwords.OldPass != "" && payload.Passwords.NewPass != "" {
		if hasher.CheckPass(profileInfo.PasswordHash, payload.Passwords.OldPass) {
			profileInfo.PasswordHash = hasher.HashPass(payload.Passwords.NewPass)
		} else {
			return nil, profile.ErrBadUpdateData
		}
	}

	if payload.Phone != "" {
		profileInfo.Phone = payload.Phone
	}

	if err = uc.repo.UpdateProfile(ctx, profileInfo); err != nil {
		return nil, fmt.Errorf("error happened in repoUser.UpdateProfile: %w", err)
	}

	return profileInfo, nil
}

func (uc *ProfileUsecase) UpdatePhoto(ctx context.Context, Id uuid.UUID, filePhotoByte []byte, fileType string) (*models.Profile, error) {
	if !slices.Contains(acceptingFileTypes, fileType) {
		return nil, profile.ErrorForbiddenExtension
	}

	_, err := os.Stat(uc.photoPath)
	if err != nil {
		if os.IsNotExist(err) {
			err = fmt.Errorf("photos upload file path [%s] is not exist, error: %w", uc.photoPath, err)

			return nil, err
		}

		return nil, fmt.Errorf("bad photos file path [%s], error: %w", uc.photoPath, err)
	}

	fileExtension := strings.TrimPrefix(fileType, "image/")

	photoName := uuid.NewV4().String() + "." + fileExtension

	file, err := os.Create(uc.photoPath + photoName)
	if err != nil {
		err = fmt.Errorf("error happened in create file: %w", err)

		return nil, err
	}
	defer file.Close()

	_, err = file.Write(filePhotoByte)
	if err != nil {
		err = fmt.Errorf("error happened in write to file: %w", err)

		return nil, err
	}

	err = uc.repo.UpdatePhoto(ctx, Id, photoName)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdatePhoto: %w", err)

		return nil, err
	}

	profileInfo, err := uc.repo.ReadProfile(ctx, Id)

	return profileInfo, nil
}
