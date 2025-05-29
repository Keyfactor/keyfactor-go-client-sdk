# CertificateStoresTypesCertificateStoreTypeCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**CertificateFormat** | Pointer to [**CSSCMSCoreEnumsCertificateFormat**](CSSCMSCoreEnumsCertificateFormat.md) |  | [optional] 
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
**JobProperties** | Pointer to **[]string** |  | [optional] 
**EntryParameters** | Pointer to [**[]CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 

## Methods

### NewCertificateStoresTypesCertificateStoreTypeCreationRequest

`func NewCertificateStoresTypesCertificateStoreTypeCreationRequest(name string, shortName string, ) *CertificateStoresTypesCertificateStoreTypeCreationRequest`

NewCertificateStoresTypesCertificateStoreTypeCreationRequest instantiates a new CertificateStoresTypesCertificateStoreTypeCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults

`func NewCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults() *CertificateStoresTypesCertificateStoreTypeCreationRequest`

NewCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults instantiates a new CertificateStoresTypesCertificateStoreTypeCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetCapability

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetSupportedOperations() CSSCMSDataModelModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetSupportedOperationsOk() (*CSSCMSDataModelModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetSupportedOperations(v CSSCMSDataModelModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetProperties() []CSSCMSDataModelModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPropertiesOk() (*[]CSSCMSDataModelModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetProperties(v []CSSCMSDataModelModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPasswordOptions() CSSCMSDataModelModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPasswordOptionsOk() (*CSSCMSDataModelModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetPasswordOptions(v CSSCMSDataModelModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### SetStorePathTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathTypeNil(b bool)`

 SetStorePathTypeNil sets the value for StorePathType to be an explicit nil

### UnsetStorePathType
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetStorePathType()`

UnsetStorePathType ensures that no value is present for StorePathType, not even an explicit nil
### GetStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### SetStorePathValueNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathValueNil(b bool)`

 SetStorePathValueNil sets the value for StorePathValue to be an explicit nil

### UnsetStorePathValue
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetStorePathValue()`

UnsetStorePathValue ensures that no value is present for StorePathValue, not even an explicit nil
### GetPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPrivateKeyAllowed() CSSCMSCoreEnumsCertStorePrivateKey`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPrivateKeyAllowedOk() (*CSSCMSCoreEnumsCertStorePrivateKey, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetPrivateKeyAllowed(v CSSCMSCoreEnumsCertStorePrivateKey)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetCertificateFormat() CSSCMSCoreEnumsCertificateFormat`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetCertificateFormatOk() (*CSSCMSCoreEnumsCertificateFormat, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetCertificateFormat(v CSSCMSCoreEnumsCertificateFormat)`

SetCertificateFormat sets CertificateFormat field to given value.

### HasCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasCertificateFormat() bool`

HasCertificateFormat returns a boolean if a field has been set.

### GetServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetCustomAliasAllowed() KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetCustomAliasAllowedOk() (*KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetCustomAliasAllowed(v KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetServerRegistration

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRegistration() int32`

GetServerRegistration returns the ServerRegistration field if non-nil, zero value otherwise.

### GetServerRegistrationOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRegistrationOk() (*int32, bool)`

GetServerRegistrationOk returns a tuple with the ServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistration

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRegistration(v int32)`

SetServerRegistration sets ServerRegistration field to given value.

### HasServerRegistration

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasServerRegistration() bool`

HasServerRegistration returns a boolean if a field has been set.

### SetServerRegistrationNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRegistrationNil(b bool)`

 SetServerRegistrationNil sets the value for ServerRegistration to be an explicit nil

### UnsetServerRegistration
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetServerRegistration()`

UnsetServerRegistration ensures that no value is present for ServerRegistration, not even an explicit nil
### GetInventoryEndpoint

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryEndpoint() string`

GetInventoryEndpoint returns the InventoryEndpoint field if non-nil, zero value otherwise.

### GetInventoryEndpointOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryEndpointOk() (*string, bool)`

GetInventoryEndpointOk returns a tuple with the InventoryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryEndpoint

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryEndpoint(v string)`

SetInventoryEndpoint sets InventoryEndpoint field to given value.

### HasInventoryEndpoint

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasInventoryEndpoint() bool`

HasInventoryEndpoint returns a boolean if a field has been set.

### SetInventoryEndpointNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryEndpointNil(b bool)`

 SetInventoryEndpointNil sets the value for InventoryEndpoint to be an explicit nil

### UnsetInventoryEndpoint
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetInventoryEndpoint()`

UnsetInventoryEndpoint ensures that no value is present for InventoryEndpoint, not even an explicit nil
### GetInventoryJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryJobTypeId() string`

GetInventoryJobTypeId returns the InventoryJobTypeId field if non-nil, zero value otherwise.

### GetInventoryJobTypeIdOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryJobTypeIdOk() (*string, bool)`

GetInventoryJobTypeIdOk returns a tuple with the InventoryJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryJobTypeId(v string)`

SetInventoryJobTypeId sets InventoryJobTypeId field to given value.

### HasInventoryJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasInventoryJobTypeId() bool`

HasInventoryJobTypeId returns a boolean if a field has been set.

### GetManagementJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetManagementJobTypeId() string`

GetManagementJobTypeId returns the ManagementJobTypeId field if non-nil, zero value otherwise.

### GetManagementJobTypeIdOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetManagementJobTypeIdOk() (*string, bool)`

GetManagementJobTypeIdOk returns a tuple with the ManagementJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetManagementJobTypeId(v string)`

SetManagementJobTypeId sets ManagementJobTypeId field to given value.

### HasManagementJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasManagementJobTypeId() bool`

HasManagementJobTypeId returns a boolean if a field has been set.

### SetManagementJobTypeIdNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetManagementJobTypeIdNil(b bool)`

 SetManagementJobTypeIdNil sets the value for ManagementJobTypeId to be an explicit nil

### UnsetManagementJobTypeId
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetManagementJobTypeId()`

UnsetManagementJobTypeId ensures that no value is present for ManagementJobTypeId, not even an explicit nil
### GetDiscoveryJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetDiscoveryJobTypeId() string`

GetDiscoveryJobTypeId returns the DiscoveryJobTypeId field if non-nil, zero value otherwise.

### GetDiscoveryJobTypeIdOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetDiscoveryJobTypeIdOk() (*string, bool)`

GetDiscoveryJobTypeIdOk returns a tuple with the DiscoveryJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetDiscoveryJobTypeId(v string)`

SetDiscoveryJobTypeId sets DiscoveryJobTypeId field to given value.

### HasDiscoveryJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasDiscoveryJobTypeId() bool`

HasDiscoveryJobTypeId returns a boolean if a field has been set.

### SetDiscoveryJobTypeIdNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetDiscoveryJobTypeIdNil(b bool)`

 SetDiscoveryJobTypeIdNil sets the value for DiscoveryJobTypeId to be an explicit nil

### UnsetDiscoveryJobTypeId
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetDiscoveryJobTypeId()`

UnsetDiscoveryJobTypeId ensures that no value is present for DiscoveryJobTypeId, not even an explicit nil
### GetEnrollmentJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetEnrollmentJobTypeId() string`

GetEnrollmentJobTypeId returns the EnrollmentJobTypeId field if non-nil, zero value otherwise.

### GetEnrollmentJobTypeIdOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetEnrollmentJobTypeIdOk() (*string, bool)`

GetEnrollmentJobTypeIdOk returns a tuple with the EnrollmentJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetEnrollmentJobTypeId(v string)`

SetEnrollmentJobTypeId sets EnrollmentJobTypeId field to given value.

### HasEnrollmentJobTypeId

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasEnrollmentJobTypeId() bool`

HasEnrollmentJobTypeId returns a boolean if a field has been set.

### SetEnrollmentJobTypeIdNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetEnrollmentJobTypeIdNil(b bool)`

 SetEnrollmentJobTypeIdNil sets the value for EnrollmentJobTypeId to be an explicit nil

### UnsetEnrollmentJobTypeId
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetEnrollmentJobTypeId()`

UnsetEnrollmentJobTypeId ensures that no value is present for EnrollmentJobTypeId, not even an explicit nil
### GetJobProperties

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetJobProperties() []string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetJobPropertiesOk() (*[]string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetJobProperties(v []string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetEntryParameters() []CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) GetEntryParametersOk() (*[]CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetEntryParameters(v []CSSCMSDataModelModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *CertificateStoresTypesCertificateStoreTypeCreationRequest) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


