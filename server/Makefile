MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"

CONFIG_FILE ?= ./config/local.yml
APP_DSN ?= $(shell sed -n 's/^ports_dsn: *\(.*\)/\1/p' $(CONFIG_FILE))

PID_FILE := './.pid'
FSWATCH_FILE := './fswatch.cfg'

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: run
run: ## run the gRPC server
	go run ${LDFLAGS} cmd/server/main.go

.PHONY: run-restart
run-restart: ## restart the API server
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Source file changed. Restarting server..."
	@go run ${LDFLAGS} cmd/server/main.go & echo $$! > $(PID_FILE)
	@printf '%*s\n' "80" '' | tr ' ' -

run-live: ## run the gRPC server with live reload support (requires fswatch)
	@go run ${LDFLAGS} cmd/server/main.go & echo $$! > $(PID_FILE)
	@fswatch -x -o --event Created --event Updated --event Renamed -r internal pkg cmd config | xargs -n1 -I {} make run-restart

.PHONY: build
build:  ## build the gRPC server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o ./cmd/build/server ./cmd/server

.PHONY: build-docker
build-docker: ## build the gRPC server as a docker image
	docker build -f cmd/docker/Dockerfile -t eop09-server ..

.PHONY: clean
clean: ## remove temporary files
	rm -rf cmd/build/* coverage.out coverage-all.out

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: db-start
db-start: ## start the database server
	docker run -d --rm --name eop09-db \
		-v eop09-db:/data/db \
		-v eop09-config:/data/config \
		-e MONGO_INITDB_ROOT_USERNAME=admin \
		-e MONGO_INITDB_ROOT_PASSWORD=Masa1818 \
		-p 27017:27017 \
		mongo:4.4.8

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop eop09-db

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)
