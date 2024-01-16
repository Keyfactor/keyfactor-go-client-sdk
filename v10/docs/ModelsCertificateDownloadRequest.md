# ModelsCertificateDownloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertID** | Pointer to **int64** |  | [optional] 
**SerialNumber** | Pointer to **string** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsCertificateDownloadRequest

`func NewModelsCertificateDownloadRequest() *ModelsCertificateDownloadRequest`

NewModelsCertificateDownloadRequest instantiates a new ModelsCertificateDownloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateDownloadRequestWithDefaults

`func NewModelsCertificateDownloadRequestWithDefaults() *ModelsCertificateDownloadRequest`

NewModelsCertificateDownloadRequestWithDefaults instantiates a new ModelsCertificateDownloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertID

`func (o *ModelsCertificateDownloadRequest) GetCertID() int64`

GetCertID returns the CertID field if non-nil, zero value otherwise.

### GetCertIDOk

`func (o *ModelsCertificateDownloadRequest) GetCertIDOk() (*int64, bool)`

GetCertIDOk returns a tuple with the CertID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertID

`func (o *ModelsCertificateDownloadRequest) SetCertID(v int64)`

SetCertID sets CertID field to given value.

### HasCertID

`func (o *ModelsCertificateDownloadRequest) HasCertID() bool`

HasCertID returns a boolean if a field has been set.

### GetSerialNumber

`func (o *ModelsCertificateDownloadRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsCertificateDownloadRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsCertificateDownloadRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsCertificateDownloadRequest) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetIssuerDN

`func (o *ModelsCertificateDownloadRequest) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsCertificateDownloadRequest) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsCertificateDownloadRequest) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsCertificateDownloadRequest) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsCertificateDownloadRequest) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsCertificateDownloadRequest) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *ModelsCertificateDownloadRequest) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsCertificateDownloadRequest) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsCertificateDownloadRequest) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsCertificateDownloadRequest) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetIncludeChain

`func (o *ModelsCertificateDownloadRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *ModelsCertificateDownloadRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *ModelsCertificateDownloadRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *ModelsCertificateDownloadRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


