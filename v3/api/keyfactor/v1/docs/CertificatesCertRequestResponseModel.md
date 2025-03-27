# CertificatesCertRequestResponseModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CARequestId** | Pointer to **NullableString** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**DistinguishedName** | Pointer to **NullableString** |  | [optional] 
**SubmissionDate** | Pointer to **NullableTime** |  | [optional] 
**CertificateAuthority** | Pointer to **NullableString** |  | [optional] 
**Template** | Pointer to **NullableString** |  | [optional] 
**Requester** | Pointer to **NullableString** |  | [optional] 
**State** | Pointer to [**KeyfactorPKIEnumsCertificateState**](KeyfactorPKIEnumsCertificateState.md) |  | [optional] 
**StateString** | Pointer to **NullableString** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**DenialComment** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **NullableString** |  | [optional] 
**SANs** | Pointer to **[]string** |  | [optional] 
**CertStores** | Pointer to [**[]CSSCMSDataModelModelsWorkflowCertificateRequestCertStoreModel**](CSSCMSDataModelModelsWorkflowCertificateRequestCertStoreModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**SubjectAltNames** | Pointer to [**[]CertificatesSubjectAlternativeName**](CertificatesSubjectAlternativeName.md) |  | [optional] 

## Methods

### NewCertificatesCertRequestResponseModel

`func NewCertificatesCertRequestResponseModel() *CertificatesCertRequestResponseModel`

NewCertificatesCertRequestResponseModel instantiates a new CertificatesCertRequestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertRequestResponseModelWithDefaults

`func NewCertificatesCertRequestResponseModelWithDefaults() *CertificatesCertRequestResponseModel`

NewCertificatesCertRequestResponseModelWithDefaults instantiates a new CertificatesCertRequestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesCertRequestResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesCertRequestResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesCertRequestResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesCertRequestResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCARequestId

`func (o *CertificatesCertRequestResponseModel) GetCARequestId() string`

GetCARequestId returns the CARequestId field if non-nil, zero value otherwise.

### GetCARequestIdOk

`func (o *CertificatesCertRequestResponseModel) GetCARequestIdOk() (*string, bool)`

GetCARequestIdOk returns a tuple with the CARequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARequestId

`func (o *CertificatesCertRequestResponseModel) SetCARequestId(v string)`

SetCARequestId sets CARequestId field to given value.

### HasCARequestId

`func (o *CertificatesCertRequestResponseModel) HasCARequestId() bool`

HasCARequestId returns a boolean if a field has been set.

### SetCARequestIdNil

`func (o *CertificatesCertRequestResponseModel) SetCARequestIdNil(b bool)`

 SetCARequestIdNil sets the value for CARequestId to be an explicit nil

### UnsetCARequestId
`func (o *CertificatesCertRequestResponseModel) UnsetCARequestId()`

UnsetCARequestId ensures that no value is present for CARequestId, not even an explicit nil
### GetCommonName

`func (o *CertificatesCertRequestResponseModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *CertificatesCertRequestResponseModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *CertificatesCertRequestResponseModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *CertificatesCertRequestResponseModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *CertificatesCertRequestResponseModel) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *CertificatesCertRequestResponseModel) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetDistinguishedName

`func (o *CertificatesCertRequestResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *CertificatesCertRequestResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *CertificatesCertRequestResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *CertificatesCertRequestResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *CertificatesCertRequestResponseModel) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *CertificatesCertRequestResponseModel) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetSubmissionDate

`func (o *CertificatesCertRequestResponseModel) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *CertificatesCertRequestResponseModel) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *CertificatesCertRequestResponseModel) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *CertificatesCertRequestResponseModel) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### SetSubmissionDateNil

`func (o *CertificatesCertRequestResponseModel) SetSubmissionDateNil(b bool)`

 SetSubmissionDateNil sets the value for SubmissionDate to be an explicit nil

### UnsetSubmissionDate
`func (o *CertificatesCertRequestResponseModel) UnsetSubmissionDate()`

UnsetSubmissionDate ensures that no value is present for SubmissionDate, not even an explicit nil
### GetCertificateAuthority

`func (o *CertificatesCertRequestResponseModel) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *CertificatesCertRequestResponseModel) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *CertificatesCertRequestResponseModel) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *CertificatesCertRequestResponseModel) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *CertificatesCertRequestResponseModel) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *CertificatesCertRequestResponseModel) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetTemplate

