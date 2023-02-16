# End to end test

### Running End-to-end Tests locally

E2E tests are currently setup via `e2e` Python module. This setup does the following:

1. Create e2e test environment in k3d cluster.
2. Pytests would reference the k3d cluster via kubeconfig parameter in `tests/e2e/config/dataset-service.local.yaml`.
3. Dataset service will be started from a local environment via the test suites.

Running the e2e tests locally are as follows:

First, spin up K8s and other services required for Dataset and Observation Services. Additionally, it will build dataset service binary.

```bash
make setup-e2e
```

Second, install all python dependencies. It's advisible to create virtual environment

```bash
virtualenv env
source env/bin/activate
make dep
```

Lastly, execute the end to end test

```bash
make e2e
```
