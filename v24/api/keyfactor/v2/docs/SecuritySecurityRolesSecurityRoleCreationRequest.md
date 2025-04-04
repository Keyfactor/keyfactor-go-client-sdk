# SecuritySecurityRolesSecurityRoleCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the security role to create | 
**Description** | **string** | The description to be used on the created security role | 
**EmailAddress** | Pointer to **NullableString** | The email address to be used on the created security role | [optional] 
**PermissionSetId** | **string** | The Id of the permission set this role is in. | 
**Permissions** | Pointer to **[]string** | The permissions to include in the role. These must be supplied in the format of access control strings such as \&quot;/security/modify/\&quot;. | [optional] 
**Claims** | Pointer to [**[]SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest**](SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest.md) | The claim definitions to assign to the created role | [optional] 

## Methods

### NewSecuritySecurityRolesSecurityRoleCreationRequest

`func NewSecuritySecurityRolesSecurityRoleCreationRequest(name string, description string, permissionSetId string, ) *SecuritySecurityRolesSecurityRoleCreationRequest`

NewSecuritySecurityRolesSecurityRoleCreationRequest instantiates a new SecuritySecurityRolesSecurityRoleCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults

`func NewSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults() *SecuritySecurityRolesSecurityRoleCreationRequest`

NewSecuritySecurityRolesSecurityRoleCreationRequestWithDefaults instantiates a new SecuritySecurityRolesSecurityRoleCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### SetEmailAddressNil

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetEmailAddressNil(b bool)`

 SetEmailAddressNil sets the value for EmailAddress to be an explicit nil

### UnsetEmailAddress
`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) UnsetEmailAddress()`

UnsetEmailAddress ensures that no value is present for EmailAddress, not even an explicit nil
### GetPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetPermissionSetId() string`

GetPermissionSetId returns the PermissionSetId field if non-nil, zero value otherwise.

### GetPermissionSetIdOk

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetPermissionSetIdOk() (*string, bool)`

GetPermissionSetIdOk returns a tuple with the PermissionSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissionSetId

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetPermissionSetId(v string)`

SetPermissionSetId sets PermissionSetId field to given value.


### GetPermissions

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil
### GetClaims

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetClaims() []SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest`

GetClaims returns the Claims field if non-nil, zero value otherwise.

### GetClaimsOk

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) GetClaimsOk() (*[]SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest, bool)`

GetClaimsOk returns a tuple with the Claims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaims

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetClaims(v []SecurityRoleClaimDefinitionsRoleClaimDefinitionRequest)`

SetClaims sets Claims field to given value.

### HasClaims

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) HasClaims() bool`

HasClaims returns a boolean if a field has been set.

### SetClaimsNil

`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) SetClaimsNil(b bool)`

 SetClaimsNil sets the value for Claims to be an explicit nil

### UnsetClaims
`func (o *SecuritySecurityRolesSecurityRoleCreationRequest) UnsetClaims()`

UnsetClaims ensures that no value is present for Claims, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


