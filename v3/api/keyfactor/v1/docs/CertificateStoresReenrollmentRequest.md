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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


