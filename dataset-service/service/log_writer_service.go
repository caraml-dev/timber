package service

import (
	"errors"
	"fmt"

	"github.com/jinzhu/copier"
	"helm.sh/helm/v3/pkg/chart"

	timberv1 "github.com/caraml-dev/timber/dataset-service/api"
	"github.com/caraml-dev/timber/dataset-service/bq"
	"github.com/caraml-dev/timber/dataset-service/config"
	"github.com/caraml-dev/timber/dataset-service/helm"
	"github.com/caraml-dev/timber/dataset-service/helm/values"
	"github.com/caraml-dev/timber/dataset-service/model"
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
	// InstallOrUpgrade creates a new log writer deployment
	InstallOrUpgrade(projectName string, logWriter *model.LogWriter) (*model.LogWriter, error)
	// Uninstall uninstalls an existing log writer deployment
	Uninstall(projectName string, logWriter *model.LogWriter) (*model.LogWriter, error)
}

type logWriterService struct {
	helmClient         helm.Client
	helmChart          *chart.Chart
	commonDeployConfig *config.CommonDeploymentConfig
	defaults           *values.FluentdHelmValues
}

// NewLogWriterService create a new instance of log writer service
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

func (l *logWriterService) InstallOrUpgrade(projectName string, logWriter *model.LogWriter) (*model.LogWriter, error) {
	// TODO: create BQ dataset and/or table before deploying the log writer
	releaseName := createReleaseName(logWriter)
	val, err := l.createHelmValues(projectName, logWriter)
	if err != nil {
		return nil, fmt.Errorf("error creating helm values: %w", err)
	}

	// Trigger helm installation
	r, err := l.helmClient.InstallOrUpgrade(releaseName, projectName, l.helmChart, val, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating observation service: %w", err)
	}

	logWriter.Status = helm.ConvertStatus(r.Info.Status)
	logWriter.Error = ""
	return logWriter, nil
}

func (l *logWriterService) Uninstall(projectName string, logWriter *model.LogWriter) (*model.LogWriter, error) {
	releaseName := createReleaseName(logWriter)

	err := l.helmClient.Uninstall(releaseName, projectName, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating observation service: %w", err)
	}

	logWriter.Status = model.StatusUninstalled
	logWriter.Error = ""
	return logWriter, nil
}

func (l *logWriterService) createHelmValues(projectName string, logWriter *model.LogWriter) (map[string]any, error) {
	val := &values.FluentdHelmValues{}
	err := copier.CopyWithOption(val, l.defaults, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	if err != nil {
		return nil, err
	}

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

	return values.ToRaw(val)
}

func (l *logWriterService) configureSource(val *values.FluentdHelmValues,
	kafkaConfig *timberv1.KafkaConfig,
	logType timberv1.LogWriterSourceType) (*values.FluentdHelmValues, error) {

	protoName := predictionLogProto
	if logType == timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG {
		protoName = routerLogProto
	}

	val.ExtraEnvs = values.MergeEnvs(
		val.ExtraEnvs,
		[]values.Env{
			{
				Name:  values.FluentdKafkaBrokerEnv,
				Value: kafkaConfig.Brokers,
			},
			{
				Name:  values.FluentdKafkaTopicEnv,
				Value: kafkaConfig.Topic,
			},
			{
				Name:  values.FluentdProtoClassNameEnv,
				Value: protoName,
			},
			{
				Name:  values.FluentdTagEnv,
				Value: kafkaConfig.Topic,
			},
		},
	)

	return val, nil
}

func (l *logWriterService) configureSink(val *values.FluentdHelmValues,
	projectName string,
	kafkaConfig *timberv1.KafkaConfig) (*values.FluentdHelmValues, error) {

	datasetName := bq.DatasetFromProject(l.commonDeployConfig.BQConfig, projectName)
	tableName := bq.TableFromKafkaTopic(kafkaConfig.Topic)
	val.ExtraEnvs = values.MergeEnvs(
		val.ExtraEnvs,
		[]values.Env{
			{
				Name:  values.FluentdGCPProjectEnv,
				Value: l.commonDeployConfig.BQConfig.GCPProject,
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

func createReleaseName(logWriter *model.LogWriter) string {
	switch logWriter.Source.Type {
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG:
		return fmt.Sprintf("%s-%s", predictionLogWriterReleaseNamePrefix, logWriter.Name)
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG:
		return fmt.Sprintf("%s-%s", routerLogWriterReleaseNamePrefix, logWriter.Name)
	default:
		return ""
	}
}

func getKafkaConfig(logWriter *model.LogWriter) (*timberv1.KafkaConfig, error) {
	switch logWriter.Source.Type {
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_PREDICTION_LOG:
		return logWriter.Source.PredictionLogSource.Kafka, nil
	case timberv1.LogWriterSourceType_LOG_WRITER_SOURCE_TYPE_ROUTER_LOG:
		return logWriter.Source.RouterLogSource.Kafka, nil
	default:
		return nil, errors.New("Source.LogType is not specified")
	}
}
