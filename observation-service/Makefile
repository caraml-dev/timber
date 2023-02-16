export

OBSERVATION_SVC_PATH=observation-service
OBSERVATION_SVC_BIN_NAME=$(if $(OBSERVATION_SVC_NAME),$(OBSERVATION_SVC_NAME),observation_service)

# ==================================
# Build recipes
# ==================================
## Build observation service binary
.PHONY: build
build: version
	@echo "Building binary..."
	go build -o ./bin/${OBSERVATION_SVC_BIN_NAME} ./cmd/${OBSERVATION_SVC_PATH}/main.go

## Build observation service docker image
.PHONY: build-image
build-image: vendor version
	@$(eval IMAGE_TAG = $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/,)${BIN_NAME}:${VERSION})
	@echo "Building docker image: ${IMAGE_TAG}"
	docker build --tag ${IMAGE_TAG} . -f ${DOCKER_FILE}

## Build fluentd image
.PHONY: build-fluentd-image
build-fluentd-image:
	$(MAKE) -C ../images/fluentd build-image

# ==================================
# General
# ==================================

.PHONY: vendor
vendor:
	@echo "Fetching dependencies..."
	go mod vendor

.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell ../scripts/vertagen/vertagen.sh)))
	@echo "observation-service version:" $(VERSION)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: format
format: 
	gofmt -s -w .

.PHONY: dev-env
dev-env:
	cd ../infra/local/observation-service && docker-compose up -d


.PHONY: clean-dev-env
clean-dev-env:
	cd ../infra/local/observation-service && docker-compose down

# ==================================
# Run Services
# ==================================

.PHONY: run
run:
	go run cmd/observation-service/main.go serve --config="config/example.yaml"

# ==================================
# Test recipes
# ==================================
.PHONY: test
test: tidy vendor
	@echo "> Running Observation Service tests ..."
	go test -v ./... -coverpkg ./... -gcflags=-l -race -coverprofile cover.out.tmp -tags unit,integration
	cat cover.out.tmp | grep -v "api/api.go\|cmd\|.pb.go\|mock\|testutils\|server" > cover.out
	go tool cover -func cover.out

.PHONY: lint
lint: 
	golangci-lint run --timeout 5m
