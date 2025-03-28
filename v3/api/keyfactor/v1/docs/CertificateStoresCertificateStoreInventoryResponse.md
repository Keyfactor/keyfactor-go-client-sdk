# CertificateStoresCertificateStoreInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Certificates** | Pointer to [**[]CertificateStoresCertificateStoreInventoryCertificateResponse**](CertificateStoresCertificateStoreInventoryCertificateResponse.md) |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewCertificateStoresCertificateStoreInventoryResponse

`func NewCertificateStoresCertificateStoreInventoryResponse() *CertificateStoresCertificateStoreInventoryResponse`

NewCertificateStoresCertificateStoreInventoryResponse instantiates a new CertificateStoresCertificateStoreInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresCertificateStoreInventoryResponseWithDefaults

`func NewCertificateStoresCertificateStoreInventoryResponseWithDefaults() *CertificateStoresCertificateStoreInventoryResponse`

NewCertificateStoresCertificateStoreInventoryResponseWithDefaults instantiates a new CertificateStoresCertificateStoreInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CertificateStoresCertificateStoreInventoryResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoresCertificateStoreInventoryResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoresCertificateStoreInventoryResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStoresCertificateStoreInventoryResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateStoresCertificateStoreInventoryResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateStoresCertificateStoreInventoryResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCertificates

`func (o *CertificateStoresCertificateStoreInventoryResponse) GetCertificates() []CertificateStoresCertificateStoreInventoryCertificateResponse`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *CertificateStoresCertificateStoreInventoryResponse) GetCertificatesOk() (*[]CertificateStoresCertificateStoreInventoryCertificateResponse, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *CertificateStoresCertificateStoreInventoryResponse) SetCertificates(v []CertificateStoresCertificateStoreInventoryCertificateResponse)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *CertificateStoresCertificateStoreInventoryResponse) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *CertificateStoresCertificateStoreInventoryResponse) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *CertificateStoresCertificateStoreInventoryResponse) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetParameters

`func (o *CertificateStoresCertificateStoreInventoryResponse) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *CertificateStoresCertificateStoreInventoryResponse) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *CertificateStoresCertificateStoreInventoryResponse) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *CertificateStoresCertificateStoreInventoryResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *CertificateStoresCertificateStoreInventoryResponse) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *CertificateStoresCertificateStoreInventoryResponse) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


