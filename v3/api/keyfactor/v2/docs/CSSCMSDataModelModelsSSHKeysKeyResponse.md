# CSSCMSDataModelModelsSSHKeysKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**CreationDate** | Pointer to **NullableTime** |  | [optional] 
**StaleDate** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **[]string** |  | [optional] 
**LogonCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsSSHKeysKeyResponse

`func NewCSSCMSDataModelModelsSSHKeysKeyResponse() *CSSCMSDataModelModelsSSHKeysKeyResponse`

NewCSSCMSDataModelModelsSSHKeysKeyResponse instantiates a new CSSCMSDataModelModelsSSHKeysKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHKeysKeyResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHKeysKeyResponseWithDefaults() *CSSCMSDataModelModelsSSHKeysKeyResponse`

NewCSSCMSDataModelModelsSSHKeysKeyResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHKeysKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetPublicKey

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKey

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetKeyType

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCreationDate

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### SetCreationDateNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetCreationDateNil(b bool)`

 SetCreationDateNil sets the value for CreationDate to be an explicit nil

### UnsetCreationDate
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetCreationDate()`

UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
### GetStaleDate

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetStaleDate() time.Time`

GetStaleDate returns the StaleDate field if non-nil, zero value otherwise.

### GetStaleDateOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetStaleDateOk() (*time.Time, bool)`

GetStaleDateOk returns a tuple with the StaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleDate

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetStaleDate(v time.Time)`

SetStaleDate sets StaleDate field to given value.

### HasStaleDate

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasStaleDate() bool`

HasStaleDate returns a boolean if a field has been set.

### SetStaleDateNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetStaleDateNil(b bool)`

 SetStaleDateNil sets the value for StaleDate to be an explicit nil

### UnsetStaleDate
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetStaleDate()`

UnsetStaleDate ensures that no value is present for StaleDate, not even an explicit nil
### GetEmail

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetComments

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetComments() []string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetCommentsOk() (*[]string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetComments(v []string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetLogonCount

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetLogonCount() int32`

GetLogonCount returns the LogonCount field if non-nil, zero value otherwise.

### GetLogonCountOk

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) GetLogonCountOk() (*int32, bool)`

GetLogonCountOk returns a tuple with the LogonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonCount

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) SetLogonCount(v int32)`

SetLogonCount sets LogonCount field to given value.

### HasLogonCount

`func (o *CSSCMSDataModelModelsSSHKeysKeyResponse) HasLogonCount() bool`

HasLogonCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


