# V1ObservationServiceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of an Observation Service deployed by Dataset Service. | [optional] 
**Source** | Pointer to [**V1ObservationServiceDataSource**](V1ObservationServiceDataSource.md) |  | [optional] 
**Sink** | Pointer to [**V1ObservationServiceDataSink**](V1ObservationServiceDataSink.md) |  | [optional] 

## Methods

### NewV1ObservationServiceConfig

`func NewV1ObservationServiceConfig() *V1ObservationServiceConfig`

NewV1ObservationServiceConfig instantiates a new V1ObservationServiceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ObservationServiceConfigWithDefaults

`func NewV1ObservationServiceConfigWithDefaults() *V1ObservationServiceConfig`

NewV1ObservationServiceConfigWithDefaults instantiates a new V1ObservationServiceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1ObservationServiceConfig) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1ObservationServiceConfig) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1ObservationServiceConfig) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1ObservationServiceConfig) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSource

`func (o *V1ObservationServiceConfig) GetSource() V1ObservationServiceDataSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *V1ObservationServiceConfig) GetSourceOk() (*V1ObservationServiceDataSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *V1ObservationServiceConfig) SetSource(v V1ObservationServiceDataSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *V1ObservationServiceConfig) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSink

`func (o *V1ObservationServiceConfig) GetSink() V1ObservationServiceDataSink`

GetSink returns the Sink field if non-nil, zero value otherwise.

### GetSinkOk

`func (o *V1ObservationServiceConfig) GetSinkOk() (*V1ObservationServiceDataSink, bool)`

GetSinkOk returns a tuple with the Sink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSink

`func (o *V1ObservationServiceConfig) SetSink(v V1ObservationServiceDataSink)`

SetSink sets Sink field to given value.

### HasSink

`func (o *V1ObservationServiceConfig) HasSink() bool`

HasSink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


