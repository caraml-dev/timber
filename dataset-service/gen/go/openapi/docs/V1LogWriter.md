# V1LogWriter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**V1LogWriterType**](V1LogWriterType.md) |  | [optional] [default to V1LOGWRITERTYPE_UNSPECIFIED]
**FluentdConfig** | Pointer to [**V1FluentdConfig**](V1FluentdConfig.md) |  | [optional] 

## Methods

### NewV1LogWriter

`func NewV1LogWriter() *V1LogWriter`

NewV1LogWriter instantiates a new V1LogWriter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LogWriterWithDefaults

`func NewV1LogWriterWithDefaults() *V1LogWriter`

NewV1LogWriterWithDefaults instantiates a new V1LogWriter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *V1LogWriter) GetType() V1LogWriterType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *V1LogWriter) GetTypeOk() (*V1LogWriterType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *V1LogWriter) SetType(v V1LogWriterType)`

SetType sets Type field to given value.

### HasType

`func (o *V1LogWriter) HasType() bool`

HasType returns a boolean if a field has been set.

### GetFluentdConfig

`func (o *V1LogWriter) GetFluentdConfig() V1FluentdConfig`

GetFluentdConfig returns the FluentdConfig field if non-nil, zero value otherwise.

### GetFluentdConfigOk

`func (o *V1LogWriter) GetFluentdConfigOk() (*V1FluentdConfig, bool)`

GetFluentdConfigOk returns a tuple with the FluentdConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFluentdConfig

`func (o *V1LogWriter) SetFluentdConfig(v V1FluentdConfig)`

SetFluentdConfig sets FluentdConfig field to given value.

### HasFluentdConfig

`func (o *V1LogWriter) HasFluentdConfig() bool`

HasFluentdConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


