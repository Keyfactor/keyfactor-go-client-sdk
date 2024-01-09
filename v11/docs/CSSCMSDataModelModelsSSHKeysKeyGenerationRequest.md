# CSSCMSDataModelModelsSSHKeysKeyGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyType** | **string** |  | 
**KeyTypeEnum** | Pointer to [**CSSCMSDataModelEnumsSshKeyAlgorithm**](CSSCMSDataModelEnumsSshKeyAlgorithm.md) |  | [optional] 
**PrivateKeyFormat** | **string** |  | 
**PrivateKeyFormatEnum** | Pointer to [**CSSCMSDataModelEnumsSshPrivateKeyFormat**](CSSCMSDataModelEnumsSshPrivateKeyFormat.md) |  | [optional] 
**KeyLength** | **int32** |  | 
**Email** | **string** |  | 
**Password** | **string** |  | 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest

`func NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest(keyType string, privateKeyFormat string, keyLength int32, email string, password string, ) *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest`

NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequest instantiates a new CSSCMSDataModelModelsSSHKeysKeyGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequestWithDefaults

`func NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequestWithDefaults() *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest`

NewCSSCMSDataModelModelsSSHKeysKeyGenerationRequestWithDefaults instantiates a new CSSCMSDataModelModelsSSHKeysKeyGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyType

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.


### GetKeyTypeEnum

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetKeyTypeEnum() CSSCMSDataModelEnumsSshKeyAlgorithm`

GetKeyTypeEnum returns the KeyTypeEnum field if non-nil, zero value otherwise.

### GetKeyTypeEnumOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetKeyTypeEnumOk() (*CSSCMSDataModelEnumsSshKeyAlgorithm, bool)`

GetKeyTypeEnumOk returns a tuple with the KeyTypeEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyTypeEnum

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetKeyTypeEnum(v CSSCMSDataModelEnumsSshKeyAlgorithm)`

SetKeyTypeEnum sets KeyTypeEnum field to given value.

### HasKeyTypeEnum

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) HasKeyTypeEnum() bool`

HasKeyTypeEnum returns a boolean if a field has been set.

### GetPrivateKeyFormat

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormat() string`

GetPrivateKeyFormat returns the PrivateKeyFormat field if non-nil, zero value otherwise.

### GetPrivateKeyFormatOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormatOk() (*string, bool)`

GetPrivateKeyFormatOk returns a tuple with the PrivateKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormat

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetPrivateKeyFormat(v string)`

SetPrivateKeyFormat sets PrivateKeyFormat field to given value.


### GetPrivateKeyFormatEnum

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormatEnum() CSSCMSDataModelEnumsSshPrivateKeyFormat`

GetPrivateKeyFormatEnum returns the PrivateKeyFormatEnum field if non-nil, zero value otherwise.

### GetPrivateKeyFormatEnumOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetPrivateKeyFormatEnumOk() (*CSSCMSDataModelEnumsSshPrivateKeyFormat, bool)`

GetPrivateKeyFormatEnumOk returns a tuple with the PrivateKeyFormatEnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyFormatEnum

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetPrivateKeyFormatEnum(v CSSCMSDataModelEnumsSshPrivateKeyFormat)`

SetPrivateKeyFormatEnum sets PrivateKeyFormatEnum field to given value.

### HasPrivateKeyFormatEnum

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) HasPrivateKeyFormatEnum() bool`

HasPrivateKeyFormatEnum returns a boolean if a field has been set.

### GetKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.


### GetEmail

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetComment

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CSSCMSDataModelModelsSSHKeysKeyGenerationRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


