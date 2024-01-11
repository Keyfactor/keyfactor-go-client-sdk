# KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest

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

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest() *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeystoreId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetKeystoreId() string`

GetKeystoreId returns the KeystoreId field if non-nil, zero value otherwise.

### GetKeystoreIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetKeystoreIdOk() (*string, bool)`

GetKeystoreIdOk returns a tuple with the KeystoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetKeystoreId(v string)`

SetKeystoreId sets KeystoreId field to given value.

### HasKeystoreId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasKeystoreId() bool`

HasKeystoreId returns a boolean if a field has been set.

### GetSubjectName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### SetSubjectNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetSubjectNameNil(b bool)`

 SetSubjectNameNil sets the value for SubjectName to be an explicit nil

### UnsetSubjectName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) UnsetSubjectName()`

UnsetSubjectName ensures that no value is present for SubjectName, not even an explicit nil
### GetAgentGuid

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAgentGuid() string`

GetAgentGuid returns the AgentGuid field if non-nil, zero value otherwise.

### GetAgentGuidOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAgentGuidOk() (*string, bool)`

GetAgentGuidOk returns a tuple with the AgentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentGuid

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetAgentGuid(v string)`

SetAgentGuid sets AgentGuid field to given value.

### HasAgentGuid

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasAgentGuid() bool`

HasAgentGuid returns a boolean if a field has been set.

### GetAlias

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetJobProperties() map[string]interface{}`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetJobPropertiesOk() (*map[string]interface{}, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetJobProperties(v map[string]interface{})`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### SetJobPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetJobPropertiesNil(b bool)`

 SetJobPropertiesNil sets the value for JobProperties to be an explicit nil

### UnsetJobProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) UnsetJobProperties()`

UnsetJobProperties ensures that no value is present for JobProperties, not even an explicit nil
### GetCertificateAuthority

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetCertificateTemplate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateTemplate() string`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateTemplateOk() (*string, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetCertificateTemplate(v string)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.

### SetCertificateTemplateNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) SetCertificateTemplateNil(b bool)`

 SetCertificateTemplateNil sets the value for CertificateTemplate to be an explicit nil

### UnsetCertificateTemplate
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresReenrollmentRequest) UnsetCertificateTemplate()`

UnsetCertificateTemplate ensures that no value is present for CertificateTemplate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


