### Example of Helm Charts testing

Timber Observation Service and Fluentd Helm Charts are available at https://github.com/caraml-dev/helm-charts/tree/main/charts

#### Pre-requisite  
- [Minikube](https://minikube.sigs.k8s.io/docs/start/)
- Helm 3   
- Caraml repo added to helm  
`helm add caraml https://caraml-dev.github.io/helm-charts`

#### Steps

1. Start minikube, chart was tested with kubernetes version v1.22.15
```
minikube start --kubernetes-version=v1.22.15
```
2. Install Kafka (If testing with Kafka)
```
helm install kafka bitnami/kafka --values={kafka-values.yaml}
```
Sample kafka values.yaml. Persistency are disabled to prevent cache issue upon restarting kafka
```
zookeeper:
  persistence:
    enabled: false

persistence:
  enabled: false

autoCreateTopicsEnable: true

resources:
  requests:
    cpu: 500m
    memory: 1Gi
```
3. Install Logwriter that reads from Kafka and save to BQ. For this template, fluentd config should be static, with configurable in extraEnvs.
```
helm install logwriter caraml/timber-fluentd --values={log-writer-values}.yaml
```
Sample values.yaml. GCP cred, BQ and Kafka data required under extraEnvs
```
gcpServiceAccount:
  credentialsData: #Fill in encoded base64 service-account.json 

extraEnvs: 
  - name: FLUENTD_WORKER_COUNT
    value: "1"
  - name: FLUENTD_LOG_LEVEL
    value: debug
  - name: FLUENTD_BUFFER_LIMIT
    value: 1g
  - name: FLUENTD_FLUSH_INTERVAL_SECONDS
    value: "30"
  - name: FLUENTD_LOG_PATH
    value: /fluentd/cache/log
  - name: FLUENTD_TAG
    value: observation-service.log
  - name: FLUENTD_GCP_JSON_KEY_PATH
    value: /etc/gcp_service_account/service-account.json
  - name: FLUENTD_GCP_PROJECT
    value: #Fill in GCP PROJECT
  - name: FLUENTD_BQ_DATASET
    value: #Fill in BQ Dataset
  - name: FLUENTD_BQ_TABLE
    value: #Fill in BQ Table
  - name: FLUENTD_KAFKA_BROKER
    value: kafka.default.svc.cluster.local #Modify Accordingly
  - name: FLUENTD_KAFKA_TOPIC
    value: quickstart #Modify Accordingly
  - name: FLUENTD_KAFKA_PROTO_CLASS_NAME
    value: caraml.upi.v1.PredictionLog #Prediction or Observation logs supported
    
fluentdConfig: |-
  # Set fluentd log level to error
  <system>
    log_level "#{ENV['FLUENTD_LOG_LEVEL']}"
    workers "#{ENV['FLUENTD_WORKER_COUNT']}"
  </system>

  <source>
    @type kafka

    brokers "#{ENV['FLUENTD_KAFKA_BROKER']}" #broker:9092
    topics "#{ENV['FLUENTD_KAFKA_TOPIC']}" #quickstart

    format upi_logs
    class_name "#{ENV['FLUENTD_KAFKA_PROTO_CLASS_NAME']}" #"caraml.upi.v1.PredictionLog"
  </source>

  # Buffer and output to multiple sinks
  <match "#{ENV['FLUENTD_KAFKA_TOPIC']}">
    @type copy
    <store>
      @type stdout
    </store>
    <store>
      @type bigquery_load

      <buffer>
        @type file

        path "#{ENV['FLUENTD_LOG_PATH']}"
        timekey_use_utc
        
        flush_at_shutdown true
        flush_mode interval
        flush_interval "#{ENV['FLUENTD_FLUSH_INTERVAL_SECONDS']}"
        retry_max_times 3

        chunk_limit_size 1g
        compress gzip
        total_limit_size "#{ENV['FLUENTD_BUFFER_LIMIT']}"

        delayed_commit_timeout 150
        disable_chunk_backup true
      </buffer>

      # Authenticate with BigQuery using a json key
      auth_method json_key
      json_key "#{ENV['FLUENTD_GCP_JSON_KEY_PATH']}"
      project "#{ENV['FLUENTD_GCP_PROJECT']}"
      dataset "#{ENV['FLUENTD_BQ_DATASET']}"
      table "#{ENV['FLUENTD_BQ_TABLE']}"
      fetch_schema true
    </store>
  </match>
  ```

4. Install Observation Service
```
helm install logwriter caraml/timber-observation-service --values={observation-service-values}.yaml
```
Sample values.yaml. GCP cred, BQ and Kafka data required. For this template, fluentd config should be static, with configurable in extraEnvs and observationService.apiConfig.
```
observationService:
  apiConfig:
    logConsumerConfig:
      Kind: kafka
      KafkaConfig:
        Brokers: kafka.default.svc.cluster.local # Modify accordingly
        Topic: test-topic # Modify accordingly
    LogProducerConfig:
      FlushIntervalSeconds: 10
      Kind: fluentd
      FluentdConfig:
        Host: obs-timber-observation-service-fluentd # Modify accordingly, convention {RELEASE-NAME}-timber-observation-service-fluent
        Port: 24224
        Kind: bq
        Tag: observation-service.log 
        BQConfig:
          Project: #Fill in GCP PROJECT
          Dataset: #Fill in BQ DATASET
          Table: #Fill in BQ Table

fluentd:
  enabled: true

  gcpServiceAccount:
    credentialsData: #Fill in encoded base64 service-account.json 

  extraEnvs: 
    - name: FLUENTD_WORKER_COUNT
      value: "1"
    - name: FLUENTD_LOG_LEVEL
      value: debug
    - name: FLUENTD_BUFFER_LIMIT
      value: 1g
    - name: FLUENTD_FLUSH_INTERVAL_SECONDS
      value: "30"
    - name: FLUENTD_LOG_PATH
      value: /fluentd/cache/log
    - name: FLUENTD_TAG
      value: observation-service.log
    - name: FLUENTD_GCP_JSON_KEY_PATH
      value: /etc/gcp_service_account/service-account.json
    - name: FLUENTD_GCP_PROJECT
      value: #Fill in GCP PROJECT
    - name: FLUENTD_BQ_DATASET
      value: #Fill in BQ DATASET
    - name: FLUENTD_BQ_TABLE
      value: #Fill in BQ Table

  # -- Fluentd config
  fluentdConfig: |-
    # Set fluentd log level to error
    <system>
      log_level "#{ENV['FLUENTD_LOG_LEVEL']}"
      workers "#{ENV['FLUENTD_WORKER_COUNT']}"
    </system>

    # Accept HTTP input
    <source>
      @type http
      port 9880
      bind 0.0.0.0
      body_size_limit 32m
      keepalive_timeout 10s
    </source>

    # Accept events on tcp socket
    <source>
      @type forward
      port 24224
      bind 0.0.0.0
    </source>

    # Buffer and output to multiple sinks
    <match "#{ENV['FLUENTD_TAG']}">
      @type copy
      <store>
        @type stdout
      </store>
      <store>
        @type bigquery_load

        <buffer>
          @type file

          path "#{ENV['FLUENTD_LOG_PATH']}"
          timekey_use_utc
          
          flush_at_shutdown true
          flush_mode interval
          flush_interval "#{ENV['FLUENTD_FLUSH_INTERVAL_SECONDS']}"
          retry_max_times 3

          chunk_limit_size 1g
          compress gzip
          total_limit_size "#{ENV['FLUENTD_BUFFER_LIMIT']}"

          delayed_commit_timeout 150
          disable_chunk_backup true
        </buffer>

        # Authenticate with BigQuery using a json key
        auth_method json_key
        json_key "#{ENV['FLUENTD_GCP_JSON_KEY_PATH']}"
        project "#{ENV['FLUENTD_GCP_PROJECT']}"
        dataset "#{ENV['FLUENTD_BQ_DATASET']}"
        table "#{ENV['FLUENTD_BQ_TABLE']}"
        fetch_schema true
      </store>
    </match>
```

5. Write and deploy Kafka Producer. The container is deployed on minikube as there are much issues with using kafka client outside minikube to connect to the kafka deployed there.

Create you own go project
```go
package main

import (
	"fmt"
	"os"
	
	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	kafkaBrokers := "kafka.default.svc.cluster.local"
	topic := "test-topic"

	// Create Producer instance
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %s", err)
		os.Exit(1)
	}
	
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)
	
	//Generate your own prediction/observation/router log message
	message := &upiv1.ObservationLog{PredictionId: "1", ObservationTimestamp: timestamppb.Now()} 
	valueBytes, err := proto.Marshal(message)
	if err != nil {
		fmt.Println(fmt.Errorf("unable to marshal log entry value, %s", err))
	}
	
	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny},
		Value: valueBytes,
	}, deliveryChan)
	if err != nil {
		fmt.Println(err)
	}

	// Get delivery response
	event := <-deliveryChan
	msg := event.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		err = fmt.Errorf("Delivery failed: %v\n", msg.TopicPartition.Error)
		fmt.Println(err)
	}

	producer.Close()
	fmt.Println("Message sent!")
	os.Exit(0)
}
```

Create a dockerfile
```
FROM golang:alpine

RUN apk update && apk add build-base

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN go mod tidy

RUN go build \
    -tags musl \
    -o main .

CMD ["/app/main"]
```
Before building the image, run the below cli for minikube to be able to access local docker registry
```
 eval $(minikube docker-env)
```

Build the image
```
docker build . -t log-producer
```

Create the manifest and apply it. Note that imagePullPolicy must be set to never, so that it will pull from local registry
```
apiVersion: v1
kind: Pod
metadata:
  name: log-producer
spec:
  containers:
  - name: log-producer
    image: log-producer
    imagePullPolicy: Never
    ports:
    - containerPort: 80
```

Logs should be persisted in BQ.