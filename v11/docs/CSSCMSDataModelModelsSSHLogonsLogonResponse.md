# CSSCMSDataModelModelsSSHLogonsLogonResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**LastLogon** | Pointer to **NullableTime** |  | [optional] 
**Server** | Pointer to [**CSSCMSDataModelModelsSSHServersServerResponse**](CSSCMSDataModelModelsSSHServersServerResponse.md) |  | [optional] 
**KeyCount** | Pointer to **int32** |  | [optional] 
**Access** | Pointer to [**[]CSSCMSDataModelModelsSSHUsersSshUserResponse**](CSSCMSDataModelModelsSSHUsersSshUserResponse.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHLogonsLogonResponse

`func NewCSSCMSDataModelModelsSSHLogonsLogonResponse() *CSSCMSDataModelModelsSSHLogonsLogonResponse`

NewCSSCMSDataModelModelsSSHLogonsLogonResponse instantiates a new CSSCMSDataModelModelsSSHLogonsLogonResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHLogonsLogonResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHLogonsLogonResponseWithDefaults() *CSSCMSDataModelModelsSSHLogonsLogonResponse`

NewCSSCMSDataModelModelsSSHLogonsLogonResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHLogonsLogonResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLastLogon

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetLastLogon() time.Time`

GetLastLogon returns the LastLogon field if non-nil, zero value otherwise.

### GetLastLogonOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetLastLogonOk() (*time.Time, bool)`

GetLastLogonOk returns a tuple with the LastLogon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLogon

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetLastLogon(v time.Time)`

SetLastLogon sets LastLogon field to given value.

### HasLastLogon

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) HasLastLogon() bool`

HasLastLogon returns a boolean if a field has been set.

### SetLastLogonNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetLastLogonNil(b bool)`

 SetLastLogonNil sets the value for LastLogon to be an explicit nil

### UnsetLastLogon
`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) UnsetLastLogon()`

UnsetLastLogon ensures that no value is present for LastLogon, not even an explicit nil
### GetServer

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetServer() CSSCMSDataModelModelsSSHServersServerResponse`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetServerOk() (*CSSCMSDataModelModelsSSHServersServerResponse, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetServer(v CSSCMSDataModelModelsSSHServersServerResponse)`

SetServer sets Server field to given value.

### HasServer

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetKeyCount

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetKeyCount() int32`

GetKeyCount returns the KeyCount field if non-nil, zero value otherwise.

### GetKeyCountOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetKeyCountOk() (*int32, bool)`

GetKeyCountOk returns a tuple with the KeyCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCount

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetKeyCount(v int32)`

SetKeyCount sets KeyCount field to given value.

### HasKeyCount

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) HasKeyCount() bool`

HasKeyCount returns a boolean if a field has been set.

### GetAccess

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetAccess() []CSSCMSDataModelModelsSSHUsersSshUserResponse`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) GetAccessOk() (*[]CSSCMSDataModelModelsSSHUsersSshUserResponse, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetAccess(v []CSSCMSDataModelModelsSSHUsersSshUserResponse)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *CSSCMSDataModelModelsSSHLogonsLogonResponse) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


