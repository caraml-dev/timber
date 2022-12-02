package integration

import (
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	upiv1 "github.com/caraml-dev/universal-prediction-interface/gen/go/grpc/caraml/upi/v1"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	testcontainers "github.com/testcontainers/testcontainers-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/caraml-dev/observation-service/observation-service/log"
	"github.com/caraml-dev/observation-service/observation-service/server"
)

type KafkaTopic string

const (
	KafkaBrokers          = "localhost:9092"
	KafkaMaxMessageBytes  = 1048588
	KafkaCompressionType  = "none"
	ObservationServerPort = 9001
)

type KafkaKey struct {
	EventTimestamp *timestamppb.Timestamp
}

type ObservationServiceTestSuite struct {
	suite.Suite

	observationServiceServer *server.Server

	terminationChannel chan bool
	ctx                context.Context
	kafka              testcontainers.DockerCompose
}

func produceToKafka(timestamp *timestamppb.Timestamp) {
	deliveryChan := make(chan kafka.Event, 1)
	defer close(deliveryChan)

	kafkaTopic := "integration-test-source"
	producer, err := kafka.NewProducer(
		&kafka.ConfigMap{
			"bootstrap.servers": KafkaBrokers,
			"message.max.bytes": KafkaMaxMessageBytes,
			"compression.type":  KafkaCompressionType,
		},
	)
	if err != nil {
		log.Glob().Error(err)
	}

	// Generate record
	// Create the Kafka key
	key := &upiv1.ObservationLogKey{
		ObservationBatchId: uuid.New().String(),
		PredictionId:       "integration-test-prediction-id",
		RowId:              "integration-test-row-id",
	}
	// Marshal the key
	keyBytes := []byte(fmt.Sprintf("%v", key))
	if err != nil {
		log.Glob().Errorf("unable to marshal log entry key, %s", err)
	}

	// Create the Kafka message
	message := &upiv1.ObservationLog{
		PredictionId: "integration-test-prediction-id",
		RowId:        "integration-test-row-id",
		TargetName:   "target-name",
		ObservationValues: []*upiv1.Variable{
			{
				Name:        "integration-test-variable",
				Type:        upiv1.Type_TYPE_STRING,
				StringValue: "integration-test-variable-value",
			},
		},
		ObservationContext: []*upiv1.Variable{
			{
				Name:        "project",
				Type:        upiv1.Type_TYPE_STRING,
				StringValue: "integration-test",
			},
		},
		ObservationTimestamp: timestamp,
	}
	// Marshal the message
	valueBytes, err := proto.Marshal(message)
	if err != nil {
		log.Glob().Errorf("unable to marshal log entry value, %s", err)
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &kafkaTopic,
			Partition: kafka.PartitionAny},
		Value: valueBytes,
		Key:   keyBytes,
	}, deliveryChan)
	if err != nil {
		log.Glob().Error(err)
	}

	// Get delivery response
	event := <-deliveryChan
	msg := event.(*kafka.Message)
	if msg.TopicPartition.Error != nil {
		err = fmt.Errorf("delivery failed: %v", msg.TopicPartition.Error)
		log.Glob().Error(err)
	}
	producer.Close()
}

func setupObservationService() (chan bool, *server.Server) {
	observationServer, err := server.NewServer([]string{"test.yaml"})
	if err != nil {
		log.Glob().Panicf("fail to instantiate observation service server: %s", err.Error())
	}

	c := make(chan bool, 1)

	go func() {
		for {
			select {
			case <-c:
				close(c)
				return
			default:
				observationServer.Start()
			}
		}
	}()

	return c, observationServer
}

