package http

import (
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/search"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/search/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	"log/slog"
	"net/http"
)

type SearchHandler struct {
	log *slog.Logger
	uc  search.SearchUsecase
}

func NewSearchHandler(log *slog.Logger, uc search.SearchUsecase) *SearchHandler {
	return &SearchHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	SearchProducts
// @Tags Search
// @Description	Search products by name
// @Accept json
// @Produce json
// @Param product query int true "Search products by name"
// @Success	200	{object} []models.Product "Products array"
// @Failure	400	{object} responser.Response	"error message"
// @Failure	429
// @Router	/api/search/ [get]
func (h *SearchHandler) SearchProducts(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	productName := r.URL.Query().Get("product")

	products, err := h.uc.SearchProducts(r.Context(), productName)
	if err != nil {
		h.log.Error("failed to get products", sl.Err(err))
		if errors.Is(err, repo.ErrProductNotFound) {
			resp.JSON(w, http.StatusBadRequest, resp.Err("products not found"))

			return
		}
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("got products", "len", len(products))
	resp.JSON(w, http.StatusOK, products)
}
