# SecuritySecurityRolePermissionsContainerPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContainerId** | Pointer to **int32** |  | [optional] 
**Permission** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewSecuritySecurityRolePermissionsContainerPermissionRequest

`func NewSecuritySecurityRolePermissionsContainerPermissionRequest() *SecuritySecurityRolePermissionsContainerPermissionRequest`

NewSecuritySecurityRolePermissionsContainerPermissionRequest instantiates a new SecuritySecurityRolePermissionsContainerPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySecurityRolePermissionsContainerPermissionRequestWithDefaults

`func NewSecuritySecurityRolePermissionsContainerPermissionRequestWithDefaults() *SecuritySecurityRolePermissionsContainerPermissionRequest`

NewSecuritySecurityRolePermissionsContainerPermissionRequestWithDefaults instantiates a new SecuritySecurityRolePermissionsContainerPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContainerId

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### GetPermission

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) SetPermission(v string)`

SetPermission sets Permission field to given value.

### HasPermission

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) HasPermission() bool`

HasPermission returns a boolean if a field has been set.

### SetPermissionNil

`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) SetPermissionNil(b bool)`

 SetPermissionNil sets the value for Permission to be an explicit nil

### UnsetPermission
`func (o *SecuritySecurityRolePermissionsContainerPermissionRequest) UnsetPermission()`

UnsetPermission ensures that no value is present for Permission, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


