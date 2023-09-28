go test -json ./... -coverprofile coverprofile.tmp -coverpkg=./...
cat coverprofile.tmp | grep -v _mock.go > coverprofile2.tmp
go tool cover -html coverprofile2.tmp
go tool cover -func coverprofile2.tmp
