package http

import (
	"net/http"
	"strconv"

	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/grpc/gen"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
)

type ProductHandler struct {
	client gen.ProductsClient
	log    *slog.Logger
}

func NewProductsHandler(client gen.ProductsClient, log *slog.Logger) ProductHandler {
	return ProductHandler{
		client: client,
		log:    log,
	}
}

// @Summary	Product
// @Tags Products
// @Description	Get product
// @Accept json
// @Produce json
// @Param id path string true "Product UUID"
// @Success	200	{object} models.Product "Product info"
// @Failure	400	{object} responser.Response	"error message"
// @Failure	429
// @Router	/api/products/{id} [get]
func (h *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	id, err := uuid.FromString(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	gproduct, err := h.client.GetProduct(r.Context(),
		&gen.ProductRequest{Id: id.String()})
	if err != nil { // TODO: check not found
		h.log.Error("failed to get product", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}
	productID, err := uuid.FromString(gproduct.Product.Id)
	if err != nil {
		h.log.Error("failed to cast product ID", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)
		return
	}

	product := models.Product{
		Id:          productID,
		Name:        gproduct.Product.Name,
		Description: gproduct.Product.Description,
		Price:       gproduct.Product.Price,
		ImgSrc:      gproduct.Product.ImgSrc,
		Rating:      gproduct.Product.Rating,
		Category: models.Category{
			Id:     gproduct.Product.Category.Id,
			Name:   gproduct.Product.Category.Name,
			Parent: gproduct.Product.Category.Parent,
		},
	}
	h.log.Debug("got product", slog.String("product", product.Name))
	resp.JSON(w, http.StatusOK, product)
}

// @Summary	Products
// @Tags Products
// @Description	Get products
// @Accept json
// @Produce json
// @Param paging query int false "Skip number of products"
// @Param count query int true "Display number of products"
// @Param ratingBy query string false "(Use ASC/DESC) Sort products by rating"
// @Param priceBy query string false "(Use ASC/DESC) Sort products by price"
// @Success	200	{object} []models.Product "Products array"
// @Failure	400	{object} responser.Response	"error message"
// @Failure	429
// @Router	/api/products/get_all [get]
func (h *ProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	// count - обязателен
	// paging - ситуативно(тот же offset)

	var (
		paging int64
		count  int64
		err    error
	)
	pagingStr := r.URL.Query().Get("paging")
	if len(pagingStr) != 0 {
		paging, err = strconv.ParseInt(pagingStr, 10, 64)
		if err != nil {
			h.log.Error("paging is invalid", sl.Err(err))
			resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

			return
		}
	}

	countStr := r.URL.Query().Get("count")
	count, err = strconv.ParseInt(countStr, 10, 64)
	if err != nil {
		h.log.Error("count is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	//ratingBy := r.URL.Query().Get("ratingBy")

	//sortingBy := r.URL.Query().Get("priceBy")

	//products, err := h.uc.GetProducts(r.Context(), paging, count, ratingBy, sortingBy)

	response, err := h.client.GetProducts(r.Context(),
		&gen.ProductsRequest{Paging: paging, Count: count})
	if err != nil {
		h.log.Error("failed to get products", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)
		return
	}
	gproducts := response.Products

	products := make([]models.Product, len(gproducts))
	for i, gproduct := range gproducts {
		productID, err := uuid.FromString(gproduct.Id)
		if err != nil {
			h.log.Error("failed to cast product ID", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
			return
		}

		product := models.Product{
			Id:          productID,
			Name:        gproduct.Name,
			Description: gproduct.Description,
			Price:       gproduct.Price,
			ImgSrc:      gproduct.ImgSrc,
			Rating:      gproduct.Rating,
			Category: models.Category{
				Id:     gproduct.Category.Id,
				Name:   gproduct.Category.Name,
				Parent: gproduct.Category.Parent,
			},
		}
		products[i] = product
	}

	h.log.Debug("got products", "len", len(products))
	resp.JSON(w, http.StatusOK, products)
}

// @Summary	Products
// @Tags Products
// @Description	Get products by category
// @Accept json
// @Produce json
// @Param category_id query int true "Category ID"
// @Param paging query int false "Skip number of products"
// @Param count query int true "Display number of products"
// @Param ratingBy query string false "(Use ASC/DESC) Sort products by rating"
// @Param priceBy query string false "(Use ASC/DESC) Sort products by price"
// @Success	200	{object} []models.Product "Products by category id"
// @Failure	400	{object} responser.Response	"error message"
// @Failure	429
// @Router	/api/products/category [get]
func (h *ProductHandler) Category(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	// count - обязателен
	// paging - ситуативно(тот же offset)

	var (
		paging int64
		count  int64
		err    error
	)
	idStr := r.URL.Query().Get("category_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	pagingStr := r.URL.Query().Get("paging")
	if pagingStr != "" {
		paging, err = strconv.ParseInt(pagingStr, 10, 64)
		if err != nil {
			h.log.Error("paging is invalid", sl.Err(err))
			resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

			return
		}
	}

	countStr := r.URL.Query().Get("count")
	count, err = strconv.ParseInt(countStr, 10, 64)
	if err != nil {
		h.log.Error("count is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	//ratingBy := r.URL.Query().Get("ratingBy")

	//sortingBy := r.URL.Query().Get("priceBy")

	//products, err := h.uc.GetCategory(r.Context(), id, paging, count, ratingBy, sortingBy)
	response, err := h.client.GetCategory(r.Context(),
		&gen.CategoryRequest{Id: int32(id), Paging: paging, Count: count})
	if err != nil {
		h.log.Error("failed to get products by category", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	gproducts := response.Products

	products := make([]models.Product, len(gproducts))
	for i, gproduct := range gproducts {
		productID, err := uuid.FromString(gproduct.Id)
		if err != nil {
			h.log.Error("failed to cast product ID", sl.Err(err))
			resp.JSONStatus(w, http.StatusTooManyRequests)
			return
		}

		product := models.Product{
			Id:          productID,
			Name:        gproduct.Name,
			Description: gproduct.Description,
			Price:       gproduct.Price,
			ImgSrc:      gproduct.ImgSrc,
			Rating:      gproduct.Rating,
			Category: models.Category{
				Id:     gproduct.Category.Id,
				Name:   gproduct.Category.Name,
				Parent: gproduct.Category.Parent,
			},
		}
		products[i] = product
	}

	h.log.Debug("got products by category", "len", len(products))
	resp.JSON(w, http.StatusOK, products)
}
