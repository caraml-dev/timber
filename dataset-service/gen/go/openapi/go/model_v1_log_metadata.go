/*
 * caraml/timber/v1/dataset_service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type V1LogMetadata struct {

	// Unique identifier of a log generated by a LogProducer.
	Id string `json:"id,omitempty"`

	// Name of the log, generated by Dataset Service.
	Name string `json:"name,omitempty"`

	Type V1LogType `json:"type,omitempty"`

	// List of target names associated with a log.
	TargetNames []string `json:"targetNames,omitempty"`

	// BQ table ID where the data is stored.
	BqTable string `json:"bqTable,omitempty"`

	LogProducer V1LogProducer `json:"logProducer,omitempty"`
}

// AssertV1LogMetadataRequired checks if the required fields are not zero-ed
func AssertV1LogMetadataRequired(obj V1LogMetadata) error {
	if err := AssertV1LogProducerRequired(obj.LogProducer); err != nil {
		return err
	}
	return nil
}

// AssertRecurseV1LogMetadataRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1LogMetadata (e.g. [][]V1LogMetadata), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1LogMetadataRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1LogMetadata, ok := obj.(V1LogMetadata)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1LogMetadataRequired(aV1LogMetadata)
	})
}
