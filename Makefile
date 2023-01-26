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
	cd tests ; isort e2e/
	cd tests ; flake8 e2e/
	cd tests ; black e2e/

.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell scripts/vertagen/vertagen.sh)))
	@echo "API version:" $(VERSION)

# ==================================
# Build recipes
# ==================================

.PHONY: build-dataset-service
build-dataset-service: version
	@echo "Building binary..."
	@cd ${DATASET_SVC_PATH}/cmd/${DATASET_SVC_PATH} && go build -o ./bin/${DATASET_SVC_PATH}

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
lint: lint-go lint-python

lint-go:
	@echo "> Linting code..."
	cd ${DATASET_SVC_PATH} && golangci-lint run --timeout 5m
	cd ${OBSERVATION_SVC_PATH} && golangci-lint run --timeout 5m
	cd ${COMMON_MODULE_PATH} && golangci-lint run --timeout 5m

lint-python:
	cd tests ; isort e2e/ --check-only
	cd tests ; flake8 e2e/
	cd tests ; black e2e/ --check

# ==================================
# Setup Services
# ==================================

.PHONY: dependency-services
dependency-services:
	cd infra/local && docker-compose up -d

# ==================================
# Python E2E tests
# ==================================

install-python-ci-dependencies:
	pip install -r tests/requirements.txt

e2e-clean-up:
	cd infra/tests/e2e && docker-compose down

e2e: build-dataset-service e2e-clean-up
	cd infra/tests/e2e && docker-compose down
	cd infra/tests/e2e && docker-compose up -d
	cd tests/e2e; python -m pytest -s -v

e2e-ci:
	cd tests/e2e; python -m pytest -s -v
