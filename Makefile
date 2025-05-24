
ifneq (,$(wildcard .env))
include .env
export
endif
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
APP_VERSION := $(shell cat VERSION)
ECR_REP_URL := ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${APP_ECR_REPO}
DB_DSN = postgres://$(APP_DB_USER):$(APP_DB_PASSWORD)@$(APP_DB_HOST):$(APP_DB_PORT)/$(APP_DB_NAME)?sslmode=$(APP_DB_SSL_MODE)
MIGRATIONS_DIR=./db/schema
SEED_DIR=./db/seed

clean:
	@find ./bin -mindepth 1 -delete

build-styles:
	rm web/static/css/tailwind-styles.css
	@npx tailwindcss -i ./web/static/css/main.css -o web/static/css/tailwind-styles.css
	@echo "built styles: web/static/css/tailwind-styles.css"

local-run: build-styles
	$(GORUN) ./cmd/app/main.go
	
build: clean
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64/app cmd/app/main.go
	@echo "built binary: ./bin/linux-amd64/app"

	@GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin-arm64/app cmd/app/main.go
	@echo "built binary: ./bin/darwin-arm64/app"

ci-image: build-styles build
	docker build --platform linux/amd64 -t ${ECR_REP_URL}:latest -t ${ECR_REP_URL}:${APP_VERSION} .

ci-push-image: ci-image
	aws --profile=${AWS_PROFILE} ecr get-login-password --region ${AWS_REGION} | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.us-east-1.amazonaws.com
	docker push ${ECR_REP_URL}:${APP_VERSION}
	docker push ${ECR_REP_URL}:latest

shell:
	@docker exec -it ${APP_LOCAL_CONTAINER_NAME} bash

app-version:
	@echo app version: ${APP_VERSION}


migrate-status:
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) status

# Create new migration file
migrate-create:
	@goose -dir $(MIGRATIONS_DIR) -s create new sql

# Migrate the DB to the most recent version available
migrate:
	@goose -dir $(MIGRATIONS_DIR) postgres $(DB_DSN) up


# Roll back the version by 1
migrate-rollback:
	@goose -dir $(MIGRATIONS_DIR) postgres "$(DB_DSN)" down

seed-create:
	@goose -dir $(SEED_DIR) -s create new sql

seed:
	@goose -dir $(SEED_DIR) -no-versioning postgres $(DB_DSN) up

seed-truncate:
	@goose -dir $(SEED_DIR) -no-versioning postgres $(DB_DSN) down

.PHONY: build-styles local-run air shell build ci-image ci-push-image migrate-status migrate-create migrate migrate-rollback seed-create seed seed-truncate