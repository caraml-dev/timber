export

COMMON_MODULE_PATH=common
DATASET_SVC_PATH=dataset-service
OBSERVATION_SVC_PATH=observation-service

# ==================================
# General
# ==================================

.PHONY: format
format: format-go format-python

format-go:
	@echo "> Formatting code"
	gofmt -s -w ${OBSERVATION_SVC_PATH}

format-python:
	$(MAKE) -C tests format

.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell scripts/vertagen/vertagen.sh)))
	@echo "API version:" $(VERSION)

# ==================================
# Build recipes
# ==================================

.PHONY: build-dataset-service
build-dataset-service: version
	@echo "Building dataset service binary..."
	$(MAKE) -C dataset-service build

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
## Run all linter
.PHONY: lint
lint: lint-go lint-python

## Run golang linter
.PHONY: lint-go
lint-go:
	@echo "> Linting code..."
	cd ${DATASET_SVC_PATH} && golangci-lint run --timeout 5m
	cd ${OBSERVATION_SVC_PATH} && golangci-lint run --timeout 5m
	cd ${COMMON_MODULE_PATH} && golangci-lint run --timeout 5m

## Run python linter
.PHONY: lint-python
lint-python:
	$(MAKE) -C tests lint

# ==================================
# Development environment
# ==================================

## Setup development environment, the same environment is also used for end to end test
dev-env:
	$(MAKE) -C infra/local/dataset-service dev-env

## Tear down development environment
clean-dev-env:
	$(MAKE) -C infra/local/dataset-service clean-dev-env

# ==================================
# E2E Test
# ==================================

## Setup e2e test environment and dependencies
.PHONY: setup-e2e
setup-e2e: dev-env build-dataset-service
	$(MAKE) -C tests dep

## Run e2e test for dataset-service
.PHONY: e2e
e2e:
	$(MAKE) -C tests test
