# CertificateStoresCertificateStoreResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **NullableInt32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
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
**ReenrollmentStatus** | Pointer to [**CSSCMSDataModelModelsReenrollmentStatus**](CSSCMSDataModelModelsReenrollmentStatus.md) |  | [optional] 
**SetNewPasswordAllowed** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorSecret**](CSSCMSDataModelModelsKeyfactorSecret.md) |  | [optional] 

## Methods

### NewCertificateStoresCertificateStoreResponse

`func NewCertificateStoresCertificateStoreResponse() *CertificateStoresCertificateStoreResponse`

NewCertificateStoresCertificateStoreResponse instantiates a new CertificateStoresCertificateStoreResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresCertificateStoreResponseWithDefaults

`func NewCertificateStoresCertificateStoreResponseWithDefaults() *CertificateStoresCertificateStoreResponse`

NewCertificateStoresCertificateStoreResponseWithDefaults instantiates a new CertificateStoresCertificateStoreResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateStoresCertificateStoreResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateStoresCertificateStoreResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateStoresCertificateStoreResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateStoresCertificateStoreResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContainerId

`func (o *CertificateStoresCertificateStoreResponse) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *CertificateStoresCertificateStoreResponse) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *CertificateStoresCertificateStoreResponse) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *CertificateStoresCertificateStoreResponse) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *CertificateStoresCertificateStoreResponse) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *CertificateStoresCertificateStoreResponse) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetDisplayName

`func (o *CertificateStoresCertificateStoreResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CertificateStoresCertificateStoreResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CertificateStoresCertificateStoreResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CertificateStoresCertificateStoreResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CertificateStoresCertificateStoreResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CertificateStoresCertificateStoreResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetClientMachine

`func (o *CertificateStoresCertificateStoreResponse) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CertificateStoresCertificateStoreResponse) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CertificateStoresCertificateStoreResponse) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CertificateStoresCertificateStoreResponse) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CertificateStoresCertificateStoreResponse) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CertificateStoresCertificateStoreResponse) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetStorepath

`func (o *CertificateStoresCertificateStoreResponse) GetStorepath() string`

GetStorepath returns the Storepath field if non-nil, zero value otherwise.

### GetStorepathOk

`func (o *CertificateStoresCertificateStoreResponse) GetStorepathOk() (*string, bool)`

GetStorepathOk returns a tuple with the Storepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorepath

`func (o *CertificateStoresCertificateStoreResponse) SetStorepath(v string)`

SetStorepath sets Storepath field to given value.

### HasStorepath

`func (o *CertificateStoresCertificateStoreResponse) HasStorepath() bool`

HasStorepath returns a boolean if a field has been set.

### SetStorepathNil

`func (o *CertificateStoresCertificateStoreResponse) SetStorepathNil(b bool)`

 SetStorepathNil sets the value for Storepath to be an explicit nil

### UnsetStorepath
`func (o *CertificateStoresCertificateStoreResponse) UnsetStorepath()`

UnsetStorepath ensures that no value is present for Storepath, not even an explicit nil
### GetCertStoreInventoryJobId

`func (o *CertificateStoresCertificateStoreResponse) GetCertStoreInventoryJobId() string`

GetCertStoreInventoryJobId returns the CertStoreInventoryJobId field if non-nil, zero value otherwise.

### GetCertStoreInventoryJobIdOk

`func (o *CertificateStoresCertificateStoreResponse) GetCertStoreInventoryJobIdOk() (*string, bool)`

GetCertStoreInventoryJobIdOk returns a tuple with the CertStoreInventoryJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryJobId

`func (o *CertificateStoresCertificateStoreResponse) SetCertStoreInventoryJobId(v string)`

SetCertStoreInventoryJobId sets CertStoreInventoryJobId field to given value.

### HasCertStoreInventoryJobId

`func (o *CertificateStoresCertificateStoreResponse) HasCertStoreInventoryJobId() bool`

HasCertStoreInventoryJobId returns a boolean if a field has been set.

### SetCertStoreInventoryJobIdNil

`func (o *CertificateStoresCertificateStoreResponse) SetCertStoreInventoryJobIdNil(b bool)`

 SetCertStoreInventoryJobIdNil sets the value for CertStoreInventoryJobId to be an explicit nil

### UnsetCertStoreInventoryJobId
`func (o *CertificateStoresCertificateStoreResponse) UnsetCertStoreInventoryJobId()`

UnsetCertStoreInventoryJobId ensures that no value is present for CertStoreInventoryJobId, not even an explicit nil
### GetCertStoreType

`func (o *CertificateStoresCertificateStoreResponse) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *CertificateStoresCertificateStoreResponse) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *CertificateStoresCertificateStoreResponse) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *CertificateStoresCertificateStoreResponse) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetApproved

