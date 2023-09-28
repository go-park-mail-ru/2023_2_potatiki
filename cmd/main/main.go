package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"log/slog"

	"github.com/gorilla/mux"

	authHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
	productsHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/http"
	productsRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/repo"
	productsUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/usecase"
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

func run() (err error) {
	cfg := config.MustLoad() // TODO : dev-config.yaml -> readme

	log := logger.Set(cfg.Enviroment)

	log.Info(
		"starting zuzu",
		slog.String("env", cfg.Enviroment),
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
		err = errors.Join(err, db.Close())
	}(db)

	if err := db.Ping(); err != nil {
		slog.Error("fail ping postgres", sl.Err(err))
		return err
	}

	authRepo := authRepo.NewAuthRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo)
	authHandler := authHandler.NewAuthHandler(authUsecase)

	productsRepo := productsRepo.NewProductsRepo(db)
	productsUsecase := productsUsecase.NewProductsUsecase(productsRepo)
	productsHandler := productsHandler.NewProductsHandler(productsUsecase)

	r := mux.NewRouter().PathPrefix("/api").Subrouter()
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", authHandler.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/logout", authHandler.LogOut).Methods(http.MethodGet, http.MethodOptions)
	}

	products := r.PathPrefix("/products").Subrouter()
	{
		products.HandleFunc("/{id:[0-9]+}", productsHandler.Product).Methods(http.MethodGet, http.MethodOptions)
		products.HandleFunc("/", productsHandler.Products).Methods(http.MethodGet, http.MethodOptions)
	}

	//user := r.PathPrefix("/user").Subrouter()
	{

	}

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	http.Handle("/", r)

	srv := http.Server{
		Handler:           r,
		Addr:              cfg.Address,
		ReadTimeout:       cfg.Timeout,
		WriteTimeout:      cfg.Timeout,
		IdleTimeout:       cfg.IdleTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen and serve returned err: ", sl.Err(err))
		}
	}()

	log.Info("server started")
	sig := <-quit // wait for interrupt signal
	log.Debug("handle quit os/signal: ", sig)
	log.Info("server stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("server shutdown returned an err: ", sl.Err(err))
		return err
	}

	log.Info("server stopped")
	return nil
}
