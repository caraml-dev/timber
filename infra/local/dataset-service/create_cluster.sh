CLUSTER_NAME=timber-dev

setup_cluster() {
  echo "Setting up k3d cluster"
  if [[ $(k3d cluster list | grep $CLUSTER_NAME | wc -l) -eq 0 ]]
  then
    k3d cluster create $CLUSTER_NAME --image rancher/k3s:v1.23.16-k3s1 --k3s-arg 'metrics-server@server:*' --port 80:80@loadbalancer
  else
    kubectx k3d-${CLUSTER_NAME}
  fi

  k3d kubeconfig get ${CLUSTER_NAME} > /tmp/kubeconfig-${CLUSTER_NAME}.yaml
}

setup_cluster
