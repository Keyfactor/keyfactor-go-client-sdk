# SecuritySecurityRolesSecurityRoleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | The Keyfactor identifier of the role to update | 
**Name** | **string** | The new name of the security role | 
**Description** | **string** | The new description to be used for the security role | 
**EmailAddress** | Pointer to **NullableString** | The new email address to be used for the security role | [optional] 
**PermissionSetId** | **string** |  | 
**Permissions** | Pointer to **[]string** | The complete list of permissions to include in the role. These must be supplied in the format of access control strings such as \&quot;/security/modify/\&quot;. | [optional] 
**Claims** | Pointer to [**[]SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest**](SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest.md) | The complete list of claim definitions to assign to the role | [optional] 

## Methods

### NewSecuritySecurityRolesSecurityRoleUpdateRequest

`func NewSecuritySecurityRolesSecurityRoleUpdateRequest(id int32, name string, description string, permissionSetId string, ) *SecuritySecurityRolesSecurityRoleUpdateRequest`

NewSecuritySecurityRolesSecurityRoleUpdateRequest instantiates a new SecuritySecurityRolesSecurityRoleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityRolesSecurityRoleUpdateRequestWithDefaults

`func NewSecuritySecurityRolesSecurityRoleUpdateRequestWithDefaults() *SecuritySecurityRolesSecurityRoleUpdateRequest`

NewSecuritySecurityRolesSecurityRoleUpdateRequestWithDefaults instantiates a new SecuritySecurityRolesSecurityRoleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.


### GetPermissions

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetClaims

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetClaims() []SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) GetClaimsOk() (*[]SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetClaims(v []SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *SecuritySecurityRolesSecurityRoleUpdateRequest) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


