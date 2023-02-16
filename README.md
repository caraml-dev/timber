# Timber

[![License](https://img.shields.io/badge/License-Apache%202.0-blue)](https://github.com/caraml-dev/timber/blob/master/LICENSE)

## Overview

Timber is CaraML component responsible for log collection, model monitoring, and dataset creation. It consists of several components:

### Dataset Service

Dataset service is the orchestrator within Timber ecosystem. It manages deployment of dataset service and observation service.

### Observation Service

Observation Service provides an interface for reporting observations from Client-owned systems back to CaraML platform. The logged observations will be used for evaluating the effectiveness of a given prediction instance.

### Log Writer

Log writer is the component responsible for collecting prediction log and router log.


## Contributing

Timber is still under active development. Please have a look at our contributing and development guides if you want to contribute to the project:

- [Contribution Process for Timber](https://github.com/caraml-dev/timber/blob/main/CONTRIBUTING.md)

## Project Structure

```                               
├── common                              # Common code to be reused both in Dataset service and Observation service
├── dataset-service                     # Dataset service implementation.
├── images                              # Docker images
│   └── fluentd                         # FluentD docker image for logwriter and observation service
|   └── observation-service             # Dockerfile for building Observation service
├── infra                               # Infrastructure setup for testing and deployment.
├── logwriter                           # Logwriter implementation.
│   └── fluent-plugin-upi-logs          # FluentD plugin for parsing Universal Prediction Interface protobuf.
├── observation-service                 # Observation service implementation.
├── scripts                             # Miscellaneous scripts.
|   └── vertagen                        # Helper for version string generation.
├── tests                               # Integration and end to end test suites.
├── CONTRIBUTING.md                     # Contributing guide.
├── LICENSE
├── Makefile
├── README.md    
```
