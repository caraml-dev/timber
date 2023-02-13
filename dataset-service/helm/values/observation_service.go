package values

import obsconfig "github.com/caraml-dev/timber/observation-service/config"

// Type for configuring observation service deployment via helm
type ObservationServiceHelmValues struct {
	// Full name override
	FullNameOverride string `json:"fullNameOverride,omitempty"`
	// Observation service deployment configurations
	ObservationService ObservationService `json:"observationService,omitempty"`
	// Fluentd deployment configuration
	Fluentd FluentdHelmValues `json:"fluentd,omitempty"`
}

// Observation service deployment configurations
type ObservationService struct {
	// Observation service image to be deployed
	Image ImageConfig `json:"image,omitempty"`
	// Annotations to be added to the deployment
	Annotations map[string]string `json:"annotations,omitempty"`
	// Additional labels to be added to the deployment
	ExtraLabels map[string]string `json:"extraLabels,omitempty"`
	// Number of replica
	ReplicaCount int `json:"replicaCount,omitempty"`
	// Resource configuration (i.e. CPU & Memory requests and limits)
	Resources ResourcesConfig `json:"resources,omitempty"`
	// Autoscaling configuration of the deployment
	Autoscaling AutoscalingConfig `json:"autoscaling,omitempty"`
	// Additional environment variables to be added to fluentd deployment
	ExtraEnvs []Env `json:"extraEnvs,omitempty"`
	// Observation service configuration
	APIConfig obsconfig.Config `json:"apiConfig,omitempty"`
	// Service k8s service configuration
	Service Service `json:"service,omitempty"`
}
