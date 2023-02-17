export

OBSERVATION_SVC_PATH=observation-service
OBSERVATION_SVC_BIN_NAME=$(if $(OBSERVATION_SVC_NAME),$(OBSERVATION_SVC_NAME),observation_service)

# ==================================
# Build recipes
# ==================================
.PHONY: build
## Build observation service binary
build: version
	@echo "Building binary..."
	go build -o ./bin/${OBSERVATION_SVC_BIN_NAME} ./cmd/${OBSERVATION_SVC_PATH}/main.go

.PHONY: build-fluentd-image
## Build fluentd image
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
	@gofmt -s -w .
	@goimports -w -local github.com/caraml-dev/timber/common,github.com/caraml-dev/timber/dataset-service,github.com/caraml-dev/timber/observation-service \
		$$(find . -type f -name '*.go' -not -path "**/vendor/*")

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