`func (o *CertificateStoresCertificateStoreResponse) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *CertificateStoresCertificateStoreResponse) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *CertificateStoresCertificateStoreResponse) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *CertificateStoresCertificateStoreResponse) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *CertificateStoresCertificateStoreResponse) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *CertificateStoresCertificateStoreResponse) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *CertificateStoresCertificateStoreResponse) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *CertificateStoresCertificateStoreResponse) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateStoresCertificateStoreResponse) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateStoresCertificateStoreResponse) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateStoresCertificateStoreResponse) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateStoresCertificateStoreResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateStoresCertificateStoreResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateStoresCertificateStoreResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAgentId

`func (o *CertificateStoresCertificateStoreResponse) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CertificateStoresCertificateStoreResponse) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CertificateStoresCertificateStoreResponse) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CertificateStoresCertificateStoreResponse) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CertificateStoresCertificateStoreResponse) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CertificateStoresCertificateStoreResponse) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentAssigned

`func (o *CertificateStoresCertificateStoreResponse) GetAgentAssigned() bool`

GetAgentAssigned returns the AgentAssigned field if non-nil, zero value otherwise.

### GetAgentAssignedOk

`func (o *CertificateStoresCertificateStoreResponse) GetAgentAssignedOk() (*bool, bool)`

GetAgentAssignedOk returns a tuple with the AgentAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAssigned

`func (o *CertificateStoresCertificateStoreResponse) SetAgentAssigned(v bool)`

SetAgentAssigned sets AgentAssigned field to given value.

### HasAgentAssigned

`func (o *CertificateStoresCertificateStoreResponse) HasAgentAssigned() bool`

HasAgentAssigned returns a boolean if a field has been set.

### GetContainerName

`func (o *CertificateStoresCertificateStoreResponse) GetContainerName() string`

GetContainerName returns the ContainerName field if non-nil, zero value otherwise.

### GetContainerNameOk

`func (o *CertificateStoresCertificateStoreResponse) GetContainerNameOk() (*string, bool)`

GetContainerNameOk returns a tuple with the ContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerName

`func (o *CertificateStoresCertificateStoreResponse) SetContainerName(v string)`

SetContainerName sets ContainerName field to given value.

### HasContainerName

`func (o *CertificateStoresCertificateStoreResponse) HasContainerName() bool`

HasContainerName returns a boolean if a field has been set.

### SetContainerNameNil

`func (o *CertificateStoresCertificateStoreResponse) SetContainerNameNil(b bool)`

 SetContainerNameNil sets the value for ContainerName to be an explicit nil

### UnsetContainerName
`func (o *CertificateStoresCertificateStoreResponse) UnsetContainerName()`

UnsetContainerName ensures that no value is present for ContainerName, not even an explicit nil
### GetInventorySchedule

`func (o *CertificateStoresCertificateStoreResponse) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *CertificateStoresCertificateStoreResponse) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *CertificateStoresCertificateStoreResponse) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *CertificateStoresCertificateStoreResponse) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetReenrollmentStatus

`func (o *CertificateStoresCertificateStoreResponse) GetReenrollmentStatus() CSSCMSDataModelModelsReenrollmentStatus`

GetReenrollmentStatus returns the ReenrollmentStatus field if non-nil, zero value otherwise.

### GetReenrollmentStatusOk

`func (o *CertificateStoresCertificateStoreResponse) GetReenrollmentStatusOk() (*CSSCMSDataModelModelsReenrollmentStatus, bool)`

GetReenrollmentStatusOk returns a tuple with the ReenrollmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenrollmentStatus

`func (o *CertificateStoresCertificateStoreResponse) SetReenrollmentStatus(v CSSCMSDataModelModelsReenrollmentStatus)`

SetReenrollmentStatus sets ReenrollmentStatus field to given value.

### HasReenrollmentStatus

`func (o *CertificateStoresCertificateStoreResponse) HasReenrollmentStatus() bool`

HasReenrollmentStatus returns a boolean if a field has been set.

### GetSetNewPasswordAllowed

`func (o *CertificateStoresCertificateStoreResponse) GetSetNewPasswordAllowed() bool`

GetSetNewPasswordAllowed returns the SetNewPasswordAllowed field if non-nil, zero value otherwise.

### GetSetNewPasswordAllowedOk

`func (o *CertificateStoresCertificateStoreResponse) GetSetNewPasswordAllowedOk() (*bool, bool)`

GetSetNewPasswordAllowedOk returns a tuple with the SetNewPasswordAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNewPasswordAllowed

`func (o *CertificateStoresCertificateStoreResponse) SetSetNewPasswordAllowed(v bool)`

SetSetNewPasswordAllowed sets SetNewPasswordAllowed field to given value.

### HasSetNewPasswordAllowed

`func (o *CertificateStoresCertificateStoreResponse) HasSetNewPasswordAllowed() bool`

HasSetNewPasswordAllowed returns a boolean if a field has been set.

### GetPassword

`func (o *CertificateStoresCertificateStoreResponse) GetPassword() CSSCMSDataModelModelsKeyfactorSecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CertificateStoresCertificateStoreResponse) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorSecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CertificateStoresCertificateStoreResponse) SetPassword(v CSSCMSDataModelModelsKeyfactorSecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CertificateStoresCertificateStoreResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


