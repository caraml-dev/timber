# ==================================
# General
# ==================================
.PHONY: build
build:
	@echo "> Building dataset service"
	go build -o ./bin/dataset-service ./cmd/dataset-service

.PHONY: test
test:
	@echo "> Running Dataset Service tests ..."
	go test -v ./... -coverpkg ./... -gcflags=-l -race -coverprofile cover.out.tmp -tags unit,integration
	cat cover.out.tmp | grep -v "api/api.go\|cmd\|.pb.go\|.pb.gw.go\|mock\|testutils\|server" > cover.out
	go tool cover -func cover.out

.PHONY: run
run:
	@echo "> Running Dataset Service locally"
	go run ./cmd/dataset-service/main.go serve --config config/local.yaml

.PHONY: version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell ../scripts/vertagen/vertagen.sh)))
	@echo "dataset-service version:" $(VERSION)

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: vendor
vendor:
	@echo "Fetching dependencies..."
	go mod vendor

.PHONY: setup
setup: tidy vendor
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

# ==================================
# Codegen recipes
# ==================================
.PHONY: generate
generate: lint
	@buf generate

# Protobuf linting
.PHONY: lint
lint: 
	@buf lint


# ==================================
# Development Environment
# ==================================
.PHONY: dev-env
dev-env:
	$(MAKE) -C ../infra/local/dataset-service dev-env
