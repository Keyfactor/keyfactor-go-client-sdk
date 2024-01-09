# CSSCMSDataModelModelsWorkflowCertificateRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CaRequestId** | Pointer to **NullableString** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**DistinguishedName** | Pointer to **NullableString** |  | [optional] 
**SubmissionDate** | Pointer to **NullableTime** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Requester** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**KeyfactorPKIEnumsCertificateState**](KeyfactorPKIEnumsCertificateState.md) |  | [optional] 
**StateString** | Pointer to **NullableString** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsWorkflowCertificateRequestModel

`func NewCSSCMSDataModelModelsWorkflowCertificateRequestModel() *CSSCMSDataModelModelsWorkflowCertificateRequestModel`

NewCSSCMSDataModelModelsWorkflowCertificateRequestModel instantiates a new CSSCMSDataModelModelsWorkflowCertificateRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsWorkflowCertificateRequestModelWithDefaults

`func NewCSSCMSDataModelModelsWorkflowCertificateRequestModelWithDefaults() *CSSCMSDataModelModelsWorkflowCertificateRequestModel`

NewCSSCMSDataModelModelsWorkflowCertificateRequestModelWithDefaults instantiates a new CSSCMSDataModelModelsWorkflowCertificateRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaRequestId

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetCaRequestId() string`

GetCaRequestId returns the CaRequestId field if non-nil, zero value otherwise.

### GetCaRequestIdOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetCaRequestIdOk() (*string, bool)`

GetCaRequestIdOk returns a tuple with the CaRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRequestId

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetCaRequestId(v string)`

SetCaRequestId sets CaRequestId field to given value.

### HasCaRequestId

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasCaRequestId() bool`

HasCaRequestId returns a boolean if a field has been set.

### SetCaRequestIdNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetCaRequestIdNil(b bool)`

 SetCaRequestIdNil sets the value for CaRequestId to be an explicit nil

### UnsetCaRequestId
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetCaRequestId()`

UnsetCaRequestId ensures that no value is present for CaRequestId, not even an explicit nil
### GetCommonName

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetDistinguishedName

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetSubmissionDate

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### SetSubmissionDateNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetSubmissionDateNil(b bool)`

 SetSubmissionDateNil sets the value for SubmissionDate to be an explicit nil

### UnsetSubmissionDate
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetSubmissionDate()`

UnsetSubmissionDate ensures that no value is present for SubmissionDate, not even an explicit nil
### GetCertificateAuthority

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetTemplate

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetRequester

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetState

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetState() KeyfactorPKIEnumsCertificateState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetStateOk() (*KeyfactorPKIEnumsCertificateState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetState(v KeyfactorPKIEnumsCertificateState)`

SetState sets State field to given value.

### HasState

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateString

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetStateString() string`

GetStateString returns the StateString field if non-nil, zero value otherwise.

### GetStateStringOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetStateStringOk() (*string, bool)`

GetStateStringOk returns a tuple with the StateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateString

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetStateString(v string)`

SetStateString sets StateString field to given value.

### HasStateString

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasStateString() bool`

HasStateString returns a boolean if a field has been set.

### SetStateStringNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetStateStringNil(b bool)`

 SetStateStringNil sets the value for StateString to be an explicit nil

### UnsetStateString
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetStateString()`

UnsetStateString ensures that no value is present for StateString, not even an explicit nil
### GetMetadata

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsWorkflowCertificateRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


