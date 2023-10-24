package usecase

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/user"
	"github.com/google/uuid"
)

type UserUsecase struct {
	repoUser user.UserRepo
	repoAuth auth.AuthRepo
}

func NewUserUsecase(repoUser user.UserRepo, repoAuth auth.AuthRepo) *UserUsecase {
	return &UserUsecase{
		repoUser: repoUser,
		repoAuth: repoAuth,
	}
}

func (uc *UserUsecase) UpdatePhoto(ctx context.Context, photoID uuid.UUID) error {
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
