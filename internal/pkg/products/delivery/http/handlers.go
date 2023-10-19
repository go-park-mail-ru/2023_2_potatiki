package http

import (
	"net/http"
	"strconv"

	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	log *slog.Logger
	uc  products.ProductsUsecase
}

func NewProductsHandler(log *slog.Logger, uc products.ProductsUsecase) ProductHandler {
	return ProductHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	Product
// @Tags Products
// @Description	Get product
// @Accept json
// @Produce json
// @Param id path string true "Product UUID"
// @Success	200	{object} models.Product "Product info"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/products/{id} [get]
func (h *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))

		return
	}

	product, err := h.uc.GetProduct(r.Context(), id)
	if err != nil { // TODO: check not found
		h.log.Error("failed to get product", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
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
// @Success	200	{object} []models.Product "Product info"
// @Failure	400	{object} response.Response	"invalid request"
// @Failure	429
// @Router	/api/products/get_all [get]
func (h *ProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
	)
	// count - обязателен
	// paging - ситуативно(тот же offset)

	var (
		paging int64
		count  int64
		err    error
	)
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

	products, err := h.uc.GetProducts(r.Context(), paging, count)
	if err != nil {
		h.log.Error("failed to get products", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("got products", "len", len(products))
	resp.JSON(w, http.StatusOK, products)
}
