package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	generatedAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/generated"
)

//go:generate mockgen -source=./generated/auth_grpc.pb.go -destination=../../mocks/auth_grpc.go -package=mock

type GrpcAuthHandler struct {
	uc auth.AuthUsecase

	generatedAuth.AuthServiceServer
}

func NewGrpcAuthHandler(uc auth.AuthUsecase) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		uc: uc,
	}
}

func (h GrpcAuthHandler) SignIn(ctx context.Context, in *generatedAuth.SignInPayload) (*generatedAuth.ProfileAndCookie, error) {
	userSignIn := models.SignInPayload{
		Login:    in.Login,
		Password: in.Password,
	}

	profile, token, expires, err := h.uc.SignIn(ctx, &userSignIn)
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
