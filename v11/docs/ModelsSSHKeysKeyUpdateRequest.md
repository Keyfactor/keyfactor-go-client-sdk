# ModelsSSHKeysKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Email** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsSSHKeysKeyUpdateRequest

`func NewModelsSSHKeysKeyUpdateRequest(id int32, email string, ) *ModelsSSHKeysKeyUpdateRequest`

NewModelsSSHKeysKeyUpdateRequest instantiates a new ModelsSSHKeysKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHKeysKeyUpdateRequestWithDefaults

`func NewModelsSSHKeysKeyUpdateRequestWithDefaults() *ModelsSSHKeysKeyUpdateRequest`

NewModelsSSHKeysKeyUpdateRequestWithDefaults instantiates a new ModelsSSHKeysKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHKeysKeyUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHKeysKeyUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHKeysKeyUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetEmail

`func (o *ModelsSSHKeysKeyUpdateRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsSSHKeysKeyUpdateRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsSSHKeysKeyUpdateRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetComment

`func (o *ModelsSSHKeysKeyUpdateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelsSSHKeysKeyUpdateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelsSSHKeysKeyUpdateRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelsSSHKeysKeyUpdateRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelsSSHKeysKeyUpdateRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelsSSHKeysKeyUpdateRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


