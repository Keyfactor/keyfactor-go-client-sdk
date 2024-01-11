# KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** | The name of the created role | [optional] 
**Permissions** | Pointer to **[]string** | The permissions included in the created security role | [optional] 
**Id** | Pointer to **NullableInt32** | The Id of the created role | [optional] 
**Description** | Pointer to **NullableString** | The description of the created role | [optional] 
**Enabled** | Pointer to **NullableBool** | A boolean indicating whether or not the created role is enabled | [optional] 
**Immutable** | Pointer to **NullableBool** | A boolean indicating whther or not the security role will be read-only | [optional] 
**Valid** | Pointer to **NullableBool** | A boolean that indicates whether or not the Audit XML was able to be verified | [optional] 
**Private** | Pointer to **NullableBool** | A boolean that indicates whether or not the created security role is private | [optional] 
**PermissionSetId** | Pointer to **string** | The Id of the permission set the role belongs to. | [optional] 
**Identities** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse**](KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse.md) | The identities assigned to the created security role | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse`

NewKeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### SetEnabledNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetEnabledNil(b bool)`

 SetEnabledNil sets the value for Enabled to be an explicit nil

### UnsetEnabled
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetEnabled()`

UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
### GetImmutable

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### SetImmutableNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetImmutableNil(b bool)`

 SetImmutableNil sets the value for Immutable to be an explicit nil

### UnsetImmutable
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetImmutable()`

UnsetImmutable ensures that no value is present for Immutable, not even an explicit nil
### GetValid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### SetValidNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetValidNil(b bool)`

 SetValidNil sets the value for Valid to be an explicit nil

### UnsetValid
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetValid()`

UnsetValid ensures that no value is present for Valid, not even an explicit nil
### GetPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### SetPrivateNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetPrivateNil(b bool)`

 SetPrivateNil sets the value for Private to be an explicit nil

### UnsetPrivate
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetPrivate()`

UnsetPrivate ensures that no value is present for Private, not even an explicit nil
### GetPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetIdentities() []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse`

GetIdentities returns the Identities field if non-nil, zero value otherwise.

### GetIdentitiesOk

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) GetIdentitiesOk() (*[]KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse, bool)`

GetIdentitiesOk returns a tuple with the Identities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetIdentities(v []KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityIdentityResponse)`

SetIdentities sets Identities field to given value.

### HasIdentities

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) HasIdentities() bool`

HasIdentities returns a boolean if a field has been set.

### SetIdentitiesNil

`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) SetIdentitiesNil(b bool)`

 SetIdentitiesNil sets the value for Identities to be an explicit nil

### UnsetIdentities
`func (o *KeyfactorWebKeyfactorApiModelsSecurityLegacySecurityRolesSecurityRoleResponse) UnsetIdentities()`

UnsetIdentities ensures that no value is present for Identities, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


