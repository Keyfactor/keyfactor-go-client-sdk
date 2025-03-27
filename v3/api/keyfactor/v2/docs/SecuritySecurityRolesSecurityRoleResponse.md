# SecuritySecurityRolesSecurityRoleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**EmailAddress** | Pointer to **NullableString** |  | [optional] 
**Immutable** | Pointer to **bool** |  | [optional] 
**PermissionSetId** | Pointer to **string** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 
**Claims** | Pointer to [**[]SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse**](SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse.md) |  | [optional] 

## Methods

### NewSecuritySecurityRolesSecurityRoleResponse

`func NewSecuritySecurityRolesSecurityRoleResponse() *SecuritySecurityRolesSecurityRoleResponse`

NewSecuritySecurityRolesSecurityRoleResponse instantiates a new SecuritySecurityRolesSecurityRoleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityRolesSecurityRoleResponseWithDefaults

`func NewSecuritySecurityRolesSecurityRoleResponseWithDefaults() *SecuritySecurityRolesSecurityRoleResponse`

NewSecuritySecurityRolesSecurityRoleResponseWithDefaults instantiates a new SecuritySecurityRolesSecurityRoleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *SecuritySecurityRolesSecurityRoleResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetDescription

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *SecuritySecurityRolesSecurityRoleResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SecuritySecurityRolesSecurityRoleResponse) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetImmutable

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetImmutable() bool`

GetImmutable returns the Immutable field if non-nil, zero value otherwise.

### GetImmutableOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetImmutableOk() (*bool, bool)`

GetImmutableOk returns a tuple with the Immutable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImmutable

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetImmutable(v bool)`

SetImmutable sets Immutable field to given value.

### HasImmutable

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasImmutable() bool`

HasImmutable returns a boolean if a field has been set.

### GetPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.

### HasPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasPermissionSetId() bool`

HasPermissionSetId returns a boolean if a field has been set.

### GetPermissions

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecuritySecurityRolesSecurityRoleResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetClaims

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetClaims() []SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SecuritySecurityRolesSecurityRoleResponse) GetClaimsOk() (*[]SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetClaims(v []SecurityRoleClaimDefinitionsRoleClaimDefinitionResponse)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *SecuritySecurityRolesSecurityRoleResponse) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *SecuritySecurityRolesSecurityRoleResponse) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *SecuritySecurityRolesSecurityRoleResponse) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


