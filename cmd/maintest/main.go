package main

import (
	"context"
	"fmt"
	"github.com/go-park-mail-ru/2023_2_potatiki/internal/models/protomodels"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	clientConn, err := grpc.Dial("0.0.0.0:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer clientConn.Close()

	client := protomodels.NewAuthServiceClient(clientConn)

	ctx := context.Background()
	returnedHello, err := client.SayHello(ctx, &protomodels.Hello{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(returnedHello)
}
