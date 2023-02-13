package service

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/helm"
	"github.com/caraml-dev/timber/dataset-service/helm/values"
	osconfig "github.com/caraml-dev/timber/observation-service/config"
	"helm.sh/helm/v3/pkg/chart"
)

const (
	// releaseNamePrefix is the helm release name prefix to be used when deploying Observation Service.
	releaseNamePrefix = "os"
)

// ObservationService provides a set of methods to interact with the MLP APIs
type ObservationService interface {
	// CreateService creates new Observation Service Helm release and returns ID of created Observation Service
	CreateService(projectName string, config *timberv1.ObservationService) (*timberv1.ObservationService, error)
	// UpdateService updates existing Observation Service Helm release and returns ID of updated Observation Service
	UpdateService(projectName string, observationServiceID int, config *timberv1.ObservationService) (*timberv1.ObservationService, error)
}

type observationService struct {
	helmClient         helm.Client
	helmChart          *chart.Chart
	commonDeployConfig *config.CommonDeploymentConfig
	defaults           *values.ObservationServiceHelmValues
}

// NewObservationService instantiates ObservationService
func NewObservationService(
	commonDeployConfig *config.CommonDeploymentConfig,
	observationServiceConfig *config.ObservationServiceConfig,
) (ObservationService, error) {
	helmClient := helm.NewClient(commonDeployConfig.KubeConfig)
	helmChart, err := helmClient.ReadChart(observationServiceConfig.HelmChartPath)
	if err != nil {
		return nil, fmt.Errorf("failed initializing observation service %w", err)
	}

	return &observationService{
		helmClient:         helmClient,
		helmChart:          helmChart,
		commonDeployConfig: commonDeployConfig,
		defaults:           observationServiceConfig.DefaultValues,
	}, nil
}

func (o *observationService) CreateService(projectName string, svc *timberv1.ObservationService) (*timberv1.ObservationService, error) {
	//TODO: create BQ dataset and/or table before deploying the observation service, although observation service has that privileges

	releaseName := fmt.Sprintf("%s-%s", releaseNamePrefix, svc.GetName())
	val, err := o.createHelmValues(releaseName, projectName, svc)
	if err != nil {
		return nil, fmt.Errorf("error creating helm values: %w", err)
	}
	// Trigger helm installation
	r, err := o.helmClient.Install(releaseName, projectName, o.helmChart, val, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating observation service: %w", err)
	}

	svc.Status = helm.ConvertStatus(r.Info.Status)
	// TODO: store observation service in DB and update the status based on the final release status
	return svc, nil
}

func (o *observationService) UpdateService(projectName string, observationServiceID int, svc *timberv1.ObservationService) (*timberv1.ObservationService, error) {

	releaseName := fmt.Sprintf("%s-%s", releaseNamePrefix, svc.GetName())
	val, err := o.createHelmValues(releaseName, projectName, svc)
	if err != nil {
		return nil, fmt.Errorf("error creating helm values: %w", err)
	}

	// Trigger helm release update
	r, err := o.helmClient.Upgrade(releaseName, projectName, o.helmChart, val, nil)
	if err != nil {
		return nil, fmt.Errorf("error upgrading observation service: %w", err)
	}

	log.Debugf("deployment manifest %s", r.Manifest)

	svc.Status = helm.ConvertStatus(r.Info.Status)

	// TODO: store observation service in DB and update the status based on the final release status
	return svc, nil
}

// createHelmValues create helm values for deployment based on the default values and the configuration given by the request
func (o *observationService) createHelmValues(releaseName string, projectName string, svc *timberv1.ObservationService) (map[string]any, error) {
	val := o.defaults

	val.ObservationService.APIConfig.DeploymentConfig.ServiceName = svc.Name
	val.ObservationService.APIConfig.DeploymentConfig.ProjectName = projectName

	val, err := setLogConsumerConfig(val, svc)
	if err != nil {
		return nil, err
	}

	val, err = setLogProducerConfig(releaseName, projectName, val, svc, o.commonDeployConfig.BQConfig)
	if err != nil {
		return nil, err
	}

	// Convert type of values for merging
	var interfaceValues map[string]any
	byteArr, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(byteArr, &interfaceValues)
	if err != nil {
		return nil, err
	}

	return interfaceValues, nil
}

// setLogConsumerConfig configures custom values for LogConsumerConfig
func setLogConsumerConfig(
	val *values.ObservationServiceHelmValues,
	svc *timberv1.ObservationService,
) (*values.ObservationServiceHelmValues, error) {
	switch svc.GetSource().GetType() {
	case timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_EAGER:
		return nil, fmt.Errorf("source type (eager) is currently unsupported")
	case timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_KAFKA:
		val.ObservationService.APIConfig.LogConsumerConfig.Kind = osconfig.LoggerKafkaConsumer
		val.ObservationService.APIConfig.LogConsumerConfig.KafkaConfig.Topic = svc.Source.Kafka.Topic
		val.ObservationService.APIConfig.LogConsumerConfig.KafkaConfig.Brokers = svc.Source.Kafka.Brokers
	case timberv1.ObservationServiceSourceType_OBSERVATION_SERVICE_SOURCE_TYPE_UNSPECIFIED:
		log.Infof("No source type specified for Observation Service deployment")
	default:
		return nil, fmt.Errorf("invalid source type (%s) was provided", svc.GetSource().GetType())
	}

	return val, nil
}

// setLogProducerConfig configures destination to which the observation service (in particular fluentd) write into
// to reduce complexity, it's limited to the bigquery
func setLogProducerConfig(releaseName string, projectName string, val *values.ObservationServiceHelmValues, svc *timberv1.ObservationService, bqConfig *config.BQConfig) (*values.ObservationServiceHelmValues, error) {
	// configure fluentd and BQ as default
	val.ObservationService.APIConfig.LogProducerConfig.Kind = osconfig.LoggerFluentdProducer
	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.Kind = osconfig.LoggerBQSinkFluentdProducer

	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.Host = fmt.Sprintf("%s-%s.%s", releaseName, val.Fluentd.NameOverride, projectName)

	// TODO: extract BQ table/dataset naming into separate functions
	datasetName := fmt.Sprintf("%s_%s", bqConfig.BQDatasetPrefix, projectName)
	tableName := fmt.Sprintf("%s_%s", bqConfig.ObservationBQTablePrefix, strings.ReplaceAll(svc.GetName(), "-", "_"))

	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.BQConfig.Project = bqConfig.GCPProject
	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.BQConfig.Dataset = datasetName
	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.BQConfig.Table = tableName

	val.Fluentd.ExtraEnvs = values.MerveEnvs(
		val.Fluentd.ExtraEnvs,
		[]values.Env{
			{
				Name:  "FLUENTD_GCP_PROJECT",
				Value: bqConfig.GCPProject,
			},
			{
				Name:  "FLUENTD_BQ_DATASET",
				Value: datasetName,
			},
			{
				Name:  "FLUENTD_BQ_TABLE",
				Value: tableName,
			},
		},
	)

	return val, nil
}
