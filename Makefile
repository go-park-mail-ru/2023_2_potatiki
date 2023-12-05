.PHONY: all clean install uninstall

compose:
	docker compose -f "local-docker-compose.yaml" up -d

build_:
	go build -o ./.bin cmd/main/main.go

run: build_
	./.bin

rund: build_
	./.bin &

runa: build_
	go run ./cmd/auth/main.go
	go run ./cmd/order/main.go
	./.bin

lint:
	golangci-lint run

swag:
	swag init -g ./cmd/main/main.go
	
cover:
	go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
	cat coverprofile_.tmp | grep -v _mock.go | grep -v _easyjson.go | grep -v .pb.go | grep -v _grpc.go > coverprofile.tmp ; \
	rm coverprofile_.tmp ; \
	go tool cover -html coverprofile.tmp ; \
	go tool cover -func coverprofile.tmp

test:
	go test ./...

done: lint test swag run
	git add .

protoc:
#go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	protoc -I proto proto/gmodels/*.proto --go_out=./proto/ --go_opt=paths=source_relative
	protoc -I proto proto/*.proto --go_out=./ --go-grpc_out=./

json:
#go get github.com/mailru/easyjson
#go install github.com/mailru/easyjson/...@latest
	easyjson -pkg ./internal/models/

#gotests -all -w internal/pkg/cart/usecase/usecase.go

