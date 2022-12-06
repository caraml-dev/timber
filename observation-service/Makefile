export

OBSERVATION_SVC_PATH=observation-service
OBSERVATION_SVC_BIN_NAME=$(if $(OBSERVATION_SVC_NAME),$(OBSERVATION_SVC_NAME),observation_service)

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

tidy:
	go mod tidy

# ==================================
# Build recipes
# ==================================

build: version
	@echo "Building binary..."
	go build -o ./bin/${OBSERVATION_SVC_BIN_NAME} ./cmd/${OBSERVATION_SVC_PATH}/main.go

.PHONY: build-image
build-image: vendor version
	@$(eval IMAGE_TAG = $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/,)${BIN_NAME}:${VERSION})
	@echo "Building docker image: ${IMAGE_TAG}"
	docker build --tag ${IMAGE_TAG} . -f ${DOCKER_FILE}

# ==================================
# Setup Services
# ==================================

.PHONY: observation-service
observation-service:
	go run cmd/observation-service/main.go serve --config="config/example.yaml"

# ==================================
# Test recipes
# ==================================

test-observation-service: tidy vendor
	go mod vendor
	@echo "> Running Observation Service tests ..."
	go test -v ./... -coverpkg ./... -gcflags=-l -race -coverprofile cover.out.tmp -tags unit,integration
	cat cover.out.tmp | grep -v "api/api.go\|cmd\|.pb.go\|mock\|testutils\|server" > cover.out
	go tool cover -func cover.out
