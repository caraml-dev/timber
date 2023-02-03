package services

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	commonconfig "github.com/caraml-dev/timber/common/config"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/models"
	os "github.com/caraml-dev/timber/observation-service/config"
)

type ObservationServiceTestSuite struct {
	suite.Suite
	ObservationService
}

func (s *ObservationServiceTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up ObservationServiceTestSuite")
}

func (s *ObservationServiceTestSuite) TearDownSuite() {
	s.Suite.T().Log("Cleaning up ObservationServiceTestSuite")
}

func TestObservationService(t *testing.T) {
	suite.Run(t, new(ObservationServiceTestSuite))
}

func (s *ObservationServiceTestSuite) TestSetDefaultHelmValues() {
	helmValues := &models.ObservationServiceHelmValues{}
	gcpProject := "test-gcp-project"
	caramlProjectName := "test-caraml-project"
	serviceName := "test-pricing"
	deploymentConfig := commonconfig.DeploymentConfig{
		EnvironmentType: "test",
		ProjectName:     "",
		ServiceName:     "dataset-service",
		LogLevel:        commonconfig.InfoLevel,
		MaxGoRoutines:   200,
	}
	imageTag := "v0.0.1"
	observationServiceConfig := config.ObservationServiceConfig{
		GCPProject:                 gcpProject,
		ObservationServiceImageTag: imageTag,
		FluentdImageTag:            imageTag,
	}

	actualHelmValues := setDefaultValues(
		helmValues,
		gcpProject,
		caramlProjectName,
		serviceName,
		deploymentConfig,
		observationServiceConfig,
	)

	// Check some hardcoded default values and configurable (empty and override) values
	s.Suite.Assert().Equal(9001, actualHelmValues.ObservationServiceConfig.ApiConfig.Port)
	s.Suite.Assert().Equal(imageTag, actualHelmValues.ObservationServiceConfig.Image.Tag)
	s.Suite.Assert().Equal(caramlProjectName, actualHelmValues.ObservationServiceConfig.ApiConfig.DeploymentConfig.ProjectName)
	s.Suite.Assert().Equal("observation-service-test-pricing", actualHelmValues.ObservationServiceConfig.ApiConfig.DeploymentConfig.ServiceName)
}

func (s *ObservationServiceTestSuite) TestSetConsumerConfigValues() {
	helmValues := &models.ObservationServiceHelmValues{}
	eagerDeploymentConfig := &timberv1.ObservationServiceConfig{
		Source: &timberv1.ObservationServiceDataSource{
			Type: timberv1.ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER,
		},
	}

	// Test Eager Config
	_, err := setLogConsumerConfig(helmValues, eagerDeploymentConfig)
	s.Suite.Require().Error(fmt.Errorf("source type (eager) is currently unsupported"), err)

	// Test Kafka Config
	kafkaDeploymentConfig := &timberv1.ObservationServiceConfig{
		Source: &timberv1.ObservationServiceDataSource{
			Type: timberv1.ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA,
		},
	}
	actual, err := setLogConsumerConfig(helmValues, kafkaDeploymentConfig)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal("kafka", actual.ObservationServiceConfig.ApiConfig.LogConsumerConfig.Kind)
	s.Suite.Assert().Equal(
		models.NewKafkaConfig(kafkaDeploymentConfig.GetSource().GetKafkaConfig()),
		actual.ObservationServiceConfig.ApiConfig.LogConsumerConfig.KafkaConfig,
	)
}

func (s *ObservationServiceTestSuite) TestSetProducerConfigValues() {
	helmValues := &models.ObservationServiceHelmValues{}
	gcpProject := "test-gcp-project"
	caramlProject := "pricing-test"
	fluentdTag := "observation-service"
	observationServiceConfig := config.ObservationServiceConfig{}
	fluentdDeploymentConfig := &timberv1.ObservationServiceConfig{
		Sink: &timberv1.ObservationServiceDataSink{
			Type: timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD,
			FluentdConfig: &timberv1.FluentdConfig{
				Tag: fluentdTag,
			},
		},
	}

	// Test Fluentd Config
	expectedFluentDConfig := &os.FluentdConfig{
		Tag:  fluentdTag,
		Host: fmt.Sprintf("observation-service-fluentd.%s.svc.cluster.local", caramlProject),
		Port: 24224,
		Kind: os.LoggerBQSinkFluentdProducer,
		BQConfig: &os.BQConfig{
			Project: gcpProject,
			Dataset: caramlProject,
			Table:   fmt.Sprintf("%s_observation_log", caramlProject),
		},
	}
	actual, err := setLogProducerConfig(helmValues, fluentdDeploymentConfig, observationServiceConfig, gcpProject, caramlProject)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal("fluentd", actual.ObservationServiceConfig.ApiConfig.LogProducerConfig.Kind)
	s.Suite.Assert().Equal(expectedFluentDConfig, actual.ObservationServiceConfig.ApiConfig.LogProducerConfig.FluentdConfig)

	// Test Kafka Config
	kafkaDeploymentConfig := &timberv1.ObservationServiceConfig{
		Sink: &timberv1.ObservationServiceDataSink{
			Type: timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA,
		},
	}
	actual, err = setLogProducerConfig(helmValues, kafkaDeploymentConfig, observationServiceConfig, gcpProject, caramlProject)
	s.Suite.Require().NoError(err)
	s.Suite.Assert().Equal("kafka", actual.ObservationServiceConfig.ApiConfig.LogProducerConfig.Kind)
	s.Suite.Assert().Equal(
		models.NewKafkaConfig(kafkaDeploymentConfig.GetSink().GetKafkaConfig()),
		actual.ObservationServiceConfig.ApiConfig.LogProducerConfig.KafkaConfig,
	)
}
