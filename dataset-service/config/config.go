package config

import (
	"fmt"
	"time"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"

	"github.com/caraml-dev/timber/dataset-service/helm/values"

	commonconfig "github.com/caraml-dev/timber/common/config"
)

// Config captures the config related to starting Dataset Service
type Config struct {
	// Configuration for dataset service deployment
	DatasetServiceConfig *DatasetServiceConfig
	// Common deployment configuration for both log writer and observation service
	CommonDeploymentConfig *CommonDeploymentConfig
	// Observation service specific configuration
	ObservationServiceConfig *ObservationServiceConfig
	// Log writer specific configuration
	LogWriterConfig *LogWriterConfig
}

// DatasetServiceConfig configuration for dataset-service
type DatasetServiceConfig struct {
	// Port to be used by dataset service
	Port int `envconfig:"PORT" default:"8080"`
	// LogLevel captures the selected supported logging level
	LogLevel commonconfig.LogLevel `envconfig:"LOG_LEVEL" split_words:"false" default:"INFO"`
	// MlpURL is URL for connecting to MLP API
	MlpURL string `default:"localhost:3000"`
	// DatabaseConfig Database configuration
	DatabaseConfig *DatabaseConfig
	// New relic configuration
	NewRelicConfig *newrelic.Config
	// Sentry configuration
	SentryConfig *sentry.Config
}

// CommonDeploymentConfig configuration common to both observation-service and log writer deployment
type CommonDeploymentConfig struct {
	// KubeConfig specifies the file path to the configuration for which Kubernetes cluster to connect to
	KubeConfig string
	// BQ Config
	BQConfig *BQConfig
}

// BQConfig BigQuery configuration
type BQConfig struct {
	// GCPProject specifies the GCP project where BQ logs will be written to via FluentdHelmValues
	GCPProject string
	// BigQuery dataset prefix
	BQDatasetPrefix string `default:"caraml"`
	// Table name prefix for table storing observation logs
	ObservationBQTablePrefix string `default:"os"`
}

// ObservationServiceConfig configuration for deploying observation service
type ObservationServiceConfig struct {
	// link to Observation Service Helm chart for deployment
	HelmChartPath string
	// Default helm values to be used when deploying observation service
	DefaultValues *values.ObservationServiceHelmValues
}

// LogWriterConfig configuration for deploying log writer
type LogWriterConfig struct {
	// link to Log Writer Helm chart for deployment
	HelmChartPath string
	// Default helm values to be used when deploying log writer
	DefaultValues *values.FluentdHelmValues
}

// DatabaseConfig database configuration
type DatabaseConfig struct {
	// Host database host URL
	Host string `envconfig:"DATABASE_HOST" required:"true"`
	// Port database port
	Port int `envconfig:"DATABASE_PORT" default:"5432"`
	// User database user name
	User string `envconfig:"DATABASE_USER" required:"true"`
	// Password database password
	Password string `envconfig:"DATABASE_PASSWORD" required:"true"`
	// Database database name
	Database string `envconfig:"DATABASE_NAME" default:"timber"`

	// ConnMaxIdleTime maximum connection idle time
	ConnMaxIdleTime time.Duration `default:"0s"`
	// ConnMaxIdleTime maximum connection lifetime
	ConnMaxLifetime time.Duration `default:"0s"`
	// MaxIdleConns maximum number of idle connection
	MaxIdleConns int `default:"0"`
	// MaxIdleConns maximum number of open connection
	MaxOpenConns int `default:"0"`
}

// ListenAddress returns the Dataset API app's port
func (c *Config) ListenAddress() string {
	return fmt.Sprintf(":%d", c.DatasetServiceConfig.Port)
}

// Load parses multiple file configs specified via filepaths and returns a Config struct
func Load(filepaths ...string) (*Config, error) {
	var cfg Config
	err := commonconfig.ParseConfig(&cfg, filepaths)
	if err != nil {
		return nil, fmt.Errorf("failed to update config: %s", err)
	}

	return &cfg, nil
}
