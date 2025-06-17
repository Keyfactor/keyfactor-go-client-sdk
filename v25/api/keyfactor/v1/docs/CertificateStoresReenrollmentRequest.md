# CertificateStoresReenrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeystoreId** | Pointer to **string** |  | [optional] 
**SubjectName** | Pointer to **NullableString** |  | [optional] 
**AgentGuid** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**JobProperties** | Pointer to **map[string]interface{}** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**CertificateTemplate** | Pointer to **NullableString** |  | [optional] 
**EnrollmentPatternId** | Pointer to **NullableInt32** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]string** |  | [optional] 
**OwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**OwnerRoleName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCertificateStoresReenrollmentRequest

`func NewCertificateStoresReenrollmentRequest() *CertificateStoresReenrollmentRequest`

NewCertificateStoresReenrollmentRequest instantiates a new CertificateStoresReenrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresReenrollmentRequestWithDefaults

`func NewCertificateStoresReenrollmentRequestWithDefaults() *CertificateStoresReenrollmentRequest`

NewCertificateStoresReenrollmentRequestWithDefaults instantiates a new CertificateStoresReenrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeystoreId

`func (o *CertificateStoresReenrollmentRequest) GetKeystoreId() string`

GetKeystoreId returns the KeystoreId field if non-nil, zero value otherwise.

### GetKeystoreIdOk

`func (o *CertificateStoresReenrollmentRequest) GetKeystoreIdOk() (*string, bool)`

GetKeystoreIdOk returns a tuple with the KeystoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreId

`func (o *CertificateStoresReenrollmentRequest) SetKeystoreId(v string)`

SetKeystoreId sets KeystoreId field to given value.

### HasKeystoreId

`func (o *CertificateStoresReenrollmentRequest) HasKeystoreId() bool`

HasKeystoreId returns a boolean if a field has been set.

### GetSubjectName

`func (o *CertificateStoresReenrollmentRequest) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *CertificateStoresReenrollmentRequest) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *CertificateStoresReenrollmentRequest) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *CertificateStoresReenrollmentRequest) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *CertificateStoresReenrollmentRequest) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *CertificateStoresReenrollmentRequest) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetAgentGuid

`func (o *CertificateStoresReenrollmentRequest) GetAgentGuid() string`

GetAgentGuid returns the AgentGuid field if non-nil, zero value otherwise.

### GetAgentGuidOk

`func (o *CertificateStoresReenrollmentRequest) GetAgentGuidOk() (*string, bool)`

GetAgentGuidOk returns a tuple with the AgentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentGuid

`func (o *CertificateStoresReenrollmentRequest) SetAgentGuid(v string)`

SetAgentGuid sets AgentGuid field to given value.

### HasAgentGuid

`func (o *CertificateStoresReenrollmentRequest) HasAgentGuid() bool`

HasAgentGuid returns a boolean if a field has been set.

### GetAlias

`func (o *CertificateStoresReenrollmentRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *CertificateStoresReenrollmentRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *CertificateStoresReenrollmentRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *CertificateStoresReenrollmentRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *CertificateStoresReenrollmentRequest) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *CertificateStoresReenrollmentRequest) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetJobProperties

`func (o *CertificateStoresReenrollmentRequest) GetJobProperties() map[string]interface{}`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *CertificateStoresReenrollmentRequest) GetJobPropertiesOk() (*map[string]interface{}, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *CertificateStoresReenrollmentRequest) SetJobProperties(v map[string]interface{})`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *CertificateStoresReenrollmentRequest) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *CertificateStoresReenrollmentRequest) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *CertificateStoresReenrollmentRequest) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetCertificateAuthority

`func (o *CertificateStoresReenrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *CertificateStoresReenrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *CertificateStoresReenrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *CertificateStoresReenrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *CertificateStoresReenrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *CertificateStoresReenrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetCertificateTemplate

`func (o *CertificateStoresReenrollmentRequest) GetCertificateTemplate() string`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *CertificateStoresReenrollmentRequest) GetCertificateTemplateOk() (*string, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *CertificateStoresReenrollmentRequest) SetCertificateTemplate(v string)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *CertificateStoresReenrollmentRequest) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *CertificateStoresReenrollmentRequest) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *CertificateStoresReenrollmentRequest) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil
### GetEnrollmentPatternId

