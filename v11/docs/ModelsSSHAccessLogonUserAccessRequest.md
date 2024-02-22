# ModelsSSHAccessLogonUserAccessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LogonName** | Pointer to **NullableString** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelsSSHAccessLogonUserAccessRequest

`func NewModelsSSHAccessLogonUserAccessRequest() *ModelsSSHAccessLogonUserAccessRequest`

NewModelsSSHAccessLogonUserAccessRequest instantiates a new ModelsSSHAccessLogonUserAccessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHAccessLogonUserAccessRequestWithDefaults

`func NewModelsSSHAccessLogonUserAccessRequestWithDefaults() *ModelsSSHAccessLogonUserAccessRequest`

NewModelsSSHAccessLogonUserAccessRequestWithDefaults instantiates a new ModelsSSHAccessLogonUserAccessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLogonName

`func (o *ModelsSSHAccessLogonUserAccessRequest) GetLogonName() string`

GetLogonName returns the LogonName field if non-nil, zero value otherwise.

### GetLogonNameOk

`func (o *ModelsSSHAccessLogonUserAccessRequest) GetLogonNameOk() (*string, bool)`

GetLogonNameOk returns a tuple with the LogonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonName

`func (o *ModelsSSHAccessLogonUserAccessRequest) SetLogonName(v string)`

SetLogonName sets LogonName field to given value.

### HasLogonName

`func (o *ModelsSSHAccessLogonUserAccessRequest) HasLogonName() bool`

HasLogonName returns a boolean if a field has been set.

### SetLogonNameNil

`func (o *ModelsSSHAccessLogonUserAccessRequest) SetLogonNameNil(b bool)`

 SetLogonNameNil sets the value for LogonName to be an explicit nil

### UnsetLogonName
`func (o *ModelsSSHAccessLogonUserAccessRequest) UnsetLogonName()`

UnsetLogonName ensures that no value is present for LogonName, not even an explicit nil
### GetUsers

`func (o *ModelsSSHAccessLogonUserAccessRequest) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *ModelsSSHAccessLogonUserAccessRequest) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *ModelsSSHAccessLogonUserAccessRequest) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *ModelsSSHAccessLogonUserAccessRequest) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### SetUsersNil

`func (o *ModelsSSHAccessLogonUserAccessRequest) SetUsersNil(b bool)`

 SetUsersNil sets the value for Users to be an explicit nil

### UnsetUsers
`func (o *ModelsSSHAccessLogonUserAccessRequest) UnsetUsers()`

UnsetUsers ensures that no value is present for Users, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


