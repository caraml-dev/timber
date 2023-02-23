
.PHONY: dev-env
dev-env: setup-db create-cluster install-dependencies

.PHONY: create-cluster
create-cluster:
	@./create_cluster.sh

.PHONY: install-dependencies
install-dependencies:
	@./install_dependencies.sh

.PHONY: setup-db
setup-db:
	@docker-compose up -d

.PHONY: clean-dev-env
clean-dev-env:
	@docker-compose down -v
	@k3d cluster delete timber-dev

