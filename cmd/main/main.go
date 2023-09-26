package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"

	"log/slog"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/config"
	authHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/handlers/slogpretty"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
)

func main() {
	if err := run(); err != nil {
		//slog.Error("main error:", sl.Err(err))
		os.Exit(1)
	}
}

func run() error {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Enviroment)

	log.Info(
		"starting zuzu",
		slog.String("env", cfg.Enviroment),
		slog.String("version", cfg.Version),
	)
	log.Debug("debug messages are enabled")

	psqlInfo := fmt.Sprintf("port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Error("fail open postgres", sl.Err(err))
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		slog.Error("fail ping postgres", sl.Err(err))
		return err
	}

	authRepo := authRepo.New(db)
	authUsecase := authUsecase.New(authRepo)
	authHandler := authHandler.New(authUsecase)

	r := mux.NewRouter().PathPrefix("/api/v" + cfg.Version).Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signUp", authHandler.SignUp).Methods(http.MethodPost)
		auth.HandleFunc("/signIn", authHandler.SignIn).Methods(http.MethodPost)
		auth.HandleFunc("/logOut", authHandler.LogOut).Methods(http.MethodGet)
	}

	//user := r.PathPrefix("/user").Subrouter()
	{

	}

	//products := r.PathPrefix("/products").Subrouter()
	{

	}

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	http.Handle("/", r)

	srv := http.Server{
		Handler:      r,
		Addr:         cfg.Address,
		ReadTimeout:  cfg.Timeout,
		WriteTimeout: cfg.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	sigs := make(chan os.Signal, 1)
	done := make(chan struct{})
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		if err := srv.Shutdown(context.TODO()); err != nil {
			log.Error("server shutdown returned an err: %v\n", sl.Err(err))
		}
		close(done)
	}()

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Error("listen and serve returned err: %v", sl.Err(err))
		return err
	}
	<-done
	log.Info("server stop")
	return nil
}

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
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

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}
