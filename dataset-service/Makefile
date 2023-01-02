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
	@echo "dataset-service version:" $(VERSION)

tidy:
	go mod tidy

# ==================================
# Codegen recipes
# ==================================

.PHONY: generate
generate: lint
	@buf generate

.PHONY: setup
setup:
	@go mod tidy
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc

# Protobuf linting
.PHONY: lint
lint: 
	@buf lint

# ==================================
# Test recipes
# ==================================

test-dataset-service: tidy vendor
	@echo "> Running Dataset Service tests ..."
	go test -v ./... -coverpkg ./... -gcflags=-l -race -coverprofile cover.out.tmp -tags unit,integration
	cat cover.out.tmp | grep -v "api/api.go\|cmd\|.pb.go\|.pb.gw.go\|mock\|testutils\|server" > cover.out
	go tool cover -func cover.out
