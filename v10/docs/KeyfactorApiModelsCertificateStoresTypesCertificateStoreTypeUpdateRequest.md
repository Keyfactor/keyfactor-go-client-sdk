# KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreType** | **int64** |  | 
**Name** | **string** |  | 
**ShortName** | **string** |  | 
**Capability** | Pointer to **string** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**ModelsCertStoreTypeSupportedOperations**](ModelsCertStoreTypeSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]ModelsCertificateStoreTypeProperty**](ModelsCertificateStoreTypeProperty.md) |  | [optional] 
**PasswordOptions** | Pointer to [**ModelsCertStoreTypePasswordOptions**](ModelsCertStoreTypePasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **string** |  | [optional] 
**StorePathValue** | Pointer to **string** |  | [optional] 
**PrivateKeyAllowed** | Pointer to **string** |  | [optional] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to **string** |  | [optional] 
**EntryParameters** | Pointer to [**[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest

`func NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest(storeType int64, name string, shortName string, ) *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest`

NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest instantiates a new KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults

`func NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults() *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest`

NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStoreType() int64`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStoreTypeOk() (*int64, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStoreType(v int64)`

SetStoreType sets StoreType field to given value.


### GetName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetSupportedOperations() ModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetSupportedOperationsOk() (*ModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetSupportedOperations(v ModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetProperties() []ModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPropertiesOk() (*[]ModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetProperties(v []ModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPasswordOptions() ModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPasswordOptionsOk() (*ModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPasswordOptions(v ModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### GetStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### GetPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPrivateKeyAllowed() string`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPrivateKeyAllowedOk() (*string, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPrivateKeyAllowed(v string)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCustomAliasAllowed() string`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCustomAliasAllowedOk() (*string, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCustomAliasAllowed(v string)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetEntryParameters() []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) GetEntryParametersOk() (*[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) SetEntryParameters(v []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeUpdateRequest) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


