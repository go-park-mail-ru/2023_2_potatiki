package grpc

import (
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
	"google.golang.org/grpc"
)

//go:generate mockgen -source=./generated/products_grpc.pb.go -destination=../../mocks/products_grpc.go -package=mock

type serverAPI struct {
	log *slog.Logger
	uc  products.ProductsUsecase

	//gen.ProductsServer
	gen.UnimplementedProductsServer
}

func Register(gRPCServer *grpc.Server, log *slog.Logger, uc products.ProductsUsecase) {
	gen.RegisterProductsServer(gRPCServer, &serverAPI{log: log, uc: uc})
}
