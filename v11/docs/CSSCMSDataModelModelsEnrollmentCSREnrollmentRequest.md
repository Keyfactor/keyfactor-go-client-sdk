# CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest

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

### NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequest

`func NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequest(csr string, ) *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest`

NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequest instantiates a new CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequestWithDefaults

`func NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequestWithDefaults() *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest`

NewCSSCMSDataModelModelsEnrollmentCSREnrollmentRequestWithDefaults instantiates a new CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplate

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSaNs

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetSaNs() map[string][]string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetSaNsOk() (*map[string][]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetSaNs(v map[string][]string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetCertificateAuthority

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetIncludeChain

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetMetadataInput

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetMetadataInput() map[string]interface{}`

GetMetadataInput returns the MetadataInput field if non-nil, zero value otherwise.

### GetMetadataInputOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetMetadataInputOk() (*map[string]interface{}, bool)`

GetMetadataInputOk returns a tuple with the MetadataInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataInput

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetMetadataInput(v map[string]interface{})`

SetMetadataInput sets MetadataInput field to given value.

### HasMetadataInput

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasMetadataInput() bool`

HasMetadataInput returns a boolean if a field has been set.

### SetMetadataInputNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetMetadataInputNil(b bool)`

 SetMetadataInputNil sets the value for MetadataInput to be an explicit nil

### UnsetMetadataInput
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetMetadataInput()`

UnsetMetadataInput ensures that no value is present for MetadataInput, not even an explicit nil
### GetAdditionalEnrollmentFieldsInput

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsInput() map[string]interface{}`

GetAdditionalEnrollmentFieldsInput returns the AdditionalEnrollmentFieldsInput field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsInputOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsInputOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsInputOk returns a tuple with the AdditionalEnrollmentFieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFieldsInput

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsInput(v map[string]interface{})`

SetAdditionalEnrollmentFieldsInput sets AdditionalEnrollmentFieldsInput field to given value.

### HasAdditionalEnrollmentFieldsInput

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasAdditionalEnrollmentFieldsInput() bool`

HasAdditionalEnrollmentFieldsInput returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsInputNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsInputNil(b bool)`

 SetAdditionalEnrollmentFieldsInputNil sets the value for AdditionalEnrollmentFieldsInput to be an explicit nil

### UnsetAdditionalEnrollmentFieldsInput
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetAdditionalEnrollmentFieldsInput()`

UnsetAdditionalEnrollmentFieldsInput ensures that no value is present for AdditionalEnrollmentFieldsInput, not even an explicit nil
### GetTimestamp

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFields() map[string]string`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]string, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]string)`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetCsr

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetCsr(v string)`

SetCsr sets Csr field to given value.


### GetPrivateKey

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *CSSCMSDataModelModelsEnrollmentCSREnrollmentRequest) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


