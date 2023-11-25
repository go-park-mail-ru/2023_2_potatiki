package grpc

import (
	"log/slog"

	generatedAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/generated"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"google.golang.org/grpc"
)

//go:generate mockgen -source=./generated/products_grpc.pb.go -destination=../../mocks/products_grpc.go -package=mock

type ProductHandler struct {
	log *slog.Logger
	uc  products.ProductsUsecase

	generatedAuth.AuthServiceServer
}

func NewProductsHandler(log *slog.Logger, uc products.ProductsUsecase) ProductHandler {
	return ProductHandler{
		log: log,
		uc:  uc,
	}
}

func Register(gRPCServer *grpc.Server, uc products.ProductsUsecase) {
	//ssov1.RegisterAuthServer(gRPCServer, &serverAPI{uc: uc})
}

type serverAPI struct {
	//ssov1.UnimplementedAuthServer
	uc products.ProductsUsecase
}
