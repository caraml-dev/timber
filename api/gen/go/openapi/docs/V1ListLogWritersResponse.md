# V1ListLogWritersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogWriters** | Pointer to [**[]V1LogWriter**](V1LogWriter.md) | Log Writers deployed by Dataset Service for a particular CaraML project. | [optional] 

## Methods

### NewV1ListLogWritersResponse

`func NewV1ListLogWritersResponse() *V1ListLogWritersResponse`

NewV1ListLogWritersResponse instantiates a new V1ListLogWritersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListLogWritersResponseWithDefaults

`func NewV1ListLogWritersResponseWithDefaults() *V1ListLogWritersResponse`

NewV1ListLogWritersResponseWithDefaults instantiates a new V1ListLogWritersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogWriters

`func (o *V1ListLogWritersResponse) GetLogWriters() []V1LogWriter`

GetLogWriters returns the LogWriters field if non-nil, zero value otherwise.

### GetLogWritersOk

`func (o *V1ListLogWritersResponse) GetLogWritersOk() (*[]V1LogWriter, bool)`

GetLogWritersOk returns a tuple with the LogWriters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogWriters

`func (o *V1ListLogWritersResponse) SetLogWriters(v []V1LogWriter)`

SetLogWriters sets LogWriters field to given value.

### HasLogWriters

`func (o *V1ListLogWritersResponse) HasLogWriters() bool`

HasLogWriters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


