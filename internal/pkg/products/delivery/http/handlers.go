package http

import (
	"errors"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/repo"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/response"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	usecase products.ProductsUsecase
}

func NewProductsHandler(usecase products.ProductsUsecase) ProductHandler {
	return ProductHandler{
		usecase: usecase,
	}
}

func (h *ProductHandler) Product(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		//log.Info("id is empty")
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		//log.Info("id is nvalid")
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}

	product, err := h.usecase.GetProduct(r.Context(), id)
	if errors.Is(err, repo.ErrPoductNotFound) {
		//log.Info("product not found", "id", id)

		resp.JSON(w, http.StatusBadRequest, resp.Err("not found"))

		return
	}
	if err != nil {
		//log.Error("failed to get product", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}

	//log.Info("got product", slog.String("product", product.Name))

	resp.JSON(w, http.StatusOK, product)
}

func (h *ProductHandler) Products(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok || idStr == "" {
		//log.Info("id is empty")
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}
	id, err := uuid.Parse(idStr)
	if err != nil {
		//log.Info("id is nvalid")
		resp.JSON(w, http.StatusAccepted, resp.Err("invalid request"))
		return
	}

	product, err := h.usecase.GetProduct(r.Context(), id)
	if errors.Is(err, repo.ErrPoductNotFound) {
		//log.Info("product not found", "id", id)

		resp.JSON(w, http.StatusBadRequest, resp.Err("not found"))

		return
	}
	if err != nil {
		//log.Error("failed to get product", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("internal error"))
		return
	}

	//log.Info("got product", slog.String("product", product.Name))

	resp.JSON(w, http.StatusOK, product)
}
