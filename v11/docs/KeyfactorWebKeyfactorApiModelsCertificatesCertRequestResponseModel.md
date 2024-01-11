# KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel

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
**State** | Pointer to **int32** |  | [optional] 
**StateString** | Pointer to **NullableString** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**DenialComment** | Pointer to **NullableString** |  | [optional] 
**KeyLength** | Pointer to **NullableString** |  | [optional] 
**SaNs** | Pointer to **[]string** |  | [optional] 
**CertStores** | Pointer to [**[]ModelsWorkflowCertificateRequestCertStoreModel**](ModelsWorkflowCertificateRequestCertStoreModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**SubjectAltNames** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsCertificatesSubjectAlternativeName**](KeyfactorWebKeyfactorApiModelsCertificatesSubjectAlternativeName.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel

`func NewKeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel() *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel`

NewKeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModelWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModelWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel`

NewKeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModelWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaRequestId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCaRequestId() string`

GetCaRequestId returns the CaRequestId field if non-nil, zero value otherwise.

### GetCaRequestIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCaRequestIdOk() (*string, bool)`

GetCaRequestIdOk returns a tuple with the CaRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRequestId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCaRequestId(v string)`

SetCaRequestId sets CaRequestId field to given value.

### HasCaRequestId

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasCaRequestId() bool`

HasCaRequestId returns a boolean if a field has been set.

### SetCaRequestIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCaRequestIdNil(b bool)`

 SetCaRequestIdNil sets the value for CaRequestId to be an explicit nil

### UnsetCaRequestId
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetCaRequestId()`

UnsetCaRequestId ensures that no value is present for CaRequestId, not even an explicit nil
### GetCommonName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetDistinguishedName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### SetDistinguishedNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetDistinguishedNameNil(b bool)`

 SetDistinguishedNameNil sets the value for DistinguishedName to be an explicit nil

### UnsetDistinguishedName
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetDistinguishedName()`

UnsetDistinguishedName ensures that no value is present for DistinguishedName, not even an explicit nil
### GetSubmissionDate

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### SetSubmissionDateNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetSubmissionDateNil(b bool)`

 SetSubmissionDateNil sets the value for SubmissionDate to be an explicit nil

### UnsetSubmissionDate
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetSubmissionDate()`

UnsetSubmissionDate ensures that no value is present for SubmissionDate, not even an explicit nil
### GetCertificateAuthority

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### SetCertificateAuthorityNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCertificateAuthorityNil(b bool)`

 SetCertificateAuthorityNil sets the value for CertificateAuthority to be an explicit nil

### UnsetCertificateAuthority
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetCertificateAuthority()`

UnsetCertificateAuthority ensures that no value is present for CertificateAuthority, not even an explicit nil
### GetTemplate

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### SetTemplateNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetTemplateNil(b bool)`

 SetTemplateNil sets the value for Template to be an explicit nil

### UnsetTemplate
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetTemplate()`

UnsetTemplate ensures that no value is present for Template, not even an explicit nil
### GetRequester

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### SetRequesterNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetRequesterNil(b bool)`

 SetRequesterNil sets the value for Requester to be an explicit nil

### UnsetRequester
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetRequester()`

UnsetRequester ensures that no value is present for Requester, not even an explicit nil
### GetState

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetState() int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetStateOk() (*int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetState(v int32)`

SetState sets State field to given value.

### HasState

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateString

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetStateString() string`

GetStateString returns the StateString field if non-nil, zero value otherwise.

### GetStateStringOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetStateStringOk() (*string, bool)`

GetStateStringOk returns a tuple with the StateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateString

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetStateString(v string)`

SetStateString sets StateString field to given value.

### HasStateString

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasStateString() bool`

HasStateString returns a boolean if a field has been set.

### SetStateStringNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetStateStringNil(b bool)`

 SetStateStringNil sets the value for StateString to be an explicit nil

### UnsetStateString
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetStateString()`

UnsetStateString ensures that no value is present for StateString, not even an explicit nil
### GetMetadata

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDenialComment

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetDenialComment() string`

GetDenialComment returns the DenialComment field if non-nil, zero value otherwise.

### GetDenialCommentOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetDenialCommentOk() (*string, bool)`

GetDenialCommentOk returns a tuple with the DenialComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialComment

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetDenialComment(v string)`

SetDenialComment sets DenialComment field to given value.

### HasDenialComment

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasDenialComment() bool`

HasDenialComment returns a boolean if a field has been set.

### SetDenialCommentNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetDenialCommentNil(b bool)`

 SetDenialCommentNil sets the value for DenialComment to be an explicit nil

### UnsetDenialComment
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetDenialComment()`

UnsetDenialComment ensures that no value is present for DenialComment, not even an explicit nil
### GetKeyLength

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetKeyLength() string`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetKeyLengthOk() (*string, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetKeyLength(v string)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### SetKeyLengthNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetKeyLengthNil(b bool)`

 SetKeyLengthNil sets the value for KeyLength to be an explicit nil

### UnsetKeyLength
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetKeyLength()`

UnsetKeyLength ensures that no value is present for KeyLength, not even an explicit nil
### GetSaNs

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetSaNs() []string`

GetSaNs returns the SaNs field if non-nil, zero value otherwise.

### GetSaNsOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetSaNsOk() (*[]string, bool)`

GetSaNsOk returns a tuple with the SaNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaNs

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetSaNs(v []string)`

SetSaNs sets SaNs field to given value.

### HasSaNs

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasSaNs() bool`

HasSaNs returns a boolean if a field has been set.

### SetSaNsNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetSaNsNil(b bool)`

 SetSaNsNil sets the value for SaNs to be an explicit nil

### UnsetSaNs
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetSaNs()`

UnsetSaNs ensures that no value is present for SaNs, not even an explicit nil
### GetCertStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCertStores() []ModelsWorkflowCertificateRequestCertStoreModel`

GetCertStores returns the CertStores field if non-nil, zero value otherwise.

### GetCertStoresOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCertStoresOk() (*[]ModelsWorkflowCertificateRequestCertStoreModel, bool)`

GetCertStoresOk returns a tuple with the CertStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCertStores(v []ModelsWorkflowCertificateRequestCertStoreModel)`

SetCertStores sets CertStores field to given value.

### HasCertStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasCertStores() bool`

HasCertStores returns a boolean if a field has been set.

### SetCertStoresNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCertStoresNil(b bool)`

 SetCertStoresNil sets the value for CertStores to be an explicit nil

### UnsetCertStores
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetCertStores()`

UnsetCertStores ensures that no value is present for CertStores, not even an explicit nil
### GetCurve

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetSubjectAltNames

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetSubjectAltNames() []KeyfactorWebKeyfactorApiModelsCertificatesSubjectAlternativeName`

GetSubjectAltNames returns the SubjectAltNames field if non-nil, zero value otherwise.

### GetSubjectAltNamesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) GetSubjectAltNamesOk() (*[]KeyfactorWebKeyfactorApiModelsCertificatesSubjectAlternativeName, bool)`

GetSubjectAltNamesOk returns a tuple with the SubjectAltNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectAltNames

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetSubjectAltNames(v []KeyfactorWebKeyfactorApiModelsCertificatesSubjectAlternativeName)`

SetSubjectAltNames sets SubjectAltNames field to given value.

### HasSubjectAltNames

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) HasSubjectAltNames() bool`

HasSubjectAltNames returns a boolean if a field has been set.

### SetSubjectAltNamesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) SetSubjectAltNamesNil(b bool)`

 SetSubjectAltNamesNil sets the value for SubjectAltNames to be an explicit nil

### UnsetSubjectAltNames
`func (o *KeyfactorWebKeyfactorApiModelsCertificatesCertRequestResponseModel) UnsetSubjectAltNames()`

UnsetSubjectAltNames ensures that no value is present for SubjectAltNames, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


