# ModelsSSHKeysUnmanagedKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** |  | [optional] 
**KeyType** | Pointer to **string** |  | [optional] 
**KeyLength** | Pointer to **int32** |  | [optional] 
**DiscoveredDate** | Pointer to **time.Time** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**Comments** | Pointer to **[]string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
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


