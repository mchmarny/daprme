SERVICE_NAME    ?=demoapp
RELEASE_VERSION ?=v0.1.1
DOCKER_USERNAME ?=$(DOCKER_USER)

.PHONY: all
all: help

.PHONY: clean
clean:  ## Cleans project
	npm run clean

.PHONY: install
install: ## Installs dependencies 
	npm install

.PHONY: test
test: ## Tests the entire project 
	npm test

.PHONY: run
run: tidy ## Runs un-compiled code in Dapr
	dapr run \
		 --app-id $(SERVICE_NAME) \
		 --app-port 3000 \
		 --app-protocol http \
		 --dapr-http-port 3500 \
         --components-path ./config \
         node main.go
.PHONY: eventhazelcast-pubsub
eventhazelcast-pubsub: ## Publishes sample JSON message to Dapr pubsub API 
	curl -d '{ "from": "John", "to": "Lary", "message": "hi" }' \
     -H "Content-type: application/json" \
     "http://localhost:3500/v1.0/publish/hazelcast-pubsub/messages"

.PHONY: image
image: tidy ## Builds and publish docker image 
	docker build -t "$(DOCKER_USERNAME)/$(SERVICE_NAME):$(RELEASE_VERSION)" .
	docker push "$(DOCKER_USERNAME)/$(SERVICE_NAME):$(RELEASE_VERSION)"

.PHONY: tag
tag: ## Creates release tag 
	git tag $(RELEASE_VERSION)
	git push origin $(RELEASE_VERSION)


.PHONY: help
help: ## Display available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk \
		'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'