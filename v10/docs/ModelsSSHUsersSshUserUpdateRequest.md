# ModelsSSHUsersSshUserUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** |  | 
**LogonIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewModelsSSHUsersSshUserUpdateRequest

`func NewModelsSSHUsersSshUserUpdateRequest(id int64, ) *ModelsSSHUsersSshUserUpdateRequest`

NewModelsSSHUsersSshUserUpdateRequest instantiates a new ModelsSSHUsersSshUserUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHUsersSshUserUpdateRequestWithDefaults

`func NewModelsSSHUsersSshUserUpdateRequestWithDefaults() *ModelsSSHUsersSshUserUpdateRequest`

NewModelsSSHUsersSshUserUpdateRequestWithDefaults instantiates a new ModelsSSHUsersSshUserUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHUsersSshUserUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHUsersSshUserUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHUsersSshUserUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.


### GetLogonIds

`func (o *ModelsSSHUsersSshUserUpdateRequest) GetLogonIds() []int64`

GetLogonIds returns the LogonIds field if non-nil, zero value otherwise.

### GetLogonIdsOk

`func (o *ModelsSSHUsersSshUserUpdateRequest) GetLogonIdsOk() (*[]int64, bool)`

GetLogonIdsOk returns a tuple with the LogonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonIds

`func (o *ModelsSSHUsersSshUserUpdateRequest) SetLogonIds(v []int64)`

SetLogonIds sets LogonIds field to given value.

### HasLogonIds

`func (o *ModelsSSHUsersSshUserUpdateRequest) HasLogonIds() bool`

HasLogonIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


