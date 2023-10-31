package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/go-park-mail-ru/2023_2_potatiki/docs"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/config"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/authmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/middleware/logmw"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"

	authHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/http"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"

	cartHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/delivery/http"
	cartRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	cartUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/usecase"

	categoryHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category/delivery/http"
	categoryRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category/repo"
	categoryUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/category/usecase"

	productsHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/delivery/http"
	productsRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/repo"
	productsUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/products/usecase"

	profileHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/delivery/http"
	profileRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/repo"
	profileUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/usecase"

	orderHandler "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/http"
	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
	orderUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/usecase"
)

// @title ZuZu Backend API
// @description API server for ZuZu.

// @contact.name Dima
// @contact.url http://t.me/belozerovmsk

// @securityDefinitions	AuthKey
// @in					header
// @name				Authorization
func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	cfg := config.MustLoad() // TODO : dev-config.yaml -> readme.

	logFile, err := os.OpenFile(cfg.LogFilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("fail open logFile", sl.Err(err))
		err = fmt.Errorf("fail open logFile: %w", err)

		return err
	}
	defer func() { err = errors.Join(err, logFile.Close()) }()

	log := logger.Set(cfg.Enviroment, logFile)

	log.Info(
		"starting zuzu-main",
		slog.String("env", cfg.Enviroment),
		slog.String("addr", cfg.Address),
	)
	log.Debug("debug messages are enabled")

	// ============================Database============================ //
	//nolint:lll
	// docker run --name zuzu-postgres -v zuzu-db-data:/var/lib/postgresql/data -v -e 'PGDATA:/var/lib/postgresql/data/pgdata' './build/sql/injection_db.sql:/docker-entrypoint-initdb.d/init.sql' -p 8079:5432 --env-file .env --restart always postgres:latest
	db, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName))
	if err != nil {
		log.Error("fail open postgres", sl.Err(err))
		err = fmt.Errorf("error happened in sql.Open: %w", err)

		return err
	}
	defer db.Close()

	if err = db.Ping(context.Background()); err != nil {
		log.Error("fail ping postgres", sl.Err(err))
		err = fmt.Errorf("error happened in db.Ping: %w", err)

		return err
	}
	// ----------------------------Database---------------------------- //
	//
	//
	// ============================Init layers============================ //
	profileRepo := profileRepo.NewProfileRepo(db)
	profileUsecase := profileUsecase.NewProfileUsecase(log, profileRepo)
	profileHandler := profileHandler.NewProfileHandler(log, profileUsecase)

	authUsecase := authUsecase.NewAuthUsecase(profileRepo, cfg.Auther)
	authHandler := authHandler.NewAuthHandler(log, authUsecase)

	cartRepo := cartRepo.NewCartRepo(db)
	cartUsecase := cartUsecase.NewCartUsecase(cartRepo)
	cartHandler := cartHandler.NewCartHandler(log, cartUsecase)

	productsRepo := productsRepo.NewProductsRepo(db)
	productsUsecase := productsUsecase.NewProductsUsecase(productsRepo)
	productsHandler := productsHandler.NewProductsHandler(log, productsUsecase)

	categoryRepo := categoryRepo.NewCategoryRepo(db)
	categoryUsecase := categoryUsecase.NewCategoryUsecase(categoryRepo)
	categoryHandler := categoryHandler.NewCategoryHandler(log, categoryUsecase)

	orderRepo := orderRepo.NewOrderRepo(db)
	orderUsecase := orderUsecase.NewOrderUsecase(orderRepo, cartRepo)
	orderHandler := orderHandler.NewOrderHandler(log, orderUsecase)
	// ----------------------------Init layers---------------------------- //
	//
	//
	// ============================Create router============================ //

	r := mux.NewRouter().PathPrefix("/api").Subrouter()

	r.Use(middleware.Recover(log), middleware.CORSMiddleware, logmw.New(log))

	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not Found", http.StatusNotFound)
	})

	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	// ----------------------------Create router---------------------------- //
	//
	//
	// ============================Setup endpoints============================ //
	authMW := authmw.New(log, authUsecase.Auther)
	auth := r.PathPrefix("/auth").Subrouter()
	{
		auth.HandleFunc("/signup", authHandler.SignUp).
			Methods(http.MethodPost, http.MethodOptions)

		auth.HandleFunc("/signin", authHandler.SignIn).
			Methods(http.MethodPost, http.MethodOptions)

		auth.Handle("/logout", authMW(http.HandlerFunc(authHandler.LogOut))).
			Methods(http.MethodGet, http.MethodOptions)

		auth.Handle("/check_auth", authMW(http.HandlerFunc(authHandler.CheckAuth))).
			Methods(http.MethodGet, http.MethodOptions)
	}

	profile := r.PathPrefix("/profile").Subrouter()
	{
		profile.HandleFunc("/{id:[0-9a-fA-F-]+}", profileHandler.GetProfile).
			Methods(http.MethodGet, http.MethodOptions)

		profile.Handle("/update-photo", authMW(http.HandlerFunc(profileHandler.UpdatePhoto))).
			Methods(http.MethodPost, http.MethodOptions)

		profile.Handle("/update-data", authMW(http.HandlerFunc(profileHandler.UpdateProfileData))).
			Methods(http.MethodPost, http.MethodOptions)
	}

	cart := r.PathPrefix("/cart").Subrouter()
	{
		cart.Handle("/update", authMW(http.HandlerFunc(cartHandler.UpdateCart))).
			Methods(http.MethodPost, http.MethodOptions)

		cart.Handle("/summary", authMW(http.HandlerFunc(cartHandler.GetCart))).
			Methods(http.MethodGet, http.MethodOptions)

		cart.Handle("/add_product", authMW(http.HandlerFunc(cartHandler.AddProduct))).
			Methods(http.MethodPost, http.MethodOptions)

		cart.Handle("/delete_product", authMW(http.HandlerFunc(cartHandler.DeleteProduct))).
			Methods(http.MethodPost, http.MethodOptions)
	}

	products := r.PathPrefix("/products").Subrouter()
	{
		products.HandleFunc("/{id:[0-9a-fA-F-]+}", productsHandler.Product).
			Methods(http.MethodGet, http.MethodOptions)

		products.HandleFunc("/get_all", productsHandler.Products).
			Methods(http.MethodGet, http.MethodOptions)

		products.HandleFunc("/category", productsHandler.Category).
			Methods(http.MethodGet, http.MethodOptions)
	}

	category := r.PathPrefix("/category").Subrouter()
	{
		category.HandleFunc("/get_all", categoryHandler.Categories).
			Methods(http.MethodGet, http.MethodOptions)
	}

	order := r.PathPrefix("/order").Subrouter()
	{
		order.Handle("/create", authMW(http.HandlerFunc(orderHandler.CreateOrder))).
			Methods(http.MethodPost, http.MethodOptions)

		order.Handle("/get_current", authMW(http.HandlerFunc(orderHandler.GetCurrentOrder))).
			Methods(http.MethodGet, http.MethodOptions)

		order.Handle("/get_all", authMW(http.HandlerFunc(orderHandler.GetOrders))).
			Methods(http.MethodGet, http.MethodOptions)
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
	// SIGINT = ctrl+c; SIGTERM = kill; Interrupt = аппаратное прерывание, в Windows даст ошибку
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
		err = fmt.Errorf("error happened in srv.Shutdown: %w", err)

		return err
	}

	log.Info("server stopped")

	return nil
}
