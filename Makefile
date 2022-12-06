export

COMMON_MODULE_PATH=common
OBSERVATION_SVC_PATH=observation-service

# ==================================
# General
# ==================================

.PHONY: format
format: format-go

format-go:
	@echo "> Formatting code"
	gofmt -s -w ${OBSERVATION_SVC_PATH}

.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell scripts/vertagen/vertagen.sh)))
	@echo "API version:" $(VERSION)

# ==================================
# Build recipes
# ==================================

.PHONY: build-image
build-image: version
	@$(eval IMAGE_TAG = $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/,)${BIN_NAME}:${VERSION})
	@echo "Building docker image: ${IMAGE_TAG}"
	docker build --tag ${IMAGE_TAG} . -f ${DOCKER_FILE}

# ==================================
# Code dependencies recipes
# ==================================
.PHONY: setup
setup:
	@echo "> Initializing dependencies ..."
	@test -x ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1

	@echo "Setting up dev tools..."
	@test -x "$(which pre-commit)" || pip install pre-commit
	@pre-commit install
	@pre-commit install-hooks

# ==================================
# Linting recipes
# ==================================

.PHONY: lint
lint: lint-go

lint-go:
	@echo "> Linting code..."
	cd ${OBSERVATION_SVC_PATH} && golangci-lint run --timeout 5m
	cd ${COMMON_MODULE_PATH} && golangci-lint run --timeout 5m

# ==================================
# Setup Services
# ==================================

.PHONY: dependency-services
dependency-services:
	cd infra/local && docker-compose up -d
