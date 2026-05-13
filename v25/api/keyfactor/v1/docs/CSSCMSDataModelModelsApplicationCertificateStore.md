# CSSCMSDataModelModelsApplicationCertificateStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] [readonly] 
**ApplicationId** | Pointer to **NullableInt32** |  | [optional] 
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**Storepath** | Pointer to **NullableString** |  | [optional] 
**CertStoreInventoryJobId** | Pointer to **NullableString** |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**Approved** | Pointer to **bool** |  | [optional] 
**CreateIfMissing** | Pointer to **bool** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**AgentAssigned** | Pointer to **bool** |  | [optional] 
**ApplicationName** | Pointer to **NullableString** |  | [optional] 
**InventorySchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**ReenrollmentStatus** | Pointer to [**CSSCMSDataModelModelsReenrollmentStatus**](CSSCMSDataModelModelsReenrollmentStatus.md) |  | [optional] 
**SetNewPasswordAllowed** | Pointer to **bool** |  | [optional] 
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorSecret**](CSSCMSDataModelModelsKeyfactorSecret.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsApplicationCertificateStore

`func NewCSSCMSDataModelModelsApplicationCertificateStore() *CSSCMSDataModelModelsApplicationCertificateStore`

NewCSSCMSDataModelModelsApplicationCertificateStore instantiates a new CSSCMSDataModelModelsApplicationCertificateStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsApplicationCertificateStoreWithDefaults

`func NewCSSCMSDataModelModelsApplicationCertificateStoreWithDefaults() *CSSCMSDataModelModelsApplicationCertificateStore`

NewCSSCMSDataModelModelsApplicationCertificateStoreWithDefaults instantiates a new CSSCMSDataModelModelsApplicationCertificateStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetApplicationId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### SetApplicationIdNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetApplicationIdNil(b bool)`

 SetApplicationIdNil sets the value for ApplicationId to be an explicit nil

### UnsetApplicationId
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetApplicationId()`

UnsetApplicationId ensures that no value is present for ApplicationId, not even an explicit nil
### GetClientMachine

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetStorepath

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetStorepath() string`

GetStorepath returns the Storepath field if non-nil, zero value otherwise.

### GetStorepathOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetStorepathOk() (*string, bool)`

GetStorepathOk returns a tuple with the Storepath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorepath

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetStorepath(v string)`

SetStorepath sets Storepath field to given value.

### HasStorepath

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasStorepath() bool`

HasStorepath returns a boolean if a field has been set.

### SetStorepathNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetStorepathNil(b bool)`

 SetStorepathNil sets the value for Storepath to be an explicit nil

### UnsetStorepath
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetStorepath()`

UnsetStorepath ensures that no value is present for Storepath, not even an explicit nil
### GetCertStoreInventoryJobId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetCertStoreInventoryJobId() string`

GetCertStoreInventoryJobId returns the CertStoreInventoryJobId field if non-nil, zero value otherwise.

### GetCertStoreInventoryJobIdOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetCertStoreInventoryJobIdOk() (*string, bool)`

GetCertStoreInventoryJobIdOk returns a tuple with the CertStoreInventoryJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryJobId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetCertStoreInventoryJobId(v string)`

SetCertStoreInventoryJobId sets CertStoreInventoryJobId field to given value.

### HasCertStoreInventoryJobId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasCertStoreInventoryJobId() bool`

HasCertStoreInventoryJobId returns a boolean if a field has been set.

### SetCertStoreInventoryJobIdNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetCertStoreInventoryJobIdNil(b bool)`

 SetCertStoreInventoryJobIdNil sets the value for CertStoreInventoryJobId to be an explicit nil

### UnsetCertStoreInventoryJobId
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetCertStoreInventoryJobId()`

UnsetCertStoreInventoryJobId ensures that no value is present for CertStoreInventoryJobId, not even an explicit nil
### GetCertStoreType

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetApproved

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetApproved() bool`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetApprovedOk() (*bool, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetApproved(v bool)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetCreateIfMissing

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetCreateIfMissing() bool`

GetCreateIfMissing returns the CreateIfMissing field if non-nil, zero value otherwise.

### GetCreateIfMissingOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetCreateIfMissingOk() (*bool, bool)`

GetCreateIfMissingOk returns a tuple with the CreateIfMissing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateIfMissing

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetCreateIfMissing(v bool)`

SetCreateIfMissing sets CreateIfMissing field to given value.

### HasCreateIfMissing

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasCreateIfMissing() bool`

HasCreateIfMissing returns a boolean if a field has been set.

### GetProperties

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetAgentId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetAgentAssigned

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetAgentAssigned() bool`

GetAgentAssigned returns the AgentAssigned field if non-nil, zero value otherwise.

### GetAgentAssignedOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetAgentAssignedOk() (*bool, bool)`

GetAgentAssignedOk returns a tuple with the AgentAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentAssigned

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetAgentAssigned(v bool)`

SetAgentAssigned sets AgentAssigned field to given value.

### HasAgentAssigned

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasAgentAssigned() bool`

HasAgentAssigned returns a boolean if a field has been set.

### GetApplicationName

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### SetApplicationNameNil

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetApplicationNameNil(b bool)`

 SetApplicationNameNil sets the value for ApplicationName to be an explicit nil

### UnsetApplicationName
`func (o *CSSCMSDataModelModelsApplicationCertificateStore) UnsetApplicationName()`

UnsetApplicationName ensures that no value is present for ApplicationName, not even an explicit nil
### GetInventorySchedule

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetReenrollmentStatus

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetReenrollmentStatus() CSSCMSDataModelModelsReenrollmentStatus`

GetReenrollmentStatus returns the ReenrollmentStatus field if non-nil, zero value otherwise.

### GetReenrollmentStatusOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetReenrollmentStatusOk() (*CSSCMSDataModelModelsReenrollmentStatus, bool)`

GetReenrollmentStatusOk returns a tuple with the ReenrollmentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenrollmentStatus

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetReenrollmentStatus(v CSSCMSDataModelModelsReenrollmentStatus)`

SetReenrollmentStatus sets ReenrollmentStatus field to given value.

### HasReenrollmentStatus

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasReenrollmentStatus() bool`

HasReenrollmentStatus returns a boolean if a field has been set.

### GetSetNewPasswordAllowed

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetSetNewPasswordAllowed() bool`

GetSetNewPasswordAllowed returns the SetNewPasswordAllowed field if non-nil, zero value otherwise.

### GetSetNewPasswordAllowedOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetSetNewPasswordAllowedOk() (*bool, bool)`

GetSetNewPasswordAllowedOk returns a tuple with the SetNewPasswordAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetNewPasswordAllowed

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetSetNewPasswordAllowed(v bool)`

SetSetNewPasswordAllowed sets SetNewPasswordAllowed field to given value.

### HasSetNewPasswordAllowed

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasSetNewPasswordAllowed() bool`

HasSetNewPasswordAllowed returns a boolean if a field has been set.

### GetPassword

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetPassword() CSSCMSDataModelModelsKeyfactorSecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorSecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) SetPassword(v CSSCMSDataModelModelsKeyfactorSecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CSSCMSDataModelModelsApplicationCertificateStore) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


