# EnrollmentCSREnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CSR** | **string** |  | 
**PrivateKey** | Pointer to **NullableString** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**IncludeSubjectHeader** | Pointer to **bool** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**OwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**OwnerRoleName** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewEnrollmentCSREnrollmentRequest

`func NewEnrollmentCSREnrollmentRequest(cSR string, ) *EnrollmentCSREnrollmentRequest`

NewEnrollmentCSREnrollmentRequest instantiates a new EnrollmentCSREnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentCSREnrollmentRequestWithDefaults

`func NewEnrollmentCSREnrollmentRequestWithDefaults() *EnrollmentCSREnrollmentRequest`

NewEnrollmentCSREnrollmentRequestWithDefaults instantiates a new EnrollmentCSREnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCSR

`func (o *EnrollmentCSREnrollmentRequest) GetCSR() string`

GetCSR returns the CSR field if non-nil, zero value otherwise.

### GetCSROk

`func (o *EnrollmentCSREnrollmentRequest) GetCSROk() (*string, bool)`

GetCSROk returns a tuple with the CSR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSR

`func (o *EnrollmentCSREnrollmentRequest) SetCSR(v string)`

SetCSR sets CSR field to given value.


### GetPrivateKey

`func (o *EnrollmentCSREnrollmentRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *EnrollmentCSREnrollmentRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *EnrollmentCSREnrollmentRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *EnrollmentCSREnrollmentRequest) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *EnrollmentCSREnrollmentRequest) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *EnrollmentCSREnrollmentRequest) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetCertificateAuthority

`func (o *EnrollmentCSREnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *EnrollmentCSREnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *EnrollmentCSREnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *EnrollmentCSREnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *EnrollmentCSREnrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *EnrollmentCSREnrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetIncludeChain

`func (o *EnrollmentCSREnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *EnrollmentCSREnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *EnrollmentCSREnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *EnrollmentCSREnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetIncludeSubjectHeader

`func (o *EnrollmentCSREnrollmentRequest) GetIncludeSubjectHeader() bool`

GetIncludeSubjectHeader returns the IncludeSubjectHeader field if non-nil, zero value otherwise.

### GetIncludeSubjectHeaderOk

`func (o *EnrollmentCSREnrollmentRequest) GetIncludeSubjectHeaderOk() (*bool, bool)`

GetIncludeSubjectHeaderOk returns a tuple with the IncludeSubjectHeader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubjectHeader

`func (o *EnrollmentCSREnrollmentRequest) SetIncludeSubjectHeader(v bool)`

SetIncludeSubjectHeader sets IncludeSubjectHeader field to given value.

### HasIncludeSubjectHeader

`func (o *EnrollmentCSREnrollmentRequest) HasIncludeSubjectHeader() bool`

HasIncludeSubjectHeader returns a boolean if a field has been set.

### GetMetadata

`func (o *EnrollmentCSREnrollmentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *EnrollmentCSREnrollmentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *EnrollmentCSREnrollmentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *EnrollmentCSREnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *EnrollmentCSREnrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *EnrollmentCSREnrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *EnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFields() map[string]interface{}`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *EnrollmentCSREnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *EnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]interface{})`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *EnrollmentCSREnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *EnrollmentCSREnrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *EnrollmentCSREnrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetTimestamp

`func (o *EnrollmentCSREnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *EnrollmentCSREnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *EnrollmentCSREnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *EnrollmentCSREnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetOwnerRoleId

`func (o *EnrollmentCSREnrollmentRequest) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *EnrollmentCSREnrollmentRequest) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *EnrollmentCSREnrollmentRequest) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *EnrollmentCSREnrollmentRequest) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *EnrollmentCSREnrollmentRequest) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *EnrollmentCSREnrollmentRequest) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *EnrollmentCSREnrollmentRequest) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *EnrollmentCSREnrollmentRequest) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *EnrollmentCSREnrollmentRequest) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *EnrollmentCSREnrollmentRequest) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *EnrollmentCSREnrollmentRequest) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *EnrollmentCSREnrollmentRequest) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil
### GetTemplate

`func (o *EnrollmentCSREnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *EnrollmentCSREnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *EnrollmentCSREnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *EnrollmentCSREnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *EnrollmentCSREnrollmentRequest) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *EnrollmentCSREnrollmentRequest) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetSANs

`func (o *EnrollmentCSREnrollmentRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *EnrollmentCSREnrollmentRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *EnrollmentCSREnrollmentRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *EnrollmentCSREnrollmentRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### SetSANsNil

`func (o *EnrollmentCSREnrollmentRequest) SetSANsNil(b bool)`

 SetSANsNil sets the value for SANs to be an explicit nil

### UnsetSANs
`func (o *EnrollmentCSREnrollmentRequest) UnsetSANs()`

UnsetSANs ensures that no value is present for SANs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


