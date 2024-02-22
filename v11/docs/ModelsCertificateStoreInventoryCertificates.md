# ModelsCertificateStoreInventoryCertificates

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

### NewModelsCertificateStoreInventoryCertificates

`func NewModelsCertificateStoreInventoryCertificates() *ModelsCertificateStoreInventoryCertificates`

NewModelsCertificateStoreInventoryCertificates instantiates a new ModelsCertificateStoreInventoryCertificates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreInventoryCertificatesWithDefaults

`func NewModelsCertificateStoreInventoryCertificatesWithDefaults() *ModelsCertificateStoreInventoryCertificates`

NewModelsCertificateStoreInventoryCertificatesWithDefaults instantiates a new ModelsCertificateStoreInventoryCertificates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateStoreInventoryCertificates) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateStoreInventoryCertificates) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateStoreInventoryCertificates) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedDN

`func (o *ModelsCertificateStoreInventoryCertificates) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *ModelsCertificateStoreInventoryCertificates) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *ModelsCertificateStoreInventoryCertificates) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *ModelsCertificateStoreInventoryCertificates) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *ModelsCertificateStoreInventoryCertificates) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetSerialNumber

`func (o *ModelsCertificateStoreInventoryCertificates) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *ModelsCertificateStoreInventoryCertificates) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *ModelsCertificateStoreInventoryCertificates) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### SetSerialNumberNil

`func (o *ModelsCertificateStoreInventoryCertificates) SetSerialNumberNil(b bool)`

 SetSerialNumberNil sets the value for SerialNumber to be an explicit nil

### UnsetSerialNumber
`func (o *ModelsCertificateStoreInventoryCertificates) UnsetSerialNumber()`

UnsetSerialNumber ensures that no value is present for SerialNumber, not even an explicit nil
### GetNotBefore

`func (o *ModelsCertificateStoreInventoryCertificates) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ModelsCertificateStoreInventoryCertificates) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *ModelsCertificateStoreInventoryCertificates) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetNotAfter

`func (o *ModelsCertificateStoreInventoryCertificates) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *ModelsCertificateStoreInventoryCertificates) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *ModelsCertificateStoreInventoryCertificates) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetSigningAlgorithm

`func (o *ModelsCertificateStoreInventoryCertificates) GetSigningAlgorithm() string`

GetSigningAlgorithm returns the SigningAlgorithm field if non-nil, zero value otherwise.

### GetSigningAlgorithmOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetSigningAlgorithmOk() (*string, bool)`

GetSigningAlgorithmOk returns a tuple with the SigningAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningAlgorithm

`func (o *ModelsCertificateStoreInventoryCertificates) SetSigningAlgorithm(v string)`

SetSigningAlgorithm sets SigningAlgorithm field to given value.

### HasSigningAlgorithm

`func (o *ModelsCertificateStoreInventoryCertificates) HasSigningAlgorithm() bool`

HasSigningAlgorithm returns a boolean if a field has been set.

### SetSigningAlgorithmNil

`func (o *ModelsCertificateStoreInventoryCertificates) SetSigningAlgorithmNil(b bool)`

 SetSigningAlgorithmNil sets the value for SigningAlgorithm to be an explicit nil

### UnsetSigningAlgorithm
`func (o *ModelsCertificateStoreInventoryCertificates) UnsetSigningAlgorithm()`

UnsetSigningAlgorithm ensures that no value is present for SigningAlgorithm, not even an explicit nil
### GetIssuerDN

`func (o *ModelsCertificateStoreInventoryCertificates) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsCertificateStoreInventoryCertificates) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsCertificateStoreInventoryCertificates) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsCertificateStoreInventoryCertificates) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsCertificateStoreInventoryCertificates) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *ModelsCertificateStoreInventoryCertificates) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsCertificateStoreInventoryCertificates) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsCertificateStoreInventoryCertificates) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *ModelsCertificateStoreInventoryCertificates) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *ModelsCertificateStoreInventoryCertificates) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetCertStoreInventoryItemId

`func (o *ModelsCertificateStoreInventoryCertificates) GetCertStoreInventoryItemId() int32`

GetCertStoreInventoryItemId returns the CertStoreInventoryItemId field if non-nil, zero value otherwise.

### GetCertStoreInventoryItemIdOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetCertStoreInventoryItemIdOk() (*int32, bool)`

GetCertStoreInventoryItemIdOk returns a tuple with the CertStoreInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryItemId

`func (o *ModelsCertificateStoreInventoryCertificates) SetCertStoreInventoryItemId(v int32)`

SetCertStoreInventoryItemId sets CertStoreInventoryItemId field to given value.

### HasCertStoreInventoryItemId

`func (o *ModelsCertificateStoreInventoryCertificates) HasCertStoreInventoryItemId() bool`

HasCertStoreInventoryItemId returns a boolean if a field has been set.

### GetCertState

`func (o *ModelsCertificateStoreInventoryCertificates) GetCertState() int32`

GetCertState returns the CertState field if non-nil, zero value otherwise.

### GetCertStateOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetCertStateOk() (*int32, bool)`

GetCertStateOk returns a tuple with the CertState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertState

`func (o *ModelsCertificateStoreInventoryCertificates) SetCertState(v int32)`

SetCertState sets CertState field to given value.

### HasCertState

`func (o *ModelsCertificateStoreInventoryCertificates) HasCertState() bool`

HasCertState returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsCertificateStoreInventoryCertificates) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsCertificateStoreInventoryCertificates) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsCertificateStoreInventoryCertificates) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsCertificateStoreInventoryCertificates) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ModelsCertificateStoreInventoryCertificates) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ModelsCertificateStoreInventoryCertificates) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


