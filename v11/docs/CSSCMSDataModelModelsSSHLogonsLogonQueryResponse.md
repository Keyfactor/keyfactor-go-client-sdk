# CSSCMSDataModelModelsSSHLogonsLogonQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**LastLogon** | Pointer to **NullableTime** |  | [optional] 
**ServerId** | Pointer to **int32** |  | [optional] 
**ServerName** | Pointer to **NullableString** |  | [optional] 
**GroupName** | Pointer to **NullableString** |  | [optional] 
**KeyCount** | Pointer to **int32** |  | [optional] 
**ServerUnderManagement** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHLogonsLogonQueryResponse

`func NewCSSCMSDataModelModelsSSHLogonsLogonQueryResponse() *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse`

NewCSSCMSDataModelModelsSSHLogonsLogonQueryResponse instantiates a new CSSCMSDataModelModelsSSHLogonsLogonQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHLogonsLogonQueryResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHLogonsLogonQueryResponseWithDefaults() *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse`

NewCSSCMSDataModelModelsSSHLogonsLogonQueryResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHLogonsLogonQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLastLogon

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetLastLogon() time.Time`

GetLastLogon returns the LastLogon field if non-nil, zero value otherwise.

### GetLastLogonOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetLastLogonOk() (*time.Time, bool)`

GetLastLogonOk returns a tuple with the LastLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogon

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetLastLogon(v time.Time)`

SetLastLogon sets LastLogon field to given value.

### HasLastLogon

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasLastLogon() bool`

HasLastLogon returns a boolean if a field has been set.

### SetLastLogonNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetLastLogonNil(b bool)`

 SetLastLogonNil sets the value for LastLogon to be an explicit nil

### UnsetLastLogon
`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) UnsetLastLogon()`

UnsetLastLogon ensures that no value is present for LastLogon, not even an explicit nil
### GetServerId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### SetServerNameNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetServerNameNil(b bool)`

 SetServerNameNil sets the value for ServerName to be an explicit nil

### UnsetServerName
`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) UnsetServerName()`

UnsetServerName ensures that no value is present for ServerName, not even an explicit nil
### GetGroupName

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### SetGroupNameNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetGroupNameNil(b bool)`

 SetGroupNameNil sets the value for GroupName to be an explicit nil

### UnsetGroupName
`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) UnsetGroupName()`

UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
### GetKeyCount

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetKeyCount() int32`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetKeyCountOk() (*int32, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetKeyCount(v int32)`

SetKeyCount sets KeyCount field to given value.

### HasKeyCount

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasKeyCount() bool`

HasKeyCount returns a boolean if a field has been set.

### GetServerUnderManagement

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetServerUnderManagement() bool`

GetServerUnderManagement returns the ServerUnderManagement field if non-nil, zero value otherwise.

### GetServerUnderManagementOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) GetServerUnderManagementOk() (*bool, bool)`

GetServerUnderManagementOk returns a tuple with the ServerUnderManagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUnderManagement

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) SetServerUnderManagement(v bool)`

SetServerUnderManagement sets ServerUnderManagement field to given value.

### HasServerUnderManagement

`func (o *CSSCMSDataModelModelsSSHLogonsLogonQueryResponse) HasServerUnderManagement() bool`

HasServerUnderManagement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


