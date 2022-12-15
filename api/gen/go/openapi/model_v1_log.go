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

// V1Log struct for V1Log
type V1Log struct {
	// Unique identifier of a log generated by a LogProducer.
	Id *string `json:"id,omitempty"`
	// Name of the log, generated by Dataset Service.
	Name *string `json:"name,omitempty"`
	Type *V1LogType `json:"type,omitempty"`
	// List of target names associated with a log.
	TargetNames []string `json:"targetNames,omitempty"`
	// BQ table ID where the data is stored.
	BqTable *string `json:"bqTable,omitempty"`
	LogProducer *V1LogProducer `json:"logProducer,omitempty"`
}

// NewV1Log instantiates a new V1Log object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1Log() *V1Log {
	this := V1Log{}
	var type_ V1LogType = V1LOGTYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// NewV1LogWithDefaults instantiates a new V1Log object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1LogWithDefaults() *V1Log {
	this := V1Log{}
	var type_ V1LogType = V1LOGTYPE_UNSPECIFIED
	this.Type = &type_
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *V1Log) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Log) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *V1Log) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *V1Log) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *V1Log) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Log) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *V1Log) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *V1Log) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *V1Log) GetType() V1LogType {
	if o == nil || o.Type == nil {
		var ret V1LogType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Log) GetTypeOk() (*V1LogType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *V1Log) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given V1LogType and assigns it to the Type field.
func (o *V1Log) SetType(v V1LogType) {
	o.Type = &v
}

// GetTargetNames returns the TargetNames field value if set, zero value otherwise.
func (o *V1Log) GetTargetNames() []string {
	if o == nil || o.TargetNames == nil {
		var ret []string
		return ret
	}
	return o.TargetNames
}

// GetTargetNamesOk returns a tuple with the TargetNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Log) GetTargetNamesOk() ([]string, bool) {
	if o == nil || o.TargetNames == nil {
		return nil, false
	}
	return o.TargetNames, true
}

// HasTargetNames returns a boolean if a field has been set.
func (o *V1Log) HasTargetNames() bool {
	if o != nil && o.TargetNames != nil {
		return true
	}

	return false
}

// SetTargetNames gets a reference to the given []string and assigns it to the TargetNames field.
func (o *V1Log) SetTargetNames(v []string) {
	o.TargetNames = v
}

// GetBqTable returns the BqTable field value if set, zero value otherwise.
func (o *V1Log) GetBqTable() string {
	if o == nil || o.BqTable == nil {
		var ret string
		return ret
	}
	return *o.BqTable
}

// GetBqTableOk returns a tuple with the BqTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Log) GetBqTableOk() (*string, bool) {
	if o == nil || o.BqTable == nil {
		return nil, false
	}
	return o.BqTable, true
}

// HasBqTable returns a boolean if a field has been set.
func (o *V1Log) HasBqTable() bool {
	if o != nil && o.BqTable != nil {
		return true
	}

	return false
}

// SetBqTable gets a reference to the given string and assigns it to the BqTable field.
func (o *V1Log) SetBqTable(v string) {
	o.BqTable = &v
}

// GetLogProducer returns the LogProducer field value if set, zero value otherwise.
func (o *V1Log) GetLogProducer() V1LogProducer {
	if o == nil || o.LogProducer == nil {
		var ret V1LogProducer
		return ret
	}
	return *o.LogProducer
}

// GetLogProducerOk returns a tuple with the LogProducer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1Log) GetLogProducerOk() (*V1LogProducer, bool) {
	if o == nil || o.LogProducer == nil {
		return nil, false
	}
	return o.LogProducer, true
}

// HasLogProducer returns a boolean if a field has been set.
func (o *V1Log) HasLogProducer() bool {
	if o != nil && o.LogProducer != nil {
		return true
	}

	return false
}

// SetLogProducer gets a reference to the given V1LogProducer and assigns it to the LogProducer field.
func (o *V1Log) SetLogProducer(v V1LogProducer) {
	o.LogProducer = &v
}

func (o V1Log) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TargetNames != nil {
		toSerialize["targetNames"] = o.TargetNames
	}
	if o.BqTable != nil {
		toSerialize["bqTable"] = o.BqTable
	}
	if o.LogProducer != nil {
		toSerialize["logProducer"] = o.LogProducer
	}
	return json.Marshal(toSerialize)
}

type NullableV1Log struct {
	value *V1Log
	isSet bool
}

func (v NullableV1Log) Get() *V1Log {
	return v.value
}

func (v *NullableV1Log) Set(val *V1Log) {
	v.value = val
	v.isSet = true
}

func (v NullableV1Log) IsSet() bool {
	return v.isSet
}

func (v *NullableV1Log) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1Log(val *V1Log) *NullableV1Log {
	return &NullableV1Log{value: val, isSet: true}
}

func (v NullableV1Log) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1Log) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


