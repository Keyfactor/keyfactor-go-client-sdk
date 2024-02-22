# ModelsSSHServiceAccountsServiceAccountUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyUpdateRequest** | [**ModelsSSHKeysKeyUpdateRequest**](ModelsSSHKeysKeyUpdateRequest.md) |  | 
**Id** | **int32** |  | 

## Methods

### NewModelsSSHServiceAccountsServiceAccountUpdateRequest

`func NewModelsSSHServiceAccountsServiceAccountUpdateRequest(keyUpdateRequest ModelsSSHKeysKeyUpdateRequest, id int32, ) *ModelsSSHServiceAccountsServiceAccountUpdateRequest`

NewModelsSSHServiceAccountsServiceAccountUpdateRequest instantiates a new ModelsSSHServiceAccountsServiceAccountUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHServiceAccountsServiceAccountUpdateRequestWithDefaults

`func NewModelsSSHServiceAccountsServiceAccountUpdateRequestWithDefaults() *ModelsSSHServiceAccountsServiceAccountUpdateRequest`

NewModelsSSHServiceAccountsServiceAccountUpdateRequestWithDefaults instantiates a new ModelsSSHServiceAccountsServiceAccountUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyUpdateRequest

`func (o *ModelsSSHServiceAccountsServiceAccountUpdateRequest) GetKeyUpdateRequest() ModelsSSHKeysKeyUpdateRequest`

GetKeyUpdateRequest returns the KeyUpdateRequest field if non-nil, zero value otherwise.

### GetKeyUpdateRequestOk

`func (o *ModelsSSHServiceAccountsServiceAccountUpdateRequest) GetKeyUpdateRequestOk() (*ModelsSSHKeysKeyUpdateRequest, bool)`

GetKeyUpdateRequestOk returns a tuple with the KeyUpdateRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyUpdateRequest

`func (o *ModelsSSHServiceAccountsServiceAccountUpdateRequest) SetKeyUpdateRequest(v ModelsSSHKeysKeyUpdateRequest)`

SetKeyUpdateRequest sets KeyUpdateRequest field to given value.


### GetId

`func (o *ModelsSSHServiceAccountsServiceAccountUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHServiceAccountsServiceAccountUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHServiceAccountsServiceAccountUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


