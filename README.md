# observation-service

[![License](https://img.shields.io/badge/License-Apache%202.0-blue)](https://github.com/caraml-dev/observation-service/blob/master/LICENSE)

## Overview

Observation Service provides an interface for reporting observations from Client-owned systems back to CaraML platform. The logged observations will be used for evaluating the effectiveness of a given prediction instance.

## Development Environment

### Quick Start

#### a. Build dependencies
```bash
# Build Fluentd image
make build-fluentd-image
```

#### b. Setup dependencies

Fluentd BQ logging requires a GCP service account to be present at infra/local directory to work properly. Without it, logs will only be flushed to stdout.

```bash
# Starts Kafka, Fluentd services
make dependency-services
```

#### c. Setup Observation Service
```bash
make observation-service
```

#### d. Sample Requests

`GRPC requests`
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

## Contributing

Observation Service is still under active development. Please have a look at our contributing and development guides if you want to contribute to the project:

- [Contribution Process for Observation Service](https://github.com/caraml-dev/observation-service/blob/main/CONTRIBUTING.md)
