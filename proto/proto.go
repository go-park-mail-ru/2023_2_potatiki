package proto

//go:generate protoc -I proto proto/gmodels/*.proto --go_out=./proto/ --go_opt=paths=source_relative
//go:generate protoc -I proto proto/*.proto --go_out=./ --go-grpc_out=./
