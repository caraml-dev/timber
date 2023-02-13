package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/helm"
	"github.com/caraml-dev/timber/dataset-service/helm/values"
	"helm.sh/helm/v3/pkg/chart"
)

const (
	// release name prefix for all prediction log writer deployments
	predictionLogWriterReleaseNamePrefix = "pl"
	// release name prefix for all router log writer deployments
	routerLogWriterReleaseNamePrefix = "rl"
	// protobuf name of prediction log
	predictionLogProto = "caraml.upi.v1.PredictionLog"
	// protobuf name of router log
	routerLogProto = "caraml.upi.v1.RouterLog"
)

// LogWriterService provides a set of methods for controlling the log writer's deployment
type LogWriterService interface {
	// Create creates a new log writer deployment
	Create(projectName string, logWriter *timberv1.LogWriter) (*timberv1.LogWriter, error)
	// Update updates an existing log writer deployment
	Update(projectName string, logWriter *timberv1.LogWriter) (*timberv1.LogWriter, error)
}

type logWriterService struct {
	helmClient         helm.Client
	helmChart          *chart.Chart
	commonDeployConfig *config.CommonDeploymentConfig
	defaults           *values.FluentdHelmValues
}

func NewLogWriterService(commonDeployConfig *config.CommonDeploymentConfig, logWriterConfig *config.LogWriterConfig) (LogWriterService, error) {
	helmClient := helm.NewClient(commonDeployConfig.KubeConfig)
	helmChart, err := helmClient.ReadChart(logWriterConfig.HelmChartPath)
	if err != nil {
		return nil, fmt.Errorf("failed initializing log writer service %w", err)
	}

	return &logWriterService{
		helmClient:         helmClient,
		helmChart:          helmChart,
		commonDeployConfig: commonDeployConfig,
		defaults:           logWriterConfig.DefaultValues,
	}, nil
}

func (l *logWriterService) Create(projectName string, logWriter *timberv1.LogWriter) (*timberv1.LogWriter, error) {
	// TODO: create BQ dataset and/or table before deploying the log writer
	releaseName := createReleaseName(logWriter)
	val, err := l.createHelmValues(projectName, logWriter)
	if err != nil {
		return nil, fmt.Errorf("error creating helm values: %w", err)
	}

	// Trigger helm installation
	r, err := l.helmClient.Install(releaseName, projectName, l.helmChart, val, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating observation service: %w", err)
	}

	logWriter.Status = helm.ConvertStatus(r.Info.Status)
	// TODO: store log writer in DB and update the status based on the final release status
	return logWriter, nil
}

func (l *logWriterService) Update(projectName string, logWriter *timberv1.LogWriter) (*timberv1.LogWriter, error) {
	releaseName := createReleaseName(logWriter)
	val, err := l.createHelmValues(projectName, logWriter)
	if err != nil {
		return nil, fmt.Errorf("error creating helm values: %w", err)
	}

	// Trigger helm installation
	r, err := l.helmClient.Upgrade(releaseName, projectName, l.helmChart, val, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating observation service: %w", err)
	}

	logWriter.Status = helm.ConvertStatus(r.Info.Status)
	// TODO: store log writer in DB and update the status based on the final release status
	return logWriter, nil
}

func (l *logWriterService) createHelmValues(projectName string, logWriter *timberv1.LogWriter) (map[string]any, error) {
	val := l.defaults
	kafkaConfig, err := getKafkaConfig(logWriter)
	if err != nil {
		return nil, err
	}

	val, err = l.configureSource(val, kafkaConfig, logWriter.Source.Type)
	if err != nil {
		return nil, fmt.Errorf("error configuring source: %w", err)
	}

	val, err = l.configureSink(val, projectName, kafkaConfig)
	if err != nil {
		return nil, fmt.Errorf("error configuring sink: %w", err)
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

func (l *logWriterService) configureSource(val *values.FluentdHelmValues, kafkaConfig *timberv1.KafkaConfig, logType timberv1.LogWriterSourceType) (*values.FluentdHelmValues, error) {
	protoName := predictionLogProto
	if logType == timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG {
		protoName = routerLogProto
	}

	val.ExtraEnvs = values.MerveEnvs(
		val.ExtraEnvs,
		[]values.Env{
			{
				Name:  "FLUENTD_KAFKA_BROKER",
				Value: kafkaConfig.Brokers,
			},
			{
				Name:  "FLUENTD_KAFKA_TOPIC",
				Value: kafkaConfig.Topic,
			},
			{
				Name:  "FLUENTD_KAFKA_PROTO_CLASS_NAME",
				Value: protoName,
			},
			{
				Name:  "FLUENTD_TAG",
				Value: kafkaConfig.Topic,
			},
		},
	)

	return val, nil
}

func (l *logWriterService) configureSink(val *values.FluentdHelmValues, projectName string, kafkaConfig *timberv1.KafkaConfig) (*values.FluentdHelmValues, error) {
	datasetName := l.createDatasetName(projectName)
	tableName := l.createTableName(kafkaConfig)
	val.ExtraEnvs = values.MerveEnvs(
		val.ExtraEnvs,
		[]values.Env{
			{
				Name:  "FLUENTD_GCP_PROJECT",
				Value: l.commonDeployConfig.BQConfig.GCPProject,
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

func (l *logWriterService) createDatasetName(projectName string) string {
	return strings.ReplaceAll(fmt.Sprintf("%s_%s", l.commonDeployConfig.BQConfig.BQDatasetPrefix, projectName), "-", "_")
}

func (l *logWriterService) createTableName(kafkaConfig *timberv1.KafkaConfig) string {
	return strings.ReplaceAll(kafkaConfig.Topic, "-", "_")
}

func createReleaseName(logWriter *timberv1.LogWriter) string {
	switch logWriter.Source.Type {
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG:
		return fmt.Sprintf("%s-%s", predictionLogWriterReleaseNamePrefix, logWriter.Name)
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG:
		return fmt.Sprintf("%s-%s", routerLogWriterReleaseNamePrefix, logWriter.Name)
	default:
		return ""
	}
}

func getKafkaConfig(logWriter *timberv1.LogWriter) (*timberv1.KafkaConfig, error) {
	switch logWriter.Source.Type {
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG:
		return logWriter.Source.PredictionLogSource.Kafka, nil
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG:
		return logWriter.Source.RouterLogSource.Kafka, nil
	default:
		return nil, errors.New("Source.LogType is not specified")
	}
}
