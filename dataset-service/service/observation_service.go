package service

import (
	"fmt"

	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/bq"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/helm"
	"github.com/caraml-dev/timber/dataset-service/helm/values"
	osconfig "github.com/caraml-dev/timber/observation-service/config"
	"github.com/jinzhu/copier"
	"helm.sh/helm/v3/pkg/chart"
)

const (
	// releaseNamePrefix is the helm release name prefix to be used when deploying Observation Service.
	releaseNamePrefix = "os"
)

// ObservationService provides a set of methods for controlling observation log's deployment
type ObservationService interface {
	// Create creates new Observation Service Helm release and returns ID of created Observation Service
	Create(projectName string, svc *timberv1.ObservationService) (*timberv1.ObservationService, error)
	// Update updates existing Observation Service Helm release and returns ID of updated Observation Service
	Update(projectName string, svc *timberv1.ObservationService) (*timberv1.ObservationService, error)
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

func (o *observationService) Create(projectName string, svc *timberv1.ObservationService) (*timberv1.ObservationService, error) {
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

func (o *observationService) Update(projectName string, svc *timberv1.ObservationService) (*timberv1.ObservationService, error) {

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
	val := &values.ObservationServiceHelmValues{}
	err := copier.CopyWithOption(val, o.defaults, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}

	val.ObservationService.APIConfig.DeploymentConfig.ServiceName = svc.Name
	val.ObservationService.APIConfig.DeploymentConfig.ProjectName = projectName

	val, err = setLogConsumerConfig(val, svc)
	if err != nil {
		return nil, err
	}

	val, err = setLogProducerConfig(releaseName, projectName, val, svc, o.commonDeployConfig.BQConfig)
	if err != nil {
		return nil, err
	}

	return values.ToRaw(val)
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
func setLogProducerConfig(releaseName string,
	projectName string,
	val *values.ObservationServiceHelmValues,
	svc *timberv1.ObservationService,
	bqConfig *config.BQConfig) (*values.ObservationServiceHelmValues, error) {

	// configure fluentd and BQ as default
	val.ObservationService.APIConfig.LogProducerConfig.Kind = osconfig.LoggerFluentdProducer
	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.Kind = osconfig.LoggerBQSinkFluentdProducer

	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.Host = fmt.Sprintf("%s-%s.%s", releaseName, val.Fluentd.NameOverride, projectName)

	// TODO: extract BQ table/dataset naming into separate functions
	datasetName := bq.DatasetFromProject(bqConfig, projectName)
	tableName := bq.TableFromObservationService(bqConfig, svc.GetName())

	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.BQConfig.Project = bqConfig.GCPProject
	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.BQConfig.Dataset = datasetName
	val.ObservationService.APIConfig.LogProducerConfig.FluentdConfig.BQConfig.Table = tableName

	val.Fluentd.ExtraEnvs = values.MergeEnvs(
		val.Fluentd.ExtraEnvs,
		[]values.Env{
			{
				Name:  values.FluentdGCPProjectEnv,
				Value: bqConfig.GCPProject,
			},
			{
				Name:  values.FluentdBQDatasetEnv,
				Value: datasetName,
			},
			{
				Name:  values.FluentdBQTableEnv,
				Value: tableName,
			},
		},
	)

	return val, nil
}
