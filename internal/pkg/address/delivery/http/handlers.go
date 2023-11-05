package http

import (
	"encoding/json"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
)

type AddressHandler struct {
	log *slog.Logger
	uc  address.AddressUsecase
}

func NewAddressHandler(log *slog.Logger, uc address.AddressUsecase) AddressHandler {
	return AddressHandler{
		log: log,
		uc:  uc,
	}
}

// @Summary	CreateOrder
// @Tags Order
// @Description	Create Order using profile ID from cookies
// @Accept json
// @Produce json
// @Success	200	{object} models.Order "New order info"
// @Failure	401	"User unauthorized"
// @Failure	404	"Cart not found"
// @Failure	429
// @Router	/api/order/create [post]

func (h *AddressHandler) AddAddress(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	defer r.Body.Close()
	h.log.Debug("request body decoded", slog.Any("request", r))

	addressInfo := models.AddressInfo{}
	err = json.Unmarshal(body, &addressInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	address, err := h.uc.AddAddress(r.Context(), userID, addressInfo)
	if err != nil {
		h.log.Error("failed to add address", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.AddAddress", "address", address)
	resp.JSON(w, http.StatusOK, address)
}

func (h *AddressHandler) UpdateAddress(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	defer r.Body.Close()
	h.log.Debug("request body decoded", slog.Any("request", r))

	adressInfo := models.Address{}
	err = json.Unmarshal(body, &adressInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	adressInfo.ProfileId = userID
	address, err := h.uc.UpdateAddress(r.Context(), adressInfo)
	if err != nil {
		h.log.Error("failed to update address", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.UpdateAddress", "address", address)
	resp.JSON(w, http.StatusOK, address)
}

func (h *AddressHandler) DeleteAddress(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	defer r.Body.Close()
	h.log.Debug("request body decoded", slog.Any("request", r))

	adressInfo := models.AddressDelete{}
	err = json.Unmarshal(body, &adressInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	adressInfo.ProfileId = userID
	err = h.uc.DeleteAddress(r.Context(), adressInfo)
	if err != nil {
		h.log.Error("failed to update address", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.UpdateAddress")
	resp.JSONStatus(w, http.StatusOK)
}

func (h *AddressHandler) MakeCurrentAddress(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)
	userID, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	defer r.Body.Close()
	h.log.Debug("request body decoded", slog.Any("request", r))

	adressInfo := models.AddressMakeCurrent{}
	err = json.Unmarshal(body, &adressInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	adressInfo.ProfileId = userID
	err = h.uc.MakeCurrentAddress(r.Context(), adressInfo)
	if err != nil {
		h.log.Error("failed to make current address", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.MakeCurrentAddress")
	resp.JSONStatus(w, http.StatusOK)
}
