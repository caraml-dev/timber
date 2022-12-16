# V1ListLogsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Logs** | Pointer to [**[]V1Log**](V1Log.md) | Logs stored in configured Dataset Service storage sink. | [optional] 

## Methods

### NewV1ListLogsResponse

`func NewV1ListLogsResponse() *V1ListLogsResponse`

NewV1ListLogsResponse instantiates a new V1ListLogsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListLogsResponseWithDefaults

`func NewV1ListLogsResponseWithDefaults() *V1ListLogsResponse`

NewV1ListLogsResponseWithDefaults instantiates a new V1ListLogsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogs

`func (o *V1ListLogsResponse) GetLogs() []V1Log`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *V1ListLogsResponse) GetLogsOk() (*[]V1Log, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *V1ListLogsResponse) SetLogs(v []V1Log)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *V1ListLogsResponse) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


