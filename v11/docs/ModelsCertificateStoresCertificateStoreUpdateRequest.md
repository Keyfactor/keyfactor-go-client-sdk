# ModelsCertificateStoresCertificateStoreUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **NullableInt32** |  | [optional] 
**CreateIfMissing** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**InventorySchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**Password** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewModelsCertificateStoresCertificateStoreUpdateRequest

`func NewModelsCertificateStoresCertificateStoreUpdateRequest() *ModelsCertificateStoresCertificateStoreUpdateRequest`

NewModelsCertificateStoresCertificateStoreUpdateRequest instantiates a new ModelsCertificateStoresCertificateStoreUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoresCertificateStoreUpdateRequestWithDefaults

`func NewModelsCertificateStoresCertificateStoreUpdateRequestWithDefaults() *ModelsCertificateStoresCertificateStoreUpdateRequest`

NewModelsCertificateStoresCertificateStoreUpdateRequestWithDefaults instantiates a new ModelsCertificateStoresCertificateStoreUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContainerId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetCreateIfMissing

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAgentId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetInventorySchedule

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetPassword() ModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) SetPassword(v ModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsCertificateStoresCertificateStoreUpdateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


