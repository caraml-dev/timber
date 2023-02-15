package values

import (
	"encoding/json"

	"github.com/jinzhu/copier"
)

// ImageConfig Docker image configuration
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

// ResourcesConfig resource configurations for controlling requests/limits
type ResourcesConfig struct {
	// Resource requests
	Requests Resource `json:"requests,omitempty"`
	// Resource limits
	Limits Resource `json:"limits,omitempty"`
}

// Resource CPU and Memory resource configuration
type Resource struct {
	// CPU resource
	CPU string `json:"cpu,omitempty"`
	// Memory resource
	Memory string `json:"memory,omitempty"`
}

// AutoscalingConfig Autoscaling configuration
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

// PVCConfig Persistent volume claim configuration
type PVCConfig struct {
	// Name of pvc
	Name string `json:"name,omitempty"`
	// PVC mount path
	MountPath string `json:"mountPath,omitempty"`
	// Storage type
	Storage string `json:"storage,omitempty"`
}

// Env environment variable
type Env struct {
	// Environment variable name
	Name string `json:"name,omitempty"`
	// Environment variable value
	Value string `json:"value,omitempty"`
}

// GCPServiceAccount configuration for setting the GCP service account to use
type GCPServiceAccount struct {
	// String containing base64 of the GCP service account json
	CredentialsData string `json:"credentialsData,omitempty"`
	// Credentials allow mounting an existing secret
	Credentials Credentials `json:"credentials,omitempty"`
}

// Credentials existing secret
type Credentials struct {
	// Name of secret
	Name string `json:"name,omitempty"`
	// Key of secret
	Key string `json:"key,omitempty"`
}

// Service K8S service configuration
type Service struct {
	// Type of the service
	Type string `json:"type,omitempty"`
	// Port exposed by the service
	ExternalPort string `json:"externalPort,omitempty"`
	// Pod's port being mapped to the external port
	InternalPort string `json:"internalPort,omitempty"`
}

// MerveEnvs merge 2 slices of Env and give priority for right slice
func MerveEnvs(left []Env, right []Env) []Env {
	newSlice := make([]Env, len(left))

	_ = copier.CopyWithOption(&newSlice, &left, copier.Option{IgnoreEmpty: true, DeepCopy: true})
	for _, e := range right {
		found := false
		for i, d := range newSlice {
			if d.Name == e.Name {
				newSlice[i].Value = e.Value
				found = true
				break
			}
		}

		if !found {
			newSlice = append(newSlice, e)
		}
	}

	return newSlice
}

// ToRaw converts struct value to map[string]any
func ToRaw(val any) (map[string]any, error) {
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
