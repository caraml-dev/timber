/*
caraml/timber/v1/dataset_service.proto

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: version not set
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// V1CreateObservationServiceResponse Response message for CreateObservationService method.
type V1CreateObservationServiceResponse struct {
	ObservationService *V1ObservationServiceConfig `json:"observationService,omitempty"`
}

// NewV1CreateObservationServiceResponse instantiates a new V1CreateObservationServiceResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1CreateObservationServiceResponse() *V1CreateObservationServiceResponse {
	this := V1CreateObservationServiceResponse{}
	return &this
}

// NewV1CreateObservationServiceResponseWithDefaults instantiates a new V1CreateObservationServiceResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1CreateObservationServiceResponseWithDefaults() *V1CreateObservationServiceResponse {
	this := V1CreateObservationServiceResponse{}
	return &this
}

// GetObservationService returns the ObservationService field value if set, zero value otherwise.
func (o *V1CreateObservationServiceResponse) GetObservationService() V1ObservationServiceConfig {
	if o == nil || o.ObservationService == nil {
		var ret V1ObservationServiceConfig
		return ret
	}
	return *o.ObservationService
}

// GetObservationServiceOk returns a tuple with the ObservationService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1CreateObservationServiceResponse) GetObservationServiceOk() (*V1ObservationServiceConfig, bool) {
	if o == nil || o.ObservationService == nil {
		return nil, false
	}
	return o.ObservationService, true
}

// HasObservationService returns a boolean if a field has been set.
func (o *V1CreateObservationServiceResponse) HasObservationService() bool {
	if o != nil && o.ObservationService != nil {
		return true
	}

	return false
}

// SetObservationService gets a reference to the given V1ObservationServiceConfig and assigns it to the ObservationService field.
func (o *V1CreateObservationServiceResponse) SetObservationService(v V1ObservationServiceConfig) {
	o.ObservationService = &v
}

func (o V1CreateObservationServiceResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ObservationService != nil {
		toSerialize["observationService"] = o.ObservationService
	}
	return json.Marshal(toSerialize)
}

type NullableV1CreateObservationServiceResponse struct {
	value *V1CreateObservationServiceResponse
	isSet bool
}

func (v NullableV1CreateObservationServiceResponse) Get() *V1CreateObservationServiceResponse {
	return v.value
}

func (v *NullableV1CreateObservationServiceResponse) Set(val *V1CreateObservationServiceResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableV1CreateObservationServiceResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableV1CreateObservationServiceResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1CreateObservationServiceResponse(val *V1CreateObservationServiceResponse) *NullableV1CreateObservationServiceResponse {
	return &NullableV1CreateObservationServiceResponse{value: val, isSet: true}
}

func (v NullableV1CreateObservationServiceResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1CreateObservationServiceResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


