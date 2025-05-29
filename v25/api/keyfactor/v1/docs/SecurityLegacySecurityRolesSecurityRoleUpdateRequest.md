# SecurityLegacySecurityRolesSecurityRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Id of the security role to update | 
**Name** | **string** | The name of the security role to update | 
**Description** | **string** | The description to be used on the updated security role | 
**EmailAddress** | Pointer to **NullableString** | The email address to be used on the updated security role | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the security role should be enabled | [optional] 
**Private** | Pointer to **bool** | Whether or not the security role should be private | [optional] 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format \&quot;Area:Permission\&quot; | [optional] 
**PermissionSetId** | Pointer to **string** | The Id of the permission set the role belongs to. | [optional] 
**Identities** | Pointer to [**[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier**](CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier.md) | The Keyfactor identities to assign to the updated role | [optional] 

## Methods

### NewSecurityLegacySecurityRolesSecurityRoleUpdateRequest

`func NewSecurityLegacySecurityRolesSecurityRoleUpdateRequest(id int32, name string, description string, ) *SecurityLegacySecurityRolesSecurityRoleUpdateRequest`

NewSecurityLegacySecurityRolesSecurityRoleUpdateRequest instantiates a new SecurityLegacySecurityRolesSecurityRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults

`func NewSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults() *SecurityLegacySecurityRolesSecurityRoleUpdateRequest`

NewSecurityLegacySecurityRolesSecurityRoleUpdateRequestWithDefaults instantiates a new SecurityLegacySecurityRolesSecurityRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdentities() []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) GetIdentitiesOk() (*[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetIdentities(v []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *SecurityLegacySecurityRolesSecurityRoleUpdateRequest) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


