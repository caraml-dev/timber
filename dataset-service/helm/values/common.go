package values

// Docker image configuration
type ImageConfig struct {
	// docker registry
	Registry string `json:"registry,omitempty"`
	// docker repository
	Repository string `json:"repository,omitempty"`
	// docker tag
	Tag string `json:"tag,omitempty"`
	// image pull policy
	PullPolicy string `json:"pullPolicy,omitempty"`
}

// K8S Resource configuration
type ResourcesConfig struct {
	// Resource requests
	Requests Resource `json:"requests,omitempty"`
	// Resource limits
	Limits Resource `json:"limits,omitempty"`
}

// Resource
type Resource struct {
	// CPU resource
	CPU string `json:"cpu,omitempty"`
	// Memory resource
	Memory string `json:"memory,omitempty"`
}

// Autoscaling configuration
type AutoscalingConfig struct {
	// Enable/disable autoscaling flag
	Enabled bool `json:"enabled,omitempty"`
	// Minimum number of replicas for the deployment
	MinReplicas int `json:"minReplicas,omitempty"`
	// Maximum number of replicas for the deployment
	MaxReplicas int `json:"maxReplicas,omitempty"`
	// CPU target utilization in percentage
	TargetCPUUtilizationPercentage int `json:"targetCPUUtilizationPercentage,omitempty"`
}

// Persistent volume claim configuration
type PVCConfig struct {
	// Name of pvc
	Name string `json:"name,omitempty"`
	// PVC mount path
	MountPath string `json:"mountPath,omitempty"`
	// Storage type
	Storage string `json:"storage,omitempty"`
}

// Environment variable
type Env struct {
	// Environment variable name
	Name string `json:"name,omitempty"`
	// Environment variable value
	Value string `json:"value,omitempty"`
}

type GCPServiceAccount struct {
	CredentialsData string      `json:"credentialsData,omitempty"`
	Credentials     Credentials `json:"credentials,omitempty"`
}

type Credentials struct {
	Name string `json:"name,omitempty"`
	Key  string `json:"key,omitempty"`
}

type Service struct {
	Type         string `json:"type,omitempty"`
	ExternalPort string `json:"externalPort,omitempty"`
	InternalPort string `json:"internalPort,omitempty"`
}

// MerveEnvs merge 2 slices of Env and give priority for right slice
func MerveEnvs(left []Env, right []Env) []Env {
	for _, e := range right {
		found := false
		for _, d := range left {
			if d.Name == e.Name {
				d.Value = e.Value
				found = true
				break
			}
		}

		if !found {
			left = append(left, e)
		}
	}

	return left
}
