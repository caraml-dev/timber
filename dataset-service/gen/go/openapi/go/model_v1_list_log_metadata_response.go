/*
 * caraml/timber/v1/dataset_service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// V1ListLogMetadataResponse - Response message for ListLogMetadata method.
type V1ListLogMetadataResponse struct {

	// Log metadata stored in configured Dataset Service storage sink.
	LogMetadata []V1LogMetadata `json:"logMetadata,omitempty"`
}

// AssertV1ListLogMetadataResponseRequired checks if the required fields are not zero-ed
func AssertV1ListLogMetadataResponseRequired(obj V1ListLogMetadataResponse) error {
	for _, el := range obj.LogMetadata {
		if err := AssertV1LogMetadataRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRecurseV1ListLogMetadataResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1ListLogMetadataResponse (e.g. [][]V1ListLogMetadataResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1ListLogMetadataResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1ListLogMetadataResponse, ok := obj.(V1ListLogMetadataResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1ListLogMetadataResponseRequired(aV1ListLogMetadataResponse)
	})
}
