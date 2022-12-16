# V1LogProducer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the producer. | [optional] 
**Name** | Pointer to **string** | Name of the producer, dependent on the type of the log. | [optional] 
**Project** | Pointer to **string** | Name of the CaraML project that hosts the producer. | [optional] 

## Methods

### NewV1LogProducer

`func NewV1LogProducer() *V1LogProducer`

NewV1LogProducer instantiates a new V1LogProducer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1LogProducerWithDefaults

`func NewV1LogProducerWithDefaults() *V1LogProducer`

NewV1LogProducerWithDefaults instantiates a new V1LogProducer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *V1LogProducer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *V1LogProducer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *V1LogProducer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *V1LogProducer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *V1LogProducer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *V1LogProducer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *V1LogProducer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *V1LogProducer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProject

`func (o *V1LogProducer) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *V1LogProducer) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *V1LogProducer) SetProject(v string)`

SetProject sets Project field to given value.

### HasProject

`func (o *V1LogProducer) HasProject() bool`

HasProject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


