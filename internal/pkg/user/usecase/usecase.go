package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/google/uuid"
	"log/slog"
	"os"
	"slices"
	"strings"
)

type UserUsecase struct {
	log      *slog.Logger
	repoUser user.UserRepo
	repoAuth auth.AuthRepo
}

var (
	acceptingFileTypes      = []string{"image/webp", "image/png", "image/jpeg"}
	ErrorForbiddenExtension = errors.New("this file extension is not allowed")
)

// TODO: nginx path from env to save imgs

func NewUserUsecase(repoUser user.UserRepo, repoAuth auth.AuthRepo, log *slog.Logger) *UserUsecase {
	return &UserUsecase{
		log:      log,
		repoUser: repoUser,
		repoAuth: repoAuth,
	}
}

func (uc *UserUsecase) UpdatePhoto(ctx context.Context, userID uuid.UUID, filePhotoByte []byte, fileType string) error {
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

	err = uc.repoUser.UpdatePhoto(ctx, userID, photoName)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdatePhoto: %w", err)

		return err
	}

	return nil
}

func (uc *UserUsecase) UpdateInfo(ctx context.Context, userID uuid.UUID, profileInfo models.ProfileInfo) error {
	if !profileInfo.User.IsValid() {
		err := errors.New("user is not valid")

		return err
	}

	if profileInfo.Description == profileInfo.NewDescription {
		err := errors.New("user description and new description are the same")

		return err
	}

	if profileInfo.PasswordHash == profileInfo.NewPasswordHash {
		err := errors.New("user password and new password are the same")

		return err
	}

	profile, err := uc.repoAuth.CheckUser(ctx, profileInfo.User)
	if err != nil {
		err = fmt.Errorf("error happened in repo.CheckUser: %w", err)

		return err
	}
	if profile.Id != userID {
		err = errors.New("user id in db and id from token does not match")

		return err
	}
	if profile.Login != profileInfo.Login {
		err = errors.New("user login in db and login from client does not match")

		return err
	}

	err = uc.repoUser.UpdateInfo(ctx, userID, profileInfo.UserInfo)
	if err != nil {
		err = fmt.Errorf("error happened in repoUser.UpdateInfo: %w", err)

		return err
	}

	return nil
}
