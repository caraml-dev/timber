package config

import (
	"strings"
	"testing"

	"github.com/caraml-dev/timber/dataset-service/helm/values"
	"github.com/caraml-dev/timber/observation-service/config"
	"github.com/gojek/mlp/api/pkg/instrumentation/newrelic"
	"github.com/gojek/mlp/api/pkg/instrumentation/sentry"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var obsFluentdConf = `# Set fluentd log level to error
<system>
  log_level "#{ENV['FLUENTD_LOG_LEVEL']}"
  workers "#{ENV['FLUENTD_WORKER_COUNT']}"
</system>

# Accept HTTP input
<source>
  @type http
  port 9880
  bind 0.0.0.0
  body_size_limit 32m
  keepalive_timeout 10s
</source>

# Accept events on tcp socket
<source>
  @type forward
  port 24224
  bind 0.0.0.0
</source>

# Buffer and output to multiple sinks
<match "#{ENV['FLUENTD_TAG']}">
  @type copy
  <store>
    @type stdout
  </store>
  <store>
    @type bigquery_load

    <buffer>
      @type file

      path "#{ENV['FLUENTD_LOG_PATH']}"
      timekey_use_utc

      flush_at_shutdown true
      flush_mode interval
      flush_interval "#{ENV['FLUENTD_FLUSH_INTERVAL_SECONDS']}"
      retry_max_times 3

      chunk_limit_size 1g
      compress gzip
      total_limit_size "#{ENV['FLUENTD_BUFFER_LIMIT']}"

      delayed_commit_timeout 150
      disable_chunk_backup true
    </buffer>

    # Authenticate with BigQuery using a json key
    auth_method json_key
    json_key "#{ENV['FLUENTD_GCP_JSON_KEY_PATH']}"
    project "#{ENV['FLUENTD_GCP_PROJECT']}"
    dataset "#{ENV['FLUENTD_BQ_DATASET']}"
    table "#{ENV['FLUENTD_BQ_TABLE']}"
    fetch_schema true
  </store>
</match>`

var logWriterFluentdConf = `# Set fluentd log level to error
<system>
  log_level "#{ENV['FLUENTD_LOG_LEVEL']}"
  workers "#{ENV['FLUENTD_WORKER_COUNT']}"
</system>

<source>
  @type kafka

  brokers "#{ENV['FLUENTD_KAFKA_BROKER']}" #broker:9092
  topics "#{ENV['FLUENTD_KAFKA_TOPIC']}" #quickstart

  format upi_logs
  class_name "#{ENV['FLUENTD_KAFKA_PROTO_CLASS_NAME']}" #"caraml.upi.v1.PredictionLog"
</source>

# Buffer and output to multiple sinks
<match "#{ENV['FLUENTD_KAFKA_TOPIC']}">
  @type copy
  <store>
    @type stdout
  </store>
  <store>
    @type bigquery_load

    <buffer>
      @type file

      path "#{ENV['FLUENTD_LOG_PATH']}"
      timekey_use_utc

      flush_at_shutdown true
      flush_mode interval
      flush_interval "#{ENV['FLUENTD_FLUSH_INTERVAL_SECONDS']}"
      retry_max_times 3

      chunk_limit_size 1g
      compress gzip
      total_limit_size "#{ENV['FLUENTD_BUFFER_LIMIT']}"

      delayed_commit_timeout 150
      disable_chunk_backup true
    </buffer>

    # Authenticate with BigQuery using a json key
    auth_method json_key
    json_key "#{ENV['FLUENTD_GCP_JSON_KEY_PATH']}"
    project "#{ENV['FLUENTD_GCP_PROJECT']}"
    dataset "#{ENV['FLUENTD_BQ_DATASET']}"
    table "#{ENV['FLUENTD_BQ_TABLE']}"
    fetch_schema true
  </store>
</match>`

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
			name:        "success: load multiple config files",
			configFiles: []string{"testdata/config1.yaml", "testdata/config2.yaml"},
			expected: Config{
				DatasetServiceConfig: &DatasetServiceConfig{
					Port:     9000,
					LogLevel: "DEBUG",
					MlpURL:   "http://mlp.mlp.127.0.0.1.nip.io/v1",
					NewRelicConfig: &newrelic.Config{
						Enabled: false,
						AppName: "dataset-service",
						License: "newrelic-license-secret",
						IgnoreStatusCodes: []int{
							400, 401, 403, 404, 405, 412,
						},
						Labels: map[string]interface{}{"app": "dataset-service"},
					},
					SentryConfig: &sentry.Config{
						Enabled: false,
						DSN:     "xxx.xxx.xxx",
						Labels:  map[string]string{"app": "dataset-service"},
					},
				},
				CommonDeploymentConfig: &CommonDeploymentConfig{
					EnvironmentType: "local",
					KubeConfig:      "/tmp/kubeconfig-timber-dev.yaml",
					BQConfig: &BQConfig{
						GCPProject:               "my-gcp-project",
						BQDatasetPrefix:          "caraml",
						ObservationBQTablePrefix: "os",
					},
				},
				ObservationServiceConfig: &ObservationServiceConfig{
					HelmChartPath: "https://github.com/caraml-dev/helm-charts/releases/download/observation-svc-0.2.5/observation-svc-0.2.6.tgz",
					DefaultValues: &values.ObservationServiceHelmValues{
						ObservationService: values.ObservationService{
							Image: values.ImageConfig{
								Registry:   "ghcr.io",
								Repository: "caraml-dev/timber/observation-service",
								Tag:        "v0.0.0-build.15-b8afdb5",
								PullPolicy: "IfNotPresent",
							},
							Annotations:  map[string]string{"key": "value"},
							ExtraLabels:  map[string]string{"key": "value"},
							ReplicaCount: 1,
							Resources: values.ResourcesConfig{
								Requests: values.Resource{CPU: "10m", Memory: "50Mi"},
								Limits:   values.Resource{CPU: "1", Memory: "100Mi"},
							},
							Autoscaling: values.AutoscalingConfig{
								Enabled:                        false,
								MinReplicas:                    1,
								MaxReplicas:                    2,
								TargetCPUUtilizationPercentage: 50,
							},
							ExtraEnvs: []values.Env{
								{
									Name:  "EXTRA1",
									Value: "VALUE1",
								},
							},
							APIConfig: config.Config{
								Port: 8080,
								DeploymentConfig: config.DeploymentConfig{
									EnvironmentType: "local",
									ProjectName:     "my-project",
									ServiceName:     "observation-service",
									LogLevel:        "DEBUG",
									MaxGoRoutines:   1000,
								},
								NewRelicConfig: newrelic.Config{
									Enabled:           false,
									Labels:            map[string]interface{}{},
									IgnoreStatusCodes: []int{},
								},
								SentryConfig: sentry.Config{
									Enabled: false,
									Labels:  map[string]string{},
								},
								LogConsumerConfig: config.LogConsumerConfig{
									Kind: "kafka",
									KafkaConfig: &config.KafkaConfig{
										Brokers:          "kafka.mlp.svc.cluster.local",
										Topic:            "test-topic",
										MaxMessageBytes:  1048589,
										CompressionType:  "none",
										ConnectTimeoutMS: 1000,
										PollInterval:     1000,
										AutoOffsetReset:  "earliest",
									},
								},
								LogProducerConfig: config.LogProducerConfig{
									Kind:        "fluentd",
									QueueLength: 100000,
									KafkaConfig: &config.KafkaConfig{
										Brokers:          "",
										Topic:            "",
										MaxMessageBytes:  1048588,
										CompressionType:  "none",
										ConnectTimeoutMS: 1000,
										PollInterval:     1000,
										AutoOffsetReset:  "latest",
									},
									FluentdConfig: &config.FluentdConfig{
										Kind:        "bq",
										Host:        "obs-timber-observation-service-fluentd",
										Port:        24224,
										Tag:         "observation-service.log",
										BufferLimit: 8192,
										BQConfig: &config.BQConfig{
											Project: "my-project",
											Dataset: "my-dataset",
											Table:   "my-table",
										},
									},
								},
								MonitoringConfig: config.MonitoringConfig{
									Kind: "prometheus",
								},
							},
						},
						Fluentd: values.FluentdHelmValues{
							NameOverride: "fluentd",
							Image: values.ImageConfig{
								Registry:   "ghcr.io",
								Repository: "caraml-dev/timber/fluentd",
								Tag:        "v0.0.0-build.16-01ac82e",
								PullPolicy: "IfNotPresent",
							},
							Annotations:  map[string]string{"key": "value"},
							ExtraLabels:  map[string]string{"key": "value"},
							ReplicaCount: 1,
							Resources: values.ResourcesConfig{
								Requests: values.Resource{CPU: "10m", Memory: "50Mi"},
								Limits:   values.Resource{CPU: "1", Memory: "100Mi"},
							},
							GCPServiceAccount: values.GCPServiceAccount{
								CredentialsData: "ZHVtbXkK",
							},
							PVCConfig: values.PVCConfig{
								Name:      "cache-volume",
								MountPath: "/cache",
								Storage:   "3Gi",
							},
							ExtraEnvs: []values.Env{
								{
									Name:  "FLUENTD_WORKER_COUNT",
									Value: "1",
								},
								{
									Name:  "FLUENTD_LOG_LEVEL",
									Value: "debug",
								},
								{
									Name:  "FLUENTD_BUFFER_LIMIT",
									Value: "1g",
								},
								{
									Name:  "FLUENTD_FLUSH_INTERVAL_SECONDS",
									Value: "30",
								},
								{
									Name:  "FLUENTD_LOG_PATH",
									Value: "/fluentd/cache/log",
								},
								{
									Name:  "FLUENTD_TAG",
									Value: "observation-service.log",
								},
								{
									Name:  "FLUENTD_GCP_JSON_KEY_PATH",
									Value: "/etc/gcp_service_account/service-account.json",
								},
								{
									Name:  "FLUENTD_GCP_PROJECT",
									Value: "my-project",
								},
								{
									Name:  "FLUENTD_BQ_DATASET",
									Value: "my-dataset",
								},
								{
									Name:  "FLUENTD_BQ_TABLE",
									Value: "my-table",
								},
							},
							Autoscaling:   values.AutoscalingConfig{},
							FluentdConfig: obsFluentdConf,
							Enabled:       true,
						},
					},
				},
				LogWriterConfig: &LogWriterConfig{
					HelmChartPath: "https://github.com/caraml-dev/helm-charts/releases/download/fluentd-0.1.4/fluentd-0.1.5.tgz",
					DefaultValues: &values.FluentdHelmValues{
						Image: values.ImageConfig{
							Registry:   "ghcr.io",
							Repository: "caraml-dev/timber/fluentd",
							Tag:        "v0.0.0-build.16-01ac82e",
							PullPolicy: "IfNotPresent",
						},
						Annotations:  map[string]string{"key": "value"},
						ExtraLabels:  map[string]string{"key": "value"},
						ReplicaCount: 1,
						Resources: values.ResourcesConfig{
							Requests: values.Resource{CPU: "10m", Memory: "50Mi"},
							Limits:   values.Resource{CPU: "1", Memory: "100Mi"},
						},
						GCPServiceAccount: values.GCPServiceAccount{
							CredentialsData: "ZHVtbXkK",
						},
						PVCConfig: values.PVCConfig{
							Name:      "cache-volume",
							MountPath: "/cache",
							Storage:   "3Gi",
						},
						ExtraEnvs: []values.Env{
							{
								Name:  "FLUENTD_WORKER_COUNT",
								Value: "1",
							},
							{
								Name:  "FLUENTD_LOG_LEVEL",
								Value: "debug",
							},
							{
								Name:  "FLUENTD_BUFFER_LIMIT",
								Value: "1g",
							},
							{
								Name:  "FLUENTD_FLUSH_INTERVAL_SECONDS",
								Value: "30",
							},
							{
								Name:  "FLUENTD_LOG_PATH",
								Value: "/fluentd/cache/log",
							},
							{
								Name:  "FLUENTD_TAG",
								Value: "observation-service.log",
							},
							{
								Name:  "FLUENTD_GCP_JSON_KEY_PATH",
								Value: "/etc/gcp_service_account/service-account.json",
							},
							{
								Name:  "FLUENTD_GCP_PROJECT",
								Value: "my-project",
							},
							{
								Name:  "FLUENTD_BQ_DATASET",
								Value: "my-dataset",
							},
							{
								Name:  "FLUENTD_BQ_TABLE",
								Value: "my-table",
							},
							{
								Name:  "FLUENTD_KAFKA_BROKER",
								Value: "kafka.default.svc.cluster.local",
							},
							{
								Name:  "FLUENTD_KAFKA_TOPIC",
								Value: "quickstart",
							},
							{
								Name:  "FLUENTD_KAFKA_PROTO_CLASS_NAME",
								Value: "caraml.upi.v1.PredictionLog",
							},
						},
						Autoscaling:   values.AutoscalingConfig{},
						FluentdConfig: logWriterFluentdConf,
						Enabled:       false,
					},
				},
			},
		},
		{
			name:        "failure: bad config",
			configFiles: []string{"testdata/config3.yaml"},
			errString: strings.Join([]string{"failed to update config: failed to unmarshal viper config values: 1 error(s) decoding:\n\n* cannot ",
				"parse 'DatasetServiceConfig.Port' as int: strconv.ParseInt: parsing \"abc\": invalid syntax"}, ""),
		},
		{
			name:        "failure: file read",
			configFiles: []string{"testdata/config4.yaml"},
			errString: strings.Join([]string{"failed to update config: failed to read viper config from file 'testdata/config4.yaml': ",
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
