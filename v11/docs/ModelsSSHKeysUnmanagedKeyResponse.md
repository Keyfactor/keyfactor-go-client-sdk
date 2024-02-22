# ModelsSSHKeysUnmanagedKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**KeyType** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**DiscoveredDate** | Pointer to **NullableTime** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Comments** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**LogonCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewModelsSSHKeysUnmanagedKeyResponse

`func NewModelsSSHKeysUnmanagedKeyResponse() *ModelsSSHKeysUnmanagedKeyResponse`

NewModelsSSHKeysUnmanagedKeyResponse instantiates a new ModelsSSHKeysUnmanagedKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsSSHKeysUnmanagedKeyResponseWithDefaults

`func NewModelsSSHKeysUnmanagedKeyResponseWithDefaults() *ModelsSSHKeysUnmanagedKeyResponse`

NewModelsSSHKeysUnmanagedKeyResponseWithDefaults instantiates a new ModelsSSHKeysUnmanagedKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetPublicKey

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKey

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetKeyType

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetDiscoveredDate

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetDiscoveredDate() time.Time`

GetDiscoveredDate returns the DiscoveredDate field if non-nil, zero value otherwise.

### GetDiscoveredDateOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetDiscoveredDateOk() (*time.Time, bool)`

GetDiscoveredDateOk returns a tuple with the DiscoveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredDate

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetDiscoveredDate(v time.Time)`

SetDiscoveredDate sets DiscoveredDate field to given value.

### HasDiscoveredDate

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasDiscoveredDate() bool`

HasDiscoveredDate returns a boolean if a field has been set.

### SetDiscoveredDateNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetDiscoveredDateNil(b bool)`

 SetDiscoveredDateNil sets the value for DiscoveredDate to be an explicit nil

### UnsetDiscoveredDate
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetDiscoveredDate()`

UnsetDiscoveredDate ensures that no value is present for DiscoveredDate, not even an explicit nil
### GetEmail

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetComments

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetComments() []string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetCommentsOk() (*[]string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetComments(v []string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetUsername

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *ModelsSSHKeysUnmanagedKeyResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLogonCount

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetLogonCount() int32`

GetLogonCount returns the LogonCount field if non-nil, zero value otherwise.

### GetLogonCountOk

`func (o *ModelsSSHKeysUnmanagedKeyResponse) GetLogonCountOk() (*int32, bool)`

GetLogonCountOk returns a tuple with the LogonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonCount

`func (o *ModelsSSHKeysUnmanagedKeyResponse) SetLogonCount(v int32)`

SetLogonCount sets LogonCount field to given value.

### HasLogonCount

`func (o *ModelsSSHKeysUnmanagedKeyResponse) HasLogonCount() bool`

HasLogonCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


