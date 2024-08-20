# ModelsSSHUsersSshUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Key** | Pointer to [**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**LogonIds** | Pointer to **[]int64** |  | [optional] 

## Methods

### NewModelsSSHUsersSshUserResponse

`func NewModelsSSHUsersSshUserResponse() *ModelsSSHUsersSshUserResponse`

NewModelsSSHUsersSshUserResponse instantiates a new ModelsSSHUsersSshUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHUsersSshUserResponseWithDefaults

`func NewModelsSSHUsersSshUserResponseWithDefaults() *ModelsSSHUsersSshUserResponse`

NewModelsSSHUsersSshUserResponseWithDefaults instantiates a new ModelsSSHUsersSshUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHUsersSshUserResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHUsersSshUserResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHUsersSshUserResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHUsersSshUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ModelsSSHUsersSshUserResponse) GetKey() ModelsSSHKeysKeyResponse`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelsSSHUsersSshUserResponse) GetKeyOk() (*ModelsSSHKeysKeyResponse, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelsSSHUsersSshUserResponse) SetKey(v ModelsSSHKeysKeyResponse)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelsSSHUsersSshUserResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsSSHUsersSshUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHUsersSshUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHUsersSshUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsSSHUsersSshUserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLogonIds

`func (o *ModelsSSHUsersSshUserResponse) GetLogonIds() []int64`

GetLogonIds returns the LogonIds field if non-nil, zero value otherwise.

### GetLogonIdsOk

`func (o *ModelsSSHUsersSshUserResponse) GetLogonIdsOk() (*[]int64, bool)`

GetLogonIdsOk returns a tuple with the LogonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonIds

`func (o *ModelsSSHUsersSshUserResponse) SetLogonIds(v []int64)`

SetLogonIds sets LogonIds field to given value.

### HasLogonIds

`func (o *ModelsSSHUsersSshUserResponse) HasLogonIds() bool`

HasLogonIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


