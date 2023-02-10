#!/usr/bin/env bash
# Bash3 Boilerplate. Copyright (c) 2014, kvz.io

set -o errexit
set -o pipefail
set -o nounset

INGRESS_HOST=127.0.0.1.nip.io
CLUSTER_NAME=timber-dev
MLP_CHART_VERSION=0.3.4
MLP_URL=http://mlp.mlp.${INGRESS_HOST}
TIMEOUT=600s
PROJECT_NAME=testing


setup_cluster() {
  echo "Setting up k3d cluster"
  if [[ $(k3d cluster list | grep $CLUSTER_NAME | wc -l) -eq 0 ]]
  then
    k3d cluster create $CLUSTER_NAME --image rancher/k3s:v1.22.15-k3s1 --k3s-arg 'metrics-server@server:*' --port 80:80@loadbalancer
  fi

  k3d kubeconfig get ${CLUSTER_NAME} > /tmp/kubeconfig-${CLUSTER_NAME}.yaml
}

add_helm_repo() {
  echo "Adding helm repo"
  helm repo add caraml https://caraml-dev.github.io/helm-charts
  helm repo add bitnami https://charts.bitnami.com/bitnami
  helm repo update
}

install_mlp() {
  echo "Installing mlp"
  helm upgrade --install mlp caraml/mlp --namespace mlp --create-namespace \
    --version ${MLP_CHART_VERSION} \
    --set fullnameOverride=mlp \
    --set deployment.apiHost=http://mlp.mlp.${INGRESS_HOST}/v1 \
    --set ingress.enabled=true \
    --set ingress.class="traefik" \
    --set ingress.host=mlp.mlp.${INGRESS_HOST} \
    --wait --timeout=${TIMEOUT}
}

install_kafka() {
  echo "Installing kafka"
  helm upgrade --install kafka bitnami/kafka --namespace mlp --create-namespace \
    --values kafka/values.yaml --wait --timeout=${TIMEOUT}
}

create_mlp_project() {
  echo "Creating merlin project: PROJECT_NAME"
  curl "${MLP_URL}/v1/projects" -d "{
    \"name\"   : \"${PROJECT_NAME}\",
    \"team\"   : \"dsp\",
    \"stream\" : \"dsp\"
  }"
}

add_helm_repo
setup_cluster
install_mlp
install_kafka
create_mlp_project