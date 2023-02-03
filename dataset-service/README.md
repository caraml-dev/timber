# dataset-service

Dataset Service will be responsible for the following:
1. Deploy/undeploy log writer deployment based on the availability of model and router deployment. 
2. Synchronize the access control list of CaraML projects to the corresponding BQ dataset to avoid maintaining separate processes to obtain access to the dataset.
3. Store and provide access to the log table metadata produced in CaraML. 
4. Store and provide read only access to the list of log writers it manages.
5. Store and provide read only access to the list of observation services it manages.

## Getting Started

### OpenAPI Specifications and Docs

To get started, install following software:

- [buf](https://docs.buf.build/installation)
- [Go 1.18 or above](https://go.dev/doc/install)
- Go dependencies and protoc plugins

    ```bash
    make setup
    ```

To generate all code and documentation
```
make generate
```

To perform lint check
```
make lint
```
