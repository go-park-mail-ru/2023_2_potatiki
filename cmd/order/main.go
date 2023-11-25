package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	addressRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/address/repo"
	cartRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/cart/repo"
	grpcOrder "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/delivery/grpc"
	orderRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/repo"
	orderUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/order/usecase"

	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/config"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		os.Exit(1)
	}
}

func run() (err error) {
	cfg := config.MustLoad() // TODO : dev-config.yaml -> readme.

	logFile, err := os.OpenFile("products.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("fail open logFile", sl.Err(err))
		return fmt.Errorf("fail open logFile: %w", err)
	}
	defer func() {
		err = errors.Join(err, logFile.Close())
	}()

	log := logger.Set(cfg.Enviroment, logFile)

	log.Info(
		"starting zuzu-order",
		slog.String("env", cfg.Enviroment),
		slog.String("addr", cfg.Address),
		slog.String("log_file_path", cfg.LogFilePath),
		slog.String("photos_file_path", cfg.PhotosFilePath),
	)
	log.Debug("debug messages are enabled")

	// ============================Database============================ //
	db, err := pgxpool.Connect(context.Background(), fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName))
	if err != nil {
		log.Error("fail open postgres", sl.Err(err))
		return fmt.Errorf("error happened in sql.Open: %w", err)
	}
	defer db.Close()

	if err = db.Ping(context.Background()); err != nil {
		log.Error("fail ping postgres", sl.Err(err))
		return fmt.Errorf("error happened in db.Ping: %w", err)
	}
	// ----------------------------Database---------------------------- //
	addressRepo := addressRepo.NewAddressRepo(db)
	cartRepo := cartRepo.NewCartRepo(db)
	orderRepo := orderRepo.NewOrderRepo(db)

	orderUsecase := orderUsecase.NewOrderUsecase(orderRepo, cartRepo, addressRepo)

	gRPCServer := grpc.NewServer()
	grpcOrder.Register(gRPCServer, log, orderUsecase)

	go func() {
		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.OrderPort))
		if err != nil {
			log.Error("listen returned err: ", sl.Err(err))
		}
		log.Info("grpc server started", slog.String("addr", listener.Addr().String()))
		if err := gRPCServer.Serve(listener); err != nil {
			log.Error("serve returned err: ", sl.Err(err))
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	gRPCServer.GracefulStop()
	log.Info("Gracefully stopped")
	return nil

}
