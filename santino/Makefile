# Load environment variables from .env file
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

run: build
	@./bin/san

build:
	@go build -o bin/san main.go

test:
	@go test -v ./...

db-down:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) down

db-up:
	@GOOSE_DRIVER=$(DB_DRIVER) GOOSE_DBSTRING=$(DB_NAME) go run github.com/pressly/goose/v3/cmd/goose@latest -dir=$(MIGRATION_DIR) up

# {
#   "message": "create user success",
#   "status": 200,
#   "user": {
#     "ID": 1,
#     "CreatedAt": "2024-07-15T09:32:17.214298208+08:00",
#     "UpdatedAt": "2024-07-15T09:32:17.214298208+08:00",
#     "DeletedAt": null,
#     "userID": "afd0d05e-bd38-4af8-9e89-879761ff7bf8",
#     "username": "zhangsan",
#     "email": "1233333@gmail.com",
#     "synopsis": "",
#     "avatar": "./avatar/1211.jpg",
#     "background_image": "./bg/1211.jpg"
#   }
# }