func (suite *ObservationServiceTestSuite) SetupSuite() {
	// Docker compose file copied from official confluentinc repository.
	// See: https://github.com/confluentinc/cp-all-in-one/blob/7.3.0-post/cp-all-in-one-kraft/docker-compose.yml
	composeFilePaths := []string{"docker-compose/docker-compose.yaml"}
	kafka := testcontainers.NewLocalDockerCompose(composeFilePaths, "")
	execError := kafka.
		WithCommand([]string{"up", "-d"}).
		Invoke()
	err := execError.Error
	if err != nil {
		panic(err)
	}
	suite.kafka = kafka

	c, observationServer := setupObservationService()
	waitForServerToListen := func() bool {
		conn, err := net.Dial("tcp", net.JoinHostPort("", strconv.Itoa(ObservationServerPort)))
		if conn != nil {
			conn.Close()
		}
		return err == nil
	}
	suite.Require().Eventuallyf(waitForServerToListen, 3*time.Second, 1*time.Second, "observation service failed to start")

	suite.terminationChannel = c
	suite.observationServiceServer = observationServer

	suite.ctx = context.Background()
}

func TestObservationServiceTestSuite(t *testing.T) {
	fmt.Printf("Start time: %s", time.Now())
	suite.Run(t, new(ObservationServiceTestSuite))
}

func (suite *ObservationServiceTestSuite) TearDownSuite() {
	fmt.Printf("End time: %s", time.Now())
	suite.terminationChannel <- true

	_ = suite.kafka.Down()
}

func (suite *ObservationServiceTestSuite) TestLogKafkaSourceToKafkaSink() {
	currentTime := time.Now()
	timestamp := &timestamppb.Timestamp{Seconds: currentTime.Unix()}
	// Produce to Kafka Source Topic
	produceToKafka(timestamp)

	// Read from Kafka Sink Topic
	consumer, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers":    "localhost:9092",
			"group.id":             "integration-test-group",
			"default.topic.config": kafka.ConfigMap{"auto.offset.reset": "earliest"},
		})
	suite.Require().NoError(err)
	err = consumer.Subscribe("integration-test-sink", nil)
	suite.Require().NoError(err)

	startTime := time.Now()
	var observationLogKeyJSON string
	var observationLogEntryJSON string
W:
	for {
		ev := consumer.Poll(1000)
		switch e := ev.(type) {
		case *kafka.Message:
			keyValue := e.Key
			observationLogKeyJSON = string(keyValue)
			fmt.Printf("%% Key on %s:\n%s\n",
				e.TopicPartition, string(keyValue))
			messageValue := e.Value
			fmt.Printf("%% Message on %s:\n%s\n",
				e.TopicPartition, string(messageValue))
			observationLogEntryJSON = string(messageValue)
			break W
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
			break W
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			break W
		default:
			if time.Since(startTime).Seconds() > 30 {
				break W
			}
		}
	}
	consumer.Close()
	suite.Require().NoError(err)
	expectedObservationKey := "\"prediction_id\":\"integration-test-prediction-id\",\"row_id\":\"integration-test-row-id\""
	expectedObservation := fmt.Sprintf(
		"{\"prediction_id\":\"integration-test-prediction-id\",\"row_id\":\"integration-test-row-id\","+
			"\"target_name\":\"target-name\",\"observation_values\":[{\"name\":\"integration-test-variable\","+
			"\"type\":\"TYPE_STRING\",\"string_value\":\"integration-test-variable-value\"}],"+
			"\"observation_context\":[{\"name\":\"project\",\"type\":\"TYPE_STRING\",\"string_value\":\"integration-test\"}],"+
			"\"observation_timestamp\":\"%s\"}",
		currentTime.UTC().Format(time.RFC3339),
	)
	// strings.ReplaceAll is required to make test output deterministic
	// https://developers.google.com/protocol-buffers/docs/reference/go/faq#unstable-json
	suite.Require().Equal(strings.ReplaceAll(observationLogEntryJSON, " ", ""), expectedObservation)
	suite.Require().Contains(strings.ReplaceAll(observationLogKeyJSON, " ", ""), expectedObservationKey)
}
