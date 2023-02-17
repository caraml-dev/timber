# Dataset Service

Dataset Service is responsible for managing deployment of log collection component (i.e. log writer and observation service), organizing the collected data, and generating dataset from the collected data.

## API Specification

Dataset Service API is defined as [protobuf](proto/caraml/timber/v1/dataset_service.proto). The service exposes the API both as gRPC API and REST API. The REST API documentation can be found in [dataset_service.swagger.json](docs/openapiv2/caraml/timber/v1/dataset_service.swagger.json).

## Getting Started

To get started, install following software:

- [buf](https://docs.buf.build/installation)
- [Go 1.18 or above](https://go.dev/doc/install)
- Go dependencies and protoc plugins
- [k3d](https://k3d.io/v5.4.7/#installation)

Setup all dependencies

```bash
make setup
```

To build dataset service binary execute

```bash
make build
```

To perform test

```bash
make test
```

To setup local development environment. This command will create a k3d cluster and install dependency services (MLP and Kafka) in the cluster.

```bash
make dev-env
```

Once the local development environment is ready, you can run dataset-service locally using following command. 
The command will use [local.yaml](config/local.yaml) configuration, which is preconfigured to target the k3d cluster for deployment of log writer and observation service. 

```bash
make run
```

And send following requests to deploy observation-service or log writer

```bash

# Deploy observation service
curl --location --request POST 'localhost:8080/v1/projects/1/observation_services' \
--header 'Content-Type: application/json' \
--data-raw '{
    "observation_service" : {
        "project_id" : 1,
        "name" : "my-observation",
        "source" : {
            "type": "OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA",
            "kafka" : {
                "brokers": "kafka.mlp.svc.cluster.local",
                "topic" : "sample"
            }
        }
    }
}'

# Deploy router log writer
curl --location --request POST 'localhost:8080/v1/projects/1/log_writers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "log_writer" : {
        "project_id" : 1,
        "name" : "my-router-log",
        "source" : {
            "type": "LOG_WRITER_SOURCE_TYPE_ROUTER_LOG",
            "router_log_source" : {
                "router_id": 1,
                "router_name" : "my-router",
                "kafka" : {
                    "brokers": "kafka.mlp.svc.cluster.local",
                    "topic" : "my-router-log"
                }
            }
        }
    }
}'

# Deploy prediction log writer
curl --location --request POST 'localhost:8080/v1/projects/1/log_writers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "log_writer" : {
        "project_id" : 1,
        "name" : "my-prediction-log",
        "source" : {
            "type": "LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG",
            "prediction_log_source" : {
                "model_id": 1,
                "model_name" : "my-model",
                "kafka" : {
                    "brokers": "kafka.mlp.svc.cluster.local",
                    "topic" : "my-prediction-log"
                }
            }
        }
    }
}'

```

You can check all available commands using

```
make
```
