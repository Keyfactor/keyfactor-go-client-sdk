# CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse

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

### NewCSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse

`func NewCSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse() *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse`

NewCSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse instantiates a new CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsSSHKeysUnmanagedKeyResponseWithDefaults

`func NewCSSCMSDataModelModelsSSHKeysUnmanagedKeyResponseWithDefaults() *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse`

NewCSSCMSDataModelModelsSSHKeysUnmanagedKeyResponseWithDefaults instantiates a new CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFingerprint

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetPublicKey

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetPrivateKey

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetKeyType

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetKeyType() string`

GetKeyType returns the KeyType field if non-nil, zero value otherwise.

### GetKeyTypeOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetKeyTypeOk() (*string, bool)`

GetKeyTypeOk returns a tuple with the KeyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyType

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetKeyType(v string)`

SetKeyType sets KeyType field to given value.

### HasKeyType

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasKeyType() bool`

HasKeyType returns a boolean if a field has been set.

### SetKeyTypeNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetKeyTypeNil(b bool)`

 SetKeyTypeNil sets the value for KeyType to be an explicit nil

### UnsetKeyType
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetKeyType()`

UnsetKeyType ensures that no value is present for KeyType, not even an explicit nil
### GetKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetDiscoveredDate

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetDiscoveredDate() time.Time`

GetDiscoveredDate returns the DiscoveredDate field if non-nil, zero value otherwise.

### GetDiscoveredDateOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetDiscoveredDateOk() (*time.Time, bool)`

GetDiscoveredDateOk returns a tuple with the DiscoveredDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveredDate

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetDiscoveredDate(v time.Time)`

SetDiscoveredDate sets DiscoveredDate field to given value.

### HasDiscoveredDate

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasDiscoveredDate() bool`

HasDiscoveredDate returns a boolean if a field has been set.

### SetDiscoveredDateNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetDiscoveredDateNil(b bool)`

 SetDiscoveredDateNil sets the value for DiscoveredDate to be an explicit nil

### UnsetDiscoveredDate
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetDiscoveredDate()`

UnsetDiscoveredDate ensures that no value is present for DiscoveredDate, not even an explicit nil
### GetEmail

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetComments

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetComments() []string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetCommentsOk() (*[]string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetComments(v []string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasComments() bool`

HasComments returns a boolean if a field has been set.

### SetCommentsNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetCommentsNil(b bool)`

 SetCommentsNil sets the value for Comments to be an explicit nil

### UnsetComments
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetComments()`

UnsetComments ensures that no value is present for Comments, not even an explicit nil
### GetUsername

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetLogonCount

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetLogonCount() int32`

GetLogonCount returns the LogonCount field if non-nil, zero value otherwise.

### GetLogonCountOk

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) GetLogonCountOk() (*int32, bool)`

GetLogonCountOk returns a tuple with the LogonCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogonCount

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) SetLogonCount(v int32)`

SetLogonCount sets LogonCount field to given value.

### HasLogonCount

`func (o *CSSCMSDataModelModelsSSHKeysUnmanagedKeyResponse) HasLogonCount() bool`

HasLogonCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


