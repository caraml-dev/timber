package config

import (
	"fmt"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"

	common_config "github.com/caraml-dev/observation-service/common/config"
)

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
	LoggerNoopConsumer  ObservationLoggerConsumerKind = ""
	LoggerKafkaConsumer ObservationLoggerConsumerKind = "kafka"
)

type LogConsumerConfig struct {
	Kind ObservationLoggerConsumerKind `default:""`

	KafkaConfig *KafkaConfig
}

type KafkaConfig struct {
	Brokers          string
	Topic            string
	MaxMessageBytes  int    `default:"1048588"`
	CompressionType  string `default:"none"`
	ConnectTimeoutMS int    `default:"1000"`
}

// ObservationLoggerProducerKind captures the producer config for flushing Observation Service logs
type ObservationLoggerProducerKind = string

const (
	LoggerNoopProducer   ObservationLoggerProducerKind = ""
	LoggerStdOutProducer ObservationLoggerProducerKind = "stdout"
	LoggerKafkaProducer  ObservationLoggerProducerKind = "kafka"
)

type LogProducerConfig struct {
	Kind                 ObservationLoggerProducerKind `default:""`
	QueueLength          int                           `default:"100"`
	FlushIntervalSeconds int                           `default:"1"`

	KafkaConfig *KafkaConfig
}

// ListenAddress returns the Observation API port
func (c *Config) ListenAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func Load(filepaths ...string) (*Config, error) {
	var cfg Config
	err := common_config.ParseConfig(&cfg, filepaths)
	if err != nil {
		return nil, fmt.Errorf("failed to update viper config: %s", err)
	}

	return &cfg, nil
}
