go test -json ./... -coverprofile coverprofile_.tmp -coverpkg=./...
	cat coverprofile_.tmp | grep -v _mock.go > coverprofile.tmp
	rm coverprofile_.tmp
	go tool cover -html coverprofile.tmp
	go tool cover -func coverprofile.tmp