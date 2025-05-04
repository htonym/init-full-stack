
ifneq (,$(wildcard .env))
include .env
export
endif
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
APP_VERSION := $(shell cat VERSION)
ECR_REP_URL := ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/${APP_ECR_REPO}

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

.PHONY: build-styles local-run air shell build ci-image ci-push-image