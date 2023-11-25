package proto

//go:generate go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
//go:generate go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

//go:generate protoc api.proto --go_out=./../internal/models --go-grpc_out=./../internal/models

//go:generate protoc auth.proto --go_out=./../ --go-grpc_out=./../
//go:generate protoc products.proto --go_out=./../ --go-grpc_out=./../

//protoc -I proto proto/*.proto --go_out=./proto/gen/go/ --go_opt=paths=source_relative --go-grpc_out=./proto/gen/go/ --go-grpc_opt=paths=source_relative
