# ModelsCertificateRecoveryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** |  | 
**CertID** | Pointer to **int32** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsCertificateRecoveryRequest

`func NewModelsCertificateRecoveryRequest(password string, ) *ModelsCertificateRecoveryRequest`

NewModelsCertificateRecoveryRequest instantiates a new ModelsCertificateRecoveryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateRecoveryRequestWithDefaults

`func NewModelsCertificateRecoveryRequestWithDefaults() *ModelsCertificateRecoveryRequest`

NewModelsCertificateRecoveryRequestWithDefaults instantiates a new ModelsCertificateRecoveryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *ModelsCertificateRecoveryRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateRecoveryRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateRecoveryRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetCertID

`func (o *ModelsCertificateRecoveryRequest) GetCertID() int32`

GetCertID returns the CertID field if non-nil, zero value otherwise.

### GetCertIDOk

`func (o *ModelsCertificateRecoveryRequest) GetCertIDOk() (*int32, bool)`

GetCertIDOk returns a tuple with the CertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertID

`func (o *ModelsCertificateRecoveryRequest) SetCertID(v int32)`

SetCertID sets CertID field to given value.

### HasCertID

`func (o *ModelsCertificateRecoveryRequest) HasCertID() bool`

HasCertID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ModelsCertificateRecoveryRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsCertificateRecoveryRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsCertificateRecoveryRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsCertificateRecoveryRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsCertificateRecoveryRequest) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsCertificateRecoveryRequest) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsCertificateRecoveryRequest) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsCertificateRecoveryRequest) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsCertificateRecoveryRequest) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsCertificateRecoveryRequest) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *ModelsCertificateRecoveryRequest) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsCertificateRecoveryRequest) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsCertificateRecoveryRequest) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsCertificateRecoveryRequest) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetIncludeChain

`func (o *ModelsCertificateRecoveryRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *ModelsCertificateRecoveryRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *ModelsCertificateRecoveryRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *ModelsCertificateRecoveryRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


