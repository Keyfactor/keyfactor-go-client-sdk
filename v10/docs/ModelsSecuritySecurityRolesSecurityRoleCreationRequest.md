# ModelsSecuritySecurityRolesSecurityRoleCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the security role to create | 
**Description** | **string** | The description to be used on the created security role | 
**Enabled** | Pointer to **bool** | Whether or not the security role should be enabled | [optional] 
**Private** | Pointer to **bool** | Whether or not the security role should be private | [optional] 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format \&quot;Area:Permission\&quot; | [optional] 
**Identities** | Pointer to [**[]ModelsSecurityIdentitiesSecurityIdentityIdentifier**](ModelsSecurityIdentitiesSecurityIdentityIdentifier.md) | The Keyfactor identities to assign to the created role | [optional] 

## Methods

### NewModelsSecuritySecurityRolesSecurityRoleCreationRequest

`func NewModelsSecuritySecurityRolesSecurityRoleCreationRequest(name string, description string, ) *ModelsSecuritySecurityRolesSecurityRoleCreationRequest`

NewModelsSecuritySecurityRolesSecurityRoleCreationRequest instantiates a new ModelsSecuritySecurityRolesSecurityRoleCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults

`func NewModelsSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults() *ModelsSecuritySecurityRolesSecurityRoleCreationRequest`

NewModelsSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults instantiates a new ModelsSecuritySecurityRolesSecurityRoleCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetIdentities

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetIdentities() []ModelsSecurityIdentitiesSecurityIdentityIdentifier`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) GetIdentitiesOk() (*[]ModelsSecurityIdentitiesSecurityIdentityIdentifier, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) SetIdentities(v []ModelsSecurityIdentitiesSecurityIdentityIdentifier)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *ModelsSecuritySecurityRolesSecurityRoleCreationRequest) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


