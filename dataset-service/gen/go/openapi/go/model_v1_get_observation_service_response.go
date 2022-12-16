/*
 * caraml/timber/v1/dataset_service.proto
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: version not set
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// V1GetObservationServiceResponse - Response message for GetObservationService method.
type V1GetObservationServiceResponse struct {

	ObservationService V1ObservationServiceConfig `json:"observationService,omitempty"`
}

// AssertV1GetObservationServiceResponseRequired checks if the required fields are not zero-ed
func AssertV1GetObservationServiceResponseRequired(obj V1GetObservationServiceResponse) error {
	if err := AssertV1ObservationServiceConfigRequired(obj.ObservationService); err != nil {
		return err
	}
	return nil
}

// AssertRecurseV1GetObservationServiceResponseRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of V1GetObservationServiceResponse (e.g. [][]V1GetObservationServiceResponse), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseV1GetObservationServiceResponseRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aV1GetObservationServiceResponse, ok := obj.(V1GetObservationServiceResponse)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertV1GetObservationServiceResponseRequired(aV1GetObservationServiceResponse)
	})
}
