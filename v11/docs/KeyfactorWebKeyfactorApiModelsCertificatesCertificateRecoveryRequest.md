# KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertID** | Pointer to **NullableInt32** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**ChainOrder** | Pointer to **NullableString** |  | [optional] 
**Password** | **string** |  | 
**UseLegacyEncryption** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest(password string, ) *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest`

NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest`

NewKeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertID

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetCertID() int32`

GetCertID returns the CertID field if non-nil, zero value otherwise.

### GetCertIDOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetCertIDOk() (*int32, bool)`

GetCertIDOk returns a tuple with the CertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertID

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetCertID(v int32)`

SetCertID sets CertID field to given value.

### HasCertID

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasCertID() bool`

HasCertID returns a boolean if a field has been set.

### SetCertIDNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetCertIDNil(b bool)`

 SetCertIDNil sets the value for CertID to be an explicit nil

### UnsetCertID
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetCertID()`

UnsetCertID ensures that no value is present for CertID, not even an explicit nil
### GetSerialNumber

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuerDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetIncludeChain

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetChainOrder

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil
### GetPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUseLegacyEncryption

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetUseLegacyEncryption() bool`

GetUseLegacyEncryption returns the UseLegacyEncryption field if non-nil, zero value otherwise.

### GetUseLegacyEncryptionOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) GetUseLegacyEncryptionOk() (*bool, bool)`

GetUseLegacyEncryptionOk returns a tuple with the UseLegacyEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseLegacyEncryption

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetUseLegacyEncryption(v bool)`

SetUseLegacyEncryption sets UseLegacyEncryption field to given value.

### HasUseLegacyEncryption

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) HasUseLegacyEncryption() bool`

HasUseLegacyEncryption returns a boolean if a field has been set.

### SetUseLegacyEncryptionNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) SetUseLegacyEncryptionNil(b bool)`

 SetUseLegacyEncryptionNil sets the value for UseLegacyEncryption to be an explicit nil

### UnsetUseLegacyEncryption
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertificateRecoveryRequest) UnsetUseLegacyEncryption()`

UnsetUseLegacyEncryption ensures that no value is present for UseLegacyEncryption, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


