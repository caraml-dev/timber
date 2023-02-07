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
	defaultCfg := Config{
		Port: 9001,
		DeploymentConfig: commonconfig.DeploymentConfig{
			EnvironmentType: "local",
			LogLevel:        commonconfig.InfoLevel,
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
		LogConsumerConfig: LogConsumerConfig{
			Kind: "",
			KafkaConfig: &KafkaConfig{
				Brokers:          "",
				Topic:            "",
				MaxMessageBytes:  1048588,
				CompressionType:  "none",
				ConnectTimeoutMS: 1000,
				PollInterval:     1000,
				AutoOffsetReset:  "latest",
			},
		},
		LogProducerConfig: LogProducerConfig{
			Kind:                 "",
			QueueLength:          100,
			FlushIntervalSeconds: 1,
			KafkaConfig: &KafkaConfig{
				Brokers:          "",
				Topic:            "",
				MaxMessageBytes:  1048588,
				CompressionType:  "none",
				ConnectTimeoutMS: 1000,
				PollInterval:     1000,
				AutoOffsetReset:  "latest",
			},
			FluentdConfig: &FluentdConfig{
				Kind:        "",
				Host:        "localhost",
				Port:        24224,
				Tag:         "observation-service",
				BufferLimit: 8192,
				BQConfig: &BQConfig{
					Project: "",
					Dataset: "",
					Table:   "",
				},
			},
		},
		MonitoringConfig: MonitoringConfig{
			Kind: "",
		},
	}
	cfg, err := Load()
	require.NoError(t, err)
	assert.Equal(t, defaultCfg, *cfg)
	assert.Equal(t, ":9001", cfg.ListenAddress())
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
				Port: 9002,
				DeploymentConfig: commonconfig.DeploymentConfig{
					EnvironmentType: "dev",
					LogLevel:        commonconfig.InfoLevel,
					MaxGoRoutines:   1000,
				},
				NewRelicConfig: newrelic.Config{
					Enabled:           true,
					AppName:           "observation-service",
					License:           "amazing-license",
					IgnoreStatusCodes: []int{403, 404, 405},
					Labels:            map[string]interface{}{"env": "dev"},
				},
				SentryConfig: sentry.Config{Enabled: false, Labels: map[string]string{"app": "observation-service", "env": "dev"}},
				LogConsumerConfig: LogConsumerConfig{
					Kind: "",
					KafkaConfig: &KafkaConfig{
						Brokers:          "",
						Topic:            "",
						MaxMessageBytes:  1048588,
						CompressionType:  "none",
						ConnectTimeoutMS: 1000,
						PollInterval:     1000,
						AutoOffsetReset:  "latest",
					},
				},
				LogProducerConfig: LogProducerConfig{
					Kind:                 "",
					QueueLength:          100,
					FlushIntervalSeconds: 1,
					KafkaConfig: &KafkaConfig{
						Brokers:          "",
						Topic:            "",
						MaxMessageBytes:  1048588,
						CompressionType:  "none",
						ConnectTimeoutMS: 1000,
						PollInterval:     1000,
						AutoOffsetReset:  "latest",
					},
					FluentdConfig: &FluentdConfig{
						Kind:        "",
						Host:        "localhost",
						Port:        24224,
						Tag:         "observation-service",
						BufferLimit: 8192,
						BQConfig: &BQConfig{
							Project: "",
							Dataset: "",
							Table:   "",
						},
					},
				},
				MonitoringConfig: MonitoringConfig{
					Kind: "",
				},
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
