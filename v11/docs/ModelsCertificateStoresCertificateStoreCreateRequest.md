# ModelsCertificateStoresCertificateStoreCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | Pointer to **NullableInt32** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**Storepath** | Pointer to **NullableString** |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**CreateIfMissing** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**AgentAssigned** | Pointer to **bool** |  | [optional] 
**InventorySchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**Password** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewModelsCertificateStoresCertificateStoreCreateRequest

`func NewModelsCertificateStoresCertificateStoreCreateRequest() *ModelsCertificateStoresCertificateStoreCreateRequest`

NewModelsCertificateStoresCertificateStoreCreateRequest instantiates a new ModelsCertificateStoresCertificateStoreCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoresCertificateStoreCreateRequestWithDefaults

`func NewModelsCertificateStoresCertificateStoreCreateRequestWithDefaults() *ModelsCertificateStoresCertificateStoreCreateRequest`

NewModelsCertificateStoresCertificateStoreCreateRequestWithDefaults instantiates a new ModelsCertificateStoresCertificateStoreCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetClientMachine

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetStorepath

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetStorepath() string`

GetStorepath returns the Storepath field if non-nil, zero value otherwise.

### GetStorepathOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetStorepathOk() (*string, bool)`

GetStorepathOk returns a tuple with the Storepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorepath

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetStorepath(v string)`

SetStorepath sets Storepath field to given value.

### HasStorepath

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasStorepath() bool`

HasStorepath returns a boolean if a field has been set.

### SetStorepathNil

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetStorepathNil(b bool)`

 SetStorepathNil sets the value for Storepath to be an explicit nil

### UnsetStorepath
`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) UnsetStorepath()`

UnsetStorepath ensures that no value is present for Storepath, not even an explicit nil
### GetCertStoreType

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAgentId

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentAssigned

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentAssigned() bool`

GetAgentAssigned returns the AgentAssigned field if non-nil, zero value otherwise.

### GetAgentAssignedOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetAgentAssignedOk() (*bool, bool)`

GetAgentAssignedOk returns a tuple with the AgentAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAssigned

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetAgentAssigned(v bool)`

SetAgentAssigned sets AgentAssigned field to given value.

### HasAgentAssigned

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasAgentAssigned() bool`

HasAgentAssigned returns a boolean if a field has been set.

### GetInventorySchedule

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetPassword() ModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) SetPassword(v ModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsCertificateStoresCertificateStoreCreateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


