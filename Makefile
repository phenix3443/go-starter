
.PHONY: lint build test config-validate config-create env-create

build:
	go build -o build/ ./app/...

lint:
	golangci-lint run --new-from-rev=HEAD~1

test:
	go test -v -cover ./...

# 配置管理
config-validate:
	@echo "Validating configuration files..."
	@if [ -f "app/your-app/config/development.yaml" ]; then \
		echo "Validating development.yaml..."; \
		go run cmd/config/validate.go -c app/your-app/config/development.yaml; \
	fi
	@if [ -f "app/your-app/config/production.yaml" ]; then \
		echo "Validating production.yaml..."; \
		go run cmd/config/validate.go -c app/your-app/config/production.yaml; \
	fi
	@if [ -f "app/your-app/config/local.yaml" ]; then \
		echo "Validating local.yaml..."; \
		go run cmd/config/validate.go -c app/your-app/config/local.yaml; \
	fi

config-create:
	@echo "Creating configuration files..."
	@mkdir -p app/your-app/config
	@if [ ! -f "app/your-app/config/development.yaml" ]; then \
		cp app/your-app/config/development.yaml.example app/your-app/config/development.yaml 2>/dev/null || \
		echo "development.yaml.example not found, creating from template..."; \
	fi
	@if [ ! -f "app/your-app/config/production.yaml" ]; then \
		cp app/your-app/config/production.yaml.example app/your-app/config/production.yaml 2>/dev/null || \
		echo "production.yaml.example not found, creating from template..."; \
	fi
	@if [ ! -f "app/your-app/config/local.yaml" ]; then \
		cp app/your-app/config/config.yaml.example app/your-app/config/local.yaml 2>/dev/null || \
		echo "config.yaml.example not found, creating from template..."; \
	fi

# 环境变量文件管理
env-create:
	@echo "Creating environment variable files..."
	@mkdir -p app/your-app/env
	@if [ ! -f "app/your-app/env/local.env" ]; then \
		cp app/your-app/env/local.env app/your-app/env/.env 2>/dev/null || \
		echo "local.env not found, creating from template..."; \
	fi
	@if [ ! -f "app/your-app/env/development.env" ]; then \
		cp app/your-app/env/development.env app/your-app/env/.env 2>/dev/null || \
		echo "development.env not found, creating from template..."; \
	fi
	@if [ ! -f "app/your-app/env/production.env" ]; then \
		cp app/your-app/env/production.env app/your-app/env/.env 2>/dev/null || \
		echo "production.env not found, creating from template..."; \
	fi

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

###############################################################################
###                                Development                              ###
###############################################################################

dev-setup: config-create
	@echo "Development environment setup complete!"
	@echo "You can now run: make start-dc"

run-api-local:
	go run app/your-app/main.go --config app/your-app/config/local.yaml api

run-api-dev:
	go run app/your-app/main.go --config app/your-app/config/development.yaml api

run-api-prod:
	go run app/your-app/main.go --config app/your-app/config/production.yaml api

.PHONY: dev-setup run-api-local run-api-dev run-api-prod