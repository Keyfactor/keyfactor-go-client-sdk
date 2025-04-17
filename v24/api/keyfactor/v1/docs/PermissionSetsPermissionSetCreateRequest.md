# PermissionSetsPermissionSetCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the permission set. | 
**Permissions** | **[]string** | The permissions within the set. | 

## Methods

### NewPermissionSetsPermissionSetCreateRequest

`func NewPermissionSetsPermissionSetCreateRequest(name string, permissions []string, ) *PermissionSetsPermissionSetCreateRequest`

NewPermissionSetsPermissionSetCreateRequest instantiates a new PermissionSetsPermissionSetCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionSetsPermissionSetCreateRequestWithDefaults

`func NewPermissionSetsPermissionSetCreateRequestWithDefaults() *PermissionSetsPermissionSetCreateRequest`

NewPermissionSetsPermissionSetCreateRequestWithDefaults instantiates a new PermissionSetsPermissionSetCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PermissionSetsPermissionSetCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PermissionSetsPermissionSetCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PermissionSetsPermissionSetCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPermissions

`func (o *PermissionSetsPermissionSetCreateRequest) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionSetsPermissionSetCreateRequest) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionSetsPermissionSetCreateRequest) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


