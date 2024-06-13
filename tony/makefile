run: build
	@./bin/tony

build:
	@go build -o bin/tony

mongo:
	@docker start mongo

test:
	@go test -v ./...
