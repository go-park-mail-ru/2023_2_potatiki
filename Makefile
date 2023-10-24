.PHONY: all clean install uninstall

lint:
	golangci-lint run

compose:
	docker compose -f "local-docker-compose.yaml" up -d

build_:
	go build -o ./.bin cmd/main/main.go

swag:
	swag init -g ./cmd/main/main.go
	
test:
	go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./... ; \
	cat coverprofile_.tmp | grep -v _mock.go > coverprofile.tmp ; \
	rm coverprofile_.tmp ; \
	go tool cover -html coverprofile.tmp ; \
	go tool cover -func coverprofile.tmp

run: build_
	./.bin

rund: build_
	./.bin &

