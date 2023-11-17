package main

import (
	"context"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models/protomodels"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	protomodels.AuthServiceServer
}

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:8001")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	protomodels.RegisterAuthServiceServer(server, &Server{})

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}

}

func (s *Server) SayHello(context.Context, *protomodels.Hello) (*protomodels.Hello, error) {

	return &protomodels.Hello{Line: "Hello world"}, nil
}
