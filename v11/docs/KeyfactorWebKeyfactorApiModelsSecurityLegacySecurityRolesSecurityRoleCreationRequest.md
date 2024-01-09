# KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the security role to create | 
**Description** | **string** | The description to be used on the created security role | 
**Enabled** | Pointer to **bool** | Whether or not the security role should be enabled | [optional] 
**Private** | Pointer to **bool** | Whether or not the security role should be private | [optional] 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format \&quot;Area:Permission\&quot; | [optional] 
**PermissionSetId** | Pointer to **string** | The Id of the permission set the role belongs to. | [optional] 
**Identities** | Pointer to [**[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier**](CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier.md) | The Keyfactor identities to assign to the created role | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest(name string, description string, ) *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetIdentities() []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) GetIdentitiesOk() (*[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetIdentities(v []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleCreationRequest) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


