.PHONY: all clean install uninstall

compose:
	docker compose -f "local-docker-compose.yaml" up -d

build_:
	go build -o ./.bin cmd/main/main.go

run: build_
	./.bin

rund: build_
	./.bin &

lint:
	golangci-lint run

swag:
	swag init -g ./cmd/main/main.go
	
cover:
	go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
	cat coverprofile_.tmp | grep -v _mock.go > coverprofile.tmp ; \
	rm coverprofile_.tmp ; \
	go tool cover -html coverprofile.tmp ; \
	go tool cover -func coverprofile.tmp

test:
	go test ./...

done: lint test swag run
	git add .

#gotests -all -w internal/pkg/cart/usecase/usecase.go

