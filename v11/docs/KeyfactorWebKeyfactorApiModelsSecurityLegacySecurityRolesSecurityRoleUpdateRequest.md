# KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Id of the security role to update | 
**Name** | **string** | The name of the security role to update | 
**Description** | **string** | The description to be used on the updated security role | 
**Enabled** | Pointer to **bool** | Whether or not the security role should be enabled | [optional] 
**Private** | Pointer to **bool** | Whether or not the security role should be private | [optional] 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format \&quot;Area:Permission\&quot; | [optional] 
**PermissionSetId** | Pointer to **string** | The Id of the permission set the role belongs to. | [optional] 
**Identities** | Pointer to [**[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier**](CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier.md) | The Keyfactor identities to assign to the updated role | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest(id int32, name string, description string, ) *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdentities() []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdentitiesOk() (*[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetIdentities(v []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleUpdateRequest) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


