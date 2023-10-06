package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	authHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/http"
	authRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/repo"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/config"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware"
	productsHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/http"
	productsRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/repo"
	productsUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/usecase"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"

	_ "github.com/lib/pq"

	_ "github.com/go-park-mail-ru/2023_2_potatiki/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title ZuZu Backend API
// @description API server for ZuZu.

// @contact.name Dima
// @contact.url http://t.me/belozerovmsk

// @host localhost:8082
// @BasePath /

// @securityDefinitions	AuthKey
// @in					header
// @name				Authorization
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	cfg := config.MustLoad() // TODO : dev-config.yaml -> readme

	log := logger.Set(cfg.Enviroment)
	log.Info(
		"starting zuzu-main",
		slog.String("env", cfg.Enviroment),
	)
	log.Debug("debug messages are enabled")

	//============================Database============================//
	//docker run --name zuzu-postgres -v zuzu-db-data:/var/lib/postgresql/data -v -e 'PGDATA:/var/lib/postgresql/data/pgdata' './build/sql/injection_db.sql:/docker-entrypoint-initdb.d/init.sql' -p 8079:5432 --env-file .env --restart always postgres:latest
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName))
	if err != nil {
		log.Error("fail open postgres", sl.Err(err))
		return err
	}
	defer func(db *sql.DB) {
		err = errors.Join(err, db.Close())
	}(db)

	if err = db.Ping(); err != nil {
		slog.Error("fail ping postgres", sl.Err(err))
		return err
	}
	//----------------------------Database----------------------------//
	//
	//
	//============================Init layers============================//
	authRepo := authRepo.NewAuthRepo(db)
	authUsecase := authUsecase.NewAuthUsecase(authRepo, cfg.Auther)
	authHandler := authHandler.NewAuthHandler(log, authUsecase)

	productsRepo := productsRepo.NewProductsRepo(db)
	productsUsecase := productsUsecase.NewProductsUsecase(productsRepo)
	productsHandler := productsHandler.NewProductsHandler(log, productsUsecase)
	//----------------------------Init layers----------------------------//
	//
	//
	//============================Create router============================//

	r := mux.NewRouter().PathPrefix("/api").Subrouter()

	r.Use(middleware.CORSMiddleware)

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	//----------------------------Create router----------------------------//
	//
	//
	//============================Setup endpoints============================//
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", authHandler.SignUp).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/signin", authHandler.SignIn).Methods(http.MethodPost, http.MethodOptions)
		auth.HandleFunc("/logout", authHandler.LogOut).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/check_auth", authHandler.CheckAuth).Methods(http.MethodGet, http.MethodOptions)
		auth.HandleFunc("/{id:[0-9a-fA-F-]+}", authHandler.GetProfile).Methods(http.MethodGet, http.MethodOptions)
	}

	products := r.PathPrefix("/products").Subrouter()
	{
		products.HandleFunc("/{id:[0-9a-fA-F-]+}", productsHandler.Product).Methods(http.MethodGet, http.MethodOptions)
		products.HandleFunc("/get_all", productsHandler.Products).Methods(http.MethodGet, http.MethodOptions)
	}
	//----------------------------Setup endpoints----------------------------//

	http.Handle("/", r)

	srv := http.Server{
		Handler:           r,
		Addr:              cfg.Address,
		ReadTimeout:       cfg.Timeout,
		WriteTimeout:      cfg.Timeout,
		IdleTimeout:       cfg.IdleTimeout,
		ReadHeaderTimeout: cfg.ReadHeaderTimeout,
	}

	quit := make(chan os.Signal, 1)
	//SIGINT = ctrl+c; SIGTERM = kill; Interrupt = аппаратное прерывание, в Windows даст ошибку
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("listen and serve returned err: ", sl.Err(err))
		}
	}()

	log.Info("server started")
	sig := <-quit
	log.Debug("handle quit chanel: ", slog.Any("os.Signal", sig.String()))
	log.Info("server stopping...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err = srv.Shutdown(ctx); err != nil {
		log.Error("server shutdown returned an err: ", sl.Err(err))
		return err
	}

	log.Info("server stopped")
	return nil
}
