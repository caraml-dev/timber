# tests

### Running End-to-end tests locally

E2E tests are currently setup via `tests/e2e` directory. This setup does the following:

1. Deploy a k3d cluster via docker-compose
2. Pytests would reference the k3d cluster via kubeconfig parameter in `tests/e2e/config/dataset-service.local.yaml`
3. Dataset service will be started from a local environment, which triggers helm deployment of Observation Service into the k3d cluster

Running the e2e tests locally are as follows:

#### <a>Spin up k8s and required dependencies</b>
---

First, spin up K8s and other services required for Dataset and Observation Services:

```bash
pushd ../infra/tests/e2e
{
    docker-compose up -d
}
popd
```

Here we have to wait for all the pods to be ready.

```bash
watch KUBECONFIG=/tmp/kubeconfig kubectl get pod -A
```

An example of ready would look something like this:
```
NAMESPACE     NAME                                      READY   STATUS    RESTARTS   AGE
kube-system   local-path-provisioner-687d6d7765-npqvv   1/1     Running   0          2m35s
kube-system   coredns-7b5bbc6644-wzrxr                  1/1     Running   0          2m35s
kube-system   metrics-server-667586758d-sjj9j           1/1     Running   0          2m35s
```

#### <b>Run E2E tests</b>
---

- Builds Dataset Service binary
- Starts k3d cluster via Docker Compose
- Run pytests against k3d cluster

```bash
make e2e
```
