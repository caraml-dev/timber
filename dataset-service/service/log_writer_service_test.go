package service

import (
	"fmt"
	"testing"

	"github.com/imdario/mergo"
	"github.com/jinzhu/copier"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/release"

	"github.com/caraml-dev/timber/dataset-service/model"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/helm/mocks"
	"github.com/caraml-dev/timber/dataset-service/helm/values"
)

type LogWriterServicetestSuite struct {
	suite.Suite
	config    *config.Config
	helmChart *chart.Chart
}

func (s *LogWriterServicetestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up LogWriterServicetestSuite")
	cfg, err := config.Load("testdata/test_config.yaml")
	s.NoError(err)

	s.config = cfg
	s.helmChart = &chart.Chart{}
}

func (s *LogWriterServicetestSuite) TearDownSuite() {
	s.Suite.T().Log("Cleaning up LogWriterServiceTestSuite")
}

func (s *LogWriterServicetestSuite) TestInstallOrUpgrade() {
	type args struct {
		projectName string
		svc         *model.LogWriter
	}

	tests := []struct {
		name string
		args args
		want *model.LogWriter
		// helm values that's being overridden by observation service
		wantOverrideHelmValues *values.FluentdHelmValues
		wantErr                bool
	}{
		{
			name: "create log writer for prediction log",
			args: args{
				projectName: "my-project",
				svc: &model.LogWriter{
					Base: model.Base{
						ProjectID: 1,
					},
					Name: "prediction-log-writer",
					Source: &model.LogWriterSource{
						LogWriterSource: &timberv1.LogWriterSource{
							Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG,
							PredictionLogSource: &timberv1.PredictionLogSource{
								ModelName: "sample-model",
								ModelId:   1,
								Kafka: &timberv1.KafkaConfig{
									Brokers: "kafka-brokers.svc",
									Topic:   "sample-model-prediction-log",
								},
							},
						},
					},
				},
			},
			wantOverrideHelmValues: &values.FluentdHelmValues{
				ExtraEnvs: values.MergeEnvs(s.config.LogWriterConfig.DefaultValues.ExtraEnvs, []values.Env{
					{
						Name:  values.FluentdKafkaBrokerEnv,
						Value: "kafka-brokers.svc",
					},
					{
						Name:  values.FluentdKafkaTopicEnv,
						Value: "sample-model-prediction-log",
					},
					{
						Name:  values.FluentdProtoClassNameEnv,
						Value: predictionLogProto,
					},
					{
						Name:  values.FluentdTagEnv,
						Value: "sample-model-prediction-log",
					},
					{
						Name:  values.FluentdGCPProjectEnv,
						Value: "my-gcp-project",
					},
					{
						Name:  values.FluentdBQDatasetEnv,
						Value: "caraml_my_project",
					},
					{
						Name:  values.FluentdBQTableEnv,
						Value: "sample_model_prediction_log",
					},
				}),
			},
			want: &model.LogWriter{
				Base: model.Base{
					ProjectID: 1,
				},
				Name: "prediction-log-writer",
				Source: &model.LogWriterSource{
					LogWriterSource: &timberv1.LogWriterSource{
						Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG,
						PredictionLogSource: &timberv1.PredictionLogSource{
							ModelName: "sample-model",
							ModelId:   1,
							Kafka: &timberv1.KafkaConfig{
								Brokers: "kafka-brokers.svc",
								Topic:   "sample-model-prediction-log",
							},
						},
					},
				},
				Status: model.StatusDeployed,
			},
		},
		{
			name: "create log writer for router log",
			args: args{
				projectName: "my-project",
				svc: &model.LogWriter{
					Base: model.Base{
						ProjectID: 1,
					},
					Name: "router-log-writer",
					Source: &model.LogWriterSource{
						LogWriterSource: &timberv1.LogWriterSource{
							Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
							RouterLogSource: &timberv1.RouterLogSource{
								RouterName: "sample-router",
								RouterId:   1,
								Kafka: &timberv1.KafkaConfig{
									Brokers: "kafka-brokers.svc",
									Topic:   "sample-router-log",
								},
							},
						},
					},
				},
			},
			wantOverrideHelmValues: &values.FluentdHelmValues{
				ExtraEnvs: values.MergeEnvs(s.config.LogWriterConfig.DefaultValues.ExtraEnvs, []values.Env{
					{
						Name:  values.FluentdKafkaBrokerEnv,
						Value: "kafka-brokers.svc",
					},
					{
						Name:  values.FluentdKafkaTopicEnv,
						Value: "sample-router-log",
					},
					{
						Name:  values.FluentdProtoClassNameEnv,
						Value: routerLogProto,
					},
					{
						Name:  values.FluentdTagEnv,
						Value: "sample-router-log",
					},
					{
						Name:  values.FluentdGCPProjectEnv,
						Value: "my-gcp-project",
					},
					{
						Name:  values.FluentdBQDatasetEnv,
						Value: "caraml_my_project",
					},
					{
						Name:  values.FluentdBQTableEnv,
						Value: "sample_router_log",
					},
				}),
			},
			want: &model.LogWriter{
				Base: model.Base{
					ProjectID: 1,
				},
				Name: "router-log-writer",
				Source: &model.LogWriterSource{
					LogWriterSource: &timberv1.LogWriterSource{
						Type: timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG,
						RouterLogSource: &timberv1.RouterLogSource{
							RouterName: "sample-router",
							RouterId:   1,
							Kafka: &timberv1.KafkaConfig{
								Brokers: "kafka-brokers.svc",
								Topic:   "sample-router-log",
							},
						},
					},
				},
				Status: model.StatusDeployed,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			releaseName := fmt.Sprintf("%s-%s", predictionLogWriterReleaseNamePrefix, tt.args.svc.Name)
			if tt.args.svc.Source.Type == timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG {
				releaseName = fmt.Sprintf("%s-%s", routerLogWriterReleaseNamePrefix, tt.args.svc.Name)
			}
			mockHelmClient := &mocks.Client{}
			mockHelmClient.On("InstallOrUpgrade",
				releaseName,
				tt.args.projectName,
				s.helmChart,
				mock.Anything,
				mock.Anything,
			).
				Return(&release.Release{
					Info: &release.Info{
						Status: release.StatusDeployed,
					},
				}, nil)

			lws := &logWriterService{
				helmClient:         mockHelmClient,
				helmChart:          s.helmChart,
				commonDeployConfig: s.config.CommonDeploymentConfig,
				defaults:           s.config.LogWriterConfig.DefaultValues,
			}

			got, err := lws.InstallOrUpgrade(tt.args.projectName, tt.args.svc)
			if (err != nil) != tt.wantErr {
				s.T().Errorf("InstallOrUpgrade() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			assert.Equal(s.T(), tt.want, got)

			// validate that the helm values passed to helm client is as expected
			// the expected data is built from default values (`s.obs.defaults`) and merged with `wantOverrideHelmValues`
			s.assertHelmValuesOverride(mockHelmClient, tt.wantOverrideHelmValues)
			mockHelmClient.AssertExpectations(s.T())
		})
	}
}

func (s *LogWriterServicetestSuite) assertHelmValuesOverride(mockHelmClient *mocks.Client, override *values.FluentdHelmValues) {
	// copy first to avoid s.config.LogWriterConfig.DefaultValues getting overwritten by test
	expHelmValues := values.FluentdHelmValues{}
	err := copier.CopyWithOption(&expHelmValues, s.config.LogWriterConfig.DefaultValues, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	s.NoError(err)

	// merge expHelmValues (that contains copy of s.config.LogWriterConfig.DefaultValues) with expected override
	err = mergo.Merge(&expHelmValues, override, mergo.WithOverride)
	s.NoError(err)

	// compare against the value received by mock helm client
	gotHelmValues := mockHelmClient.Calls[0].Arguments[3]
	wantRawValues, err := values.ToRaw(expHelmValues)
	s.NoError(err)

	assert.Equal(s.T(), wantRawValues, gotHelmValues)
}

func TestLogWriterService(t *testing.T) {
	suite.Run(t, new(LogWriterServicetestSuite))
}
