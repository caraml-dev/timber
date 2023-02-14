package bq

import (
	"fmt"
	"regexp"

	"github.com/caraml-dev/timber/dataset-service/config"
)

// Restrict dataset and table name to only contain alphanumeric and "_" .
// Note that table name limitation is not the same as dataset name, however for consistency we'll use the dataset name
// restriction. Any characters outside of these allowed character will be replaced with "_"
var invalidChars = regexp.MustCompile("[^a-zA-Z0-9_]")

// DatasetFromProject create BQ dataset name from a project and prefix it with the BQConfig.BQDatasetPrefix
func DatasetFromProject(bqConfig *config.BQConfig, project string) string {
	return invalidChars.ReplaceAllString(fmt.Sprintf("%s_%s", bqConfig.BQDatasetPrefix, project), "_")
}

// TableFromObservationService create BQ table name from observation service name and prefix it with BQConfig.ObservationBQTablePrefix
func TableFromObservationService(bqConfig *config.BQConfig, observationServiceName string) string {
	return invalidChars.ReplaceAllString(fmt.Sprintf("%s_%s", bqConfig.ObservationBQTablePrefix, observationServiceName), "_")
}

// TableFromKafkaTopic create BQ table name from a given kafka topic.
func TableFromKafkaTopic(kafkaTopic string) string {
	return invalidChars.ReplaceAllString(kafkaTopic, "_")
}
