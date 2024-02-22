# ModelsSSHKeysKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**CreationDate** | Pointer to **time.Time** |  | [optional] 
**StaleDate** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **[]string** |  | [optional] 
**LogonCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsSSHKeysKeyResponse

`func NewModelsSSHKeysKeyResponse() *ModelsSSHKeysKeyResponse`

NewModelsSSHKeysKeyResponse instantiates a new ModelsSSHKeysKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHKeysKeyResponseWithDefaults

`func NewModelsSSHKeysKeyResponseWithDefaults() *ModelsSSHKeysKeyResponse`

NewModelsSSHKeysKeyResponseWithDefaults instantiates a new ModelsSSHKeysKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHKeysKeyResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHKeysKeyResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHKeysKeyResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHKeysKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *ModelsSSHKeysKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ModelsSSHKeysKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ModelsSSHKeysKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ModelsSSHKeysKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetPublicKey

`func (o *ModelsSSHKeysKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ModelsSSHKeysKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ModelsSSHKeysKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ModelsSSHKeysKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *ModelsSSHKeysKeyResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ModelsSSHKeysKeyResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ModelsSSHKeysKeyResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ModelsSSHKeysKeyResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetKeyType

`func (o *ModelsSSHKeysKeyResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsSSHKeysKeyResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsSSHKeysKeyResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ModelsSSHKeysKeyResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### GetKeyLength

`func (o *ModelsSSHKeysKeyResponse) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *ModelsSSHKeysKeyResponse) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *ModelsSSHKeysKeyResponse) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *ModelsSSHKeysKeyResponse) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetCreationDate

`func (o *ModelsSSHKeysKeyResponse) GetCreationDate() time.Time`

GetCreationDate returns the CreationDate field if non-nil, zero value otherwise.

### GetCreationDateOk

`func (o *ModelsSSHKeysKeyResponse) GetCreationDateOk() (*time.Time, bool)`

GetCreationDateOk returns a tuple with the CreationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationDate

`func (o *ModelsSSHKeysKeyResponse) SetCreationDate(v time.Time)`

SetCreationDate sets CreationDate field to given value.

### HasCreationDate

`func (o *ModelsSSHKeysKeyResponse) HasCreationDate() bool`

HasCreationDate returns a boolean if a field has been set.

### GetStaleDate

`func (o *ModelsSSHKeysKeyResponse) GetStaleDate() time.Time`

GetStaleDate returns the StaleDate field if non-nil, zero value otherwise.

### GetStaleDateOk

`func (o *ModelsSSHKeysKeyResponse) GetStaleDateOk() (*time.Time, bool)`

GetStaleDateOk returns a tuple with the StaleDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaleDate

`func (o *ModelsSSHKeysKeyResponse) SetStaleDate(v time.Time)`

SetStaleDate sets StaleDate field to given value.

### HasStaleDate

`func (o *ModelsSSHKeysKeyResponse) HasStaleDate() bool`

HasStaleDate returns a boolean if a field has been set.

### GetEmail

`func (o *ModelsSSHKeysKeyResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsSSHKeysKeyResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsSSHKeysKeyResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsSSHKeysKeyResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetComments

`func (o *ModelsSSHKeysKeyResponse) GetComments() []string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModelsSSHKeysKeyResponse) GetCommentsOk() (*[]string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModelsSSHKeysKeyResponse) SetComments(v []string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ModelsSSHKeysKeyResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetLogonCount

`func (o *ModelsSSHKeysKeyResponse) GetLogonCount() int32`

GetLogonCount returns the LogonCount field if non-nil, zero value otherwise.

### GetLogonCountOk

`func (o *ModelsSSHKeysKeyResponse) GetLogonCountOk() (*int32, bool)`

GetLogonCountOk returns a tuple with the LogonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonCount

`func (o *ModelsSSHKeysKeyResponse) SetLogonCount(v int32)`

SetLogonCount sets LogonCount field to given value.

### HasLogonCount

`func (o *ModelsSSHKeysKeyResponse) HasLogonCount() bool`

HasLogonCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


