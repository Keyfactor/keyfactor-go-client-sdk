# CertificatesCertificateRecoveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**UseLegacyEncryption** | Pointer to **NullableBool** |  | [optional] 
**MicrosoftTargetCSP** | Pointer to **NullableString** |  | [optional] 
**CertID** | Pointer to **NullableInt32** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**IncludeSubjectHeader** | Pointer to **bool** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**ChainOrder** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificatesCertificateRecoveryRequest

`func NewCertificatesCertificateRecoveryRequest(password string, ) *CertificatesCertificateRecoveryRequest`

NewCertificatesCertificateRecoveryRequest instantiates a new CertificatesCertificateRecoveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateRecoveryRequestWithDefaults

`func NewCertificatesCertificateRecoveryRequestWithDefaults() *CertificatesCertificateRecoveryRequest`

NewCertificatesCertificateRecoveryRequestWithDefaults instantiates a new CertificatesCertificateRecoveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *CertificatesCertificateRecoveryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CertificatesCertificateRecoveryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CertificatesCertificateRecoveryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUseLegacyEncryption

`func (o *CertificatesCertificateRecoveryRequest) GetUseLegacyEncryption() bool`

GetUseLegacyEncryption returns the UseLegacyEncryption field if non-nil, zero value otherwise.

### GetUseLegacyEncryptionOk

`func (o *CertificatesCertificateRecoveryRequest) GetUseLegacyEncryptionOk() (*bool, bool)`

GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyEncryption

`func (o *CertificatesCertificateRecoveryRequest) SetUseLegacyEncryption(v bool)`

SetUseLegacyEncryption sets UseLegacyEncryption field to given value.

### HasUseLegacyEncryption

`func (o *CertificatesCertificateRecoveryRequest) HasUseLegacyEncryption() bool`

HasUseLegacyEncryption returns a boolean if a field has been set.

### SetUseLegacyEncryptionNil

`func (o *CertificatesCertificateRecoveryRequest) SetUseLegacyEncryptionNil(b bool)`

 SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil

### UnsetUseLegacyEncryption
`func (o *CertificatesCertificateRecoveryRequest) UnsetUseLegacyEncryption()`

UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil
### GetMicrosoftTargetCSP

`func (o *CertificatesCertificateRecoveryRequest) GetMicrosoftTargetCSP() string`

GetMicrosoftTargetCSP returns the MicrosoftTargetCSP field if non-nil, zero value otherwise.

### GetMicrosoftTargetCSPOk

`func (o *CertificatesCertificateRecoveryRequest) GetMicrosoftTargetCSPOk() (*string, bool)`

GetMicrosoftTargetCSPOk returns a tuple with the MicrosoftTargetCSP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMicrosoftTargetCSP

`func (o *CertificatesCertificateRecoveryRequest) SetMicrosoftTargetCSP(v string)`

SetMicrosoftTargetCSP sets MicrosoftTargetCSP field to given value.

### HasMicrosoftTargetCSP

`func (o *CertificatesCertificateRecoveryRequest) HasMicrosoftTargetCSP() bool`

HasMicrosoftTargetCSP returns a boolean if a field has been set.

### SetMicrosoftTargetCSPNil

`func (o *CertificatesCertificateRecoveryRequest) SetMicrosoftTargetCSPNil(b bool)`

 SetMicrosoftTargetCSPNil sets the value for MicrosoftTargetCSP to be an explicit nil

### UnsetMicrosoftTargetCSP
`func (o *CertificatesCertificateRecoveryRequest) UnsetMicrosoftTargetCSP()`

UnsetMicrosoftTargetCSP ensures that no value is present for MicrosoftTargetCSP, not even an explicit nil
### GetCertID

`func (o *CertificatesCertificateRecoveryRequest) GetCertID() int32`

GetCertID returns the CertID field if non-nil, zero value otherwise.

### GetCertIDOk

`func (o *CertificatesCertificateRecoveryRequest) GetCertIDOk() (*int32, bool)`

GetCertIDOk returns a tuple with the CertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertID

`func (o *CertificatesCertificateRecoveryRequest) SetCertID(v int32)`

SetCertID sets CertID field to given value.

### HasCertID

`func (o *CertificatesCertificateRecoveryRequest) HasCertID() bool`

HasCertID returns a boolean if a field has been set.

### SetCertIDNil

`func (o *CertificatesCertificateRecoveryRequest) SetCertIDNil(b bool)`

 SetCertIDNil sets the value for CertID to be an explicit nil

### UnsetCertID
`func (o *CertificatesCertificateRecoveryRequest) UnsetCertID()`

UnsetCertID ensures that no value is present for CertID, not even an explicit nil
### GetSerialNumber

`func (o *CertificatesCertificateRecoveryRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificatesCertificateRecoveryRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificatesCertificateRecoveryRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificatesCertificateRecoveryRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CertificatesCertificateRecoveryRequest) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CertificatesCertificateRecoveryRequest) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuerDN

`func (o *CertificatesCertificateRecoveryRequest) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertificatesCertificateRecoveryRequest) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertificatesCertificateRecoveryRequest) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertificatesCertificateRecoveryRequest) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CertificatesCertificateRecoveryRequest) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CertificatesCertificateRecoveryRequest) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *CertificatesCertificateRecoveryRequest) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificatesCertificateRecoveryRequest) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificatesCertificateRecoveryRequest) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CertificatesCertificateRecoveryRequest) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CertificatesCertificateRecoveryRequest) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CertificatesCertificateRecoveryRequest) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetIncludeChain

`func (o *CertificatesCertificateRecoveryRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *CertificatesCertificateRecoveryRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *CertificatesCertificateRecoveryRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *CertificatesCertificateRecoveryRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetIncludeSubjectHeader

`func (o *CertificatesCertificateRecoveryRequest) GetIncludeSubjectHeader() bool`

GetIncludeSubjectHeader returns the IncludeSubjectHeader field if non-nil, zero value otherwise.

### GetIncludeSubjectHeaderOk

`func (o *CertificatesCertificateRecoveryRequest) GetIncludeSubjectHeaderOk() (*bool, bool)`

GetIncludeSubjectHeaderOk returns a tuple with the IncludeSubjectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubjectHeader

`func (o *CertificatesCertificateRecoveryRequest) SetIncludeSubjectHeader(v bool)`

SetIncludeSubjectHeader sets IncludeSubjectHeader field to given value.

### HasIncludeSubjectHeader

`func (o *CertificatesCertificateRecoveryRequest) HasIncludeSubjectHeader() bool`

HasIncludeSubjectHeader returns a boolean if a field has been set.

### GetFriendlyName

`func (o *CertificatesCertificateRecoveryRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *CertificatesCertificateRecoveryRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *CertificatesCertificateRecoveryRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *CertificatesCertificateRecoveryRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *CertificatesCertificateRecoveryRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *CertificatesCertificateRecoveryRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetChainOrder

`func (o *CertificatesCertificateRecoveryRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *CertificatesCertificateRecoveryRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *CertificatesCertificateRecoveryRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *CertificatesCertificateRecoveryRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *CertificatesCertificateRecoveryRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *CertificatesCertificateRecoveryRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


