# ModelsEnrollmentCSREnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Template** | Pointer to **NullableString** |  | [optional] 
**SaNs** | Pointer to **map[string][]string** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**MetadataInput** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalEnrollmentFieldsInput** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]string** |  | [optional] 
**Csr** | **string** |  | 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsEnrollmentCSREnrollmentRequest

`func NewModelsEnrollmentCSREnrollmentRequest(csr string, ) *ModelsEnrollmentCSREnrollmentRequest`

NewModelsEnrollmentCSREnrollmentRequest instantiates a new ModelsEnrollmentCSREnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentCSREnrollmentRequestWithDefaults

`func NewModelsEnrollmentCSREnrollmentRequestWithDefaults() *ModelsEnrollmentCSREnrollmentRequest`

NewModelsEnrollmentCSREnrollmentRequestWithDefaults instantiates a new ModelsEnrollmentCSREnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSaNs

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetSaNs() map[string][]string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetSaNsOk() (*map[string][]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetSaNs(v map[string][]string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetCertificateAuthority

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetIncludeChain

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetMetadataInput

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetMetadataInput() map[string]interface{}`

GetMetadataInput returns the MetadataInput field if non-nil, zero value otherwise.

### GetMetadataInputOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetMetadataInputOk() (*map[string]interface{}, bool)`

GetMetadataInputOk returns a tuple with the MetadataInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInput

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetMetadataInput(v map[string]interface{})`

SetMetadataInput sets MetadataInput field to given value.

### HasMetadataInput

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasMetadataInput() bool`

HasMetadataInput returns a boolean if a field has been set.

### SetMetadataInputNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetMetadataInputNil(b bool)`

 SetMetadataInputNil sets the value for MetadataInput to be an explicit nil

### UnsetMetadataInput
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetMetadataInput()`

UnsetMetadataInput ensures that no value is present for MetadataInput, not even an explicit nil
### GetAdditionalEnrollmentFieldsInput

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsInput() map[string]interface{}`

GetAdditionalEnrollmentFieldsInput returns the AdditionalEnrollmentFieldsInput field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsInputOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsInputOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsInputOk returns a tuple with the AdditionalEnrollmentFieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFieldsInput

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsInput(v map[string]interface{})`

SetAdditionalEnrollmentFieldsInput sets AdditionalEnrollmentFieldsInput field to given value.

### HasAdditionalEnrollmentFieldsInput

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasAdditionalEnrollmentFieldsInput() bool`

HasAdditionalEnrollmentFieldsInput returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsInputNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsInputNil(b bool)`

 SetAdditionalEnrollmentFieldsInputNil sets the value for AdditionalEnrollmentFieldsInput to be an explicit nil

### UnsetAdditionalEnrollmentFieldsInput
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetAdditionalEnrollmentFieldsInput()`

UnsetAdditionalEnrollmentFieldsInput ensures that no value is present for AdditionalEnrollmentFieldsInput, not even an explicit nil
### GetTimestamp

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFields() map[string]string`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]string, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]string)`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetCsr

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.


### GetPrivateKey

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *ModelsEnrollmentCSREnrollmentRequest) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


