package grpc

import (
	"context"
	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2023_2_potatiki/proto/gmodels"
	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (h serverAPI) GetProduct(ctx context.Context,
	in *gen.ProductRequest) (*gen.ProductResponse, error) {
	id, err := uuid.FromString(in.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid ID, fail to cast uuid")
	}

	product, err := h.uc.GetProduct(ctx, id)
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get product")
	}

	gproduct := &gmodels.Product{
		Id:          product.Id.String(),
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		ImgSrc:      product.ImgSrc,
		Rating:      product.Rating,
		Category: &gmodels.Category{
			Id:     product.Category.Id,
			Name:   product.Category.Name,
			Parent: product.Parent,
		},
	}

	return &gen.ProductResponse{Product: gproduct}, nil
}

func (h serverAPI) GetProducts(ctx context.Context,
	in *gen.ProductsRequest) (*gen.ProductsResponse, error) {

	products, err := h.uc.GetProducts(ctx, in.Paging, in.Count, "", "")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get products")
	}

	gproducts := make([]*gmodels.Product, len(products))
	for i, product := range products {
		gproduct := &gmodels.Product{
			Id:          product.Id.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ImgSrc:      product.ImgSrc,
			Rating:      product.Rating,
			Category: &gmodels.Category{
				Id:     product.Category.Id,
				Name:   product.Category.Name,
				Parent: product.Parent,
			},
		}
		gproducts[i] = gproduct
	}
	return &gen.ProductsResponse{Products: gproducts}, nil
}

func (h serverAPI) GetCategory(ctx context.Context,
	in *gen.CategoryRequest) (*gen.CategoryResponse, error) {

	products, err := h.uc.GetCategory(ctx, int(in.Id), in.Paging, in.Count, "", "")
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get products")
	}

	gproducts := make([]*gmodels.Product, len(products))
	for i, product := range products {
		gproduct := &gmodels.Product{
			Id:          product.Id.String(),
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			ImgSrc:      product.ImgSrc,
			Rating:      product.Rating,
			Category: &gmodels.Category{
				Id:     product.Category.Id,
				Name:   product.Category.Name,
				Parent: product.Parent,
			},
		}
		gproducts[i] = gproduct
	}
	return &gen.CategoryResponse{Products: gproducts}, nil
}
