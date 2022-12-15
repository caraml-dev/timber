# V1ObservationServiceDataSink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**V1ObservationServiceDataSinkType**](V1ObservationServiceDataSinkType.md) |  | [optional] [default to V1OBSERVATIONSERVICEDATASINKTYPE_UNSPECIFIED]
**KafkaConfig** | Pointer to [**V1KafkaConfig**](V1KafkaConfig.md) |  | [optional] 
**FluentdConfig** | Pointer to [**V1FluentdConfig**](V1FluentdConfig.md) |  | [optional] 

## Methods

### NewV1ObservationServiceDataSink

`func NewV1ObservationServiceDataSink() *V1ObservationServiceDataSink`

NewV1ObservationServiceDataSink instantiates a new V1ObservationServiceDataSink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ObservationServiceDataSinkWithDefaults

`func NewV1ObservationServiceDataSinkWithDefaults() *V1ObservationServiceDataSink`

NewV1ObservationServiceDataSinkWithDefaults instantiates a new V1ObservationServiceDataSink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1ObservationServiceDataSink) GetType() V1ObservationServiceDataSinkType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ObservationServiceDataSink) GetTypeOk() (*V1ObservationServiceDataSinkType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ObservationServiceDataSink) SetType(v V1ObservationServiceDataSinkType)`

SetType sets Type field to given value.

### HasType

`func (o *V1ObservationServiceDataSink) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKafkaConfig

`func (o *V1ObservationServiceDataSink) GetKafkaConfig() V1KafkaConfig`

GetKafkaConfig returns the KafkaConfig field if non-nil, zero value otherwise.

### GetKafkaConfigOk

`func (o *V1ObservationServiceDataSink) GetKafkaConfigOk() (*V1KafkaConfig, bool)`

GetKafkaConfigOk returns a tuple with the KafkaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaConfig

`func (o *V1ObservationServiceDataSink) SetKafkaConfig(v V1KafkaConfig)`

SetKafkaConfig sets KafkaConfig field to given value.

### HasKafkaConfig

`func (o *V1ObservationServiceDataSink) HasKafkaConfig() bool`

HasKafkaConfig returns a boolean if a field has been set.

### GetFluentdConfig

`func (o *V1ObservationServiceDataSink) GetFluentdConfig() V1FluentdConfig`

GetFluentdConfig returns the FluentdConfig field if non-nil, zero value otherwise.

### GetFluentdConfigOk

`func (o *V1ObservationServiceDataSink) GetFluentdConfigOk() (*V1FluentdConfig, bool)`

GetFluentdConfigOk returns a tuple with the FluentdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluentdConfig

`func (o *V1ObservationServiceDataSink) SetFluentdConfig(v V1FluentdConfig)`

SetFluentdConfig sets FluentdConfig field to given value.

### HasFluentdConfig

`func (o *V1ObservationServiceDataSink) HasFluentdConfig() bool`

HasFluentdConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


