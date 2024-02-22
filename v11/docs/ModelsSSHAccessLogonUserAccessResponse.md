# ModelsSSHAccessLogonUserAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogonId** | Pointer to **NullableInt32** |  | [optional] 
**LogonName** | Pointer to **NullableString** |  | [optional] 
**Users** | Pointer to [**[]ModelsSSHUsersSshUserResponse**](ModelsSSHUsersSshUserResponse.md) |  | [optional] 

## Methods

### NewModelsSSHAccessLogonUserAccessResponse

`func NewModelsSSHAccessLogonUserAccessResponse() *ModelsSSHAccessLogonUserAccessResponse`

NewModelsSSHAccessLogonUserAccessResponse instantiates a new ModelsSSHAccessLogonUserAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHAccessLogonUserAccessResponseWithDefaults

`func NewModelsSSHAccessLogonUserAccessResponseWithDefaults() *ModelsSSHAccessLogonUserAccessResponse`

NewModelsSSHAccessLogonUserAccessResponseWithDefaults instantiates a new ModelsSSHAccessLogonUserAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogonId

`func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonId() int32`

GetLogonId returns the LogonId field if non-nil, zero value otherwise.

### GetLogonIdOk

`func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonIdOk() (*int32, bool)`

GetLogonIdOk returns a tuple with the LogonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonId

`func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonId(v int32)`

SetLogonId sets LogonId field to given value.

### HasLogonId

`func (o *ModelsSSHAccessLogonUserAccessResponse) HasLogonId() bool`

HasLogonId returns a boolean if a field has been set.

### SetLogonIdNil

`func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonIdNil(b bool)`

 SetLogonIdNil sets the value for LogonId to be an explicit nil

### UnsetLogonId
`func (o *ModelsSSHAccessLogonUserAccessResponse) UnsetLogonId()`

UnsetLogonId ensures that no value is present for LogonId, not even an explicit nil
### GetLogonName

`func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonName() string`

GetLogonName returns the LogonName field if non-nil, zero value otherwise.

### GetLogonNameOk

`func (o *ModelsSSHAccessLogonUserAccessResponse) GetLogonNameOk() (*string, bool)`

GetLogonNameOk returns a tuple with the LogonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonName

`func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonName(v string)`

SetLogonName sets LogonName field to given value.

### HasLogonName

`func (o *ModelsSSHAccessLogonUserAccessResponse) HasLogonName() bool`

HasLogonName returns a boolean if a field has been set.

### SetLogonNameNil

`func (o *ModelsSSHAccessLogonUserAccessResponse) SetLogonNameNil(b bool)`

 SetLogonNameNil sets the value for LogonName to be an explicit nil

### UnsetLogonName
`func (o *ModelsSSHAccessLogonUserAccessResponse) UnsetLogonName()`

UnsetLogonName ensures that no value is present for LogonName, not even an explicit nil
### GetUsers

`func (o *ModelsSSHAccessLogonUserAccessResponse) GetUsers() []ModelsSSHUsersSshUserResponse`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModelsSSHAccessLogonUserAccessResponse) GetUsersOk() (*[]ModelsSSHUsersSshUserResponse, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModelsSSHAccessLogonUserAccessResponse) SetUsers(v []ModelsSSHUsersSshUserResponse)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ModelsSSHAccessLogonUserAccessResponse) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *ModelsSSHAccessLogonUserAccessResponse) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ModelsSSHAccessLogonUserAccessResponse) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


