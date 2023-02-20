# ==================================
# General
# ==================================
.PHONY: build
## Build dataset-service binary
build:
	@echo "> Building dataset service"
	go build -o ./bin/dataset-service ./cmd/dataset-service

.PHONY: test
## Execute all test and generate coverage report
test:
	@echo "> Running Dataset Service tests ..."
	go test -v ./... -coverpkg ./... -gcflags=-l -race -coverprofile cover.out.tmp -tags unit,integration
	cat cover.out.tmp | grep -v "api/api.go\|cmd\|.pb.go\|.pb.gw.go\|mock\|testutils\|server" > cover.out
	go tool cover -func cover.out

.PHONY: run
## Run dataset-service locally
run:
	@echo "> Running Dataset Service locally"
	go run ./cmd/dataset-service/main.go serve --config config/local.yaml

.PHONY: version
## Generate version string
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell ../scripts/vertagen/vertagen.sh)))
	@echo "dataset-service version:" $(VERSION)

.PHONY: tidy
## Run go mod tidy
tidy:
	go mod tidy

.PHONY: vendor
## Run go mod vendor
vendor:
	@echo "Fetching dependencies..."
	go mod vendor

.PHONY: setup
## Install software dependencies
setup: tidy vendor
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	@test -x ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2

.PHONY: format
## Format code
format: 
	@gofmt -s -w .
	@goimports -w -local github.com/caraml-dev/timber/common,github.com/caraml-dev/timber/dataset-service,github.com/caraml-dev/timber/observation-service \
		$$(find . -type f -name '*.go' -not -path "**/vendor/*")

# ==================================
# Codegen recipes
# ==================================
.PHONY: generate
## Generate code from protobuf definition
generate:
	@buf generate

.PHONY: lint
## Protobuf linting
lint: 
	@buf lint
	@golangci-lint run --timeout 5m

# ==================================
# Development Environment
# ==================================
.PHONY: dev-env
## Setup development environment and its dependencies
dev-env:
	$(MAKE) -C ../infra/local/dataset-service dev-env

.PHONY: clean-dev-env
## Tear down development environment
clean-dev-env:
	$(MAKE) -C ../infra/local/dataset-service clean-dev-env

#################################################################################
# Self Documenting Commands
#################################################################################
.DEFAULT_GOAL := show-help
# Inspired by <http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html>
# sed script explained:
# /^##/:
# 	* save line in hold space
# 	* purge line
# 	* Loop:
# 		* append newline + line to hold space
# 		* go to next line
# 		* if line starts with doc comment, strip comment character off and loop
# 	* remove target prerequisites
# 	* append hold space (+ newline) to line
# 	* replace newline plus comments by `---`
# 	* print line
# Separate expressions are necessary because labels cannot be delimited by
# semicolon; see <http://stackoverflow.com/a/11799865/1968>
## Show help
show-help:
	@echo "$$(tput bold)Available rules:$$(tput sgr0)"
	@echo
	@sed -n -e "/^## / { \
		h; \
		s/.*//; \
		:doc" \
		-e "H; \
		n; \
		s/^## //; \
		t doc" \
		-e "s/:.*//; \
		G; \
		s/\\n## /---/; \
		s/\\n/ /g; \
		p; \
	}" ${MAKEFILE_LIST} \
	| LC_ALL='C' sort --ignore-case \
	| awk -F '---' \
		-v ncol=$$(tput cols) \
		-v indent=19 \
		-v col_on="$$(tput setaf 6)" \
		-v col_off="$$(tput sgr0)" \
	'{ \
		printf "%s%*s%s ", col_on, -indent, $$1, col_off; \
		n = split($$2, words, " "); \
		line_length = ncol - indent; \
		for (i = 1; i <= n; i++) { \
			line_length -= length(words[i]) + 1; \
			if (line_length <= 0) { \
				line_length = ncol - indent - length(words[i]) - 1; \
				printf "\n%*s ", -indent, " "; \
			} \
			printf "%s ", words[i]; \
		} \
		printf "\n"; \
	}' \
	| more $(shell test $(shell uname) = Darwin && echo '--no-init --raw-control-chars')
