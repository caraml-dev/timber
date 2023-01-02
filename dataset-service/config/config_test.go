package config

import (
	"strings"
	"testing"

	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	common_config "github.com/caraml-dev/timber/common/config"
)

func TestDefaultConfigs(t *testing.T) {
	emptyInterfaceMap := make(map[string]interface{})
	emptyStringMap := make(map[string]string)
	defaultCfg := Config{
		Port: 8080,
		DeploymentConfig: common_config.DeploymentConfig{
			EnvironmentType: "local",
			LogLevel:        common_config.InfoLevel,
			MaxGoRoutines:   1000,
		},
		NewRelicConfig: newrelic.Config{
			Enabled:           false,
			AppName:           "",
			License:           "",
			IgnoreStatusCodes: []int{},
			Labels:            emptyInterfaceMap,
		},
		SentryConfig: sentry.Config{Enabled: false, Labels: emptyStringMap},
	}
	cfg, err := Load()
	require.NoError(t, err)
	assert.Equal(t, defaultCfg, *cfg)
	assert.Equal(t, ":8080", cfg.ListenAddress())
}

// TestLoadConfigFiles verifies that when multiple configs are passed in
// they are consumed in the correct order
func TestLoadConfigFiles(t *testing.T) {
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
				DeploymentConfig: common_config.DeploymentConfig{
					EnvironmentType: "dev",
					LogLevel:        common_config.InfoLevel,
					MaxGoRoutines:   1000,
				},
				NewRelicConfig: newrelic.Config{
					Enabled:           true,
					AppName:           "dataset-service",
					License:           "amazing-license",
					IgnoreStatusCodes: []int{403, 404, 405},
					Labels:            map[string]interface{}{"env": "dev"},
				},
				SentryConfig: sentry.Config{Enabled: false, Labels: map[string]string{"app": "dataset-service", "env": "dev"}},
			},
		},
		{
			name:        "failure | bad config",
			configFiles: []string{"../testdata/config3.yaml"},
			errString: strings.Join([]string{"failed to update viper config: failed to unmarshal config values: 1 error(s) decoding:\n\n* cannot ",
				"parse 'Port' as int: strconv.ParseInt: parsing \"abc\": invalid syntax"}, ""),
		},
		{
			name:        "failure | file read",
			configFiles: []string{"../testdata/config4.yaml"},
			errString: strings.Join([]string{"failed to update viper config: failed to read config from file '../testdata/config4.yaml': ",
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
