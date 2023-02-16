export

COMMON_MODULE_PATH=common
DATASET_SVC_PATH=dataset-service
OBSERVATION_SVC_PATH=observation-service

# ==================================
# Build recipes
# ==================================
## Build all
.PHONY: build
build: build-dataset-service build-observation-service

## Build dataset service
.PHONY: build-dataset-service
build-dataset-service: version
	$(MAKE) -C dataset-service build

## Build observation service
.PHONY: build-observation-service
build-observation-service: version
	$(MAKE) -C observation-service build

## Generate version
.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell scripts/vertagen/vertagen.sh)))
	@echo "API version:" $(VERSION)

# ==================================
# General
# ==================================

## Format all source code
.PHONY: format
format: format-go format-python

## Format all golang source code
format-go:
	@echo "> Formatting code"
	$(MAKE) -C observation-service format
	$(MAKE) -C dataset-service format

## Format all python source code
format-python:
	$(MAKE) -C tests format

# ==================================
# Code dependencies recipes
# ==================================

## Setup dependencies
.PHONY: setup
setup:
	@echo "> Initializing dependencies ..."
	@test -x ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1

	@echo "Setting up dev tools..."
	@test -x "$(which pre-commit)" || pip install pre-commit
	@pre-commit install
	@pre-commit install-hooks

## Install Python dependencies
.PHONY: dep-python
dep-python:
	$(MAKE) -C tests dep

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
	$(MAKE) -C dataset-service lint
	$(MAKE) -C observation-service lint
	cd ${COMMON_MODULE_PATH} && golangci-lint run --timeout 5m

## Run python linter
.PHONY: lint-python
lint-python:
	$(MAKE) -C tests lint

# ==================================
# Development environment
# ==================================

## Setup development environment, the same environment is also used for end to end test
.PHONY: dev-env
dev-env:
	$(MAKE) -C infra/local/dataset-service dev-env

## Tear down development environment
.PHONY: clean-dev-env
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
