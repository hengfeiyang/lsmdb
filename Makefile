run:
	go run cmd/lsmdb/main.go
build:
	go build -o lsmdb cmd/lsmdb/main.go
test:
	go test -v ./...
