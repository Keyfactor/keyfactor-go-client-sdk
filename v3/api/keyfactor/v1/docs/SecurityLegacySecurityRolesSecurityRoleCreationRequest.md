# SecurityLegacySecurityRolesSecurityRoleCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the security role to create | 
**Description** | **string** | The description to be used on the created security role | 
**EmailAddress** | Pointer to **NullableString** | The email address to be used on the created security role | [optional] 
**Enabled** | Pointer to **bool** | Whether or not the security role should be enabled | [optional] 
**Private** | Pointer to **bool** | Whether or not the security role should be private | [optional] 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format \&quot;Area:Permission\&quot; | [optional] 
**PermissionSetId** | Pointer to **string** | The Id of the permission set the role belongs to. | [optional] 
**Identities** | Pointer to [**[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier**](CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier.md) | The Keyfactor identities to assign to the created role | [optional] 

## Methods

### NewSecurityLegacySecurityRolesSecurityRoleCreationRequest

`func NewSecurityLegacySecurityRolesSecurityRoleCreationRequest(name string, description string, ) *SecurityLegacySecurityRolesSecurityRoleCreationRequest`

NewSecurityLegacySecurityRolesSecurityRoleCreationRequest instantiates a new SecurityLegacySecurityRolesSecurityRoleCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLegacySecurityRolesSecurityRoleCreationRequestWithDefaults

`func NewSecurityLegacySecurityRolesSecurityRoleCreationRequestWithDefaults() *SecurityLegacySecurityRolesSecurityRoleCreationRequest`

NewSecurityLegacySecurityRolesSecurityRoleCreationRequestWithDefaults instantiates a new SecurityLegacySecurityRolesSecurityRoleCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetIdentities() []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) GetIdentitiesOk() (*[]CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetIdentities(v []CSSCMSDataModelModelsSecurityIdentitiesSecurityIdentityIdentifier)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *SecurityLegacySecurityRolesSecurityRoleCreationRequest) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


