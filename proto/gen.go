package proto

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.54.1
// https://grpc.io/docs/languages/go/quickstart/

//go:generate protoc api.proto --go_out=./../internal/models --go-grpc_out=./../internal/models

// For pkg auth
//go:generate protoc auth.proto --go_out=./../ --go-grpc_out=./../
