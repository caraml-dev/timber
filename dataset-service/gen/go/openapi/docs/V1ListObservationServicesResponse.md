# V1ListObservationServicesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservationServices** | Pointer to [**[]V1ObservationServiceConfig**](V1ObservationServiceConfig.md) | Observation Services deployed by Dataset Service for a particular CaraML project. | [optional] 

## Methods

### NewV1ListObservationServicesResponse

`func NewV1ListObservationServicesResponse() *V1ListObservationServicesResponse`

NewV1ListObservationServicesResponse instantiates a new V1ListObservationServicesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListObservationServicesResponseWithDefaults

`func NewV1ListObservationServicesResponseWithDefaults() *V1ListObservationServicesResponse`

NewV1ListObservationServicesResponseWithDefaults instantiates a new V1ListObservationServicesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservationServices

`func (o *V1ListObservationServicesResponse) GetObservationServices() []V1ObservationServiceConfig`

GetObservationServices returns the ObservationServices field if non-nil, zero value otherwise.

### GetObservationServicesOk

`func (o *V1ListObservationServicesResponse) GetObservationServicesOk() (*[]V1ObservationServiceConfig, bool)`

GetObservationServicesOk returns a tuple with the ObservationServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationServices

`func (o *V1ListObservationServicesResponse) SetObservationServices(v []V1ObservationServiceConfig)`

SetObservationServices sets ObservationServices field to given value.

### HasObservationServices

`func (o *V1ListObservationServicesResponse) HasObservationServices() bool`

HasObservationServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


