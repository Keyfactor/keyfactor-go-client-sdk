# ModelsCertificateStoreInventory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **NullableString** |  | [optional] 
**Certificates** | Pointer to [**[]ModelsCertificateStoreInventoryCertificates**](ModelsCertificateStoreInventoryCertificates.md) |  | [optional] 
**CertStoreInventoryItemId** | Pointer to **int32** |  | [optional] 
**Parameters** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModelsCertificateStoreInventory

`func NewModelsCertificateStoreInventory() *ModelsCertificateStoreInventory`

NewModelsCertificateStoreInventory instantiates a new ModelsCertificateStoreInventory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreInventoryWithDefaults

`func NewModelsCertificateStoreInventoryWithDefaults() *ModelsCertificateStoreInventory`

NewModelsCertificateStoreInventoryWithDefaults instantiates a new ModelsCertificateStoreInventory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ModelsCertificateStoreInventory) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCertificateStoreInventory) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCertificateStoreInventory) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsCertificateStoreInventory) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelsCertificateStoreInventory) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelsCertificateStoreInventory) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCertificates

`func (o *ModelsCertificateStoreInventory) GetCertificates() []ModelsCertificateStoreInventoryCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ModelsCertificateStoreInventory) GetCertificatesOk() (*[]ModelsCertificateStoreInventoryCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ModelsCertificateStoreInventory) SetCertificates(v []ModelsCertificateStoreInventoryCertificates)`

SetCertificates sets Certificates field to given value.

### HasCertificates

`func (o *ModelsCertificateStoreInventory) HasCertificates() bool`

HasCertificates returns a boolean if a field has been set.

### SetCertificatesNil

`func (o *ModelsCertificateStoreInventory) SetCertificatesNil(b bool)`

 SetCertificatesNil sets the value for Certificates to be an explicit nil

### UnsetCertificates
`func (o *ModelsCertificateStoreInventory) UnsetCertificates()`

UnsetCertificates ensures that no value is present for Certificates, not even an explicit nil
### GetCertStoreInventoryItemId

`func (o *ModelsCertificateStoreInventory) GetCertStoreInventoryItemId() int32`

GetCertStoreInventoryItemId returns the CertStoreInventoryItemId field if non-nil, zero value otherwise.

### GetCertStoreInventoryItemIdOk

`func (o *ModelsCertificateStoreInventory) GetCertStoreInventoryItemIdOk() (*int32, bool)`

GetCertStoreInventoryItemIdOk returns a tuple with the CertStoreInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreInventoryItemId

`func (o *ModelsCertificateStoreInventory) SetCertStoreInventoryItemId(v int32)`

SetCertStoreInventoryItemId sets CertStoreInventoryItemId field to given value.

### HasCertStoreInventoryItemId

`func (o *ModelsCertificateStoreInventory) HasCertStoreInventoryItemId() bool`

HasCertStoreInventoryItemId returns a boolean if a field has been set.

### GetParameters

`func (o *ModelsCertificateStoreInventory) GetParameters() map[string]interface{}`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *ModelsCertificateStoreInventory) GetParametersOk() (*map[string]interface{}, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *ModelsCertificateStoreInventory) SetParameters(v map[string]interface{})`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *ModelsCertificateStoreInventory) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### SetParametersNil

`func (o *ModelsCertificateStoreInventory) SetParametersNil(b bool)`

 SetParametersNil sets the value for Parameters to be an explicit nil

### UnsetParameters
`func (o *ModelsCertificateStoreInventory) UnsetParameters()`

UnsetParameters ensures that no value is present for Parameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


