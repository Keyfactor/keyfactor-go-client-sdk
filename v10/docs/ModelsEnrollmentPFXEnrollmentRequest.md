# ModelsEnrollmentPFXEnrollmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFriendlyName** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PopulateMissingValuesFromAD** | Pointer to **bool** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**IncludeChain** | Pointer to **bool** |  | [optional] 
**RenewalCertificateId** | Pointer to **int64** |  | [optional] 
**CertificateAuthority** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**AdditionalEnrollmentFields** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**SANs** | Pointer to **map[string][]string** |  | [optional] 

## Methods

### NewModelsEnrollmentPFXEnrollmentRequest

`func NewModelsEnrollmentPFXEnrollmentRequest() *ModelsEnrollmentPFXEnrollmentRequest`

NewModelsEnrollmentPFXEnrollmentRequest instantiates a new ModelsEnrollmentPFXEnrollmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsEnrollmentPFXEnrollmentRequestWithDefaults

`func NewModelsEnrollmentPFXEnrollmentRequestWithDefaults() *ModelsEnrollmentPFXEnrollmentRequest`

NewModelsEnrollmentPFXEnrollmentRequestWithDefaults instantiates a new ModelsEnrollmentPFXEnrollmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFriendlyName

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyName() string`

GetCustomFriendlyName returns the CustomFriendlyName field if non-nil, zero value otherwise.

### GetCustomFriendlyNameOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCustomFriendlyNameOk() (*string, bool)`

GetCustomFriendlyNameOk returns a tuple with the CustomFriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFriendlyName

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCustomFriendlyName(v string)`

SetCustomFriendlyName sets CustomFriendlyName field to given value.

### HasCustomFriendlyName

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCustomFriendlyName() bool`

HasCustomFriendlyName returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPopulateMissingValuesFromAD

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromAD() bool`

GetPopulateMissingValuesFromAD returns the PopulateMissingValuesFromAD field if non-nil, zero value otherwise.

### GetPopulateMissingValuesFromADOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetPopulateMissingValuesFromADOk() (*bool, bool)`

GetPopulateMissingValuesFromADOk returns a tuple with the PopulateMissingValuesFromAD field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulateMissingValuesFromAD

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetPopulateMissingValuesFromAD(v bool)`

SetPopulateMissingValuesFromAD sets PopulateMissingValuesFromAD field to given value.

### HasPopulateMissingValuesFromAD

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasPopulateMissingValuesFromAD() bool`

HasPopulateMissingValuesFromAD returns a boolean if a field has been set.

### GetSubject

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetIncludeChain

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetIncludeChain() bool`

GetIncludeChain returns the IncludeChain field if non-nil, zero value otherwise.

### GetIncludeChainOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetIncludeChainOk() (*bool, bool)`

GetIncludeChainOk returns a tuple with the IncludeChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeChain

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetIncludeChain(v bool)`

SetIncludeChain sets IncludeChain field to given value.

### HasIncludeChain

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasIncludeChain() bool`

HasIncludeChain returns a boolean if a field has been set.

### GetRenewalCertificateId

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateId() int64`

GetRenewalCertificateId returns the RenewalCertificateId field if non-nil, zero value otherwise.

### GetRenewalCertificateIdOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetRenewalCertificateIdOk() (*int64, bool)`

GetRenewalCertificateIdOk returns a tuple with the RenewalCertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRenewalCertificateId

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetRenewalCertificateId(v int64)`

SetRenewalCertificateId sets RenewalCertificateId field to given value.

### HasRenewalCertificateId

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasRenewalCertificateId() bool`

HasRenewalCertificateId returns a boolean if a field has been set.

### GetCertificateAuthority

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFields() map[string]map[string]interface{}`

GetAdditionalEnrollmentFields returns the AdditionalEnrollmentFields field if non-nil, zero value otherwise.

### GetAdditionalEnrollmentFieldsOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetAdditionalEnrollmentFieldsOk() (*map[string]map[string]interface{}, bool)`

GetAdditionalEnrollmentFieldsOk returns a tuple with the AdditionalEnrollmentFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalEnrollmentFields

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetAdditionalEnrollmentFields(v map[string]map[string]interface{})`

SetAdditionalEnrollmentFields sets AdditionalEnrollmentFields field to given value.

### HasAdditionalEnrollmentFields

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasAdditionalEnrollmentFields() bool`

HasAdditionalEnrollmentFields returns a boolean if a field has been set.

### GetTimestamp

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTemplate

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetSANs

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSANs() map[string][]string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *ModelsEnrollmentPFXEnrollmentRequest) GetSANsOk() (*map[string][]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *ModelsEnrollmentPFXEnrollmentRequest) SetSANs(v map[string][]string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *ModelsEnrollmentPFXEnrollmentRequest) HasSANs() bool`

HasSANs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


