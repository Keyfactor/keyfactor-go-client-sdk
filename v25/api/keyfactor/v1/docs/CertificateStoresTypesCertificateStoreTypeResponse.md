# CertificateStoresTypesCertificateStoreTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ShortName** | Pointer to **NullableString** |  | [optional] 
**Capability** | Pointer to **NullableString** |  | [optional] 
**StoreType** | Pointer to **NullableInt32** |  | [optional] 
**ImportType** | Pointer to **NullableInt32** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**CertificateStoresTypesSupportedOperations**](CertificateStoresTypesSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]CertificateStoresTypesStoreTypeProperty**](CertificateStoresTypesStoreTypeProperty.md) |  | [optional] 
**EntryParameters** | Pointer to [**[]CertificateStoresTypesEntryParameters**](CertificateStoresTypesEntryParameters.md) |  | [optional] 
**PasswordOptions** | Pointer to [**CertificateStoresTypesPasswordOptions**](CertificateStoresTypesPasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **NullableString** |  | [optional] 
**StorePathValue** | Pointer to **NullableString** |  | [optional] 
**PrivateKeyAllowed** | Pointer to [**CSSCMSCoreEnumsCertStorePrivateKey**](CSSCMSCoreEnumsCertStorePrivateKey.md) |  | [optional] 
**CertificateFormat** | Pointer to [**CSSCMSCoreEnumsCertificateFormat**](CSSCMSCoreEnumsCertificateFormat.md) |  | [optional] 
**JobProperties** | Pointer to **[]string** |  | [optional] [readonly] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to [**KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias**](KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias.md) |  | [optional] 
**ServerRegistration** | Pointer to **NullableInt32** |  | [optional] 
**InventoryEndpoint** | Pointer to **NullableString** |  | [optional] 
**InventoryJobType** | Pointer to **string** |  | [optional] 
**ManagementJobType** | Pointer to **NullableString** |  | [optional] 
**DiscoveryJobType** | Pointer to **NullableString** |  | [optional] 
**EnrollmentJobType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateStoresTypesCertificateStoreTypeResponse

`func NewCertificateStoresTypesCertificateStoreTypeResponse() *CertificateStoresTypesCertificateStoreTypeResponse`

NewCertificateStoresTypesCertificateStoreTypeResponse instantiates a new CertificateStoresTypesCertificateStoreTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresTypesCertificateStoreTypeResponseWithDefaults

`func NewCertificateStoresTypesCertificateStoreTypeResponseWithDefaults() *CertificateStoresTypesCertificateStoreTypeResponse`

NewCertificateStoresTypesCertificateStoreTypeResponseWithDefaults instantiates a new CertificateStoresTypesCertificateStoreTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortName

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetCapability

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetStoreType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetStoreType() int32`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetStoreTypeOk() (*int32, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetStoreType(v int32)`

SetStoreType sets StoreType field to given value.

### HasStoreType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasStoreType() bool`

HasStoreType returns a boolean if a field has been set.

### SetStoreTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetStoreTypeNil(b bool)`

 SetStoreTypeNil sets the value for StoreType to be an explicit nil

### UnsetStoreType
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetStoreType()`

UnsetStoreType ensures that no value is present for StoreType, not even an explicit nil
### GetImportType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetImportType() int32`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetImportTypeOk() (*int32, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetImportType(v int32)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### SetImportTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetImportTypeNil(b bool)`

 SetImportTypeNil sets the value for ImportType to be an explicit nil

### UnsetImportType
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetImportType()`

UnsetImportType ensures that no value is present for ImportType, not even an explicit nil
### GetLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetSupportedOperations() CertificateStoresTypesSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetSupportedOperationsOk() (*CertificateStoresTypesSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetSupportedOperations(v CertificateStoresTypesSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetProperties() []CertificateStoresTypesStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPropertiesOk() (*[]CertificateStoresTypesStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetProperties(v []CertificateStoresTypesStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetEntryParameters() []CertificateStoresTypesEntryParameters`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetEntryParametersOk() (*[]CertificateStoresTypesEntryParameters, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetEntryParameters(v []CertificateStoresTypesEntryParameters)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil
### GetPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPasswordOptions() CertificateStoresTypesPasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPasswordOptionsOk() (*CertificateStoresTypesPasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetPasswordOptions(v CertificateStoresTypesPasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### SetStorePathTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetStorePathTypeNil(b bool)`

 SetStorePathTypeNil sets the value for StorePathType to be an explicit nil

### UnsetStorePathType
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetStorePathType()`

UnsetStorePathType ensures that no value is present for StorePathType, not even an explicit nil
### GetStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### SetStorePathValueNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetStorePathValueNil(b bool)`

 SetStorePathValueNil sets the value for StorePathValue to be an explicit nil

### UnsetStorePathValue
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetStorePathValue()`

UnsetStorePathValue ensures that no value is present for StorePathValue, not even an explicit nil
### GetPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPrivateKeyAllowed() CSSCMSCoreEnumsCertStorePrivateKey`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPrivateKeyAllowedOk() (*CSSCMSCoreEnumsCertStorePrivateKey, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetPrivateKeyAllowed(v CSSCMSCoreEnumsCertStorePrivateKey)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetCertificateFormat() CSSCMSCoreEnumsCertificateFormat`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetCertificateFormatOk() (*CSSCMSCoreEnumsCertificateFormat, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetCertificateFormat(v CSSCMSCoreEnumsCertificateFormat)`

SetCertificateFormat sets CertificateFormat field to given value.

### HasCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasCertificateFormat() bool`

HasCertificateFormat returns a boolean if a field has been set.

### GetJobProperties

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetJobProperties() []string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetJobPropertiesOk() (*[]string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetJobProperties(v []string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetCustomAliasAllowed() KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetCustomAliasAllowedOk() (*KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetCustomAliasAllowed(v KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetServerRegistration

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetServerRegistration() int32`

GetServerRegistration returns the ServerRegistration field if non-nil, zero value otherwise.

### GetServerRegistrationOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetServerRegistrationOk() (*int32, bool)`

GetServerRegistrationOk returns a tuple with the ServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistration

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetServerRegistration(v int32)`

SetServerRegistration sets ServerRegistration field to given value.

### HasServerRegistration

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasServerRegistration() bool`

HasServerRegistration returns a boolean if a field has been set.

### SetServerRegistrationNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetServerRegistrationNil(b bool)`

 SetServerRegistrationNil sets the value for ServerRegistration to be an explicit nil

### UnsetServerRegistration
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetServerRegistration()`

UnsetServerRegistration ensures that no value is present for ServerRegistration, not even an explicit nil
### GetInventoryEndpoint

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetInventoryEndpoint() string`

GetInventoryEndpoint returns the InventoryEndpoint field if non-nil, zero value otherwise.

### GetInventoryEndpointOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetInventoryEndpointOk() (*string, bool)`

GetInventoryEndpointOk returns a tuple with the InventoryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryEndpoint

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetInventoryEndpoint(v string)`

SetInventoryEndpoint sets InventoryEndpoint field to given value.

### HasInventoryEndpoint

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasInventoryEndpoint() bool`

HasInventoryEndpoint returns a boolean if a field has been set.

### SetInventoryEndpointNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetInventoryEndpointNil(b bool)`

 SetInventoryEndpointNil sets the value for InventoryEndpoint to be an explicit nil

### UnsetInventoryEndpoint
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetInventoryEndpoint()`

UnsetInventoryEndpoint ensures that no value is present for InventoryEndpoint, not even an explicit nil
### GetInventoryJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetInventoryJobType() string`

GetInventoryJobType returns the InventoryJobType field if non-nil, zero value otherwise.

### GetInventoryJobTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetInventoryJobTypeOk() (*string, bool)`

GetInventoryJobTypeOk returns a tuple with the InventoryJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetInventoryJobType(v string)`

SetInventoryJobType sets InventoryJobType field to given value.

### HasInventoryJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasInventoryJobType() bool`

HasInventoryJobType returns a boolean if a field has been set.

### GetManagementJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetManagementJobType() string`

GetManagementJobType returns the ManagementJobType field if non-nil, zero value otherwise.

### GetManagementJobTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetManagementJobTypeOk() (*string, bool)`

GetManagementJobTypeOk returns a tuple with the ManagementJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetManagementJobType(v string)`

SetManagementJobType sets ManagementJobType field to given value.

### HasManagementJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasManagementJobType() bool`

HasManagementJobType returns a boolean if a field has been set.

### SetManagementJobTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetManagementJobTypeNil(b bool)`

 SetManagementJobTypeNil sets the value for ManagementJobType to be an explicit nil

### UnsetManagementJobType
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetManagementJobType()`

UnsetManagementJobType ensures that no value is present for ManagementJobType, not even an explicit nil
### GetDiscoveryJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetDiscoveryJobType() string`

GetDiscoveryJobType returns the DiscoveryJobType field if non-nil, zero value otherwise.

### GetDiscoveryJobTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetDiscoveryJobTypeOk() (*string, bool)`

GetDiscoveryJobTypeOk returns a tuple with the DiscoveryJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetDiscoveryJobType(v string)`

SetDiscoveryJobType sets DiscoveryJobType field to given value.

### HasDiscoveryJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasDiscoveryJobType() bool`

HasDiscoveryJobType returns a boolean if a field has been set.

### SetDiscoveryJobTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetDiscoveryJobTypeNil(b bool)`

 SetDiscoveryJobTypeNil sets the value for DiscoveryJobType to be an explicit nil

### UnsetDiscoveryJobType
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetDiscoveryJobType()`

UnsetDiscoveryJobType ensures that no value is present for DiscoveryJobType, not even an explicit nil
### GetEnrollmentJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetEnrollmentJobType() string`

GetEnrollmentJobType returns the EnrollmentJobType field if non-nil, zero value otherwise.

### GetEnrollmentJobTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) GetEnrollmentJobTypeOk() (*string, bool)`

GetEnrollmentJobTypeOk returns a tuple with the EnrollmentJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetEnrollmentJobType(v string)`

SetEnrollmentJobType sets EnrollmentJobType field to given value.

### HasEnrollmentJobType

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) HasEnrollmentJobType() bool`

HasEnrollmentJobType returns a boolean if a field has been set.

### SetEnrollmentJobTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeResponse) SetEnrollmentJobTypeNil(b bool)`

 SetEnrollmentJobTypeNil sets the value for EnrollmentJobType to be an explicit nil

### UnsetEnrollmentJobType
`func (o *CertificateStoresTypesCertificateStoreTypeResponse) UnsetEnrollmentJobType()`

UnsetEnrollmentJobType ensures that no value is present for EnrollmentJobType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


