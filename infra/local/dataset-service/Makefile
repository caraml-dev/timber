
.PHONY: dev-env
dev-env: create-cluster install-dependencies

.PHONY: create-cluster
create-cluster:
	@./create_cluster.sh

.PHONY: install-dependencies
install-dependencies:
	@./install_dependencies.sh

.PHONY: clean-dev-env
clean-dev-env:
	@k3d cluster delete timber-dev
