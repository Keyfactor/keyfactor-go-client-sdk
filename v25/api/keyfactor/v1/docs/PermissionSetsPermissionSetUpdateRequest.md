# PermissionSetsPermissionSetUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Id of the permission set. | 
**Permissions** | **[]string** | The permissions within the set. | 

## Methods

### NewPermissionSetsPermissionSetUpdateRequest

`func NewPermissionSetsPermissionSetUpdateRequest(id string, permissions []string, ) *PermissionSetsPermissionSetUpdateRequest`

NewPermissionSetsPermissionSetUpdateRequest instantiates a new PermissionSetsPermissionSetUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSetsPermissionSetUpdateRequestWithDefaults

`func NewPermissionSetsPermissionSetUpdateRequestWithDefaults() *PermissionSetsPermissionSetUpdateRequest`

NewPermissionSetsPermissionSetUpdateRequestWithDefaults instantiates a new PermissionSetsPermissionSetUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PermissionSetsPermissionSetUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PermissionSetsPermissionSetUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PermissionSetsPermissionSetUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetPermissions

`func (o *PermissionSetsPermissionSetUpdateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionSetsPermissionSetUpdateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionSetsPermissionSetUpdateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


