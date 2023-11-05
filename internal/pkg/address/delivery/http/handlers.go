package http

import (
	"encoding/json"
	"errors"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"
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

// @Summary	AddAddress
// @Tags Address
// @Description	Add Address
// @Accept json
// @Produce json
// @Param input body models.AddressInfo true "Address info"
// @Success	200	{object} models.Address "Address full info"
// @Failure	401	"User unauthorized"
// @Failure	429
// @Router	/api/address/add [post]
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

// @Summary	UpdateAddress
// @Tags Address
// @Description	Update Address
// @Accept json
// @Produce json
// @Param input body models.Address true "Address info"
// @Success	200	{object} models.Address "Address info"
// @Failure	401	"User unauthorized"
// @Failure	404	"Address not found"
// @Failure	429
// @Router	/api/address/update [post]
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
		if errors.Is(err, repo.ErrAddressNotFound) {
			resp.JSONStatus(w, http.StatusNotFound)

			return
		}
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.UpdateAddress", "address", address)
	resp.JSON(w, http.StatusOK, address)
}

// @Summary	DeleteAddress
// @Tags Address
// @Description	Delete Address
// @Accept json
// @Produce json
// @Param input body models.AddressDelete true "Address ID"
// @Success	200	"Address deleted info"
// @Failure	401	"User unauthorized"
// @Failure	429
// @Router	/api/address/delete [delete]
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
		h.log.Error("failed to delete address", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.DeleteAddress")
	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	MakeCurrentAddress
// @Tags Address
// @Description	Make Current Address
// @Accept json
// @Produce json
// @Param input body models.AddressMakeCurrent true "Address ID"
// @Success	200	"Address is made current"
// @Failure	401	"User unauthorized"
// @Failure	404	"Address not found"
// @Failure	429
// @Router	/api/address/make_current [post]
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
		if errors.Is(err, repo.ErrCurrentAddressNotFound) {
			resp.JSONStatus(w, http.StatusNotFound)

			return
		}
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.MakeCurrentAddress")
	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	GetCurrentAddress
// @Tags Address
// @Description	Get Current Address
// @Accept json
// @Produce json
// @Success	200	{object} models.Address "Address full info"
// @Failure	401	"User unauthorized"
// @Failure	404	"Address not found"
// @Failure	429
// @Router	/api/address/get_current [get]
func (h *AddressHandler) GetCurrentAddress(w http.ResponseWriter, r *http.Request) {
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

	address, err := h.uc.GetCurrentAddress(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get address", sl.Err(err))
		if errors.Is(err, repo.ErrAddressNotFound) {
			resp.JSONStatus(w, http.StatusNotFound)

			return
		}
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetCurrentAddress", "address", address)
	resp.JSON(w, http.StatusOK, address)
}

// @Summary	GetAllAddresses
// @Tags Address
// @Description	Get All Addresses
// @Accept json
// @Produce json
// @Success	200	{array} models.Address "Address full info"
// @Failure	401	"User unauthorized"
// @Failure	404	"Addresses not found"
// @Failure	429
// @Router	/api/address/get_all [get]
func (h *AddressHandler) GetAllAddresses(w http.ResponseWriter, r *http.Request) {
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

	addresses, err := h.uc.GetAllAddresses(r.Context(), userID)
	if err != nil {
		h.log.Error("failed to get addresses", sl.Err(err))
		if errors.Is(err, repo.ErrAddressesNotFound) {
			resp.JSONStatus(w, http.StatusNotFound)

			return
		}
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("h.uc.GetAllAddresses", "addresses", addresses)
	resp.JSON(w, http.StatusOK, addresses)
}
