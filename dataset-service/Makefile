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
