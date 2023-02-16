export

## Run end to end test
.PHONY: test
test:
	python -m pytest -s -v

## Run linter
.PHONY: lint
lint:
	@isort e2e/ --check-only
	@flake8 e2e/
	@black e2e/ --check

## Format
.PHONY: format
format:
	isort e2e/
	black e2e/
	flake8 e2e/

## Install python dependencies
.PHONY: dep
dep:
	pip install -r requirements.txt

## Setup e2e test environment and dependencies
.PHONY: setup-e2e
setup-e2e:
	$(MAKE) -C ../ dep
