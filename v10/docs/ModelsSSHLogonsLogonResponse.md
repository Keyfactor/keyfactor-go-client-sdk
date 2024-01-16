# ModelsSSHLogonsLogonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**LastLogon** | Pointer to **time.Time** |  | [optional] 
**Server** | Pointer to [**ModelsSSHServersServerResponse**](ModelsSSHServersServerResponse.md) |  | [optional] 
**KeyCount** | Pointer to **int64** |  | [optional] 
**Access** | Pointer to [**[]ModelsSSHUsersSshUserResponse**](ModelsSSHUsersSshUserResponse.md) |  | [optional] 

## Methods

### NewModelsSSHLogonsLogonResponse

`func NewModelsSSHLogonsLogonResponse() *ModelsSSHLogonsLogonResponse`

NewModelsSSHLogonsLogonResponse instantiates a new ModelsSSHLogonsLogonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHLogonsLogonResponseWithDefaults

`func NewModelsSSHLogonsLogonResponseWithDefaults() *ModelsSSHLogonsLogonResponse`

NewModelsSSHLogonsLogonResponseWithDefaults instantiates a new ModelsSSHLogonsLogonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHLogonsLogonResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHLogonsLogonResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHLogonsLogonResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHLogonsLogonResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsSSHLogonsLogonResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHLogonsLogonResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHLogonsLogonResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsSSHLogonsLogonResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLastLogon

`func (o *ModelsSSHLogonsLogonResponse) GetLastLogon() time.Time`

GetLastLogon returns the LastLogon field if non-nil, zero value otherwise.

### GetLastLogonOk

`func (o *ModelsSSHLogonsLogonResponse) GetLastLogonOk() (*time.Time, bool)`

GetLastLogonOk returns a tuple with the LastLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogon

`func (o *ModelsSSHLogonsLogonResponse) SetLastLogon(v time.Time)`

SetLastLogon sets LastLogon field to given value.

### HasLastLogon

`func (o *ModelsSSHLogonsLogonResponse) HasLastLogon() bool`

HasLastLogon returns a boolean if a field has been set.

### GetServer

`func (o *ModelsSSHLogonsLogonResponse) GetServer() ModelsSSHServersServerResponse`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *ModelsSSHLogonsLogonResponse) GetServerOk() (*ModelsSSHServersServerResponse, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *ModelsSSHLogonsLogonResponse) SetServer(v ModelsSSHServersServerResponse)`

SetServer sets Server field to given value.

### HasServer

`func (o *ModelsSSHLogonsLogonResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetKeyCount

`func (o *ModelsSSHLogonsLogonResponse) GetKeyCount() int64`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *ModelsSSHLogonsLogonResponse) GetKeyCountOk() (*int64, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *ModelsSSHLogonsLogonResponse) SetKeyCount(v int64)`

SetKeyCount sets KeyCount field to given value.

### HasKeyCount

`func (o *ModelsSSHLogonsLogonResponse) HasKeyCount() bool`

HasKeyCount returns a boolean if a field has been set.

### GetAccess

`func (o *ModelsSSHLogonsLogonResponse) GetAccess() []ModelsSSHUsersSshUserResponse`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *ModelsSSHLogonsLogonResponse) GetAccessOk() (*[]ModelsSSHUsersSshUserResponse, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *ModelsSSHLogonsLogonResponse) SetAccess(v []ModelsSSHUsersSshUserResponse)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *ModelsSSHLogonsLogonResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


