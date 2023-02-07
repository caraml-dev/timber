package services

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/cli"
	"k8s.io/apimachinery/pkg/api/resource"

	commonconfig "github.com/caraml-dev/timber/common/config"
	"github.com/caraml-dev/timber/common/log"
	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/models"
	osconfig "github.com/caraml-dev/timber/observation-service/config"
)

// TODO: Add dummy interfaces and use helm dependencies via dependency injection to skip actual installation in tests
// Context: https://github.com/caraml-dev/timber/pull/12#discussion_r1067755536

const (
	// helm_driver is the storage backend to be used. Documentation details: https://helm.sh/docs/topics/advanced/#storage-backends
	helm_driver = "secret"
	// release_name is the helm release name prefix to be used when deploying Observation Service.
	release_name = "obs"
	// fluentd_name_override is the name suffix for fluentd deployment
	fluentd_name_override = "timber-obs-fluentd"
)

// ObservationService provides a set of methods to interact with the MLP APIs
type ObservationService interface {
	// CreateService creates new Observation Service Helm release and returns ID of created Observation Service
	CreateService(projectName string, config *timberv1.ObservationServiceConfig) (*string, error)
	// UpdateService updates existing Observation Service Helm release and returns ID of updated Observation Service
	UpdateService(
		projectName string,
		observationServiceID int,
		config *timberv1.ObservationServiceConfig,
	) (*string, error)
}

type observationService struct {
	gcpProject               string
	deploymentConfig         commonconfig.DeploymentConfig
	observationServiceConfig config.ObservationServiceConfig
	observationServiceChart  *chart.Chart
}

// NewObservationService instantiates ObservationService
func NewObservationService(
	deploymentConfig commonconfig.DeploymentConfig,
	observationServiceConfig config.ObservationServiceConfig,
) (ObservationService, error) {

	settings := cli.New()
	chartPathOption := action.ChartPathOptions{}
	chartPath, err := chartPathOption.LocateChart(observationServiceConfig.ObservationServiceHelmChartPath, settings)
	if err != nil {
		return nil, err
	}
	observationServiceChart, err := loader.Load(chartPath)
	if err != nil {
		return nil, err
	}

	return &observationService{
		gcpProject:               observationServiceConfig.GCPProject,
		deploymentConfig:         deploymentConfig,
		observationServiceConfig: observationServiceConfig,
		observationServiceChart:  observationServiceChart,
	}, nil
}

func (o *observationService) CreateService(
	caramlProjectName string,
	config *timberv1.ObservationServiceConfig,
) (*string, error) {
	updatedChartValues, actionConfig, err := retrieveChartAndActionConfig(
		config,
		o.gcpProject,
		caramlProjectName,
		o.deploymentConfig,
		o.observationServiceConfig,
	)
	if err != nil {
		return nil, err
	}

	// Postfix helm release name with provided Service name as a CaraML project could have multiple different services
	projectReleaseName := fmt.Sprintf("%s-%s", release_name, config.GetServiceName())

	// Initialize helm installation
	installation := action.NewInstall(actionConfig)
	installation.ReleaseName = projectReleaseName
	installation.Namespace = caramlProjectName
	installation.CreateNamespace = true

	// Trigger helm installation
	release, err := installation.Run(o.observationServiceChart, updatedChartValues)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("%s helm release installation (version %d) underway! Status: %s", release.Name, release.Version, release.Info.Status)
	// Pretty print Helm release YAML manifest with fmt.Println
	// TODO: Store Manifest as a blob in DB
	if o.deploymentConfig.LogLevel == commonconfig.DebugLevel {
		fmt.Println(release.Manifest)
	}

	// TODO: Retrieve Observation Service ID from DB
	observationServiceID := uuid.New().String()

	return &observationServiceID, nil
}

func (o *observationService) UpdateService(
	caramlProjectName string,
	observationServiceID int,
	config *timberv1.ObservationServiceConfig,
) (*string, error) {
	updatedChartValues, actionConfig, err := retrieveChartAndActionConfig(
		config,
		o.gcpProject,
		caramlProjectName,
		o.deploymentConfig,
		o.observationServiceConfig,
	)
	if err != nil {
		return nil, err
	}

	// Initialize helm upgrade
	upgrade := action.NewUpgrade(actionConfig)
	upgrade.Namespace = caramlProjectName

	// Trigger helm upgrade
	// TODO: Get project release name based on provided Observation Service ID
	projectReleaseName := fmt.Sprintf("%s-%s", release_name, config.GetServiceName())
	release, err := upgrade.Run(projectReleaseName, o.observationServiceChart, updatedChartValues)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Infof("%s helm release upgrade (version %d) underway! Status: %s", release.Name, release.Version, release.Info.Status)
	// Pretty print Helm release YAML manifest with fmt.Println
	// TODO: Store Manifest as a blob in DB
	if o.deploymentConfig.LogLevel == commonconfig.DebugLevel {
		fmt.Println(release.Manifest)
	}

	// TODO: Retrieve Observation Service ID from DB
	retrievedObservationServiceID := strconv.Itoa(observationServiceID)

	return &retrievedObservationServiceID, nil
}

