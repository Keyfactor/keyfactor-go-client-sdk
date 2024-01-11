# KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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
**ServerRegistration** | Pointer to **int32** |  | [optional] 
**InventoryEndpoint** | Pointer to **string** |  | [optional] 
**InventoryJobTypeId** | Pointer to **string** |  | [optional] 
**ManagementJobTypeId** | Pointer to **string** |  | [optional] 
**DiscoveryJobTypeId** | Pointer to **string** |  | [optional] 
**EnrollmentJobTypeId** | Pointer to **string** |  | [optional] 
**EntryParameters** | Pointer to [**[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter**](ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest

`func NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest(name string, shortName string, ) *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest`

NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest instantiates a new KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults

`func NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults() *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest`

NewKeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequestWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetShortName(v string)`

SetShortName sets ShortName field to given value.


### GetCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCapability() string`

GetCapability returns the Capability field if non-nil, zero value otherwise.

### GetCapabilityOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCapabilityOk() (*string, bool)`

GetCapabilityOk returns a tuple with the Capability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetCapability(v string)`

SetCapability sets Capability field to given value.

### HasCapability

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasCapability() bool`

HasCapability returns a boolean if a field has been set.

### GetLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetLocalStore() bool`

GetLocalStore returns the LocalStore field if non-nil, zero value otherwise.

### GetLocalStoreOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetLocalStoreOk() (*bool, bool)`

GetLocalStoreOk returns a tuple with the LocalStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetLocalStore(v bool)`

SetLocalStore sets LocalStore field to given value.

### HasLocalStore

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasLocalStore() bool`

HasLocalStore returns a boolean if a field has been set.

### GetSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetSupportedOperations() ModelsCertStoreTypeSupportedOperations`

GetSupportedOperations returns the SupportedOperations field if non-nil, zero value otherwise.

### GetSupportedOperationsOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetSupportedOperationsOk() (*ModelsCertStoreTypeSupportedOperations, bool)`

GetSupportedOperationsOk returns a tuple with the SupportedOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetSupportedOperations(v ModelsCertStoreTypeSupportedOperations)`

SetSupportedOperations sets SupportedOperations field to given value.

### HasSupportedOperations

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasSupportedOperations() bool`

HasSupportedOperations returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetProperties() []ModelsCertificateStoreTypeProperty`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPropertiesOk() (*[]ModelsCertificateStoreTypeProperty, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetProperties(v []ModelsCertificateStoreTypeProperty)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPasswordOptions() ModelsCertStoreTypePasswordOptions`

GetPasswordOptions returns the PasswordOptions field if non-nil, zero value otherwise.

### GetPasswordOptionsOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPasswordOptionsOk() (*ModelsCertStoreTypePasswordOptions, bool)`

GetPasswordOptionsOk returns a tuple with the PasswordOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPasswordOptions(v ModelsCertStoreTypePasswordOptions)`

SetPasswordOptions sets PasswordOptions field to given value.

### HasPasswordOptions

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasPasswordOptions() bool`

HasPasswordOptions returns a boolean if a field has been set.

### GetStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathType() string`

GetStorePathType returns the StorePathType field if non-nil, zero value otherwise.

### GetStorePathTypeOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathTypeOk() (*string, bool)`

GetStorePathTypeOk returns a tuple with the StorePathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathType(v string)`

SetStorePathType sets StorePathType field to given value.

### HasStorePathType

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasStorePathType() bool`

HasStorePathType returns a boolean if a field has been set.

### GetStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathValue() string`

GetStorePathValue returns the StorePathValue field if non-nil, zero value otherwise.

### GetStorePathValueOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetStorePathValueOk() (*string, bool)`

GetStorePathValueOk returns a tuple with the StorePathValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetStorePathValue(v string)`

SetStorePathValue sets StorePathValue field to given value.

### HasStorePathValue

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasStorePathValue() bool`

HasStorePathValue returns a boolean if a field has been set.

### GetPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPrivateKeyAllowed() string`

GetPrivateKeyAllowed returns the PrivateKeyAllowed field if non-nil, zero value otherwise.

### GetPrivateKeyAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPrivateKeyAllowedOk() (*string, bool)`

GetPrivateKeyAllowedOk returns a tuple with the PrivateKeyAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPrivateKeyAllowed(v string)`

SetPrivateKeyAllowed sets PrivateKeyAllowed field to given value.

### HasPrivateKeyAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasPrivateKeyAllowed() bool`

HasPrivateKeyAllowed returns a boolean if a field has been set.

### GetServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRequired() bool`

GetServerRequired returns the ServerRequired field if non-nil, zero value otherwise.

### GetServerRequiredOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRequiredOk() (*bool, bool)`

GetServerRequiredOk returns a tuple with the ServerRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRequired(v bool)`

SetServerRequired sets ServerRequired field to given value.

### HasServerRequired

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasServerRequired() bool`

HasServerRequired returns a boolean if a field has been set.

### GetPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPowerShell() bool`

GetPowerShell returns the PowerShell field if non-nil, zero value otherwise.

### GetPowerShellOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetPowerShellOk() (*bool, bool)`

GetPowerShellOk returns a tuple with the PowerShell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetPowerShell(v bool)`

SetPowerShell sets PowerShell field to given value.

### HasPowerShell

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasPowerShell() bool`

HasPowerShell returns a boolean if a field has been set.

### GetBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetBlueprintAllowed() bool`

GetBlueprintAllowed returns the BlueprintAllowed field if non-nil, zero value otherwise.

### GetBlueprintAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetBlueprintAllowedOk() (*bool, bool)`

GetBlueprintAllowedOk returns a tuple with the BlueprintAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetBlueprintAllowed(v bool)`

SetBlueprintAllowed sets BlueprintAllowed field to given value.

### HasBlueprintAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasBlueprintAllowed() bool`

HasBlueprintAllowed returns a boolean if a field has been set.

### GetCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCustomAliasAllowed() string`

GetCustomAliasAllowed returns the CustomAliasAllowed field if non-nil, zero value otherwise.

### GetCustomAliasAllowedOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetCustomAliasAllowedOk() (*string, bool)`

GetCustomAliasAllowedOk returns a tuple with the CustomAliasAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetCustomAliasAllowed(v string)`

SetCustomAliasAllowed sets CustomAliasAllowed field to given value.

### HasCustomAliasAllowed

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasCustomAliasAllowed() bool`

HasCustomAliasAllowed returns a boolean if a field has been set.

### GetServerRegistration

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRegistration() int32`

GetServerRegistration returns the ServerRegistration field if non-nil, zero value otherwise.

### GetServerRegistrationOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetServerRegistrationOk() (*int32, bool)`

GetServerRegistrationOk returns a tuple with the ServerRegistration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRegistration

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetServerRegistration(v int32)`

SetServerRegistration sets ServerRegistration field to given value.

### HasServerRegistration

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasServerRegistration() bool`

HasServerRegistration returns a boolean if a field has been set.

### GetInventoryEndpoint

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryEndpoint() string`

GetInventoryEndpoint returns the InventoryEndpoint field if non-nil, zero value otherwise.

### GetInventoryEndpointOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryEndpointOk() (*string, bool)`

GetInventoryEndpointOk returns a tuple with the InventoryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryEndpoint

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryEndpoint(v string)`

SetInventoryEndpoint sets InventoryEndpoint field to given value.

### HasInventoryEndpoint

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasInventoryEndpoint() bool`

HasInventoryEndpoint returns a boolean if a field has been set.

### GetInventoryJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryJobTypeId() string`

GetInventoryJobTypeId returns the InventoryJobTypeId field if non-nil, zero value otherwise.

### GetInventoryJobTypeIdOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetInventoryJobTypeIdOk() (*string, bool)`

GetInventoryJobTypeIdOk returns a tuple with the InventoryJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetInventoryJobTypeId(v string)`

SetInventoryJobTypeId sets InventoryJobTypeId field to given value.

### HasInventoryJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasInventoryJobTypeId() bool`

HasInventoryJobTypeId returns a boolean if a field has been set.

### GetManagementJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetManagementJobTypeId() string`

GetManagementJobTypeId returns the ManagementJobTypeId field if non-nil, zero value otherwise.

### GetManagementJobTypeIdOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetManagementJobTypeIdOk() (*string, bool)`

GetManagementJobTypeIdOk returns a tuple with the ManagementJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetManagementJobTypeId(v string)`

SetManagementJobTypeId sets ManagementJobTypeId field to given value.

### HasManagementJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasManagementJobTypeId() bool`

HasManagementJobTypeId returns a boolean if a field has been set.

### GetDiscoveryJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetDiscoveryJobTypeId() string`

GetDiscoveryJobTypeId returns the DiscoveryJobTypeId field if non-nil, zero value otherwise.

### GetDiscoveryJobTypeIdOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetDiscoveryJobTypeIdOk() (*string, bool)`

GetDiscoveryJobTypeIdOk returns a tuple with the DiscoveryJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetDiscoveryJobTypeId(v string)`

SetDiscoveryJobTypeId sets DiscoveryJobTypeId field to given value.

### HasDiscoveryJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasDiscoveryJobTypeId() bool`

HasDiscoveryJobTypeId returns a boolean if a field has been set.

### GetEnrollmentJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEnrollmentJobTypeId() string`

GetEnrollmentJobTypeId returns the EnrollmentJobTypeId field if non-nil, zero value otherwise.

### GetEnrollmentJobTypeIdOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEnrollmentJobTypeIdOk() (*string, bool)`

GetEnrollmentJobTypeIdOk returns a tuple with the EnrollmentJobTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetEnrollmentJobTypeId(v string)`

SetEnrollmentJobTypeId sets EnrollmentJobTypeId field to given value.

### HasEnrollmentJobTypeId

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasEnrollmentJobTypeId() bool`

HasEnrollmentJobTypeId returns a boolean if a field has been set.

### GetEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEntryParameters() []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter`

GetEntryParameters returns the EntryParameters field if non-nil, zero value otherwise.

### GetEntryParametersOk

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) GetEntryParametersOk() (*[]ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter, bool)`

GetEntryParametersOk returns a tuple with the EntryParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) SetEntryParameters(v []ModelsCertificateStoreTypesCertificateStoreTypeEntryParameter)`

SetEntryParameters sets EntryParameters field to given value.

### HasEntryParameters

`func (o *KeyfactorApiModelsCertificateStoresTypesCertificateStoreTypeCreationRequest) HasEntryParameters() bool`

HasEntryParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


