# CertificatesCertificateDownloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertID** | Pointer to **NullableInt32** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**IncludeSubjectHeader** | Pointer to **bool** |  | [optional] 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**ChainOrder** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificatesCertificateDownloadRequest

`func NewCertificatesCertificateDownloadRequest() *CertificatesCertificateDownloadRequest`

NewCertificatesCertificateDownloadRequest instantiates a new CertificatesCertificateDownloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateDownloadRequestWithDefaults

`func NewCertificatesCertificateDownloadRequestWithDefaults() *CertificatesCertificateDownloadRequest`

NewCertificatesCertificateDownloadRequestWithDefaults instantiates a new CertificatesCertificateDownloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertID

`func (o *CertificatesCertificateDownloadRequest) GetCertID() int32`

GetCertID returns the CertID field if non-nil, zero value otherwise.

### GetCertIDOk

`func (o *CertificatesCertificateDownloadRequest) GetCertIDOk() (*int32, bool)`

GetCertIDOk returns a tuple with the CertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertID

`func (o *CertificatesCertificateDownloadRequest) SetCertID(v int32)`

SetCertID sets CertID field to given value.

### HasCertID

`func (o *CertificatesCertificateDownloadRequest) HasCertID() bool`

HasCertID returns a boolean if a field has been set.

### SetCertIDNil

`func (o *CertificatesCertificateDownloadRequest) SetCertIDNil(b bool)`

 SetCertIDNil sets the value for CertID to be an explicit nil

### UnsetCertID
`func (o *CertificatesCertificateDownloadRequest) UnsetCertID()`

UnsetCertID ensures that no value is present for CertID, not even an explicit nil
### GetSerialNumber

`func (o *CertificatesCertificateDownloadRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CertificatesCertificateDownloadRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CertificatesCertificateDownloadRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CertificatesCertificateDownloadRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CertificatesCertificateDownloadRequest) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CertificatesCertificateDownloadRequest) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetIssuerDN

`func (o *CertificatesCertificateDownloadRequest) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CertificatesCertificateDownloadRequest) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CertificatesCertificateDownloadRequest) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CertificatesCertificateDownloadRequest) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CertificatesCertificateDownloadRequest) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CertificatesCertificateDownloadRequest) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *CertificatesCertificateDownloadRequest) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CertificatesCertificateDownloadRequest) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CertificatesCertificateDownloadRequest) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CertificatesCertificateDownloadRequest) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CertificatesCertificateDownloadRequest) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CertificatesCertificateDownloadRequest) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetIncludeChain

`func (o *CertificatesCertificateDownloadRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *CertificatesCertificateDownloadRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *CertificatesCertificateDownloadRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *CertificatesCertificateDownloadRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetIncludeSubjectHeader

`func (o *CertificatesCertificateDownloadRequest) GetIncludeSubjectHeader() bool`

GetIncludeSubjectHeader returns the IncludeSubjectHeader field if non-nil, zero value otherwise.

### GetIncludeSubjectHeaderOk

`func (o *CertificatesCertificateDownloadRequest) GetIncludeSubjectHeaderOk() (*bool, bool)`

GetIncludeSubjectHeaderOk returns a tuple with the IncludeSubjectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubjectHeader

`func (o *CertificatesCertificateDownloadRequest) SetIncludeSubjectHeader(v bool)`

SetIncludeSubjectHeader sets IncludeSubjectHeader field to given value.

### HasIncludeSubjectHeader

`func (o *CertificatesCertificateDownloadRequest) HasIncludeSubjectHeader() bool`

HasIncludeSubjectHeader returns a boolean if a field has been set.

### GetFriendlyName

`func (o *CertificatesCertificateDownloadRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *CertificatesCertificateDownloadRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *CertificatesCertificateDownloadRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *CertificatesCertificateDownloadRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *CertificatesCertificateDownloadRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *CertificatesCertificateDownloadRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetChainOrder

`func (o *CertificatesCertificateDownloadRequest) GetChainOrder() string`

GetChainOrder returns the ChainOrder field if non-nil, zero value otherwise.

### GetChainOrderOk

`func (o *CertificatesCertificateDownloadRequest) GetChainOrderOk() (*string, bool)`

GetChainOrderOk returns a tuple with the ChainOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainOrder

`func (o *CertificatesCertificateDownloadRequest) SetChainOrder(v string)`

SetChainOrder sets ChainOrder field to given value.

### HasChainOrder

`func (o *CertificatesCertificateDownloadRequest) HasChainOrder() bool`

HasChainOrder returns a boolean if a field has been set.

### SetChainOrderNil

`func (o *CertificatesCertificateDownloadRequest) SetChainOrderNil(b bool)`

 SetChainOrderNil sets the value for ChainOrder to be an explicit nil

### UnsetChainOrder
`func (o *CertificatesCertificateDownloadRequest) UnsetChainOrder()`

UnsetChainOrder ensures that no value is present for ChainOrder, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