func retrieveChartAndActionConfig(
	config *timberv1.ObservationServiceConfig,
	gcpProject string,
	caramlProjectName string,
	deploymentConfig commonconfig.DeploymentConfig,
	observationServiceConfig config.ObservationServiceConfig,
) (map[string]any, *action.Configuration, error) {

	// Retrieve computed chart values based on default values and request body values
	updatedChartValues, err := getChartValues(config, gcpProject, caramlProjectName, deploymentConfig, observationServiceConfig)
	if err != nil {
		return nil, nil, err
	}

	// Generate configuration required to run helm install/upgrade
	settings := cli.New()
	actionConfig := new(action.Configuration)
	if observationServiceConfig.KubeConfig != "" {
		settings.KubeConfig = observationServiceConfig.KubeConfig
	}
	err = actionConfig.Init(settings.RESTClientGetter(), caramlProjectName, helm_driver, log.Infof)
	if err != nil {
		return nil, nil, err
	}

	return updatedChartValues, actionConfig, nil
}

func getChartValues(
	config *timberv1.ObservationServiceConfig,
	gcpProject string,
	caramlProjectName string,
	deploymentConfig commonconfig.DeploymentConfig,
	observationServiceConfig config.ObservationServiceConfig,
) (map[string]any, error) {
	// Initialize and set default values
	values := &models.ObservationServiceHelmValues{}
	values = setDefaultValues(values, gcpProject, caramlProjectName, config.GetServiceName(), deploymentConfig, observationServiceConfig)

	// Compute LogConsumerConfig
	values, err := setLogConsumerConfig(values, config)
	if err != nil {
		return nil, err
	}

	// Compute LogProducerConfig
	values, err = setLogProducerConfig(values, config, observationServiceConfig, gcpProject, caramlProjectName)
	if err != nil {
		return nil, err
	}

	// Convert type of values for merging
	var interfaceValues map[string]any
	byteArr, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(byteArr, &interfaceValues)
	if err != nil {
		return nil, err
	}

	return interfaceValues, nil
}

// setDefaultValues configures default values handled by Dataset Service for Observation Service deployments
func setDefaultValues(
	values *models.ObservationServiceHelmValues,
	gcpProject string,
	caramlProjectName string,
	serviceName string,
	deploymentConfig commonconfig.DeploymentConfig,
	observationServiceConfig config.ObservationServiceConfig,
) *models.ObservationServiceHelmValues {
	// --- Observation Service configs --- //
	// Override if application config is provided, else use helm values default
	if (observationServiceConfig.ObservationServiceImageTag) != "" {
		values.ObservationServiceConfig.Image = &models.Image{Tag: observationServiceConfig.ObservationServiceImageTag}
	}
	// ApiConfig
	values.ObservationServiceConfig.ApiConfig.Port = 9001
	values.ObservationServiceConfig.ApiConfig.DeploymentConfig = deploymentConfig
	values.ObservationServiceConfig.ApiConfig.DeploymentConfig.ProjectName = caramlProjectName
	values.ObservationServiceConfig.ApiConfig.DeploymentConfig.ServiceName = fmt.Sprintf("%s-%s", release_name, serviceName)

	// Resources
	values.ObservationServiceConfig.Resources.Requests.CPU = resource.Quantity{Format: "1"}
	values.ObservationServiceConfig.Resources.Requests.Memory = resource.Quantity{Format: "512Mi"}
	values.ObservationServiceConfig.Resources.Limits.CPU = resource.Quantity{Format: "1"}
	values.ObservationServiceConfig.Resources.Limits.Memory = resource.Quantity{Format: "1Gi"}
	// Autoscaling
	// TODO: Revisit if should be enabled by default
	values.ObservationServiceConfig.Autoscaling.Enabled = false

	// --- Fluentd configs --- //
	// Override if application config is provided, else use helm values default
	if observationServiceConfig.FluentdImageTag != "" {
		values.FluentdConfig.Image = &models.Image{Tag: observationServiceConfig.FluentdImageTag}
	}
	// Environment Variables
	envVars := []models.Env{
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
			Value: "/cache/log",
		},
		{
			Name:  "FLUENTD_TAG",
			Value: "observation-service",
		},
		{
			Name:  "FLUENTD_GCP_JSON_KEY_PATH",
			Value: "/etc/gcp_service_account/service-account.json",
		},
		{
			Name:  "FLUENTD_GCP_PROJECT",
			Value: gcpProject,
		},
		{
			Name:  "FLUENTD_BQ_DATASET",
			Value: caramlProjectName,
		},
		{
			Name:  "FLUENTD_BQ_TABLE",
			Value: fmt.Sprintf("%s_observation_log", serviceName),
		},
	}
	values.FluentdConfig.ExtraEnvs = envVars
	// Resources
	values.FluentdConfig.Resources.Requests.CPU = resource.Quantity{Format: "1"}
	values.FluentdConfig.Resources.Requests.Memory = resource.Quantity{Format: "512Mi"}
	values.FluentdConfig.Resources.Limits.CPU = resource.Quantity{Format: "1"}
	values.FluentdConfig.Resources.Limits.Memory = resource.Quantity{Format: "1Gi"}
	// Autoscaling
	values.FluentdConfig.Autoscaling.Enabled = false

	return values
}

