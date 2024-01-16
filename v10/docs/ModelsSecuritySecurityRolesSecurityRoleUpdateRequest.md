# ModelsSecuritySecurityRolesSecurityRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The Id of the security role to update | 
**Name** | **string** | The name of the security role to update | 
**Description** | **string** | The description to be used on the updated security role | 
**Enabled** | Pointer to **bool** | Whether or not the security role should be enabled | [optional] 
**Private** | Pointer to **bool** | Whether or not the security role should be private | [optional] 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format \&quot;Area:Permission\&quot; | [optional] 
**Identities** | Pointer to [**[]ModelsSecurityIdentitiesSecurityIdentityIdentifier**](ModelsSecurityIdentitiesSecurityIdentityIdentifier.md) | The Keyfactor identities to assign to the updated role | [optional] 

## Methods

### NewModelsSecuritySecurityRolesSecurityRoleUpdateRequest

`func NewModelsSecuritySecurityRolesSecurityRoleUpdateRequest(id int64, name string, description string, ) *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest`

NewModelsSecuritySecurityRolesSecurityRoleUpdateRequest instantiates a new ModelsSecuritySecurityRolesSecurityRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSecuritySecurityRolesSecurityRoleUpdateRequestWithDefaults

`func NewModelsSecuritySecurityRolesSecurityRoleUpdateRequestWithDefaults() *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest`

NewModelsSecuritySecurityRolesSecurityRoleUpdateRequestWithDefaults instantiates a new ModelsSecuritySecurityRolesSecurityRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetName

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetIdentities

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetIdentities() []ModelsSecurityIdentitiesSecurityIdentityIdentifier`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) GetIdentitiesOk() (*[]ModelsSecurityIdentitiesSecurityIdentityIdentifier, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) SetIdentities(v []ModelsSecurityIdentitiesSecurityIdentityIdentifier)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *ModelsSecuritySecurityRolesSecurityRoleUpdateRequest) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


