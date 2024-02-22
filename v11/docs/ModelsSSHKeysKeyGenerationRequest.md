# ModelsSSHKeysKeyGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | **string** |  | 
**KeyTypeEnum** | Pointer to **int32** |  | [optional] 
**PrivateKeyFormat** | **string** |  | 
**PrivateKeyFormatEnum** | Pointer to **int32** |  | [optional] 
**KeyLength** | **int32** |  | 
**Email** | **string** |  | 
**Password** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsSSHKeysKeyGenerationRequest

`func NewModelsSSHKeysKeyGenerationRequest(keyType string, privateKeyFormat string, keyLength int32, email string, password string, ) *ModelsSSHKeysKeyGenerationRequest`

NewModelsSSHKeysKeyGenerationRequest instantiates a new ModelsSSHKeysKeyGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHKeysKeyGenerationRequestWithDefaults

`func NewModelsSSHKeysKeyGenerationRequestWithDefaults() *ModelsSSHKeysKeyGenerationRequest`

NewModelsSSHKeysKeyGenerationRequestWithDefaults instantiates a new ModelsSSHKeysKeyGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *ModelsSSHKeysKeyGenerationRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsSSHKeysKeyGenerationRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyTypeEnum

`func (o *ModelsSSHKeysKeyGenerationRequest) GetKeyTypeEnum() int32`

GetKeyTypeEnum returns the KeyTypeEnum field if non-nil, zero value otherwise.

### GetKeyTypeEnumOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetKeyTypeEnumOk() (*int32, bool)`

GetKeyTypeEnumOk returns a tuple with the KeyTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeEnum

`func (o *ModelsSSHKeysKeyGenerationRequest) SetKeyTypeEnum(v int32)`

SetKeyTypeEnum sets KeyTypeEnum field to given value.

### HasKeyTypeEnum

`func (o *ModelsSSHKeysKeyGenerationRequest) HasKeyTypeEnum() bool`

HasKeyTypeEnum returns a boolean if a field has been set.

### GetPrivateKeyFormat

`func (o *ModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *ModelsSSHKeysKeyGenerationRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.


### GetPrivateKeyFormatEnum

`func (o *ModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormatEnum() int32`

GetPrivateKeyFormatEnum returns the PrivateKeyFormatEnum field if non-nil, zero value otherwise.

### GetPrivateKeyFormatEnumOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormatEnumOk() (*int32, bool)`

GetPrivateKeyFormatEnumOk returns a tuple with the PrivateKeyFormatEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormatEnum

`func (o *ModelsSSHKeysKeyGenerationRequest) SetPrivateKeyFormatEnum(v int32)`

SetPrivateKeyFormatEnum sets PrivateKeyFormatEnum field to given value.

### HasPrivateKeyFormatEnum

`func (o *ModelsSSHKeysKeyGenerationRequest) HasPrivateKeyFormatEnum() bool`

HasPrivateKeyFormatEnum returns a boolean if a field has been set.

### GetKeyLength

`func (o *ModelsSSHKeysKeyGenerationRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *ModelsSSHKeysKeyGenerationRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.


### GetEmail

`func (o *ModelsSSHKeysKeyGenerationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsSSHKeysKeyGenerationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *ModelsSSHKeysKeyGenerationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsSSHKeysKeyGenerationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetComment

`func (o *ModelsSSHKeysKeyGenerationRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelsSSHKeysKeyGenerationRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelsSSHKeysKeyGenerationRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelsSSHKeysKeyGenerationRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelsSSHKeysKeyGenerationRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelsSSHKeysKeyGenerationRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


