package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	generatedAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/generated"
	uuid "github.com/satori/go.uuid"
)

//go:generate mockgen -source=./generated/auth_grpc.pb.go -destination=../../mocks/auth_grpc.go -package=mock

type GrpcAuthHandler struct {
	uc auth.AuthUsecase
	// TODO: ADd logger

	generatedAuth.AuthServiceServer
}

func NewGrpcAuthHandler(uc auth.AuthUsecase) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		uc: uc,
	}
}

func (h GrpcAuthHandler) SignIn(ctx context.Context, in *generatedAuth.SignInPayload) (*generatedAuth.ProfileAndCookie,
	error) {
	userSignIn := models.SignInPayload{
		Login:    in.Login,
		Password: in.Password,
	}

	profile, token, expires, err := h.uc.SignIn(ctx, &userSignIn)
	if err != nil {
		return &generatedAuth.ProfileAndCookie{}, err //TODO: add err in model
	}

	return &generatedAuth.ProfileAndCookie{
		ProfileInfo: &generatedAuth.Profile{
			Id:          profile.Id.String(),
			Login:       profile.Login,
			Description: profile.Description,
			ImgSrc:      profile.ImgSrc,
			Phone:       profile.Phone,
		},
		CookieInfo: &generatedAuth.Cookie{
			Token:   token,
			Expires: expires.String(),
		},
	}, nil
}

func (h GrpcAuthHandler) SignUp(ctx context.Context, in *generatedAuth.SignUpPayload) (*generatedAuth.ProfileAndCookie,
	error) {
	userSignUp := models.SignUpPayload{
		Login:    in.Login,
		Password: in.Password,
		Phone:    in.Phone,
	}

	profile, token, expires, err := h.uc.SignUp(ctx, &userSignUp)
	if err != nil {
		return &generatedAuth.ProfileAndCookie{}, err
	}

	return &generatedAuth.ProfileAndCookie{
		ProfileInfo: &generatedAuth.Profile{
			Id:          profile.Id.String(),
			Login:       profile.Login,
			Description: profile.Description,
			ImgSrc:      profile.ImgSrc,
			Phone:       profile.Phone,
		},
		CookieInfo: &generatedAuth.Cookie{
			Token:   token,
			Expires: expires.String(),
		},
	}, nil
}

func (h GrpcAuthHandler) CheckAuth(ctx context.Context, in *generatedAuth.UserID) (*generatedAuth.Profile,
	error) {
	userID, err := uuid.FromString(in.ID)
	if err != nil {
		return &generatedAuth.Profile{}, err
	}

	profile, err := h.uc.CheckAuth(ctx, userID)
	if err != nil {
		return &generatedAuth.Profile{}, err
	}

	return &generatedAuth.Profile{
		Id:          profile.Id.String(),
		Login:       profile.Login,
		Description: profile.Description,
		ImgSrc:      profile.ImgSrc,
		Phone:       profile.Phone,
	}, nil
}
