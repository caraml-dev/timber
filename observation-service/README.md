# Observation Service

Observation service is a service for collecting the ground truth or result made from prediction process. 

## Getting Started

### Requirements

- Golang 1.18
- Docker
- grpcurl

### Build Fluentd

Observation service relies on fluentd to write collected observation into BQ

```bash
# Build Fluentd image
make build-fluentd-image
```

### Setup Dependencies

Fluentd BQ logging requires a GCP service account to be present at infra/local directory to work properly. 
Without it, logs will only be flushed to stdout.

```bash
# Starts Kafka, Fluentd services
make dev-env
```

### Run Observation Service

Run observation service locally using example configuration in [config/example.yaml](config/example.yaml)

```bash
make run
```

### Sample Requests

Send gRPC request

```bash
# LogObservations API
grpcurl -plaintext -d '{ "observations": [{"prediction_id": "1", "row_id": "1", "target_name": "accepted", "observation_values": [], "observation_context": []}] }' \
  localhost:9001 caraml.upi.v1.ObservationService/LogObservations

# Health checks
grpcurl -plaintext -d '{ "service": "caraml.upi.v1.ObservationService" }' \
  localhost:9001 grpc.health.v1.Health/Watch
grpcurl -plaintext -d '{ "service": "caraml.upi.v1.ObservationService" }' \
  localhost:9001 grpc.health.v1.Health/Check

# Service description
grpcurl -plaintext localhost:9001 describe
```


### Run Tests
To run tests:

```bash
make test
```

Note that this will attempt to set up a Kafka Broker using docker, to run integration tests.
