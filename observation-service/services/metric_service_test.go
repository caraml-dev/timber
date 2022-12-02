package services

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"

	"github.com/caraml-dev/observation-service/observation-service/config"
	"github.com/caraml-dev/observation-service/observation-service/internal/testutils"
	"github.com/caraml-dev/observation-service/observation-service/monitoring"
)

type MetricServiceTestSuite struct {
	suite.Suite
	MetricService
	cfg config.Config
}

func (s *MetricServiceTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up MetricServiceTestSuite")

	var err error
	s.cfg = config.Config{
		DeploymentConfig: config.DeploymentConfig{
			ProjectName: "default",
			ServiceName: "observation-service",
		},
		MonitoringConfig: config.MonitoringConfig{
			Kind: config.PrometheusMetricSink,
		},
	}
	s.MetricService, err = NewMetricService(s.cfg.DeploymentConfig, s.cfg.MonitoringConfig)
	if err != nil {
		s.Suite.T().Log("failed to initialize MetricService")
	}
}

func (s *MetricServiceTestSuite) TearDownSuite() {
	s.Suite.T().Log("Cleaning up MetricServiceTestSuite")
}

func TestMetricService(t *testing.T) {
	suite.Run(t, new(MetricServiceTestSuite))
}

func (s *MetricServiceTestSuite) TestGetLabels() {
	extraLabels := map[string]string{
		"type": "test",
	}
	labels := s.MetricService.GetLabels(extraLabels)
	expectedLabels := map[string]string{
		"type":         "test",
		"project_name": "default",
		"service_name": "observation-service",
	}

	s.Suite.Require().Equal(expectedLabels, labels)
}

func (s *MetricServiceTestSuite) TestLogLatencyHistogram() {
	statusCode := http.StatusOK

	stdout := testutils.CaptureStderrLogs(func() {
		s.MetricService.LogLatencyHistogram(time.Now(), statusCode, monitoring.RequestDurationMs)
	})
	s.Suite.Require().Equal("", stdout)
}

func (s *MetricServiceTestSuite) TestLogRequestCount() {
	statusCode := http.StatusInternalServerError
	stdout := testutils.CaptureStderrLogs(func() {
		s.MetricService.LogRequestCount(statusCode, monitoring.FlushCount)
	})
	s.Suite.Require().Equal("", stdout)
}
