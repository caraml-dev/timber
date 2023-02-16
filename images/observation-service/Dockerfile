ARG SERVICE_DIRECTORY=observation-service
# ============================================================
# Build stage 1: Build API
# ============================================================
FROM golang:1.18-alpine as go-builder
ARG API_BIN_NAME=observation-service
ARG SERVICE_DIRECTORY

RUN apk update && apk add build-base

COPY . /app
WORKDIR /app/${SERVICE_DIRECTORY}

# Build Observation Service binary
RUN go mod tidy
RUN go build \
    -tags musl \
    -o bin/${API_BIN_NAME} \
    -v ./cmd/${SERVICE_DIRECTORY}

# ============================================================
# Build stage 2: Copy binary
# ============================================================
FROM alpine:latest
ARG SERVICE_DIRECTORY

RUN mkdir -p /opt/observation_service

COPY --from=go-builder /app/${SERVICE_DIRECTORY}/bin/* /opt/observation_service/

WORKDIR /opt/observation_service
ENTRYPOINT [ "./observation-service" ]
