build_:
	go build -o ./.bin cmd/main/main.go

run: build_
	./.bin