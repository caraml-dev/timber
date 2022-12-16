/*
 * caraml/timber/v1/dataset_service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// V1LogType : - LOG_TYPE_PREDICTION: Log from Prediction service  - LOG_TYPE_OBSERVATION: Log from Observation service  - LOG_TYPE_ROUTER: Log from Router service
type V1LogType string

// List of V1LogType
const (
	V1LOGTYPE_UNSPECIFIED V1LogType = "LOG_TYPE_UNSPECIFIED"
	V1LOGTYPE_PREDICTION V1LogType = "LOG_TYPE_PREDICTION"
	V1LOGTYPE_OBSERVATION V1LogType = "LOG_TYPE_OBSERVATION"
	V1LOGTYPE_ROUTER V1LogType = "LOG_TYPE_ROUTER"
)

// AssertV1LogTypeRequired checks if the required fields are not zero-ed
func AssertV1LogTypeRequired(obj V1LogType) error {
	return nil
}

// AssertRecurseV1LogTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1LogType (e.g. [][]V1LogType), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1LogTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1LogType, ok := obj.(V1LogType)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1LogTypeRequired(aV1LogType)
	})
}
