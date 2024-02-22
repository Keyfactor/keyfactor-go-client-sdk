# KeyfactorApiModelsCertificateStoresReenrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeystoreId** | Pointer to **string** |  | [optional] 
**SubjectName** | Pointer to **string** |  | [optional] 
**AgentGuid** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**JobProperties** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**CertificateAuthority** | Pointer to **string** |  | [optional] 
**CertificateTemplate** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresReenrollmentRequest

`func NewKeyfactorApiModelsCertificateStoresReenrollmentRequest() *KeyfactorApiModelsCertificateStoresReenrollmentRequest`

NewKeyfactorApiModelsCertificateStoresReenrollmentRequest instantiates a new KeyfactorApiModelsCertificateStoresReenrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresReenrollmentRequestWithDefaults

`func NewKeyfactorApiModelsCertificateStoresReenrollmentRequestWithDefaults() *KeyfactorApiModelsCertificateStoresReenrollmentRequest`

NewKeyfactorApiModelsCertificateStoresReenrollmentRequestWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresReenrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeystoreId

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetKeystoreId() string`

GetKeystoreId returns the KeystoreId field if non-nil, zero value otherwise.

### GetKeystoreIdOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetKeystoreIdOk() (*string, bool)`

GetKeystoreIdOk returns a tuple with the KeystoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreId

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetKeystoreId(v string)`

SetKeystoreId sets KeystoreId field to given value.

### HasKeystoreId

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasKeystoreId() bool`

HasKeystoreId returns a boolean if a field has been set.

### GetSubjectName

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetSubjectName() string`

GetSubjectName returns the SubjectName field if non-nil, zero value otherwise.

### GetSubjectNameOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetSubjectNameOk() (*string, bool)`

GetSubjectNameOk returns a tuple with the SubjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectName

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetSubjectName(v string)`

SetSubjectName sets SubjectName field to given value.

### HasSubjectName

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasSubjectName() bool`

HasSubjectName returns a boolean if a field has been set.

### GetAgentGuid

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAgentGuid() string`

GetAgentGuid returns the AgentGuid field if non-nil, zero value otherwise.

### GetAgentGuidOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAgentGuidOk() (*string, bool)`

GetAgentGuidOk returns a tuple with the AgentGuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentGuid

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetAgentGuid(v string)`

SetAgentGuid sets AgentGuid field to given value.

### HasAgentGuid

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasAgentGuid() bool`

HasAgentGuid returns a boolean if a field has been set.

### GetAlias

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetJobProperties

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetJobProperties() map[string]map[string]interface{}`

GetJobProperties returns the JobProperties field if non-nil, zero value otherwise.

### GetJobPropertiesOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetJobPropertiesOk() (*map[string]map[string]interface{}, bool)`

GetJobPropertiesOk returns a tuple with the JobProperties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobProperties

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetJobProperties(v map[string]map[string]interface{})`

SetJobProperties sets JobProperties field to given value.

### HasJobProperties

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasJobProperties() bool`

HasJobProperties returns a boolean if a field has been set.

### GetCertificateAuthority

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetCertificateTemplate

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateTemplate() string`

GetCertificateTemplate returns the CertificateTemplate field if non-nil, zero value otherwise.

### GetCertificateTemplateOk

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) GetCertificateTemplateOk() (*string, bool)`

GetCertificateTemplateOk returns a tuple with the CertificateTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateTemplate

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) SetCertificateTemplate(v string)`

SetCertificateTemplate sets CertificateTemplate field to given value.

### HasCertificateTemplate

`func (o *KeyfactorApiModelsCertificateStoresReenrollmentRequest) HasCertificateTemplate() bool`

HasCertificateTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


