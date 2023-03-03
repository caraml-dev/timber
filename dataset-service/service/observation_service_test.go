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

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/helm/mocks"
	"github.com/caraml-dev/timber/dataset-service/helm/values"
	"github.com/caraml-dev/timber/dataset-service/model"
	osconfig "github.com/caraml-dev/timber/observation-service/config"
)

type ObservationServiceTestSuite struct {
	suite.Suite
	config    *config.Config
	helmChart *chart.Chart
}

func (s *ObservationServiceTestSuite) SetupSuite() {
	s.Suite.T().Log("Setting up ObservationServiceTestSuite")
	cfg, err := config.Load("testdata/test_config.yaml")
	s.NoError(err)

	s.config = cfg
	s.helmChart = &chart.Chart{}
}

func (s *ObservationServiceTestSuite) TearDownSuite() {
	s.Suite.T().Log("Cleaning up ObservationServiceTestSuite")
}

func (s *ObservationServiceTestSuite) TestInstallOrUpgrade() {
	type args struct {
		projectName string
		svc         *model.ObservationService
	}

	tests := []struct {
		name string
		args args
		want *model.ObservationService
		// helm values that's being overridden by observation service
		wantOverrideHelmValues *values.ObservationServiceHelmValues
		wantErr                bool
	}{
		{
			name: "create",
			args: args{
				projectName: "my-project",
				svc: &model.ObservationService{
					Base: model.Base{ProjectID: 1},
					Name: "my-observation-svc",
					Source: &model.ObservationServiceSource{
						ObservationServiceSource: &timberv1.ObservationServiceSource{
							Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
							Kafka: &timberv1.KafkaConfig{
								Brokers: "kafka.brokers",
								Topic:   "sample-topic",
							},
						},
					},
				},
			},
			wantOverrideHelmValues: &values.ObservationServiceHelmValues{
				ObservationService: values.ObservationService{
					APIConfig: osconfig.Config{
						DeploymentConfig: osconfig.DeploymentConfig{
							ProjectName: "my-project",
							ServiceName: "my-observation-svc",
						},
						LogConsumerConfig: osconfig.LogConsumerConfig{
							Kind: osconfig.LoggerKafkaConsumer,
							KafkaConfig: &osconfig.KafkaConfig{
								Brokers: "kafka.brokers",
								Topic:   "sample-topic",
							},
						},
						LogProducerConfig: osconfig.LogProducerConfig{
							Kind: osconfig.LoggerFluentdProducer,
							FluentdConfig: &osconfig.FluentdConfig{
								Kind: osconfig.LoggerBQSinkFluentdProducer,
								Host: "os-my-observation-svc-fluentd.my-project",
								BQConfig: &osconfig.BQConfig{
									Project: "my-gcp-project",
									Dataset: "caraml_my_project",
									Table:   "os_my_observation_svc",
								},
							},
						},
						MonitoringConfig: osconfig.MonitoringConfig{},
					},
				},
				Fluentd: values.FluentdHelmValues{
					ExtraEnvs: values.MergeEnvs(s.config.ObservationServiceConfig.DefaultValues.Fluentd.ExtraEnvs, []values.Env{
						{
							Name:  values.FluentdBQDatasetEnv,
							Value: "caraml_my_project",
						},
						{
							Name:  values.FluentdGCPProjectEnv,
							Value: "my-gcp-project",
						},
						{
							Name:  values.FluentdBQTableEnv,
							Value: "os_my_observation_svc",
						},
					}),
				},
			},
			want: &model.ObservationService{
				Base: model.Base{ProjectID: 1},
				Name: "my-observation-svc",
				Source: &model.ObservationServiceSource{
					ObservationServiceSource: &timberv1.ObservationServiceSource{
						Type: timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA,
						Kafka: &timberv1.KafkaConfig{
							Brokers: "kafka.brokers",
							Topic:   "sample-topic",
						},
					},
				},
				Status: model.StatusDeployed,
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			mockHelmClient := &mocks.Client{}
			mockHelmClient.On("InstallOrUpgrade",
				fmt.Sprintf("%s-%s", releaseNamePrefix, tt.args.svc.Name),
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

			obs := &observationService{
				helmClient:         mockHelmClient,
				helmChart:          s.helmChart,
				commonDeployConfig: s.config.CommonDeploymentConfig,
				defaults:           s.config.ObservationServiceConfig.DefaultValues,
			}

			got, err := obs.InstallOrUpgrade(tt.args.projectName, tt.args.svc)
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

func (s *ObservationServiceTestSuite) assertHelmValuesOverride(mockHelmClient *mocks.Client, override *values.ObservationServiceHelmValues) {
	// copy first to avoid s.config.ObservationServiceConfig.DefaultValues getting overwritten by test
	expHelmValues := values.ObservationServiceHelmValues{}
	err := copier.CopyWithOption(&expHelmValues, s.config.ObservationServiceConfig.DefaultValues, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	s.NoError(err)

	// merge expHelmValues (that contains copy of s.config.ObservationServiceConfig.DefaultValues) with expected override
	err = mergo.Merge(&expHelmValues, override, mergo.WithOverride)
	s.NoError(err)

	// compare against the value received by mock helm client
	gotHelmValues := mockHelmClient.Calls[0].Arguments[3]
	wantRawValues, err := values.ToRaw(expHelmValues)
	s.NoError(err)

	assert.Equal(s.T(), wantRawValues, gotHelmValues)
}

func TestObservationService(t *testing.T) {
	suite.Run(t, new(ObservationServiceTestSuite))
}
