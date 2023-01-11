package models

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	os "github.com/caraml-dev/timber/observation-service/config"
)

func TestNewFluentdConfigToApiSchema(t *testing.T) {
	testTag := "test-tag"
	projectName := "test-project"
	fluendConfig := &timberv1.FluentdConfig{
		Tag: testTag,
	}

	actual := NewFluentdConfig(fluendConfig, projectName)
	expected := &os.FluentdConfig{
		Tag:  testTag,
		Host: fmt.Sprintf("observation-service-fluentd.%s.svc.cluster.local", projectName),
		Port: 24224,
		Kind: os.LoggerBQSinkFluentdProducer,
	}
	assert.Equal(t, expected, actual)
}

func TestNewKafkaConfigToApiSchema(t *testing.T) {
	kafkaConfig := &timberv1.KafkaConfig{
		Brokers: "localhost:9092,localhost:9093,localhost:9094",
		Topic:   "test-topic",
	}

	actual := NewKafkaConfig(kafkaConfig)
	expected := &os.KafkaConfig{
		Brokers:          "localhost:9092,localhost:9093,localhost:9094",
		Topic:            "test-topic",
		AutoOffsetReset:  "latest",
		CompressionType:  "none",
		ConnectTimeoutMS: 1000,
		MaxMessageBytes:  1048588,
		PollInterval:     1000,
	}
	assert.Equal(t, expected, actual)
}
