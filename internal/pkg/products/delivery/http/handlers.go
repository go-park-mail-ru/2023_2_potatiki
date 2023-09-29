package http

import (
	"errors"
	"net/http"
	"strconv"

	"log/slog"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/repo"
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

func (h *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		h.log.Error("id is empty")
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		h.log.Error("id is invalid", sl.Err(err))
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}

	product, err := h.uc.GetProduct(r.Context(), id)
	if errors.Is(err, repo.ErrPoductNotFound) {
		h.log.Error("product not found", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("product not found"))

		return
	}
	if err != nil {
		h.log.Error("failed to get product", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}

	h.log.Debug("got product", slog.String("product", product.Name))
	resp.JSON(w, http.StatusOK, product)
}

func (h *ProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	var (
		offset int64
		page   int64
	)

	offsetStr := r.URL.Query().Get("offset")
	if offsetStr == "" {
		offset = 0
	}
	offset, err := strconv.ParseInt(offsetStr, 10, 64)
	if err != nil {
		h.log.Error("offset is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	pageStr := r.URL.Query().Get("page")
	if pageStr == "" {
		page = 0
	}
	page, err = strconv.ParseInt(pageStr, 10, 64)
	if err != nil {
		h.log.Error("page is invalid", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid request"))
		return
	}

	products, err := h.uc.GetProducts(r.Context(), offset, page)
	if err != nil {
		h.log.Error("failed to get products", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}

	h.log.Debug("got products", "len", len(products))
	resp.JSON(w, http.StatusOK, products)
}