`func (o *CertificateStoresReenrollmentRequest) GetEnrollmentPatternId() int32`

GetEnrollmentPatternId returns the EnrollmentPatternId field if non-nil, zero value otherwise.

### GetEnrollmentPatternIdOk

`func (o *CertificateStoresReenrollmentRequest) GetEnrollmentPatternIdOk() (*int32, bool)`

GetEnrollmentPatternIdOk returns a tuple with the EnrollmentPatternId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnrollmentPatternId

`func (o *CertificateStoresReenrollmentRequest) SetEnrollmentPatternId(v int32)`

SetEnrollmentPatternId sets EnrollmentPatternId field to given value.

### HasEnrollmentPatternId

`func (o *CertificateStoresReenrollmentRequest) HasEnrollmentPatternId() bool`

HasEnrollmentPatternId returns a boolean if a field has been set.

### SetEnrollmentPatternIdNil

`func (o *CertificateStoresReenrollmentRequest) SetEnrollmentPatternIdNil(b bool)`

 SetEnrollmentPatternIdNil sets the value for EnrollmentPatternId to be an explicit nil

### UnsetEnrollmentPatternId
`func (o *CertificateStoresReenrollmentRequest) UnsetEnrollmentPatternId()`

UnsetEnrollmentPatternId ensures that no value is present for EnrollmentPatternId, not even an explicit nil
### GetSANs

`func (o *CertificateStoresReenrollmentRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *CertificateStoresReenrollmentRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *CertificateStoresReenrollmentRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *CertificateStoresReenrollmentRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### SetSANsNil

`func (o *CertificateStoresReenrollmentRequest) SetSANsNil(b bool)`

 SetSANsNil sets the value for SANs to be an explicit nil

### UnsetSANs
`func (o *CertificateStoresReenrollmentRequest) UnsetSANs()`

UnsetSANs ensures that no value is present for SANs, not even an explicit nil
### GetMetadata

`func (o *CertificateStoresReenrollmentRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificateStoresReenrollmentRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificateStoresReenrollmentRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CertificateStoresReenrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CertificateStoresReenrollmentRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CertificateStoresReenrollmentRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetAdditionalEnrollmentFields

`func (o *CertificateStoresReenrollmentRequest) GetAdditionalEnrollmentFields() map[string]string`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *CertificateStoresReenrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]string, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *CertificateStoresReenrollmentRequest) SetAdditionalEnrollmentFields(v map[string]string)`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *CertificateStoresReenrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### SetAdditionalEnrollmentFieldsNil

`func (o *CertificateStoresReenrollmentRequest) SetAdditionalEnrollmentFieldsNil(b bool)`

 SetAdditionalEnrollmentFieldsNil sets the value for AdditionalEnrollmentFields to be an explicit nil

### UnsetAdditionalEnrollmentFields
`func (o *CertificateStoresReenrollmentRequest) UnsetAdditionalEnrollmentFields()`

UnsetAdditionalEnrollmentFields ensures that no value is present for AdditionalEnrollmentFields, not even an explicit nil
### GetOwnerRoleId

`func (o *CertificateStoresReenrollmentRequest) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *CertificateStoresReenrollmentRequest) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *CertificateStoresReenrollmentRequest) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *CertificateStoresReenrollmentRequest) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *CertificateStoresReenrollmentRequest) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *CertificateStoresReenrollmentRequest) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *CertificateStoresReenrollmentRequest) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *CertificateStoresReenrollmentRequest) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *CertificateStoresReenrollmentRequest) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *CertificateStoresReenrollmentRequest) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *CertificateStoresReenrollmentRequest) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *CertificateStoresReenrollmentRequest) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


