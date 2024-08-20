# ModelsCertificateStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] [readonly] 
**ContainerId** | Pointer to **NullableInt32** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**Storepath** | Pointer to **NullableString** |  | [optional] 
**CertStoreInventoryJobId** | Pointer to **NullableString** |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**CreateIfMissing** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**AgentAssigned** | Pointer to **bool** |  | [optional] 
**ContainerName** | Pointer to **NullableString** |  | [optional] 
**InventorySchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**ReenrollmentStatus** | Pointer to [**ModelsReenrollmentStatus**](ModelsReenrollmentStatus.md) |  | [optional] 
**SetNewPasswordAllowed** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to [**ModelsKeyfactorSecret**](ModelsKeyfactorSecret.md) |  | [optional] 

## Methods

### NewModelsCertificateStore

`func NewModelsCertificateStore() *ModelsCertificateStore`

NewModelsCertificateStore instantiates a new ModelsCertificateStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreWithDefaults

`func NewModelsCertificateStoreWithDefaults() *ModelsCertificateStore`

NewModelsCertificateStoreWithDefaults instantiates a new ModelsCertificateStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateStore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateStore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *ModelsCertificateStore) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelsCertificateStore) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelsCertificateStore) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelsCertificateStore) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ModelsCertificateStore) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ModelsCertificateStore) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetContainerId

`func (o *ModelsCertificateStore) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *ModelsCertificateStore) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *ModelsCertificateStore) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *ModelsCertificateStore) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *ModelsCertificateStore) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *ModelsCertificateStore) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetClientMachine

`func (o *ModelsCertificateStore) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *ModelsCertificateStore) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *ModelsCertificateStore) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *ModelsCertificateStore) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *ModelsCertificateStore) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *ModelsCertificateStore) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetStorepath

`func (o *ModelsCertificateStore) GetStorepath() string`

GetStorepath returns the Storepath field if non-nil, zero value otherwise.

### GetStorepathOk

`func (o *ModelsCertificateStore) GetStorepathOk() (*string, bool)`

GetStorepathOk returns a tuple with the Storepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorepath

`func (o *ModelsCertificateStore) SetStorepath(v string)`

SetStorepath sets Storepath field to given value.

### HasStorepath

`func (o *ModelsCertificateStore) HasStorepath() bool`

HasStorepath returns a boolean if a field has been set.

### SetStorepathNil

`func (o *ModelsCertificateStore) SetStorepathNil(b bool)`

 SetStorepathNil sets the value for Storepath to be an explicit nil

### UnsetStorepath
`func (o *ModelsCertificateStore) UnsetStorepath()`

UnsetStorepath ensures that no value is present for Storepath, not even an explicit nil
### GetCertStoreInventoryJobId

`func (o *ModelsCertificateStore) GetCertStoreInventoryJobId() string`

GetCertStoreInventoryJobId returns the CertStoreInventoryJobId field if non-nil, zero value otherwise.

### GetCertStoreInventoryJobIdOk

`func (o *ModelsCertificateStore) GetCertStoreInventoryJobIdOk() (*string, bool)`

GetCertStoreInventoryJobIdOk returns a tuple with the CertStoreInventoryJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryJobId

`func (o *ModelsCertificateStore) SetCertStoreInventoryJobId(v string)`

SetCertStoreInventoryJobId sets CertStoreInventoryJobId field to given value.

### HasCertStoreInventoryJobId

`func (o *ModelsCertificateStore) HasCertStoreInventoryJobId() bool`

HasCertStoreInventoryJobId returns a boolean if a field has been set.

### SetCertStoreInventoryJobIdNil

`func (o *ModelsCertificateStore) SetCertStoreInventoryJobIdNil(b bool)`

 SetCertStoreInventoryJobIdNil sets the value for CertStoreInventoryJobId to be an explicit nil

### UnsetCertStoreInventoryJobId
`func (o *ModelsCertificateStore) UnsetCertStoreInventoryJobId()`

