package bq

import (
	"testing"

	"github.com/caraml-dev/timber/dataset-service/config"
)

var bqConfig = &config.BQConfig{
	GCPProject:               "my-gcp-project",
	BQDatasetPrefix:          "caraml",
	ObservationBQTablePrefix: "os",
}

func TestDatasetFromProject(t *testing.T) {
	type args struct {
		bqConfig *config.BQConfig
		project  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nominal case",
			args: args{
				bqConfig: bqConfig,
				project:  "example",
			},
			want: "caraml_example",
		},
		{
			name: "with invalid characters",
			args: args{
				bqConfig: bqConfig,
				project:  "my-example-project",
			},
			want: "caraml_my_example_project",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DatasetFromProject(tt.args.bqConfig, tt.args.project); got != tt.want {
				t.Errorf("DatasetFromProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableFromKafkaTopic(t *testing.T) {
	type args struct {
		kafkaTopic string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nominal case",
			args: args{
				kafkaTopic: "topic",
			},
			want: "topic",
		},
		{
			name: "with invalid characters",
			args: args{
				"topic-with-dash",
			},
			want: "topic_with_dash",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TableFromKafkaTopic(tt.args.kafkaTopic); got != tt.want {
				t.Errorf("TableFromKafkaTopic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTableFromObservationService(t *testing.T) {
	type args struct {
		bqConfig               *config.BQConfig
		observationServiceName string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "nominal case",
			args: args{
				bqConfig:               bqConfig,
				observationServiceName: "example",
			},
			want: "os_example",
		},
		{
			name: "with invalid characters",
			args: args{
				bqConfig:               bqConfig,
				observationServiceName: "my-example-service",
			},
			want: "os_my_example_service",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TableFromObservationService(tt.args.bqConfig, tt.args.observationServiceName); got != tt.want {
				t.Errorf("TableFromObservationService() = %v, want %v", got, tt.want)
			}
		})
	}
}
