# KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**ShortName** | Pointer to **NullableString** |  | [optional] 
**Capability** | Pointer to **NullableString** |  | [optional] 
**StoreType** | Pointer to **NullableInt32** |  | [optional] 
**ImportType** | Pointer to **NullableInt32** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**ModelsCertStoreTypeSupportedOperations**](ModelsCertStoreTypeSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]ModelsCertificateStoreTypeProperty**](ModelsCertificateStoreTypeProperty.md) |  | [optional] 
**EntryParameters** | Pointer to [**[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 
**PasswordOptions** | Pointer to [**ModelsCertStoreTypePasswordOptions**](ModelsCertStoreTypePasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **NullableString** |  | [optional] 
**StorePathValue** | Pointer to **NullableString** |  | [optional] 
**PrivateKeyAllowed** | Pointer to **int32** |  | [optional] 
**JobProperties** | Pointer to **[]string** |  | [optional] [readonly] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to **int32** |  | [optional] 
**ServerRegistration** | Pointer to **NullableInt32** |  | [optional] 
**InventoryEndpoint** | Pointer to **NullableString** |  | [optional] 
**InventoryJobType** | Pointer to **string** |  | [optional] 
**ManagementJobType** | Pointer to **NullableString** |  | [optional] 
**DiscoveryJobType** | Pointer to **NullableString** |  | [optional] 
**EnrollmentJobType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse() *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStoreType() int32`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStoreTypeOk() (*int32, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStoreType(v int32)`

SetStoreType sets StoreType field to given value.

### HasStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasStoreType() bool`

HasStoreType returns a boolean if a field has been set.

### SetStoreTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStoreTypeNil(b bool)`

 SetStoreTypeNil sets the value for StoreType to be an explicit nil

### UnsetStoreType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetStoreType()`

UnsetStoreType ensures that no value is present for StoreType, not even an explicit nil
### GetImportType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetImportType() int32`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetImportTypeOk() (*int32, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetImportType(v int32)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### SetImportTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetImportTypeNil(b bool)`

 SetImportTypeNil sets the value for ImportType to be an explicit nil

### UnsetImportType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetImportType()`

UnsetImportType ensures that no value is present for ImportType, not even an explicit nil
### GetLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetSupportedOperations() ModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetSupportedOperationsOk() (*ModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetSupportedOperations(v ModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetProperties() []ModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPropertiesOk() (*[]ModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetProperties(v []ModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEntryParameters() []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEntryParametersOk() (*[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetEntryParameters(v []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil
### GetPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPasswordOptions() ModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPasswordOptionsOk() (*ModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPasswordOptions(v ModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### SetStorePathTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStorePathTypeNil(b bool)`

 SetStorePathTypeNil sets the value for StorePathType to be an explicit nil

### UnsetStorePathType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetStorePathType()`

UnsetStorePathType ensures that no value is present for StorePathType, not even an explicit nil
### GetStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### SetStorePathValueNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStorePathValueNil(b bool)`

 SetStorePathValueNil sets the value for StorePathValue to be an explicit nil

### UnsetStorePathValue
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetStorePathValue()`

UnsetStorePathValue ensures that no value is present for StorePathValue, not even an explicit nil
### GetPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPrivateKeyAllowed() int32`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPrivateKeyAllowedOk() (*int32, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPrivateKeyAllowed(v int32)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetJobProperties() []string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetJobPropertiesOk() (*[]string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetJobProperties(v []string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCustomAliasAllowed() int32`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCustomAliasAllowedOk() (*int32, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetCustomAliasAllowed(v int32)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetServerRegistration

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRegistration() int32`

GetServerRegistration returns the ServerRegistration field if non-nil, zero value otherwise.

### GetServerRegistrationOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRegistrationOk() (*int32, bool)`

GetServerRegistrationOk returns a tuple with the ServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistration

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetServerRegistration(v int32)`

SetServerRegistration sets ServerRegistration field to given value.

### HasServerRegistration

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasServerRegistration() bool`

HasServerRegistration returns a boolean if a field has been set.

### SetServerRegistrationNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetServerRegistrationNil(b bool)`

 SetServerRegistrationNil sets the value for ServerRegistration to be an explicit nil

### UnsetServerRegistration
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetServerRegistration()`

UnsetServerRegistration ensures that no value is present for ServerRegistration, not even an explicit nil
### GetInventoryEndpoint

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryEndpoint() string`

GetInventoryEndpoint returns the InventoryEndpoint field if non-nil, zero value otherwise.

### GetInventoryEndpointOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryEndpointOk() (*string, bool)`

GetInventoryEndpointOk returns a tuple with the InventoryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryEndpoint

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetInventoryEndpoint(v string)`

SetInventoryEndpoint sets InventoryEndpoint field to given value.

### HasInventoryEndpoint

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasInventoryEndpoint() bool`

HasInventoryEndpoint returns a boolean if a field has been set.

### SetInventoryEndpointNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetInventoryEndpointNil(b bool)`

 SetInventoryEndpointNil sets the value for InventoryEndpoint to be an explicit nil

### UnsetInventoryEndpoint
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetInventoryEndpoint()`

UnsetInventoryEndpoint ensures that no value is present for InventoryEndpoint, not even an explicit nil
### GetInventoryJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryJobType() string`

GetInventoryJobType returns the InventoryJobType field if non-nil, zero value otherwise.

### GetInventoryJobTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryJobTypeOk() (*string, bool)`

GetInventoryJobTypeOk returns a tuple with the InventoryJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetInventoryJobType(v string)`

SetInventoryJobType sets InventoryJobType field to given value.

### HasInventoryJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasInventoryJobType() bool`

HasInventoryJobType returns a boolean if a field has been set.

### GetManagementJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetManagementJobType() string`

GetManagementJobType returns the ManagementJobType field if non-nil, zero value otherwise.

### GetManagementJobTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetManagementJobTypeOk() (*string, bool)`

GetManagementJobTypeOk returns a tuple with the ManagementJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetManagementJobType(v string)`

SetManagementJobType sets ManagementJobType field to given value.

### HasManagementJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasManagementJobType() bool`

HasManagementJobType returns a boolean if a field has been set.

### SetManagementJobTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetManagementJobTypeNil(b bool)`

 SetManagementJobTypeNil sets the value for ManagementJobType to be an explicit nil

### UnsetManagementJobType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetManagementJobType()`

UnsetManagementJobType ensures that no value is present for ManagementJobType, not even an explicit nil
### GetDiscoveryJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetDiscoveryJobType() string`

GetDiscoveryJobType returns the DiscoveryJobType field if non-nil, zero value otherwise.

### GetDiscoveryJobTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetDiscoveryJobTypeOk() (*string, bool)`

GetDiscoveryJobTypeOk returns a tuple with the DiscoveryJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetDiscoveryJobType(v string)`

SetDiscoveryJobType sets DiscoveryJobType field to given value.

### HasDiscoveryJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasDiscoveryJobType() bool`

HasDiscoveryJobType returns a boolean if a field has been set.

### SetDiscoveryJobTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetDiscoveryJobTypeNil(b bool)`

 SetDiscoveryJobTypeNil sets the value for DiscoveryJobType to be an explicit nil

### UnsetDiscoveryJobType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetDiscoveryJobType()`

UnsetDiscoveryJobType ensures that no value is present for DiscoveryJobType, not even an explicit nil
### GetEnrollmentJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEnrollmentJobType() string`

GetEnrollmentJobType returns the EnrollmentJobType field if non-nil, zero value otherwise.

### GetEnrollmentJobTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEnrollmentJobTypeOk() (*string, bool)`

GetEnrollmentJobTypeOk returns a tuple with the EnrollmentJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetEnrollmentJobType(v string)`

SetEnrollmentJobType sets EnrollmentJobType field to given value.

### HasEnrollmentJobType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasEnrollmentJobType() bool`

HasEnrollmentJobType returns a boolean if a field has been set.

### SetEnrollmentJobTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetEnrollmentJobTypeNil(b bool)`

 SetEnrollmentJobTypeNil sets the value for EnrollmentJobType to be an explicit nil

### UnsetEnrollmentJobType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) UnsetEnrollmentJobType()`

UnsetEnrollmentJobType ensures that no value is present for EnrollmentJobType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


