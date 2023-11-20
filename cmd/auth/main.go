package main

import (
	"context"
	"errors"
	"fmt"
	grpcAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc"
	generatedAuth "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/delivery/grpc/generated"
	authUsecase "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/auth/usecase"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/config"
	profileRepo "github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/profile/repo"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/pkg/utils/logger/sl"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"os"
)

//type Server struct {
//	protomodels.AuthServiceServer
//}

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
	defer func() {
		err = errors.Join(err, logFile.Close())
	}()

	log := logger.Set(cfg.Enviroment, logFile)

	log.Info(
		"starting zuzu-main",
		slog.String("env", cfg.Enviroment),
		slog.String("addr", cfg.Address),
		slog.String("log_file_path", cfg.LogFilePath),
		slog.String("photos_file_path", cfg.PhotosFilePath),
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
	profileRepo := profileRepo.NewProfileRepo(db)

	authUsecase := authUsecase.NewAuthUsecase(profileRepo, cfg.AuthJWT)
	service := grpcAuth.NewGrpcAuthHandler(authUsecase)

	listener, err := net.Listen("tcp", "0.0.0.0:8010")
	if err != nil {
		return err
	}

	server := grpc.NewServer()

	generatedAuth.RegisterAuthServiceServer(server, service)

	return server.Serve(listener)
}

//func main1() {
//
//	listener, err := net.Listen("tcp", "0.0.0.0:8010")
//	if err != nil {
//		return err
//	}
//
//	server := grpc.NewServer()
//
//	protomodels.RegisterAuthServiceServer(server, &Server{})
//
//	err = server.Serve(listener)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//}
//
//func (s *Server) SayHello(context.Context, *protomodels.Hello) (*protomodels.Hello, error) {
//
//	return &protomodels.Hello{Line: "Hello world"}, nil
//}
