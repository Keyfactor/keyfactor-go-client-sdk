# CSSCMSDataModelModelsSSHUsersSshUserAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to [**CSSCMSDataModelModelsSSHKeysKeyResponse**](CSSCMSDataModelModelsSSHKeysKeyResponse.md) |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Access** | Pointer to [**[]CSSCMSDataModelModelsSSHLogonsLogonResponse**](CSSCMSDataModelModelsSSHLogonsLogonResponse.md) |  | [optional] 
**IsGroup** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHUsersSshUserAccessResponse

`func NewCSSCMSDataModelModelsSSHUsersSshUserAccessResponse() *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse`

NewCSSCMSDataModelModelsSSHUsersSshUserAccessResponse instantiates a new CSSCMSDataModelModelsSSHUsersSshUserAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHUsersSshUserAccessResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHUsersSshUserAccessResponseWithDefaults() *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse`

NewCSSCMSDataModelModelsSSHUsersSshUserAccessResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHUsersSshUserAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetKey() CSSCMSDataModelModelsSSHKeysKeyResponse`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetKeyOk() (*CSSCMSDataModelModelsSSHKeysKeyResponse, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetKey(v CSSCMSDataModelModelsSSHKeysKeyResponse)`

SetKey sets Key field to given value.

### HasKey

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUsername

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAccess

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetAccess() []CSSCMSDataModelModelsSSHLogonsLogonResponse`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetAccessOk() (*[]CSSCMSDataModelModelsSSHLogonsLogonResponse, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetAccess(v []CSSCMSDataModelModelsSSHLogonsLogonResponse)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### SetAccessNil

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetAccessNil(b bool)`

 SetAccessNil sets the value for Access to be an explicit nil

### UnsetAccess
`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) UnsetAccess()`

UnsetAccess ensures that no value is present for Access, not even an explicit nil
### GetIsGroup

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetIsGroup() bool`

GetIsGroup returns the IsGroup field if non-nil, zero value otherwise.

### GetIsGroupOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) GetIsGroupOk() (*bool, bool)`

GetIsGroupOk returns a tuple with the IsGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGroup

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) SetIsGroup(v bool)`

SetIsGroup sets IsGroup field to given value.

### HasIsGroup

`func (o *CSSCMSDataModelModelsSSHUsersSshUserAccessResponse) HasIsGroup() bool`

HasIsGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


