package values

// Type for configuring fluentd deployment via helm
type FluentdHelmValues struct {
	// Full name override
	NameOverride string `json:"nameOverride,omitempty"`
	// FluentD image to be deployed
	Image ImageConfig `json:"image,omitempty"`
	// Annotations to be added to the deployment
	Annotations map[string]string `json:"annotations,omitempty"`
	// Additional labels to be added to the deployment
	ExtraLabels map[string]string `json:"extraLabels,omitempty"`
	// Number of replica
	ReplicaCount int `json:"replicaCount,omitempty"`
	// Resource configuration (i.e. CPU & Memory requests and limits)
	Resources ResourcesConfig `json:"resources,omitempty"`
	// kubernetes service account to be mounted to the pod created by the deployment
	KubernatesServiceAccount string `json:"kubernatesServiceAccount,omitempty"`
	// Google Service account to be mounted to the pod created by the deployment
	// The service account is used as the identity when flusing logs to BQ
	GCPServiceAccount GCPServiceAccount `json:"gcpServiceAccount,omitempty"`
	// Persistent volume claim configuration
	PVCConfig PVCConfig `json:"pvcConfig,omitempty"`
	// Additional environment variables to be added to fluentd deployment
	ExtraEnvs []Env `json:"extraEnvs,omitempty"`
	// Autoscaling configuration of the deployment
	Autoscaling AutoscalingConfig `json:"autoscaling,omitempty"`
	// FluentdHelmValues config
	FluentdConfig string `json:"fluentdConfig,omitempty"`
	// FluentdHelmValues enable flag is used when deploying observation service
	Enabled bool `json:"enabled,omitempty"`
}