`func (o *CertificatesCertRequestResponseModel) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CertificatesCertRequestResponseModel) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CertificatesCertRequestResponseModel) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CertificatesCertRequestResponseModel) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *CertificatesCertRequestResponseModel) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *CertificatesCertRequestResponseModel) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetRequester

`func (o *CertificatesCertRequestResponseModel) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *CertificatesCertRequestResponseModel) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *CertificatesCertRequestResponseModel) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *CertificatesCertRequestResponseModel) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *CertificatesCertRequestResponseModel) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *CertificatesCertRequestResponseModel) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetState

`func (o *CertificatesCertRequestResponseModel) GetState() KeyfactorPKIEnumsCertificateState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CertificatesCertRequestResponseModel) GetStateOk() (*KeyfactorPKIEnumsCertificateState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CertificatesCertRequestResponseModel) SetState(v KeyfactorPKIEnumsCertificateState)`

SetState sets State field to given value.

### HasState

`func (o *CertificatesCertRequestResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateString

`func (o *CertificatesCertRequestResponseModel) GetStateString() string`

GetStateString returns the StateString field if non-nil, zero value otherwise.

### GetStateStringOk

`func (o *CertificatesCertRequestResponseModel) GetStateStringOk() (*string, bool)`

GetStateStringOk returns a tuple with the StateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateString

`func (o *CertificatesCertRequestResponseModel) SetStateString(v string)`

SetStateString sets StateString field to given value.

### HasStateString

`func (o *CertificatesCertRequestResponseModel) HasStateString() bool`

HasStateString returns a boolean if a field has been set.

### SetStateStringNil

`func (o *CertificatesCertRequestResponseModel) SetStateStringNil(b bool)`

 SetStateStringNil sets the value for StateString to be an explicit nil

### UnsetStateString
`func (o *CertificatesCertRequestResponseModel) UnsetStateString()`

UnsetStateString ensures that no value is present for StateString, not even an explicit nil
### GetMetadata

`func (o *CertificatesCertRequestResponseModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CertificatesCertRequestResponseModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CertificatesCertRequestResponseModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CertificatesCertRequestResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CertificatesCertRequestResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CertificatesCertRequestResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDenialComment

`func (o *CertificatesCertRequestResponseModel) GetDenialComment() string`

GetDenialComment returns the DenialComment field if non-nil, zero value otherwise.

### GetDenialCommentOk

`func (o *CertificatesCertRequestResponseModel) GetDenialCommentOk() (*string, bool)`

GetDenialCommentOk returns a tuple with the DenialComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialComment

`func (o *CertificatesCertRequestResponseModel) SetDenialComment(v string)`

SetDenialComment sets DenialComment field to given value.

### HasDenialComment

`func (o *CertificatesCertRequestResponseModel) HasDenialComment() bool`

HasDenialComment returns a boolean if a field has been set.

### SetDenialCommentNil

`func (o *CertificatesCertRequestResponseModel) SetDenialCommentNil(b bool)`

 SetDenialCommentNil sets the value for DenialComment to be an explicit nil

### UnsetDenialComment
`func (o *CertificatesCertRequestResponseModel) UnsetDenialComment()`

UnsetDenialComment ensures that no value is present for DenialComment, not even an explicit nil
### GetKeyLength

`func (o *CertificatesCertRequestResponseModel) GetKeyLength() string`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *CertificatesCertRequestResponseModel) GetKeyLengthOk() (*string, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *CertificatesCertRequestResponseModel) SetKeyLength(v string)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *CertificatesCertRequestResponseModel) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### SetKeyLengthNil

