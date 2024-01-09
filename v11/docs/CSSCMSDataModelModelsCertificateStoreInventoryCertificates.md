# CSSCMSDataModelModelsCertificateStoreInventoryCertificates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**IssuedDN** | Pointer to **NullableString** |  | [optional] 
**SerialNumber** | Pointer to **NullableString** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**SigningAlgorithm** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**CertStoreInventoryItemId** | Pointer to **int32** |  | [optional] 
**CertState** | Pointer to **int32** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreInventoryCertificates

`func NewCSSCMSDataModelModelsCertificateStoreInventoryCertificates() *CSSCMSDataModelModelsCertificateStoreInventoryCertificates`

NewCSSCMSDataModelModelsCertificateStoreInventoryCertificates instantiates a new CSSCMSDataModelModelsCertificateStoreInventoryCertificates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreInventoryCertificatesWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreInventoryCertificatesWithDefaults() *CSSCMSDataModelModelsCertificateStoreInventoryCertificates`

NewCSSCMSDataModelModelsCertificateStoreInventoryCertificatesWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreInventoryCertificates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedDN

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetSerialNumber

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetNotBefore

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetIssuerDN

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetCertStoreInventoryItemId

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetCertStoreInventoryItemId() int32`

GetCertStoreInventoryItemId returns the CertStoreInventoryItemId field if non-nil, zero value otherwise.

### GetCertStoreInventoryItemIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetCertStoreInventoryItemIdOk() (*int32, bool)`

GetCertStoreInventoryItemIdOk returns a tuple with the CertStoreInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryItemId

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetCertStoreInventoryItemId(v int32)`

SetCertStoreInventoryItemId sets CertStoreInventoryItemId field to given value.

### HasCertStoreInventoryItemId

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasCertStoreInventoryItemId() bool`

HasCertStoreInventoryItemId returns a boolean if a field has been set.

### GetCertState

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetCertState() int32`

GetCertState returns the CertState field if non-nil, zero value otherwise.

### GetCertStateOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetCertStateOk() (*int32, bool)`

GetCertStateOk returns a tuple with the CertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertState

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetCertState(v int32)`

SetCertState sets CertState field to given value.

### HasCertState

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasCertState() bool`

HasCertState returns a boolean if a field has been set.

### GetMetadata

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsCertificateStoreInventoryCertificates) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


