# ModelsEnrollmentCSREnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CSR** | **string** |  | 
**CertificateAuthority** | Pointer to **string** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewModelsEnrollmentCSREnrollmentRequest

`func NewModelsEnrollmentCSREnrollmentRequest(cSR string, ) *ModelsEnrollmentCSREnrollmentRequest`

NewModelsEnrollmentCSREnrollmentRequest instantiates a new ModelsEnrollmentCSREnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentCSREnrollmentRequestWithDefaults

`func NewModelsEnrollmentCSREnrollmentRequestWithDefaults() *ModelsEnrollmentCSREnrollmentRequest`

NewModelsEnrollmentCSREnrollmentRequestWithDefaults instantiates a new ModelsEnrollmentCSREnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCSR

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetCSR() string`

GetCSR returns the CSR field if non-nil, zero value otherwise.

### GetCSROk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetCSROk() (*string, bool)`

GetCSROk returns a tuple with the CSR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSR

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetCSR(v string)`

SetCSR sets CSR field to given value.


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

### GetMetadata

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFields() map[string]map[string]interface{}`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]map[string]interface{})`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

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

### GetSANs

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *ModelsEnrollmentCSREnrollmentRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *ModelsEnrollmentCSREnrollmentRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *ModelsEnrollmentCSREnrollmentRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


