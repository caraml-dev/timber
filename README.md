# timber

[![License](https://img.shields.io/badge/License-Apache%202.0-blue)](https://github.com/caraml-dev/timber/blob/master/LICENSE)

## Overview

Timber comprises of multiple services, central to logging in CaraML ecosystem.

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
cd observation-service && make observation-service
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

Timber is still under active development. Please have a look at our contributing and development guides if you want to contribute to the project:

- [Contribution Process for Timber](https://github.com/caraml-dev/timber/blob/main/CONTRIBUTING.md)

## Project Structure

```                               
├── common                          # Common code to be reused both in Dataset service and Observation service
├── dataset-service                 # Dataset service implementation.
├── images                          # Docker images
│   └── fluentd                     # FluentD docker image for logwriter and observation service
├── infra                           # Infrastructure setup for testing and deployment.
├── logwriter                       # Logwriter implementation.
│   └── fluent-plugin-upi-logs      # FluentD plugin for parsing Universal Prediction Interface protobuf.
├── observation-service             # Observation service implementation.
├── scripts                         # Miscellaneous scripts.
|   └── vertagen                    # Helper for version string generation.
├── tests                           # Integration and end to end test suites.
├── CONTRIBUTING.md                 # Contributing guide.
├── Dockerfile.observation_service
├── LICENSE
├── Makefile
├── README.md    
```
