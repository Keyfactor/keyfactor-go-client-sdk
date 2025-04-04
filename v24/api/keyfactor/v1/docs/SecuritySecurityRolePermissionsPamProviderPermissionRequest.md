# SecuritySecurityRolePermissionsPamProviderPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PamProviderId** | Pointer to **int32** |  | [optional] 
**Permissions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewSecuritySecurityRolePermissionsPamProviderPermissionRequest

`func NewSecuritySecurityRolePermissionsPamProviderPermissionRequest() *SecuritySecurityRolePermissionsPamProviderPermissionRequest`

NewSecuritySecurityRolePermissionsPamProviderPermissionRequest instantiates a new SecuritySecurityRolePermissionsPamProviderPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityRolePermissionsPamProviderPermissionRequestWithDefaults

`func NewSecuritySecurityRolePermissionsPamProviderPermissionRequestWithDefaults() *SecuritySecurityRolePermissionsPamProviderPermissionRequest`

NewSecuritySecurityRolePermissionsPamProviderPermissionRequestWithDefaults instantiates a new SecuritySecurityRolePermissionsPamProviderPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPamProviderId

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) GetPamProviderId() int32`

GetPamProviderId returns the PamProviderId field if non-nil, zero value otherwise.

### GetPamProviderIdOk

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) GetPamProviderIdOk() (*int32, bool)`

GetPamProviderIdOk returns a tuple with the PamProviderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPamProviderId

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) SetPamProviderId(v int32)`

SetPamProviderId sets PamProviderId field to given value.

### HasPamProviderId

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) HasPamProviderId() bool`

HasPamProviderId returns a boolean if a field has been set.

### GetPermissions

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *SecuritySecurityRolePermissionsPamProviderPermissionRequest) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


