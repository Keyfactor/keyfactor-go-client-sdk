# CertificateStoresTypesCertificateStoreTypeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreType** | **int32** |  | 
**Name** | **string** |  | 
**ShortName** | **string** |  | 
**Capability** | Pointer to **NullableString** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**CertificateStoresTypesSupportedOperations**](CertificateStoresTypesSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]CertificateStoresTypesStoreTypeProperty**](CertificateStoresTypesStoreTypeProperty.md) |  | [optional] 
**PasswordOptions** | Pointer to [**CertificateStoresTypesPasswordOptions**](CertificateStoresTypesPasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **NullableString** |  | [optional] 
**StorePathValue** | Pointer to **NullableString** |  | [optional] 
**PrivateKeyAllowed** | Pointer to [**CSSCMSCoreEnumsCertStorePrivateKey**](CSSCMSCoreEnumsCertStorePrivateKey.md) |  | [optional] 
**CertificateFormat** | Pointer to [**CSSCMSCoreEnumsCertificateFormat**](CSSCMSCoreEnumsCertificateFormat.md) |  | [optional] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to [**KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias**](KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias.md) |  | [optional] 
**EntryParameters** | Pointer to [**[]CertificateStoresTypesEntryParameters**](CertificateStoresTypesEntryParameters.md) |  | [optional] 

## Methods

### NewCertificateStoresTypesCertificateStoreTypeUpdateRequest

`func NewCertificateStoresTypesCertificateStoreTypeUpdateRequest(storeType int32, name string, shortName string, ) *CertificateStoresTypesCertificateStoreTypeUpdateRequest`

NewCertificateStoresTypesCertificateStoreTypeUpdateRequest instantiates a new CertificateStoresTypesCertificateStoreTypeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults

`func NewCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults() *CertificateStoresTypesCertificateStoreTypeUpdateRequest`

NewCertificateStoresTypesCertificateStoreTypeUpdateRequestWithDefaults instantiates a new CertificateStoresTypesCertificateStoreTypeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreType

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStoreType() int32`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStoreTypeOk() (*int32, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStoreType(v int32)`

SetStoreType sets StoreType field to given value.


### GetName

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetCapability

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### SetCapabilityNil

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCapabilityNil(b bool)`

 SetCapabilityNil sets the value for Capability to be an explicit nil

### UnsetCapability
`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetCapability()`

UnsetCapability ensures that no value is present for Capability, not even an explicit nil
### GetLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetSupportedOperations() CertificateStoresTypesSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetSupportedOperationsOk() (*CertificateStoresTypesSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetSupportedOperations(v CertificateStoresTypesSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetProperties() []CertificateStoresTypesStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPropertiesOk() (*[]CertificateStoresTypesStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetProperties(v []CertificateStoresTypesStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPasswordOptions() CertificateStoresTypesPasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPasswordOptionsOk() (*CertificateStoresTypesPasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPasswordOptions(v CertificateStoresTypesPasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### SetStorePathTypeNil

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathTypeNil(b bool)`

 SetStorePathTypeNil sets the value for StorePathType to be an explicit nil

### UnsetStorePathType
`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetStorePathType()`

UnsetStorePathType ensures that no value is present for StorePathType, not even an explicit nil
### GetStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### SetStorePathValueNil

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetStorePathValueNil(b bool)`

 SetStorePathValueNil sets the value for StorePathValue to be an explicit nil

### UnsetStorePathValue
`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetStorePathValue()`

UnsetStorePathValue ensures that no value is present for StorePathValue, not even an explicit nil
### GetPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPrivateKeyAllowed() CSSCMSCoreEnumsCertStorePrivateKey`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPrivateKeyAllowedOk() (*CSSCMSCoreEnumsCertStorePrivateKey, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPrivateKeyAllowed(v CSSCMSCoreEnumsCertStorePrivateKey)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCertificateFormat() CSSCMSCoreEnumsCertificateFormat`

GetCertificateFormat returns the CertificateFormat field if non-nil, zero value otherwise.

### GetCertificateFormatOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCertificateFormatOk() (*CSSCMSCoreEnumsCertificateFormat, bool)`

GetCertificateFormatOk returns a tuple with the CertificateFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCertificateFormat(v CSSCMSCoreEnumsCertificateFormat)`

SetCertificateFormat sets CertificateFormat field to given value.

### HasCertificateFormat

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCertificateFormat() bool`

HasCertificateFormat returns a boolean if a field has been set.

### GetServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCustomAliasAllowed() KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetCustomAliasAllowedOk() (*KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetCustomAliasAllowed(v KeyfactorOrchestratorsCommonEnumsCertStoreCustomAlias)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetEntryParameters() []CertificateStoresTypesEntryParameters`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) GetEntryParametersOk() (*[]CertificateStoresTypesEntryParameters, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetEntryParameters(v []CertificateStoresTypesEntryParameters)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### SetEntryParametersNil

`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) SetEntryParametersNil(b bool)`

 SetEntryParametersNil sets the value for EntryParameters to be an explicit nil

### UnsetEntryParameters
`func (o *CertificateStoresTypesCertificateStoreTypeUpdateRequest) UnsetEntryParameters()`

UnsetEntryParameters ensures that no value is present for EntryParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


