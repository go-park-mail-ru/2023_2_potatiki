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

	authHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/config"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
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

	log := logger.Set(cfg.Enviroment)

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
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db) // Обернуть

	if err := db.Ping(); err != nil {
		slog.Error("fail ping postgres", sl.Err(err))
		return err
	}

	authRepo := authRepo.NewAuthRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", authHandler.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/logout", authHandler.LogOut).Methods(http.MethodGet, http.MethodOptions)
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
		//ReadHeaderTimeout:
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
