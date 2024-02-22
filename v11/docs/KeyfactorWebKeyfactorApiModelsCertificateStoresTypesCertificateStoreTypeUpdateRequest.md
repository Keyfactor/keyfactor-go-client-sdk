# KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobProperties** | Pointer to **[]string** |  | [optional] 
**EntryParameters** | Pointer to [**[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 
**StoreType** | **int32** |  | 
**Name** | **string** |  | 
**ShortName** | **string** |  | 
**Capability** | Pointer to **NullableString** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**ModelsCertStoreTypeSupportedOperations**](ModelsCertStoreTypeSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]ModelsCertificateStoreTypeProperty**](ModelsCertificateStoreTypeProperty.md) |  | [optional] 
**PasswordOptions** | Pointer to [**ModelsCertStoreTypePasswordOptions**](ModelsCertStoreTypePasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **NullableString** |  | [optional] 
**StorePathValue** | Pointer to **NullableString** |  | [optional] 
**PrivateKeyAllowed** | Pointer to **int32** |  | [optional] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest(storeType int32, name string, shortName string, ) *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetJobProperties() []string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetJobPropertiesOk() (*[]string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetJobProperties(v []string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetEntryParameters() []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetEntryParametersOk() (*[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetEntryParameters(v []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil
### GetStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStoreType() int32`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStoreTypeOk() (*int32, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStoreType(v int32)`

SetStoreType sets StoreType field to given value.


### GetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetSupportedOperations() ModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetSupportedOperationsOk() (*ModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetSupportedOperations(v ModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetProperties() []ModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPropertiesOk() (*[]ModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetProperties(v []ModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPasswordOptions() ModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPasswordOptionsOk() (*ModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPasswordOptions(v ModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### SetStorePathTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathTypeNil(b bool)`

 SetStorePathTypeNil sets the value for StorePathType to be an explicit nil

### UnsetStorePathType
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetStorePathType()`

UnsetStorePathType ensures that no value is present for StorePathType, not even an explicit nil
### GetStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### SetStorePathValueNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathValueNil(b bool)`

 SetStorePathValueNil sets the value for StorePathValue to be an explicit nil

### UnsetStorePathValue
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetStorePathValue()`

UnsetStorePathValue ensures that no value is present for StorePathValue, not even an explicit nil
### GetPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPrivateKeyAllowed() int32`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPrivateKeyAllowedOk() (*int32, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPrivateKeyAllowed(v int32)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCustomAliasAllowed() int32`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCustomAliasAllowedOk() (*int32, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCustomAliasAllowed(v int32)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


