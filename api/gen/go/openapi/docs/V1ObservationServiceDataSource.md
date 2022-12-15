# V1ObservationServiceDataSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**V1ObservationServiceDataSourceType**](V1ObservationServiceDataSourceType.md) |  | [optional] [default to V1OBSERVATIONSERVICEDATASOURCETYPE_UNSPECIFIED]
**KafkaConfig** | Pointer to [**V1KafkaConfig**](V1KafkaConfig.md) |  | [optional] 

## Methods

### NewV1ObservationServiceDataSource

`func NewV1ObservationServiceDataSource() *V1ObservationServiceDataSource`

NewV1ObservationServiceDataSource instantiates a new V1ObservationServiceDataSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ObservationServiceDataSourceWithDefaults

`func NewV1ObservationServiceDataSourceWithDefaults() *V1ObservationServiceDataSource`

NewV1ObservationServiceDataSourceWithDefaults instantiates a new V1ObservationServiceDataSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1ObservationServiceDataSource) GetType() V1ObservationServiceDataSourceType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1ObservationServiceDataSource) GetTypeOk() (*V1ObservationServiceDataSourceType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1ObservationServiceDataSource) SetType(v V1ObservationServiceDataSourceType)`

SetType sets Type field to given value.

### HasType

`func (o *V1ObservationServiceDataSource) HasType() bool`

HasType returns a boolean if a field has been set.

### GetKafkaConfig

`func (o *V1ObservationServiceDataSource) GetKafkaConfig() V1KafkaConfig`

GetKafkaConfig returns the KafkaConfig field if non-nil, zero value otherwise.

### GetKafkaConfigOk

`func (o *V1ObservationServiceDataSource) GetKafkaConfigOk() (*V1KafkaConfig, bool)`

GetKafkaConfigOk returns a tuple with the KafkaConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafkaConfig

`func (o *V1ObservationServiceDataSource) SetKafkaConfig(v V1KafkaConfig)`

SetKafkaConfig sets KafkaConfig field to given value.

### HasKafkaConfig

`func (o *V1ObservationServiceDataSource) HasKafkaConfig() bool`

HasKafkaConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


