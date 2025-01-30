include .env
export $(shell sed 's/=.*//g' .env)

up:
	go run cmd/server/main.go

migrate-up:
	goose -dir migrations postgres "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up

migrate-down:
	goose -dir migrations postgres "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" reset

docker-build:
	docker-compose up --build && \
	docker run --rm --network host \
	  -e GOOSE_DRIVER="postgres" \
	  -e GOOSE_DBSTRING="postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" \
	  -v ./migrations:/migrations \
	  ghcr.io/kukymbr/goose-docker:3.19.2 -dir /migrations postgres "postgres://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)" up
