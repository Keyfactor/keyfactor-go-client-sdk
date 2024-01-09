# KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobProperties** | Pointer to **[]string** |  | [optional] 
**EntryParameters** | Pointer to [**[]CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 
**Name** | **string** |  | 
**ShortName** | **string** |  | 
**Capability** | Pointer to **NullableString** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**CSSCMSDataModelModelsCertStoreTypeSupportedOperations**](CSSCMSDataModelModelsCertStoreTypeSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]CSSCMSDataModelModelsCertificateStoreTypeProperty**](CSSCMSDataModelModelsCertificateStoreTypeProperty.md) |  | [optional] 
**PasswordOptions** | Pointer to [**CSSCMSDataModelModelsCertStoreTypePasswordOptions**](CSSCMSDataModelModelsCertStoreTypePasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **NullableString** |  | [optional] 
**StorePathValue** | Pointer to **NullableString** |  | [optional] 
**PrivateKeyAllowed** | Pointer to [**CSSCMSCoreEnumsCertStorePrivateKey**](CSSCMSCoreEnumsCertStorePrivateKey.md) |  | [optional] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to [**KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias**](KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias.md) |  | [optional] 
**ServerRegistration** | Pointer to **NullableInt32** |  | [optional] 
**InventoryEndpoint** | Pointer to **NullableString** |  | [optional] 
**InventoryJobTypeId** | Pointer to **string** |  | [optional] 
**ManagementJobTypeId** | Pointer to **NullableString** |  | [optional] 
**DiscoveryJobTypeId** | Pointer to **NullableString** |  | [optional] 
**EnrollmentJobTypeId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest(name string, shortName string, ) *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetJobProperties() []string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetJobPropertiesOk() (*[]string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetJobProperties(v []string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEntryParameters() []CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEntryParametersOk() (*[]CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetEntryParameters(v []CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil
### GetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetSupportedOperations() CSSCMSDataModelModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetSupportedOperationsOk() (*CSSCMSDataModelModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetSupportedOperations(v CSSCMSDataModelModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetProperties() []CSSCMSDataModelModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPropertiesOk() (*[]CSSCMSDataModelModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetProperties(v []CSSCMSDataModelModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPasswordOptions() CSSCMSDataModelModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPasswordOptionsOk() (*CSSCMSDataModelModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPasswordOptions(v CSSCMSDataModelModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### SetStorePathTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathTypeNil(b bool)`

 SetStorePathTypeNil sets the value for StorePathType to be an explicit nil

### UnsetStorePathType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetStorePathType()`

UnsetStorePathType ensures that no value is present for StorePathType, not even an explicit nil
### GetStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### SetStorePathValueNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathValueNil(b bool)`

 SetStorePathValueNil sets the value for StorePathValue to be an explicit nil

### UnsetStorePathValue
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetStorePathValue()`

UnsetStorePathValue ensures that no value is present for StorePathValue, not even an explicit nil
### GetPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPrivateKeyAllowed() CSSCMSCoreEnumsCertStorePrivateKey`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPrivateKeyAllowedOk() (*CSSCMSCoreEnumsCertStorePrivateKey, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPrivateKeyAllowed(v CSSCMSCoreEnumsCertStorePrivateKey)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCustomAliasAllowed() KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCustomAliasAllowedOk() (*KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetCustomAliasAllowed(v KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetServerRegistration

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRegistration() int32`

GetServerRegistration returns the ServerRegistration field if non-nil, zero value otherwise.

### GetServerRegistrationOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRegistrationOk() (*int32, bool)`

GetServerRegistrationOk returns a tuple with the ServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistration

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRegistration(v int32)`

SetServerRegistration sets ServerRegistration field to given value.

### HasServerRegistration

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasServerRegistration() bool`

HasServerRegistration returns a boolean if a field has been set.

### SetServerRegistrationNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRegistrationNil(b bool)`

 SetServerRegistrationNil sets the value for ServerRegistration to be an explicit nil

### UnsetServerRegistration
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetServerRegistration()`

UnsetServerRegistration ensures that no value is present for ServerRegistration, not even an explicit nil
### GetInventoryEndpoint

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryEndpoint() string`

GetInventoryEndpoint returns the InventoryEndpoint field if non-nil, zero value otherwise.

### GetInventoryEndpointOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryEndpointOk() (*string, bool)`

GetInventoryEndpointOk returns a tuple with the InventoryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryEndpoint

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryEndpoint(v string)`

SetInventoryEndpoint sets InventoryEndpoint field to given value.

### HasInventoryEndpoint

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasInventoryEndpoint() bool`

HasInventoryEndpoint returns a boolean if a field has been set.

### SetInventoryEndpointNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryEndpointNil(b bool)`

 SetInventoryEndpointNil sets the value for InventoryEndpoint to be an explicit nil

### UnsetInventoryEndpoint
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetInventoryEndpoint()`

UnsetInventoryEndpoint ensures that no value is present for InventoryEndpoint, not even an explicit nil
### GetInventoryJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryJobTypeId() string`

GetInventoryJobTypeId returns the InventoryJobTypeId field if non-nil, zero value otherwise.

### GetInventoryJobTypeIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryJobTypeIdOk() (*string, bool)`

GetInventoryJobTypeIdOk returns a tuple with the InventoryJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryJobTypeId(v string)`

SetInventoryJobTypeId sets InventoryJobTypeId field to given value.

### HasInventoryJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasInventoryJobTypeId() bool`

HasInventoryJobTypeId returns a boolean if a field has been set.

### GetManagementJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetManagementJobTypeId() string`

GetManagementJobTypeId returns the ManagementJobTypeId field if non-nil, zero value otherwise.

### GetManagementJobTypeIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetManagementJobTypeIdOk() (*string, bool)`

GetManagementJobTypeIdOk returns a tuple with the ManagementJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetManagementJobTypeId(v string)`

SetManagementJobTypeId sets ManagementJobTypeId field to given value.

### HasManagementJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasManagementJobTypeId() bool`

HasManagementJobTypeId returns a boolean if a field has been set.

### SetManagementJobTypeIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetManagementJobTypeIdNil(b bool)`

 SetManagementJobTypeIdNil sets the value for ManagementJobTypeId to be an explicit nil

### UnsetManagementJobTypeId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetManagementJobTypeId()`

UnsetManagementJobTypeId ensures that no value is present for ManagementJobTypeId, not even an explicit nil
### GetDiscoveryJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetDiscoveryJobTypeId() string`

GetDiscoveryJobTypeId returns the DiscoveryJobTypeId field if non-nil, zero value otherwise.

### GetDiscoveryJobTypeIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetDiscoveryJobTypeIdOk() (*string, bool)`

GetDiscoveryJobTypeIdOk returns a tuple with the DiscoveryJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetDiscoveryJobTypeId(v string)`

SetDiscoveryJobTypeId sets DiscoveryJobTypeId field to given value.

### HasDiscoveryJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasDiscoveryJobTypeId() bool`

HasDiscoveryJobTypeId returns a boolean if a field has been set.

### SetDiscoveryJobTypeIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetDiscoveryJobTypeIdNil(b bool)`

 SetDiscoveryJobTypeIdNil sets the value for DiscoveryJobTypeId to be an explicit nil

### UnsetDiscoveryJobTypeId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetDiscoveryJobTypeId()`

UnsetDiscoveryJobTypeId ensures that no value is present for DiscoveryJobTypeId, not even an explicit nil
### GetEnrollmentJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEnrollmentJobTypeId() string`

GetEnrollmentJobTypeId returns the EnrollmentJobTypeId field if non-nil, zero value otherwise.

### GetEnrollmentJobTypeIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEnrollmentJobTypeIdOk() (*string, bool)`

GetEnrollmentJobTypeIdOk returns a tuple with the EnrollmentJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetEnrollmentJobTypeId(v string)`

SetEnrollmentJobTypeId sets EnrollmentJobTypeId field to given value.

### HasEnrollmentJobTypeId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasEnrollmentJobTypeId() bool`

HasEnrollmentJobTypeId returns a boolean if a field has been set.

### SetEnrollmentJobTypeIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetEnrollmentJobTypeIdNil(b bool)`

 SetEnrollmentJobTypeIdNil sets the value for EnrollmentJobTypeId to be an explicit nil

### UnsetEnrollmentJobTypeId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetEnrollmentJobTypeId()`

UnsetEnrollmentJobTypeId ensures that no value is present for EnrollmentJobTypeId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


