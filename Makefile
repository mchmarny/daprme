APP_NAME         =daprme
RELEASE_VERSION  =v0.6.3
LINTER_VERSION   =v1.31.0
DOCKER_USERNAME ?=$(DOCKER_USER)

all: help

devenv: ## Sets up development envirnment
	go get -u github.com/go-bindata/go-bindata/...
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(go env GOPATH)/bin $(LINTER_VERSION)
	golangci-lint --version
.PHONY: devenv


tidy: ## Updates the go modules and vendors all dependancies 
	go mod tidy
	go mod vendor
.PHONY: tidy

res: ## Compiles resource files into binary data resource 
	go-bindata -pkg lang -o pkg/lang/resource.go template/...
.PHONY: res

test: clean res ## Tests the entire project 
	go test -count=1 -race ./...
.PHONY: test

cover: clean ## Displays test coverage 
	go test -coverprofile=coverage.out ./... && go tool cover -mode=atomic -html=coverage.out
.PHONY: cover

run: clean tidy res ## Runs uncompiled code 
	go run main.go
.PHONY: run

demo: clean tidy res ## Runs uncompiled code with manifest 
	go run main.go --file test-data/demo.yaml --out ./apps
.PHONY: demo

build: clean tidy res ## Builds binaries
	CGO_ENABLED=0 go build \
		-ldflags "-X main.Version=$(RELEASE_VERSION)" \
		-mod vendor -o bin/$(APP_NAME) .
.PHONY: build

install: tidy res ## Installs locally 
	CGO_ENABLED=0 go install -ldflags "-X main.Version=$(RELEASE_VERSION)" .
.PHONY: install

release: clean tidy res ## Builds releases
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-X main.Version=$(RELEASE_VERSION)" -mod vendor -o release/$(APP_NAME) .
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=$(RELEASE_VERSION)" -mod vendor -o release/$(APP_NAME)_linux .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-X main.Version=$(RELEASE_VERSION)" -mod vendor -o release/$(APP_NAME).exe .
.PHONY: release

lint: clean ## Lints the entire project 
	golangci-lint run --timeout=3m
.PHONY: lint

tag: ## Creates release tag 
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)
.PHONY: tag

goclean: clean ## Cleans bin and temp directories
	go clean
	rm -fr ./vendor
	rm -fr ./bin
.PHONY: goclean

clean: ## Cleans test dir
	rm -fr ./test
.PHONY: clean

help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
.PHONY: help