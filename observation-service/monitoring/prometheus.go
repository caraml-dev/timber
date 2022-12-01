package monitoring

import (
	"github.com/gojek/mlp/api/pkg/instrumentation/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	// Namespace is the Prometheus Namespace in all metrics published by the Observation Service
	Namespace string = "mlp"
	// Subsystem is the Prometheus Subsystem in all metrics published by the Observation Service
	Subsystem string = "observation_service"
	// FlushDurationMs is the key to measure duration for flushing observations from channel to configured Data sink
	FlushDurationMs metrics.MetricName = "flush_duration_ms"
	// RequestDurationMs is the key to measure end-to-end duration for consuming and publishing observations
	RequestDurationMs metrics.MetricName = "request_duration_ms"
	// ReadCount is the key to measure no. of logs read from configured Data source
	ReadCount metrics.MetricName = "read_count"
	// FlushCount is the key to measure no. of flushes to configured Data sink
	FlushCount metrics.MetricName = "flush_count"
	// FlushObservationCount is the key to measure no. of observations flushed to configured Data sink
	FlushObservationCount metrics.MetricName = "flush_observation_count"
)

// requestLatencyBuckets defines the buckets used in the custom Histogram metrics
var requestLatencyBuckets = []float64{
	1000, 3000, 5000, 10000,
}

// GaugeMap configures gauge metrics
var GaugeMap = map[metrics.MetricName]metrics.PrometheusGaugeVec{}

// GetCounterMap configures counter metrics
func GetCounterMap() map[metrics.MetricName]metrics.PrometheusCounterVec {
	allLabels := []string{"project_name", "response_code"}

	counterMap := map[metrics.MetricName]metrics.PrometheusCounterVec{
		ReadCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Help:      "Counter for no. of reads from configured data source",
			Name:      string(ReadCount),
		},
			allLabels,
		),
		FlushCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Help:      "Counter for no. of flushes to configured data sink",
			Name:      string(FlushCount),
		},
			allLabels,
		),
		FlushObservationCount: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Help:      "Counter for no. of observations flushed to configured data sink",
			Name:      string(FlushObservationCount),
		},
			allLabels,
		),
	}

	return counterMap
}

// GetHistogramMap configures histogram metrics
func GetHistogramMap() map[metrics.MetricName]metrics.PrometheusHistogramVec {
	allLabels := []string{"project_name", "response_code"}

	histogramMap := map[metrics.MetricName]metrics.PrometheusHistogramVec{
		RequestDurationMs: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      string(RequestDurationMs),
			Help:      "Histogram for the runtime (in milliseconds) of observation log requests",
			Buckets:   requestLatencyBuckets,
		},
			allLabels,
		),
		FlushDurationMs: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: Namespace,
			Subsystem: Subsystem,
			Name:      string(FlushDurationMs),
			Help:      "Histogram for the runtime (in milliseconds) of flushing observation to configured data sink",
			Buckets:   requestLatencyBuckets,
		},
			allLabels,
		),
	}

	return histogramMap
}
