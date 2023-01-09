package config

import (
	"fmt"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"

	common_config "github.com/caraml-dev/timber/common/config"
)

// Config captures the config related to starting Dataset Service
type Config struct {
	Port int `envconfig:"PORT" default:"8080"`

	DeploymentConfig         common_config.DeploymentConfig
	ObservationServiceConfig ObservationServiceConfig
	MLPConfig                *MLPConfig
	NewRelicConfig           newrelic.Config
	SentryConfig             sentry.Config
}

// ObservationServiceConfig captures the configuration used for log storage
type ObservationServiceConfig struct {
	GCPProject                 string
	ObservationServiceImageTag string
	FluentdImageTag            string
}

// MLPConfig captures the configuration used to connect to the MLP API server
type MLPConfig struct {
	URL string
}

// ListenAddress returns the Dataset API app's port
func (c *Config) ListenAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

// Load parses multiple file configs specified via filepaths and returns a Config struct
func Load(filepaths ...string) (*Config, error) {
	var cfg Config
	err := common_config.ParseConfig(&cfg, filepaths)
	if err != nil {
		return nil, fmt.Errorf("failed to update config: %s", err)
	}

	return &cfg, nil
}
