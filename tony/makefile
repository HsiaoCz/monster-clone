run: build
	@./bin/tony

build:
	@go build -o bin/tony

mongo:
	@docker start mongo
	@docker exec -it mongo mongosh

test:
	@go test -v ./...