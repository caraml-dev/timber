package config

import (
	"fmt"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"

	common_config "github.com/caraml-dev/observation-service/common/config"
)

// Config captures the config related to starting Observation Service
type Config struct {
	Port int `envconfig:"PORT" default:"9001"`

	DeploymentConfig  DeploymentConfig
	NewRelicConfig    newrelic.Config
	SentryConfig      sentry.Config
	LogConsumerConfig LogConsumerConfig
	LogProducerConfig LogProducerConfig
}

// DeploymentConfig captures the config related to the deployment of Observation Service
type DeploymentConfig struct {
	EnvironmentType string `default:"local"`
}

// ObservationLoggerConsumerKind captures the consumer config for reading Observation Service logs
type ObservationLoggerConsumerKind = string

const (
	// LoggerNoopConsumer is a No-Op ObservationLog Consumer
	LoggerNoopConsumer ObservationLoggerConsumerKind = ""
	// LoggerKafkaConsumer is a Kafka ObservationLog Consumer
	LoggerKafkaConsumer ObservationLoggerConsumerKind = "kafka"
)

// LogConsumerConfig captures the config related to consuming ObservationLog via a background process
type LogConsumerConfig struct {
	// The type of Data Source for Observation logs
	Kind ObservationLoggerConsumerKind `default:""`

	// KafkaConfig captures the config related to initializing a Kafka Consumer
	KafkaConfig *KafkaConfig
}

// KafkaConfig captures all configurable parameters when configuring a Kafka Consumer and Producer
type KafkaConfig struct {
	// Kafka Brokers to connect to, comma-delimited, in the form of "<broker_host>:<broker_port>"
	Brokers string
	// Kafka Topic to produce to/consume from
	Topic string
	// Largest record batch size allowed by Kafka (after compression if compression is enabled)
	MaxMessageBytes int `default:"1048588"`
	// The compression type for all data generated by the Producer
	CompressionType string `default:"none"`
	// ConnectTimeoutMS is the maximum duration (ms) the Kafka Producer/Consumer will block for to get Metadata, before timing out
	ConnectTimeoutMS int `default:"1000"`
	// PollInterval is the maximum duration (ms) the Kafka Consumer will block for, before timing out
	PollInterval int `default:"1000"`
	// What to do when there is no initial offset in Kafka or if the current offset does not exist any more on the server
	AutoOffsetReset string `default:"latest"`
}

// ObservationLoggerProducerKind captures the producer config for flushing Observation Service logs
type ObservationLoggerProducerKind = string

const (
	// LoggerNoopProducer is a No-Op ObservationLog Producer
	LoggerNoopProducer ObservationLoggerProducerKind = ""
	// LoggerStdOutProducer is a Standard Output ObservationLog Producer
	LoggerStdOutProducer ObservationLoggerProducerKind = "stdout"
	// LoggerKafkaProducer is a Kafka ObservationLog Producer
	LoggerKafkaProducer ObservationLoggerProducerKind = "kafka"
)

// LogProducerConfig captures the config related to producing ObservationLog
type LogProducerConfig struct {
	// The type of Data Sink for Observation logs
	Kind ObservationLoggerProducerKind `default:""`
	// Maximum no. of Observation logs to be stored in-memory prior to flushing to Data sink
	QueueLength int `default:"100"`
	// Duration that specifies how often in-memory Observation logs should be flushed to Data sink
	FlushIntervalSeconds int `default:"1"`

	// KafkaConfig captures the config related to initializing a Kafka Producer
	KafkaConfig *KafkaConfig
}

// ListenAddress returns the Observation API port
func (c *Config) ListenAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

// Load parses multiple file configs specified via filepaths using Viper and returns a Config struct
func Load(filepaths ...string) (*Config, error) {
	var cfg Config
	err := common_config.ParseConfig(&cfg, filepaths)
	if err != nil {
		return nil, fmt.Errorf("failed to update viper config: %s", err)
	}

	return &cfg, nil
}
