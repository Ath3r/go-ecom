build:
	@go build -o bin/go-ecom cmd/server/main.go

run: build
	@./bin/go-ecom

test:
	@go test -v ./...

version:
	@go version