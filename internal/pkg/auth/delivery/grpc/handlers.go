package grpc

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
)

//go:generate mockgen -source=./generated/auth_grpc.pb.go -destination=../../mocks/auth_grpc.go -package=mock

type serverAPI struct {
	log *slog.Logger
	uc  auth.AuthUsecase

	//gen.ProductsServer
	gen.UnimplementedAuthServer
}

func Register(gRPCServer *grpc.Server, log *slog.Logger, uc auth.AuthUsecase) {
	gen.RegisterAuthServer(gRPCServer, &serverAPI{log: log, uc: uc})
}

func (h serverAPI) SignIn(ctx context.Context, in *gen.SignInRequest) (*gen.SignInResponse, error) {
	userSignIn := models.SignInPayload{
		Login:    in.Login,
		Password: in.Password,
	}

	profile, token, expires, err := h.uc.SignIn(ctx, &userSignIn)
	if err != nil {
		return &gen.SignInResponse{}, err //TODO: add err in model
	}

	return &gen.SignInResponse{
		Profile: &gmodels.Profile{
			Id:          profile.Id.String(),
			Login:       profile.Login,
			Description: profile.Description,
			ImgSrc:      profile.ImgSrc,
			Phone:       profile.Phone,
		},
		Token:   token,
		Expires: expires.String(),
	}, nil
}

func (h serverAPI) SignUp(ctx context.Context, in *gen.SignUpRequest) (*gen.SignUpResponse, error) {
	userSignUp := models.SignUpPayload{
		Login:    in.Login,
		Password: in.Password,
		Phone:    in.Phone,
	}

	profile, token, expires, err := h.uc.SignUp(ctx, &userSignUp)
	if err != nil {
		return &gen.SignUpResponse{}, err
	}

	return &gen.SignUpResponse{
		Profile: &gmodels.Profile{
			Id:          profile.Id.String(),
			Login:       profile.Login,
			Description: profile.Description,
			ImgSrc:      profile.ImgSrc,
			Phone:       profile.Phone,
		},

		Token:   token,
		Expires: expires.String(),
	}, nil
}

func (h serverAPI) CheckAuth(ctx context.Context, in *gen.CheckAuthRequst) (*gen.CheckAuthResponse, error) {
	userID, err := uuid.FromString(in.ID)
	if err != nil {
		return &gen.CheckAuthResponse{}, err
	}

	profile, err := h.uc.CheckAuth(ctx, userID)
	if err != nil {
		return &gen.CheckAuthResponse{}, err
	}

	return &gen.CheckAuthResponse{
		Profile: &gmodels.Profile{
			Id:          profile.Id.String(),
			Login:       profile.Login,
			Description: profile.Description,
			ImgSrc:      profile.ImgSrc,
			Phone:       profile.Phone,
		},
	}, nil
}
