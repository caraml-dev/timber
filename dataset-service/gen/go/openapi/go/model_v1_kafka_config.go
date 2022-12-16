/*
 * caraml/timber/v1/dataset_service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type V1KafkaConfig struct {

	Brokers string `json:"brokers,omitempty"`

	Topic string `json:"topic,omitempty"`

	MaxMessageBytes string `json:"maxMessageBytes,omitempty"`

	CompressionType string `json:"compressionType,omitempty"`

	ConnectionTimeout int32 `json:"connectionTimeout,omitempty"`

	PollInterval int32 `json:"pollInterval,omitempty"`

	OffsetReset V1KafkaInitialOffset `json:"offsetReset,omitempty"`
}

// AssertV1KafkaConfigRequired checks if the required fields are not zero-ed
func AssertV1KafkaConfigRequired(obj V1KafkaConfig) error {
	return nil
}

// AssertRecurseV1KafkaConfigRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1KafkaConfig (e.g. [][]V1KafkaConfig), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1KafkaConfigRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1KafkaConfig, ok := obj.(V1KafkaConfig)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1KafkaConfigRequired(aV1KafkaConfig)
	})
}
