package grpc

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	generatedAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	uuid "github.com/satori/go.uuid"
	"log/slog"
)

//go:generate mockgen -source=./generated/auth_grpc.pb.go -destination=../../mocks/auth_grpc.go -package=mock

type GrpcAuthHandler struct {
	uc  auth.AuthUsecase
	log *slog.Logger

	generatedAuth.AuthServiceServer
}

func NewGrpcAuthHandler(uc auth.AuthUsecase, log *slog.Logger) *GrpcAuthHandler {
	return &GrpcAuthHandler{
		uc:  uc,
		log: log,
	}
}

func (h GrpcAuthHandler) SignIn(ctx context.Context, in *generatedAuth.SignInPayload) (*generatedAuth.ProfileAndCookie,
	error) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	userSignIn := models.SignInPayload{
		Login:    in.Login,
		Password: in.Password,
	}

	profile, token, expires, err := h.uc.SignIn(ctx, &userSignIn)
	if err != nil {
		h.log.Error("failed in uc.SignIn", sl.Err(err))

		return &generatedAuth.ProfileAndCookie{Error: err.Error()}, err
	}
	h.log.Debug("got profile", slog.Any("profile", profile.Id))

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
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	userSignUp := models.SignUpPayload{
		Login:    in.Login,
		Password: in.Password,
		Phone:    in.Phone,
	}

	profile, token, expires, err := h.uc.SignUp(ctx, &userSignUp)
	if err != nil {
		h.log.Error("failed in uc.SignUp", sl.Err(err))

		return &generatedAuth.ProfileAndCookie{Error: err.Error()}, err
	}
	h.log.Debug("got profile", slog.Any("profile", profile.Id))

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
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	userID, err := uuid.FromString(in.ID)
	if err != nil {
		h.log.Error("failed to get uuid from string", sl.Err(err))

		return &generatedAuth.Profile{Error: err.Error()}, err
	}

	profile, err := h.uc.CheckAuth(ctx, userID)
	if err != nil {
		h.log.Error("failed in uc.CheckAuth", sl.Err(err))

		return &generatedAuth.Profile{Error: err.Error()}, err
	}
	h.log.Debug("got profile", slog.Any("profile", profile.Id))

	return &generatedAuth.Profile{
		Id:          profile.Id.String(),
		Login:       profile.Login,
		Description: profile.Description,
		ImgSrc:      profile.ImgSrc,
		Phone:       profile.Phone,
	}, nil
}
