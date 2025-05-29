# SecurityLegacySecurityRolesSecurityRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableInt32** | The Id of the created role | [optional] 
**Description** | Pointer to **NullableString** | The description of the created role | [optional] 
**EmailAddress** | Pointer to **NullableString** | The email address to be used on the created security role | [optional] 
**Enabled** | Pointer to **NullableBool** | A boolean indicating whether or not the created role is enabled | [optional] 
**Immutable** | Pointer to **NullableBool** | A boolean indicating whther or not the security role will be read-only | [optional] 
**Valid** | Pointer to **NullableBool** | A boolean that indicates whether or not the Audit XML was able to be verified | [optional] 
**Private** | Pointer to **NullableBool** | A boolean that indicates whether or not the created security role is private | [optional] 
**PermissionSetId** | Pointer to **string** | The Id of the permission set the role belongs to. | [optional] 
**Identities** | Pointer to [**[]SecurityLegacySecurityRolesSecurityIdentityResponse**](SecurityLegacySecurityRolesSecurityIdentityResponse.md) | The identities assigned to the created security role | [optional] 
**Name** | Pointer to **NullableString** | The name of the created role | [optional] 
**Permissions** | Pointer to **[]string** | The permissions included in the created security role | [optional] 

## Methods

### NewSecurityLegacySecurityRolesSecurityRoleResponse

`func NewSecurityLegacySecurityRolesSecurityRoleResponse() *SecurityLegacySecurityRolesSecurityRoleResponse`

NewSecurityLegacySecurityRolesSecurityRoleResponse instantiates a new SecurityLegacySecurityRolesSecurityRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLegacySecurityRolesSecurityRoleResponseWithDefaults

`func NewSecurityLegacySecurityRolesSecurityRoleResponseWithDefaults() *SecurityLegacySecurityRolesSecurityRoleResponse`

NewSecurityLegacySecurityRolesSecurityRoleResponseWithDefaults instantiates a new SecurityLegacySecurityRolesSecurityRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetImmutable

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### SetImmutableNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetImmutableNil(b bool)`

 SetImmutableNil sets the value for Immutable to be an explicit nil

### UnsetImmutable
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetImmutable()`

UnsetImmutable ensures that no value is present for Immutable, not even an explicit nil
### GetValid

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil
### GetPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### SetPrivateNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetPrivateNil(b bool)`

 SetPrivateNil sets the value for Private to be an explicit nil

### UnsetPrivate
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetPrivate()`

UnsetPrivate ensures that no value is present for Private, not even an explicit nil
### GetPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetIdentities() []SecurityLegacySecurityRolesSecurityIdentityResponse`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetIdentitiesOk() (*[]SecurityLegacySecurityRolesSecurityIdentityResponse, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetIdentities(v []SecurityLegacySecurityRolesSecurityIdentityResponse)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil
### GetName

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecurityLegacySecurityRolesSecurityRoleResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


