
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run


build-styles:
	@npx @tailwindcss/cli -i ./web/static/css/main.css -o web/static/css/styles.css

local-run: build-styles
	$(GOBUILD) build -o ./bin/local ./cmd/app/main.go
	$(GORUN) ./cmd/app/main.go
	
air:
	air

.PHONY: build-styles local-run air