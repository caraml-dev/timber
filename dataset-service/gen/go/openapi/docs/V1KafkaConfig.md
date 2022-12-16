# V1KafkaConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brokers** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 
**MaxMessageBytes** | Pointer to **string** |  | [optional] 
**CompressionType** | Pointer to **string** |  | [optional] 
**ConnectionTimeout** | Pointer to **int32** |  | [optional] 
**PollInterval** | Pointer to **int32** |  | [optional] 
**OffsetReset** | Pointer to [**V1KafkaInitialOffset**](V1KafkaInitialOffset.md) |  | [optional] [default to V1KAFKAINITIALOFFSET_UNSPECIFIED]

## Methods

### NewV1KafkaConfig

`func NewV1KafkaConfig() *V1KafkaConfig`

NewV1KafkaConfig instantiates a new V1KafkaConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1KafkaConfigWithDefaults

`func NewV1KafkaConfigWithDefaults() *V1KafkaConfig`

NewV1KafkaConfigWithDefaults instantiates a new V1KafkaConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrokers

`func (o *V1KafkaConfig) GetBrokers() string`

GetBrokers returns the Brokers field if non-nil, zero value otherwise.

### GetBrokersOk

`func (o *V1KafkaConfig) GetBrokersOk() (*string, bool)`

GetBrokersOk returns a tuple with the Brokers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrokers

`func (o *V1KafkaConfig) SetBrokers(v string)`

SetBrokers sets Brokers field to given value.

### HasBrokers

`func (o *V1KafkaConfig) HasBrokers() bool`

HasBrokers returns a boolean if a field has been set.

### GetTopic

`func (o *V1KafkaConfig) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *V1KafkaConfig) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *V1KafkaConfig) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *V1KafkaConfig) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetMaxMessageBytes

`func (o *V1KafkaConfig) GetMaxMessageBytes() string`

GetMaxMessageBytes returns the MaxMessageBytes field if non-nil, zero value otherwise.

### GetMaxMessageBytesOk

`func (o *V1KafkaConfig) GetMaxMessageBytesOk() (*string, bool)`

GetMaxMessageBytesOk returns a tuple with the MaxMessageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMessageBytes

`func (o *V1KafkaConfig) SetMaxMessageBytes(v string)`

SetMaxMessageBytes sets MaxMessageBytes field to given value.

### HasMaxMessageBytes

`func (o *V1KafkaConfig) HasMaxMessageBytes() bool`

HasMaxMessageBytes returns a boolean if a field has been set.

### GetCompressionType

`func (o *V1KafkaConfig) GetCompressionType() string`

GetCompressionType returns the CompressionType field if non-nil, zero value otherwise.

### GetCompressionTypeOk

`func (o *V1KafkaConfig) GetCompressionTypeOk() (*string, bool)`

GetCompressionTypeOk returns a tuple with the CompressionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompressionType

`func (o *V1KafkaConfig) SetCompressionType(v string)`

SetCompressionType sets CompressionType field to given value.

### HasCompressionType

`func (o *V1KafkaConfig) HasCompressionType() bool`

HasCompressionType returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *V1KafkaConfig) GetConnectionTimeout() int32`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *V1KafkaConfig) GetConnectionTimeoutOk() (*int32, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *V1KafkaConfig) SetConnectionTimeout(v int32)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *V1KafkaConfig) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetPollInterval

`func (o *V1KafkaConfig) GetPollInterval() int32`

GetPollInterval returns the PollInterval field if non-nil, zero value otherwise.

### GetPollIntervalOk

`func (o *V1KafkaConfig) GetPollIntervalOk() (*int32, bool)`

GetPollIntervalOk returns a tuple with the PollInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollInterval

`func (o *V1KafkaConfig) SetPollInterval(v int32)`

SetPollInterval sets PollInterval field to given value.

### HasPollInterval

`func (o *V1KafkaConfig) HasPollInterval() bool`

HasPollInterval returns a boolean if a field has been set.

### GetOffsetReset

`func (o *V1KafkaConfig) GetOffsetReset() V1KafkaInitialOffset`

GetOffsetReset returns the OffsetReset field if non-nil, zero value otherwise.

### GetOffsetResetOk

`func (o *V1KafkaConfig) GetOffsetResetOk() (*V1KafkaInitialOffset, bool)`

GetOffsetResetOk returns a tuple with the OffsetReset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsetReset

`func (o *V1KafkaConfig) SetOffsetReset(v V1KafkaInitialOffset)`

SetOffsetReset sets OffsetReset field to given value.

### HasOffsetReset

`func (o *V1KafkaConfig) HasOffsetReset() bool`

HasOffsetReset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


