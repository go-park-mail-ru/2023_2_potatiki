package http

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	grpcAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/gen"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	resp "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/responser"
	uuid "github.com/satori/go.uuid"
)

type AuthHandler struct {
	client grpcAuth.AuthClient
	log    *slog.Logger
	uc     auth.AuthUsecase
}

const customTimeFormat = "2006-01-02 15:04:05.999999999 -0700 UTC"

func NewAuthHandler(cl grpcAuth.AuthClient, log *slog.Logger, uc auth.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		client: cl,
		log:    log,
		uc:     uc,
	}
}

// @Summary	SignIn
// @Tags Auth
// @Description	Login to Account
// @Accept json
// @Produce json
// @Param input body models.SignInPayload true "SignInPayload"
// @Success	200	{object} models.Profile "Profile"
// @Failure	400	{object} responser.Response	"error messege"
// @Failure	429
// @Router	/api/auth/signin [post]
func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))
	defer r.Body.Close()

	userInfo := &models.SignInPayload{}
	err = json.Unmarshal(body, userInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	profileAndCookie, err := h.client.SignIn(r.Context(), &grpcAuth.SignInRequest{
		Login:    userInfo.Login,
		Password: userInfo.Password,
	})
	if err != nil {
		h.log.Error("failed to signin", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))

		return
	}

	h.log.Debug("got profile", slog.Any("profile", profileAndCookie.Profile.Id))

	expiresTime, err := time.Parse(customTimeFormat, profileAndCookie.Expires)
	if err != nil {
		h.log.Error("failed to parse time from auth signin", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	http.SetCookie(w, authmw.MakeTokenCookie(profileAndCookie.Token, expiresTime))
	resp.JSON(w, http.StatusOK, profileAndCookie.Profile)
}

// @Summary	SignUp
// @Tags Auth
// @Description	Create Account
// @Accept json
// @Produce json
// @Param input body models.SignUpPayload true "SignUpPayload"
// @Success	200 {object} models.Profile "Profile"
// @Failure	400	{object} responser.Response	"error messege"
// @Failure	429
// @Router	/api/auth/signup [post]
func (h *AuthHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	body, err := io.ReadAll(r.Body)
	if resp.BodyErr(err, h.log, w) {
		return
	}
	h.log.Debug("request body decoded", slog.Any("request", r))

	userInfo := &models.SignUpPayload{}
	err = json.Unmarshal(body, userInfo)
	if err != nil {
		h.log.Error("failed to unmarshal request body", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	profileAndCookie, err := h.client.SignUp(r.Context(), &grpcAuth.SignUpRequest{
		Login:    userInfo.Login,
		Password: userInfo.Password,
		Phone:    userInfo.Phone,
	})
	if err != nil {
		h.log.Error("failed to signup", sl.Err(err))
		resp.JSON(w, http.StatusBadRequest, resp.Err("invalid login or password"))

		return
	}

	h.log.Debug("got profile", slog.Any("profile", profileAndCookie.Profile.Id))

	expiresTime, err := time.Parse(customTimeFormat, profileAndCookie.Expires)
	if err != nil {
		h.log.Error("failed to parse time from auth signup", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	http.SetCookie(w, authmw.MakeTokenCookie(profileAndCookie.Token, expiresTime))
	resp.JSON(w, http.StatusOK, profileAndCookie.Profile)
}

// @Summary	Logout
// @Tags Auth
// @Description	Logout from Account
// @Accept json
// @Produce json
// @Success	200
// @Failure	401
// @Router	/api/auth/logout [get]
func (h *AuthHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, authmw.MakeTokenCookie("", time.Now().UTC().AddDate(0, 0, -1)))
	h.log.Info("logout")
	resp.JSONStatus(w, http.StatusOK)
}

// @Summary	CheckAuth
// @Tags Auth
// @Description	Check user is logged in
// @Accept json
// @Produce json
// @Success	200	{object} models.Profile "Profile"
// @Failure	401
// @Failure	429
// @security AuthKey
// @Param Cookie header string  false "Token" default(zuzu-t=xxx)
// @Router	/api/auth/check_auth [get]
func (h *AuthHandler) CheckAuth(w http.ResponseWriter, r *http.Request) {
	h.log = h.log.With(
		slog.String("op", sl.GFN()),
		slog.String("request_id", r.Header.Get(logmw.RequestIDCtx)),
	)

	id, ok := r.Context().Value(authmw.AccessTokenCookieName).(uuid.UUID)
	if !ok {
		h.log.Error("failed cast uuid from context value")
		resp.JSONStatus(w, http.StatusUnauthorized)

		return
	}

	h.log.Debug("check auth success", "id", id)

	profile, err := h.client.CheckAuth(r.Context(), &grpcAuth.CheckAuthRequst{
		ID: id.String(),
	})
	if err != nil {
		h.log.Error("failed to CheckAuth", sl.Err(err))
		resp.JSONStatus(w, http.StatusTooManyRequests)

		return
	}

	h.log.Debug("got profile", slog.Any("profile", profile.Profile.Id))
	resp.JSON(w, http.StatusOK, profile)
}
