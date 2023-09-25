package main

import (
	"database/sql"
	"net/http"
	"os"

	"log/slog"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/config"
	authHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {

}

func run() error {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info(
		"starting url-shortener",
		slog.String("env", cfg.Env),
		slog.String("version", "123"),
	)
	log.Debug("debug messages are enabled")

	db, err := sql.Open("postgres", cfg.Database.URL)
	if err != nil {
		slog.Error("sdsd", err)
		return err
	}
	defer db.Close()

	authRepo := authRepo.New(db)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signUp", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
	}

	http.Handle("/", r)
	srv := http.Server{Handler: r, Addr: ":8000"}
	return srv.ListenAndServe()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stderr, nil))
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
