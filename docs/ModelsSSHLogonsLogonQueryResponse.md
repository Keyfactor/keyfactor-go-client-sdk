# ModelsSSHLogonsLogonQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**LastLogon** | Pointer to **time.Time** |  | [optional] 
**ServerId** | Pointer to **int32** |  | [optional] 
**ServerName** | Pointer to **string** |  | [optional] 
**GroupName** | Pointer to **string** |  | [optional] 
**KeyCount** | Pointer to **int32** |  | [optional] 
**ServerUnderManagement** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsSSHLogonsLogonQueryResponse

`func NewModelsSSHLogonsLogonQueryResponse() *ModelsSSHLogonsLogonQueryResponse`

NewModelsSSHLogonsLogonQueryResponse instantiates a new ModelsSSHLogonsLogonQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHLogonsLogonQueryResponseWithDefaults

`func NewModelsSSHLogonsLogonQueryResponseWithDefaults() *ModelsSSHLogonsLogonQueryResponse`

NewModelsSSHLogonsLogonQueryResponseWithDefaults instantiates a new ModelsSSHLogonsLogonQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHLogonsLogonQueryResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHLogonsLogonQueryResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHLogonsLogonQueryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsSSHLogonsLogonQueryResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHLogonsLogonQueryResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsSSHLogonsLogonQueryResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLastLogon

`func (o *ModelsSSHLogonsLogonQueryResponse) GetLastLogon() time.Time`

GetLastLogon returns the LastLogon field if non-nil, zero value otherwise.

### GetLastLogonOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetLastLogonOk() (*time.Time, bool)`

GetLastLogonOk returns a tuple with the LastLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogon

`func (o *ModelsSSHLogonsLogonQueryResponse) SetLastLogon(v time.Time)`

SetLastLogon sets LastLogon field to given value.

### HasLastLogon

`func (o *ModelsSSHLogonsLogonQueryResponse) HasLastLogon() bool`

HasLastLogon returns a boolean if a field has been set.

### GetServerId

`func (o *ModelsSSHLogonsLogonQueryResponse) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *ModelsSSHLogonsLogonQueryResponse) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *ModelsSSHLogonsLogonQueryResponse) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *ModelsSSHLogonsLogonQueryResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *ModelsSSHLogonsLogonQueryResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *ModelsSSHLogonsLogonQueryResponse) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetGroupName

`func (o *ModelsSSHLogonsLogonQueryResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *ModelsSSHLogonsLogonQueryResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *ModelsSSHLogonsLogonQueryResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetKeyCount

`func (o *ModelsSSHLogonsLogonQueryResponse) GetKeyCount() int32`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetKeyCountOk() (*int32, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *ModelsSSHLogonsLogonQueryResponse) SetKeyCount(v int32)`

SetKeyCount sets KeyCount field to given value.

### HasKeyCount

`func (o *ModelsSSHLogonsLogonQueryResponse) HasKeyCount() bool`

HasKeyCount returns a boolean if a field has been set.

### GetServerUnderManagement

`func (o *ModelsSSHLogonsLogonQueryResponse) GetServerUnderManagement() bool`

GetServerUnderManagement returns the ServerUnderManagement field if non-nil, zero value otherwise.

### GetServerUnderManagementOk

`func (o *ModelsSSHLogonsLogonQueryResponse) GetServerUnderManagementOk() (*bool, bool)`

GetServerUnderManagementOk returns a tuple with the ServerUnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnderManagement

`func (o *ModelsSSHLogonsLogonQueryResponse) SetServerUnderManagement(v bool)`

SetServerUnderManagement sets ServerUnderManagement field to given value.

### HasServerUnderManagement

`func (o *ModelsSSHLogonsLogonQueryResponse) HasServerUnderManagement() bool`

HasServerUnderManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