UnsetCertStoreInventoryJobId ensures that no value is present for CertStoreInventoryJobId, not even an explicit nil
### GetCertStoreType

`func (o *ModelsCertificateStore) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *ModelsCertificateStore) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *ModelsCertificateStore) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *ModelsCertificateStore) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetApproved

`func (o *ModelsCertificateStore) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *ModelsCertificateStore) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *ModelsCertificateStore) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *ModelsCertificateStore) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *ModelsCertificateStore) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *ModelsCertificateStore) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *ModelsCertificateStore) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *ModelsCertificateStore) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *ModelsCertificateStore) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *ModelsCertificateStore) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *ModelsCertificateStore) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *ModelsCertificateStore) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *ModelsCertificateStore) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *ModelsCertificateStore) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAgentId

`func (o *ModelsCertificateStore) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsCertificateStore) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsCertificateStore) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsCertificateStore) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ModelsCertificateStore) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ModelsCertificateStore) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentAssigned

`func (o *ModelsCertificateStore) GetAgentAssigned() bool`

GetAgentAssigned returns the AgentAssigned field if non-nil, zero value otherwise.

### GetAgentAssignedOk

`func (o *ModelsCertificateStore) GetAgentAssignedOk() (*bool, bool)`

GetAgentAssignedOk returns a tuple with the AgentAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAssigned

`func (o *ModelsCertificateStore) SetAgentAssigned(v bool)`

SetAgentAssigned sets AgentAssigned field to given value.

### HasAgentAssigned

`func (o *ModelsCertificateStore) HasAgentAssigned() bool`

HasAgentAssigned returns a boolean if a field has been set.

### GetContainerName

`func (o *ModelsCertificateStore) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *ModelsCertificateStore) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *ModelsCertificateStore) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *ModelsCertificateStore) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *ModelsCertificateStore) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *ModelsCertificateStore) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetInventorySchedule

`func (o *ModelsCertificateStore) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *ModelsCertificateStore) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *ModelsCertificateStore) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *ModelsCertificateStore) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetReenrollmentStatus

`func (o *ModelsCertificateStore) GetReenrollmentStatus() ModelsReenrollmentStatus`

GetReenrollmentStatus returns the ReenrollmentStatus field if non-nil, zero value otherwise.

### GetReenrollmentStatusOk

`func (o *ModelsCertificateStore) GetReenrollmentStatusOk() (*ModelsReenrollmentStatus, bool)`

GetReenrollmentStatusOk returns a tuple with the ReenrollmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenrollmentStatus

`func (o *ModelsCertificateStore) SetReenrollmentStatus(v ModelsReenrollmentStatus)`

SetReenrollmentStatus sets ReenrollmentStatus field to given value.

### HasReenrollmentStatus

`func (o *ModelsCertificateStore) HasReenrollmentStatus() bool`

HasReenrollmentStatus returns a boolean if a field has been set.

### GetSetNewPasswordAllowed

`func (o *ModelsCertificateStore) GetSetNewPasswordAllowed() bool`

GetSetNewPasswordAllowed returns the SetNewPasswordAllowed field if non-nil, zero value otherwise.

### GetSetNewPasswordAllowedOk

`func (o *ModelsCertificateStore) GetSetNewPasswordAllowedOk() (*bool, bool)`

GetSetNewPasswordAllowedOk returns a tuple with the SetNewPasswordAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNewPasswordAllowed

`func (o *ModelsCertificateStore) SetSetNewPasswordAllowed(v bool)`

SetSetNewPasswordAllowed sets SetNewPasswordAllowed field to given value.

### HasSetNewPasswordAllowed

`func (o *ModelsCertificateStore) HasSetNewPasswordAllowed() bool`

HasSetNewPasswordAllowed returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsCertificateStore) GetPassword() ModelsKeyfactorSecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateStore) GetPasswordOk() (*ModelsKeyfactorSecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateStore) SetPassword(v ModelsKeyfactorSecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsCertificateStore) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


