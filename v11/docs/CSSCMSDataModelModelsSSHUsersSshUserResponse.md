# CSSCMSDataModelModelsSSHUsersSshUserResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Key** | Pointer to [**CSSCMSDataModelModelsSSHKeysKeyResponse**](CSSCMSDataModelModelsSSHKeysKeyResponse.md) |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**LogonIds** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHUsersSshUserResponse

`func NewCSSCMSDataModelModelsSSHUsersSshUserResponse() *CSSCMSDataModelModelsSSHUsersSshUserResponse`

NewCSSCMSDataModelModelsSSHUsersSshUserResponse instantiates a new CSSCMSDataModelModelsSSHUsersSshUserResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHUsersSshUserResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHUsersSshUserResponseWithDefaults() *CSSCMSDataModelModelsSSHUsersSshUserResponse`

NewCSSCMSDataModelModelsSSHUsersSshUserResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHUsersSshUserResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetKey() CSSCMSDataModelModelsSSHKeysKeyResponse`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetKeyOk() (*CSSCMSDataModelModelsSSHKeysKeyResponse, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) SetKey(v CSSCMSDataModelModelsSSHKeysKeyResponse)`

SetKey sets Key field to given value.

### HasKey

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetUsername

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLogonIds

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetLogonIds() []int32`

GetLogonIds returns the LogonIds field if non-nil, zero value otherwise.

### GetLogonIdsOk

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) GetLogonIdsOk() (*[]int32, bool)`

GetLogonIdsOk returns a tuple with the LogonIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonIds

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) SetLogonIds(v []int32)`

SetLogonIds sets LogonIds field to given value.

### HasLogonIds

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) HasLogonIds() bool`

HasLogonIds returns a boolean if a field has been set.

### SetLogonIdsNil

`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) SetLogonIdsNil(b bool)`

 SetLogonIdsNil sets the value for LogonIds to be an explicit nil

### UnsetLogonIds
`func (o *CSSCMSDataModelModelsSSHUsersSshUserResponse) UnsetLogonIds()`

UnsetLogonIds ensures that no value is present for LogonIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


