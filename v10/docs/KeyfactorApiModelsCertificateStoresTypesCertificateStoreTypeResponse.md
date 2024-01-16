# KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Capability** | Pointer to **string** |  | [optional] 
**StoreType** | Pointer to **int64** |  | [optional] 
**ImportType** | Pointer to **int64** |  | [optional] 
**LocalStore** | Pointer to **bool** |  | [optional] 
**SupportedOperations** | Pointer to [**ModelsCertStoreTypeSupportedOperations**](ModelsCertStoreTypeSupportedOperations.md) |  | [optional] 
**Properties** | Pointer to [**[]ModelsCertificateStoreTypeProperty**](ModelsCertificateStoreTypeProperty.md) |  | [optional] 
**EntryParameters** | Pointer to [**[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 
**PasswordOptions** | Pointer to [**ModelsCertStoreTypePasswordOptions**](ModelsCertStoreTypePasswordOptions.md) |  | [optional] 
**StorePathType** | Pointer to **string** |  | [optional] 
**StorePathValue** | Pointer to **string** |  | [optional] 
**PrivateKeyAllowed** | Pointer to **string** |  | [optional] 
**JobProperties** | Pointer to **[]string** |  | [optional] [readonly] 
**ServerRequired** | Pointer to **bool** |  | [optional] 
**PowerShell** | Pointer to **bool** |  | [optional] 
**BlueprintAllowed** | Pointer to **bool** |  | [optional] 
**CustomAliasAllowed** | Pointer to **string** |  | [optional] 
**ServerRegistration** | Pointer to **int64** |  | [optional] 
**InventoryEndpoint** | Pointer to **string** |  | [optional] 
**InventoryJobType** | Pointer to **string** |  | [optional] 
**ManagementJobType** | Pointer to **string** |  | [optional] 
**DiscoveryJobType** | Pointer to **string** |  | [optional] 
**EnrollmentJobType** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse

`func NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse() *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse`

NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse instantiates a new KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponseWithDefaults

`func NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponseWithDefaults() *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse`

NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponseWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetStoreType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStoreType() int64`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStoreTypeOk() (*int64, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStoreType(v int64)`

SetStoreType sets StoreType field to given value.

### HasStoreType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasStoreType() bool`

HasStoreType returns a boolean if a field has been set.

### GetImportType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetImportType() int64`

GetImportType returns the ImportType field if non-nil, zero value otherwise.

### GetImportTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetImportTypeOk() (*int64, bool)`

GetImportTypeOk returns a tuple with the ImportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetImportType(v int64)`

SetImportType sets ImportType field to given value.

### HasImportType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasImportType() bool`

HasImportType returns a boolean if a field has been set.

### GetLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetSupportedOperations() ModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetSupportedOperationsOk() (*ModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetSupportedOperations(v ModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetProperties() []ModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPropertiesOk() (*[]ModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetProperties(v []ModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEntryParameters() []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEntryParametersOk() (*[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetEntryParameters(v []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPasswordOptions() ModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPasswordOptionsOk() (*ModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPasswordOptions(v ModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### GetStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### GetPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPrivateKeyAllowed() string`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPrivateKeyAllowedOk() (*string, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPrivateKeyAllowed(v string)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetJobProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetJobProperties() []string`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetJobPropertiesOk() (*[]string, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetJobProperties(v []string)`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### GetServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCustomAliasAllowed() string`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetCustomAliasAllowedOk() (*string, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetCustomAliasAllowed(v string)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetServerRegistration

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRegistration() int64`

GetServerRegistration returns the ServerRegistration field if non-nil, zero value otherwise.

### GetServerRegistrationOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetServerRegistrationOk() (*int64, bool)`

GetServerRegistrationOk returns a tuple with the ServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistration

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetServerRegistration(v int64)`

SetServerRegistration sets ServerRegistration field to given value.

### HasServerRegistration

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasServerRegistration() bool`

HasServerRegistration returns a boolean if a field has been set.

### GetInventoryEndpoint

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryEndpoint() string`

GetInventoryEndpoint returns the InventoryEndpoint field if non-nil, zero value otherwise.

### GetInventoryEndpointOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryEndpointOk() (*string, bool)`

GetInventoryEndpointOk returns a tuple with the InventoryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryEndpoint

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetInventoryEndpoint(v string)`

SetInventoryEndpoint sets InventoryEndpoint field to given value.

### HasInventoryEndpoint

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasInventoryEndpoint() bool`

HasInventoryEndpoint returns a boolean if a field has been set.

### GetInventoryJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryJobType() string`

GetInventoryJobType returns the InventoryJobType field if non-nil, zero value otherwise.

### GetInventoryJobTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetInventoryJobTypeOk() (*string, bool)`

GetInventoryJobTypeOk returns a tuple with the InventoryJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetInventoryJobType(v string)`

SetInventoryJobType sets InventoryJobType field to given value.

### HasInventoryJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasInventoryJobType() bool`

HasInventoryJobType returns a boolean if a field has been set.

### GetManagementJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetManagementJobType() string`

GetManagementJobType returns the ManagementJobType field if non-nil, zero value otherwise.

### GetManagementJobTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetManagementJobTypeOk() (*string, bool)`

GetManagementJobTypeOk returns a tuple with the ManagementJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetManagementJobType(v string)`

SetManagementJobType sets ManagementJobType field to given value.

### HasManagementJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasManagementJobType() bool`

HasManagementJobType returns a boolean if a field has been set.

### GetDiscoveryJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetDiscoveryJobType() string`

GetDiscoveryJobType returns the DiscoveryJobType field if non-nil, zero value otherwise.

### GetDiscoveryJobTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetDiscoveryJobTypeOk() (*string, bool)`

GetDiscoveryJobTypeOk returns a tuple with the DiscoveryJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetDiscoveryJobType(v string)`

SetDiscoveryJobType sets DiscoveryJobType field to given value.

### HasDiscoveryJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasDiscoveryJobType() bool`

HasDiscoveryJobType returns a boolean if a field has been set.

### GetEnrollmentJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEnrollmentJobType() string`

GetEnrollmentJobType returns the EnrollmentJobType field if non-nil, zero value otherwise.

### GetEnrollmentJobTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) GetEnrollmentJobTypeOk() (*string, bool)`

GetEnrollmentJobTypeOk returns a tuple with the EnrollmentJobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) SetEnrollmentJobType(v string)`

SetEnrollmentJobType sets EnrollmentJobType field to given value.

### HasEnrollmentJobType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeResponse) HasEnrollmentJobType() bool`

HasEnrollmentJobType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


