package config

import (
	"fmt"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"

	common_config "github.com/caraml-dev/observation-service/common/config"
)

type Config struct {
	HTTPPort int `envconfig:"CARAML_HTTP_PORT" default:"8081"`
	GRPCPort int `envconfig:"CARAML_GRPC_PORT" default:"9001"`

	DeploymentConfig DeploymentConfig
	NewRelicConfig   newrelic.Config
	SentryConfig     sentry.Config
}

// DeploymentConfig captures the config related to the deployment of Observation Service
type DeploymentConfig struct {
	EnvironmentType string `default:"local"`
}

// ListenAddress returns the Observation API port
func (c *Config) ListenAddress(portType string) string {
	if portType == "http" {
		return fmt.Sprintf(":%d", c.HTTPPort)
	}
	return fmt.Sprintf(":%d", c.GRPCPort)
}

func Load(filepaths ...string) (*Config, error) {
	var cfg Config
	err := common_config.ParseConfig(&cfg, filepaths)
	if err != nil {
		return nil, fmt.Errorf("failed to update viper config: %s", err)
	}

	return &cfg, nil
}
