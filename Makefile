
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

PACKAGE_NAME:=github.com/phenix3443/go-starter
GOLANG_CROSS_VERSION  = v1.22
GOPATH ?= '$(HOME)/go'
release-dry-run:
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-v ${GOPATH}/pkg:/go/pkg \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		--clean --skip validate --skip publish --snapshot

release:
	@if [ ! -f ".release-env" ]; then \
		echo "\033[91m.release-env is required for release\033[0m";\
		exit 1;\
	fi
	docker run \
		--rm \
		--privileged \
		-e CGO_ENABLED=1 \
		--env-file .release-env \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/$(PACKAGE_NAME) \
		-w /go/src/$(PACKAGE_NAME) \
		ghcr.io/goreleaser/goreleaser-cross:${GOLANG_CROSS_VERSION} \
		release --clean --skip validate

.PHONY: release-dry-run release