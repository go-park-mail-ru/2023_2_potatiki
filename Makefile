#compose:
#	docker compose -f "local-docker-compose.yaml" up -d

build_:
	go build -o ./.bin cmd/main/main.go

run: build_
	./.bin