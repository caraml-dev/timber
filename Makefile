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

tidy-observation-service:
	cd ${OBSERVATION_SVC_PATH} && go mod tidy

# ==================================
# Linting recipes
# ==================================

.PHONY: lint
lint: lint-go

lint-go:
	@echo "> Linting Observation Service code..."
	cd ${OBSERVATION_SVC_PATH} && golangci-lint run --timeout 5m
	cd ${COMMON_MODULE_PATH} && golangci-lint run --timeout 5m

# ==================================
# Setup Services
# ==================================

.PHONY: dependency-services
dependency-services:
	cd infra/local && docker-compose up -d

.PHONY: observation-service
observation-service:
	cd observation-service && go run cmd/observation-service/main.go serve --config="config/example.yaml"

# ==================================
# Build recipes
# ==================================

build-fluentd-image:
	cd images/fluentd && docker build -t observation-service-fluentd .

# ==================================
# Test recipes
# ==================================

test-observation-service: tidy-observation-service
	@cd ${OBSERVATION_SVC_PATH} && go mod vendor
	@echo "> Running Observation Service tests ..."
	@cd ${OBSERVATION_SVC_PATH} && go test -v ./... -coverpkg ./... -gcflags=-l -race -coverprofile cover.out.tmp -tags unit,integration
	@cd ${OBSERVATION_SVC_PATH} && cat cover.out.tmp | grep -v "api/api.go\|cmd\|.pb.go\|mock\|testutils\|server" > cover.out
	@cd ${OBSERVATION_SVC_PATH} && go tool cover -func cover.out