`func (o *CertificatesCertRequestResponseModel) SetKeyLengthNil(b bool)`

 SetKeyLengthNil sets the value for KeyLength to be an explicit nil

### UnsetKeyLength
`func (o *CertificatesCertRequestResponseModel) UnsetKeyLength()`

UnsetKeyLength ensures that no value is present for KeyLength, not even an explicit nil
### GetSANs

`func (o *CertificatesCertRequestResponseModel) GetSANs() []string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *CertificatesCertRequestResponseModel) GetSANsOk() (*[]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *CertificatesCertRequestResponseModel) SetSANs(v []string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *CertificatesCertRequestResponseModel) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### SetSANsNil

`func (o *CertificatesCertRequestResponseModel) SetSANsNil(b bool)`

 SetSANsNil sets the value for SANs to be an explicit nil

### UnsetSANs
`func (o *CertificatesCertRequestResponseModel) UnsetSANs()`

UnsetSANs ensures that no value is present for SANs, not even an explicit nil
### GetCertStores

`func (o *CertificatesCertRequestResponseModel) GetCertStores() []CSSCMSDataModelModelsWorkflowCertificateRequestCertStoreModel`

GetCertStores returns the CertStores field if non-nil, zero value otherwise.

### GetCertStoresOk

`func (o *CertificatesCertRequestResponseModel) GetCertStoresOk() (*[]CSSCMSDataModelModelsWorkflowCertificateRequestCertStoreModel, bool)`

GetCertStoresOk returns a tuple with the CertStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStores

`func (o *CertificatesCertRequestResponseModel) SetCertStores(v []CSSCMSDataModelModelsWorkflowCertificateRequestCertStoreModel)`

SetCertStores sets CertStores field to given value.

### HasCertStores

`func (o *CertificatesCertRequestResponseModel) HasCertStores() bool`

HasCertStores returns a boolean if a field has been set.

### SetCertStoresNil

`func (o *CertificatesCertRequestResponseModel) SetCertStoresNil(b bool)`

 SetCertStoresNil sets the value for CertStores to be an explicit nil

### UnsetCertStores
`func (o *CertificatesCertRequestResponseModel) UnsetCertStores()`

UnsetCertStores ensures that no value is present for CertStores, not even an explicit nil
### GetCurve

`func (o *CertificatesCertRequestResponseModel) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *CertificatesCertRequestResponseModel) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *CertificatesCertRequestResponseModel) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *CertificatesCertRequestResponseModel) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *CertificatesCertRequestResponseModel) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *CertificatesCertRequestResponseModel) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetSubjectAltNames

`func (o *CertificatesCertRequestResponseModel) GetSubjectAltNames() []CertificatesSubjectAlternativeName`

GetSubjectAltNames returns the SubjectAltNames field if non-nil, zero value otherwise.

### GetSubjectAltNamesOk

`func (o *CertificatesCertRequestResponseModel) GetSubjectAltNamesOk() (*[]CertificatesSubjectAlternativeName, bool)`

GetSubjectAltNamesOk returns a tuple with the SubjectAltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNames

`func (o *CertificatesCertRequestResponseModel) SetSubjectAltNames(v []CertificatesSubjectAlternativeName)`

SetSubjectAltNames sets SubjectAltNames field to given value.

### HasSubjectAltNames

`func (o *CertificatesCertRequestResponseModel) HasSubjectAltNames() bool`

HasSubjectAltNames returns a boolean if a field has been set.

### SetSubjectAltNamesNil

`func (o *CertificatesCertRequestResponseModel) SetSubjectAltNamesNil(b bool)`

 SetSubjectAltNamesNil sets the value for SubjectAltNames to be an explicit nil

### UnsetSubjectAltNames
`func (o *CertificatesCertRequestResponseModel) UnsetSubjectAltNames()`

UnsetSubjectAltNames ensures that no value is present for SubjectAltNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


