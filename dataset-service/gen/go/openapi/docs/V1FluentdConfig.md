# V1FluentdConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**V1FluentdOutputType**](V1FluentdOutputType.md) |  | [optional] [default to V1FLUENTDOUTPUTTYPE_UNSPECIFIED]
**Host** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**V1FluentdOutputBQConfig**](V1FluentdOutputBQConfig.md) |  | [optional] 

## Methods

### NewV1FluentdConfig

`func NewV1FluentdConfig() *V1FluentdConfig`

NewV1FluentdConfig instantiates a new V1FluentdConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1FluentdConfigWithDefaults

`func NewV1FluentdConfigWithDefaults() *V1FluentdConfig`

NewV1FluentdConfigWithDefaults instantiates a new V1FluentdConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1FluentdConfig) GetType() V1FluentdOutputType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1FluentdConfig) GetTypeOk() (*V1FluentdOutputType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1FluentdConfig) SetType(v V1FluentdOutputType)`

SetType sets Type field to given value.

### HasType

`func (o *V1FluentdConfig) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHost

`func (o *V1FluentdConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *V1FluentdConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *V1FluentdConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *V1FluentdConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *V1FluentdConfig) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *V1FluentdConfig) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *V1FluentdConfig) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *V1FluentdConfig) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetTag

`func (o *V1FluentdConfig) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *V1FluentdConfig) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *V1FluentdConfig) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *V1FluentdConfig) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetConfig

`func (o *V1FluentdConfig) GetConfig() V1FluentdOutputBQConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *V1FluentdConfig) GetConfigOk() (*V1FluentdOutputBQConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *V1FluentdConfig) SetConfig(v V1FluentdOutputBQConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *V1FluentdConfig) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


