run: build
	@./bin/luna

build:
	@go build -o bin/luna main.go

test:
	@go test -v ./...

mongo:
	@docker start mongo