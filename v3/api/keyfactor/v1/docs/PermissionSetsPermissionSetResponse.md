# PermissionSetsPermissionSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The Id of the permission set. | [optional] 
**Name** | Pointer to **NullableString** | The name of the permission set. | [optional] 
**Permissions** | Pointer to **[]string** | The permissions within the set. | [optional] 

## Methods

### NewPermissionSetsPermissionSetResponse

`func NewPermissionSetsPermissionSetResponse() *PermissionSetsPermissionSetResponse`

NewPermissionSetsPermissionSetResponse instantiates a new PermissionSetsPermissionSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSetsPermissionSetResponseWithDefaults

`func NewPermissionSetsPermissionSetResponseWithDefaults() *PermissionSetsPermissionSetResponse`

NewPermissionSetsPermissionSetResponseWithDefaults instantiates a new PermissionSetsPermissionSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionSetsPermissionSetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionSetsPermissionSetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionSetsPermissionSetResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PermissionSetsPermissionSetResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PermissionSetsPermissionSetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionSetsPermissionSetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionSetsPermissionSetResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PermissionSetsPermissionSetResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *PermissionSetsPermissionSetResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PermissionSetsPermissionSetResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPermissions

`func (o *PermissionSetsPermissionSetResponse) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionSetsPermissionSetResponse) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionSetsPermissionSetResponse) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PermissionSetsPermissionSetResponse) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### SetPermissionsNil

`func (o *PermissionSetsPermissionSetResponse) SetPermissionsNil(b bool)`

 SetPermissionsNil sets the value for Permissions to be an explicit nil

### UnsetPermissions
`func (o *PermissionSetsPermissionSetResponse) UnsetPermissions()`

UnsetPermissions ensures that no value is present for Permissions, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


