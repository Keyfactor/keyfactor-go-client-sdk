# CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest

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
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest

`func NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest() *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest`

NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest instantiates a new CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequestWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequestWithDefaults() *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest`

NewCSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequestWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetClientMachine

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetStorepath

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetStorepath() string`

GetStorepath returns the Storepath field if non-nil, zero value otherwise.

### GetStorepathOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetStorepathOk() (*string, bool)`

GetStorepathOk returns a tuple with the Storepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorepath

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetStorepath(v string)`

SetStorepath sets Storepath field to given value.

### HasStorepath

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasStorepath() bool`

HasStorepath returns a boolean if a field has been set.

### SetStorepathNil

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetStorepathNil(b bool)`

 SetStorepathNil sets the value for Storepath to be an explicit nil

### UnsetStorepath
`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) UnsetStorepath()`

UnsetStorepath ensures that no value is present for Storepath, not even an explicit nil
### GetCertStoreType

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAgentId

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentAssigned

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetAgentAssigned() bool`

GetAgentAssigned returns the AgentAssigned field if non-nil, zero value otherwise.

### GetAgentAssignedOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetAgentAssignedOk() (*bool, bool)`

GetAgentAssignedOk returns a tuple with the AgentAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAssigned

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetAgentAssigned(v bool)`

SetAgentAssigned sets AgentAssigned field to given value.

### HasAgentAssigned

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasAgentAssigned() bool`

HasAgentAssigned returns a boolean if a field has been set.

### GetInventorySchedule

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetPassword

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) SetPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CSSCMSDataModelModelsCertificateStoresCertificateStoreCreateRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


