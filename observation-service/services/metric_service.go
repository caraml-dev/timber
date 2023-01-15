package services

import (
	"errors"
	"strconv"
	"time"

	"github.com/gojek/mlp/api/pkg/instrumentation/metrics"

	commonconfig "github.com/caraml-dev/timber/common/config"
	"github.com/caraml-dev/timber/common/log"
	"github.com/caraml-dev/timber/observation-service/config"
	"github.com/caraml-dev/timber/observation-service/monitoring"
)

// MetricService captures the exposed methods for logging performance/health metrics
type MetricService interface {
	LogLatencyHistogram(begin time.Time, statusCode int, loggingMetric metrics.MetricName)
	LogRequestCount(statusCode int, loggingMetric metrics.MetricName)

	GetLabels(labels map[string]string) map[string]string
}

type metricService struct {
	Kind             config.MetricSinkKind
	DeploymentConfig commonconfig.DeploymentConfig
}

// NewMetricService initializes a metricService struct
func NewMetricService(deploymentCfg commonconfig.DeploymentConfig, monitoringCfg config.MonitoringConfig) (MetricService, error) {
	switch monitoringCfg.Kind {
	case config.NoopMetricSink:
	case config.PrometheusMetricSink:
		// Init metrics collector
		histogramMap := monitoring.GetHistogramMap()
		counterMap := monitoring.GetCounterMap()
		err := metrics.InitPrometheusMetricsCollector(monitoring.GaugeMap, histogramMap, counterMap)
		if err != nil {
			return nil, errors.New("failed to initialize Prometheus-based MetricService")
		}
	}

	svc := &metricService{
		Kind:             monitoringCfg.Kind,
		DeploymentConfig: deploymentCfg,
	}

	return svc, nil
}

// LogLatencyHistogram tracks histogram metrics
func (ms *metricService) LogLatencyHistogram(begin time.Time, statusCode int, loggingMetric metrics.MetricName) {
	baseLabels := map[string]string{
		"response_code": strconv.Itoa(statusCode),
	}
	labels := ms.GetLabels(baseLabels)

	var err error
	switch ms.Kind {
	case config.NoopMetricSink:
	case config.PrometheusMetricSink:
		switch loggingMetric {
		case monitoring.RequestDurationMs:
			err = metrics.Glob().MeasureDurationMsSince(
				monitoring.RequestDurationMs, begin, labels,
			)
		case monitoring.FlushDurationMs:
			err = metrics.Glob().MeasureDurationMsSince(
				monitoring.FlushDurationMs, begin, labels,
			)
		}
		if err != nil {
			log.Errorf("error while logging %s metrics (latency): %s", loggingMetric, err)
		}
	}
}

// LogLatencyHistogram tracks count metrics
func (ms *metricService) LogRequestCount(statusCode int, loggingMetric metrics.MetricName) {
	baseLabels := map[string]string{
		"response_code": strconv.Itoa(statusCode),
	}
	labels := ms.GetLabels(baseLabels)

	var err error
	switch ms.Kind {
	case config.NoopMetricSink:
	case config.PrometheusMetricSink:
		switch loggingMetric {
		case monitoring.ReadCount:
			err = metrics.Glob().Inc(
				monitoring.ReadCount, labels,
			)
		case monitoring.FlushCount:
			err = metrics.Glob().Inc(
				monitoring.FlushCount, labels,
			)
		case monitoring.FlushObservationCount:
			err = metrics.Glob().Inc(
				monitoring.FlushObservationCount, labels,
			)
		}
		if err != nil {
			log.Errorf("error while logging metrics (request_count): %s", err)
		}
	}
}

// GetLabels adds default system labels to provided labels
func (ms *metricService) GetLabels(
	labels map[string]string,
) map[string]string {
	defaultLabels := map[string]string{
		"project_name": ms.DeploymentConfig.ProjectName,
		"service_name": ms.DeploymentConfig.ServiceName,
	}
	for k, v := range defaultLabels {
		labels[k] = v
	}

	return labels
}
