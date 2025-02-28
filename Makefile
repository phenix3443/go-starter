
.PHONY: lint build

build:
	go build -o build/ ./app/...
lint:
	golangci-lint run --new-from-rev=HEAD~1

###############################################################################
###                        Docker                                           ###
###############################################################################
DOCKER := $(shell which docker)
DOCKER_IMAGE := phenix3443/go-starter
COMMIT_HASH := $(shell git rev-parse --short=7 HEAD)
DOCKER_TAG := $(COMMIT_HASH)

build-docker:
	$(DOCKER) build -t ${DOCKER_IMAGE}:${DOCKER_TAG} .
	$(DOCKER) tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:latest
	$(DOCKER) tag ${DOCKER_IMAGE}:${DOCKER_TAG} ${DOCKER_IMAGE}:${COMMIT_HASH}

.PHONY: build-docker
###############################################################################
###                        Docker Compose                                   ###
###############################################################################

start-dc:
	docker compose up -d
	docker compose ps
	
stop-dc:
	docker compose down --volumes

.PHONY: start-dc stop-dc

###############################################################################
###                                Releasing                                ###
###############################################################################
release:
	goreleaser release --snapshot --clean 

.PHONY: release-dry-run release