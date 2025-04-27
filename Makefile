
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run


build-styles:
	@npx tailwindcss -i ./web/static/css/main.css -o web/static/css/tailwind-styles.css
	@echo "built styles: web/static/css/tailwind-styles.css"

local-run: build-styles
	$(GORUN) ./cmd/app/main.go
	
build-bin:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/linux-amd64/app cmd/app/main.go
	@echo "built binary: ./bin/linux-amd64/app"

	@GOOS=darwin GOARCH=arm64 go build -o ./bin/darwin-arm64/app cmd/app/main.go
	@echo "built binary: ./bin/darwin-arm64/app"


.PHONY: build-styles local-run air