// setLogConsumerConfig configures custom values for LogConsumerConfig
func setLogConsumerConfig(
	values *models.ObservationServiceHelmValues,
	config *timberv1.ObservationServiceConfig,
) (*models.ObservationServiceHelmValues, error) {
	switch config.GetSource().GetType() {
	case timberv1.ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_EAGER:
		return nil, fmt.Errorf("source type (eager) is currently unsupported")
	case timberv1.ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_KAFKA:
		kafkaConfig := config.GetSource().GetKafkaConfig()
		values.ObservationServiceConfig.ApiConfig.LogConsumerConfig.Kind = "kafka"
		values.ObservationServiceConfig.ApiConfig.LogConsumerConfig.KafkaConfig = models.NewKafkaConfig(kafkaConfig)
	case timberv1.ObservationServiceDataSourceType_OBSERVATION_SERVICE_DATA_SOURCE_TYPE_UNSPECIFIED:
		log.Infof("No source type specified for Observation Service deployment")
	default:
		return nil, fmt.Errorf("invalid source type (%s) was provided", config.GetSource().GetType())
	}

	return values, nil
}

// setLogProducerConfig configures custom values for LogProducerConfig
func setLogProducerConfig(
	values *models.ObservationServiceHelmValues,
	config *timberv1.ObservationServiceConfig,
	appConfig config.ObservationServiceConfig,
	gcpProject string,
	projectName string,
) (*models.ObservationServiceHelmValues, error) {
	switch config.GetSink().GetType() {
	case timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_FLUENTD:
		fluentdConfig := config.GetSink().GetFluentdConfig()
		//Prefix and postfix are added to service name as convention for identify fluentd
		fluentdConfig.Host = fmt.Sprintf("%s-%s-%s", release_name, config.GetServiceName(), fluentd_name_override)
		values.ObservationServiceConfig.ApiConfig.LogProducerConfig.Kind = "fluentd"
		values.ObservationServiceConfig.ApiConfig.LogProducerConfig.FluentdConfig = models.NewFluentdConfig(fluentdConfig)
		values.ObservationServiceConfig.ApiConfig.LogProducerConfig.FluentdConfig.BQConfig = &osconfig.BQConfig{
			//TODO to re-evaluate this if request should determine dataset and table
			// Project will be app configured, dataset and table are user input
			Project: appConfig.GCPProject,
			Dataset: config.GetSink().GetFluentdConfig().GetConfig().GetDataset(),
			Table:   config.GetSink().GetFluentdConfig().GetConfig().GetTable(),
		}
		values.FluentdConfig.Enabled = true
		//TODO Revisit service account workflow or manually created, for now take from appConfig
		values.FluentdConfig.GCPServiceAccount.Credentials.Name = appConfig.GCPServiceAccountSecret
		values.FluentdConfig.GCPServiceAccount.Credentials.Key = appConfig.GCPServiceAccountKey
	case timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_KAFKA:
		kafkaConfig := config.GetSink().GetKafkaConfig()
		values.ObservationServiceConfig.ApiConfig.LogProducerConfig.Kind = "kafka"
		values.ObservationServiceConfig.ApiConfig.LogProducerConfig.KafkaConfig = models.NewKafkaConfig(kafkaConfig)
	case timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_STDOUT:
		values.ObservationServiceConfig.ApiConfig.LogProducerConfig.Kind = "stdout"
		log.Infof("Standard output sink type specified for Observation Service deployment")
	case timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_NOOP:
		log.Infof("No-Op sink type specified for Observation Service deployment")
	case timberv1.ObservationServiceDataSinkType_OBSERVATION_SERVICE_DATA_SINK_TYPE_UNSPECIFIED:
		log.Infof("No sink type specified for Observation Service deployment")
	default:
		return nil, fmt.Errorf("invalid sink type (%s) was provided", config.GetSink().GetType())
	}

	return values, nil
}
