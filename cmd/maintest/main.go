package main

//func main() {
//	clientConn, err := grpc.Dial("0.0.0.0:8001", grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer clientConn.Close()
//
//	client := protomodels.NewAuthServiceClient(clientConn)
//
//	ctx := context.Background()
//	returnedHello, err := client.SayHello(ctx, &protomodels.Hello{})
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println(returnedHello)
//}
