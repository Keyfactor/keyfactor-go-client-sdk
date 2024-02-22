# ModelsSSHUsersSshUserAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Key** | Pointer to [**ModelsSSHKeysKeyResponse**](ModelsSSHKeysKeyResponse.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Access** | Pointer to [**[]ModelsSSHLogonsLogonResponse**](ModelsSSHLogonsLogonResponse.md) |  | [optional] 
**IsGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsSSHUsersSshUserAccessResponse

`func NewModelsSSHUsersSshUserAccessResponse() *ModelsSSHUsersSshUserAccessResponse`

NewModelsSSHUsersSshUserAccessResponse instantiates a new ModelsSSHUsersSshUserAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHUsersSshUserAccessResponseWithDefaults

`func NewModelsSSHUsersSshUserAccessResponseWithDefaults() *ModelsSSHUsersSshUserAccessResponse`

NewModelsSSHUsersSshUserAccessResponseWithDefaults instantiates a new ModelsSSHUsersSshUserAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHUsersSshUserAccessResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHUsersSshUserAccessResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHUsersSshUserAccessResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHUsersSshUserAccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ModelsSSHUsersSshUserAccessResponse) GetKey() ModelsSSHKeysKeyResponse`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ModelsSSHUsersSshUserAccessResponse) GetKeyOk() (*ModelsSSHKeysKeyResponse, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ModelsSSHUsersSshUserAccessResponse) SetKey(v ModelsSSHKeysKeyResponse)`

SetKey sets Key field to given value.

### HasKey

`func (o *ModelsSSHUsersSshUserAccessResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsSSHUsersSshUserAccessResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHUsersSshUserAccessResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHUsersSshUserAccessResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsSSHUsersSshUserAccessResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAccess

`func (o *ModelsSSHUsersSshUserAccessResponse) GetAccess() []ModelsSSHLogonsLogonResponse`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ModelsSSHUsersSshUserAccessResponse) GetAccessOk() (*[]ModelsSSHLogonsLogonResponse, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ModelsSSHUsersSshUserAccessResponse) SetAccess(v []ModelsSSHLogonsLogonResponse)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ModelsSSHUsersSshUserAccessResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetIsGroup

`func (o *ModelsSSHUsersSshUserAccessResponse) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *ModelsSSHUsersSshUserAccessResponse) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *ModelsSSHUsersSshUserAccessResponse) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *ModelsSSHUsersSshUserAccessResponse) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


