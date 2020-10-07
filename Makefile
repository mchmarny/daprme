APP_NAME         =daprme
RELEASE_VERSION  =v0.1.1
DOCKER_USERNAME ?=$(DOCKER_USER)

all: help

.PHONY: tidy
tidy: ## Updates the go modules and vendors all dependancies 
	go mod tidy
	go mod vendor

.PHONY: test
test: test ## Tests the entire project 
	go test -v -count=1 -race ./...
	# go test -v -count=1 -run SpecificTestName ./...

.PHONY: run
run: tidy ## Runs uncompiled code 
	go run main.go

.PHONY: build
build: tidy ## Builds binaries
	CGO_ENABLED=0 go build -ldflags "-X main.Version=$(RELEASE_COMMIT)" \
    -mod vendor -o ../../demo/bin/$(APP_NAME) .
	bash -c "cp -r ./template/ ../../demo/template/"

.PHONY: lint
lint: ## Lints the entire project 
	golangci-lint run --timeout=3m

.PHONY: tag
tag: ## Creates release tag 
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)

.PHONY: clean
clean: ## Cleans bin and temp directories
	go clean
	rm -fr ./vendor
	rm -fr ./bin

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'



