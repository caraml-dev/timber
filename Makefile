export

COMMON_MODULE_PATH=common
DATASET_SVC_PATH=dataset-service
OBSERVATION_SVC_PATH=observation-service

# ==================================
# Build recipes
# ==================================
.PHONY: build
## Build all
build: build-dataset-service build-observation-service

.PHONY: build-dataset-service
## Build dataset service
build-dataset-service: version
	$(MAKE) -C dataset-service build

.PHONY: build-observation-service
## Build observation service
build-observation-service: version
	$(MAKE) -C observation-service build

.PHONY: build-image
## Build docker image
build-image: version
	@$(eval IMAGE_TAG = $(if $(DOCKER_REGISTRY),$(DOCKER_REGISTRY)/,)${BIN_NAME}:${VERSION})
	@echo "Building docker image: ${IMAGE_TAG}"
	docker build --tag ${IMAGE_TAG} . -f ${DOCKER_FILE}

.PHONY: version
## Generate version
version:
	$(eval VERSION=$(if $(OVERWRITE_VERSION),$(OVERWRITE_VERSION),v$(shell scripts/vertagen/vertagen.sh)))
	@echo "API version:" $(VERSION)

# ==================================
# General
# ==================================

.PHONY: format
## Format all source codes
format: format-go format-python

.PHONY: format-go
## Format all golang source code
format-go:
	@echo "> Formatting code"
	$(MAKE) -C observation-service format
	$(MAKE) -C dataset-service format

.PHONY: format-python
## Format all python source code
format-python:
	$(MAKE) -C tests format

# ==================================
# Code dependencies recipes
# ==================================

.PHONY: setup
## Setup dependencies
setup:
	@echo "> Initializing dependencies ..."
	@test -x ${GOPATH}/bin/golangci-lint || go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.40.1

	@echo "Setting up dev tools..."
	@test -x "$(which pre-commit)" || pip install pre-commit
	@pre-commit install
	@pre-commit install-hooks

.PHONY: dep-python
## Install Python dependencies
dep-python:
	$(MAKE) -C tests dep

# ==================================
# Linting recipes
# ==================================
.PHONY: lint
## Run all linter
lint: lint-go lint-python

.PHONY: lint-go
## Run golang linter
lint-go:
	@echo "> Linting code..."
	$(MAKE) -C dataset-service lint
	$(MAKE) -C observation-service lint
	cd ${COMMON_MODULE_PATH} && golangci-lint run --timeout 5m

.PHONY: lint-python
## Run python linter
lint-python:
	$(MAKE) -C tests lint

# ==================================
# Development environment
# ==================================

.PHONY: dev-env
## Setup development environment, the same environment is also used for end to end test
dev-env:
	$(MAKE) -C infra/local/dataset-service dev-env

.PHONY: clean-dev-env
## Tear down development environment
clean-dev-env:
	$(MAKE) -C infra/local/dataset-service clean-dev-env

# ==================================
# E2E Test
# ==================================
.PHONY: setup-e2e
## Setup e2e test environment and dependencies
setup-e2e: dev-env build-dataset-service
	$(MAKE) -C tests dep

.PHONY: e2e
## Run e2e test for dataset-service
e2e:
	$(MAKE) -C tests test

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
