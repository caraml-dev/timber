package config

import (
	"strings"
	"testing"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	commonconfig "github.com/caraml-dev/timber/common/config"
)

func TestDefaultConfigs(t *testing.T) {
	emptyInterfaceMap := make(map[string]interface{})
	emptyStringMap := make(map[string]string)
	mlpConfig := &MLPConfig{URL: ""}
	defaultCfg := Config{
		Port: 8080,
		DeploymentConfig: commonconfig.DeploymentConfig{
			EnvironmentType: "local",
			LogLevel:        commonconfig.InfoLevel,
			MaxGoRoutines:   1000,
		},
		MLPConfig: mlpConfig,
		NewRelicConfig: newrelic.Config{
			Enabled:           false,
			AppName:           "",
			License:           "",
			IgnoreStatusCodes: []int{},
			Labels:            emptyInterfaceMap,
		},
		ObservationServiceConfig: ObservationServiceConfig{},
		SentryConfig:             sentry.Config{Enabled: false, Labels: emptyStringMap},
	}
	cfg, err := Load()
	require.NoError(t, err)
	assert.Equal(t, defaultCfg, *cfg)
	assert.Equal(t, ":8080", cfg.ListenAddress())
}

// TestLoadConfigFiles verifies that when multiple configs are passed in
// they are consumed in the correct order
func TestLoadConfigFiles(t *testing.T) {
	mlpConfig := &MLPConfig{URL: ""}
	tests := []struct {
		name        string
		configFiles []string
		errString   string
		expected    Config
	}{
		{
			name:        "success | load multiple config files",
			configFiles: []string{"../testdata/config1.yaml", "../testdata/config2.yaml"},
			expected: Config{
				Port: 8081,
				DeploymentConfig: commonconfig.DeploymentConfig{
					EnvironmentType: "dev",
					LogLevel:        commonconfig.InfoLevel,
					MaxGoRoutines:   1000,
				},
				MLPConfig: mlpConfig,
				NewRelicConfig: newrelic.Config{
					Enabled:           true,
					AppName:           "dataset-service",
					License:           "amazing-license",
					IgnoreStatusCodes: []int{403, 404, 405},
					Labels:            map[string]interface{}{"env": "dev"},
				},
				ObservationServiceConfig: ObservationServiceConfig{
					GCPProject:                      "test-project",
					ObservationServiceImageTag:      "v0.0.0",
					ObservationServiceHelmChartPath: "www.chart.url",
					FluentdImageTag:                 "v0.0.0",
					KubeConfig:                      "/tmp/kubeconfig",
				},
				SentryConfig: sentry.Config{Enabled: false, Labels: map[string]string{"app": "dataset-service", "env": "dev"}},
			},
		},
		{
			name:        "failure | bad config",
			configFiles: []string{"../testdata/config3.yaml"},
			errString: strings.Join([]string{"failed to update config: failed to unmarshal viper config values: 1 error(s) decoding:\n\n* cannot ",
				"parse 'Port' as int: strconv.ParseInt: parsing \"abc\": invalid syntax"}, ""),
		},
		{
			name:        "failure | file read",
			configFiles: []string{"../testdata/config4.yaml"},
			errString: strings.Join([]string{"failed to update config: failed to read viper config from file '../testdata/config4.yaml': ",
				"While parsing config: yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `bad_config` ",
				"into map[string]interface {}"}, ""),
		},
	}

	for _, data := range tests {
		t.Run(data.name, func(t *testing.T) {
			cfg, err := Load(data.configFiles...)
			if data.errString == "" {
				// Success
				require.NoError(t, err)
				assert.Equal(t, data.expected, *cfg)
			} else {
				assert.EqualError(t, err, data.errString)
				assert.Nil(t, cfg)
			}
		})
	}
}
