# ModelsWorkflowCertificateRequestDetailsModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DenialComment** | Pointer to **string** |  | [optional] 
**KeyLength** | Pointer to **string** |  | [optional] 
**SANs** | Pointer to **[]string** |  | [optional] 
**CertStores** | Pointer to [**[]ModelsWorkflowCertificateRequestCertStoreModel**](ModelsWorkflowCertificateRequestCertStoreModel.md) |  | [optional] 
**Curve** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**CARequestId** | Pointer to **string** |  | [optional] 
**CommonName** | Pointer to **string** |  | [optional] 
**DistinguishedName** | Pointer to **string** |  | [optional] 
**SubmissionDate** | Pointer to **time.Time** |  | [optional] 
**CertificateAuthority** | Pointer to **string** |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 
**Requester** | Pointer to **string** |  | [optional] 
**State** | Pointer to **int64** |  | [optional] 
**StateString** | Pointer to **string** |  | [optional] [readonly] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsWorkflowCertificateRequestDetailsModel

`func NewModelsWorkflowCertificateRequestDetailsModel() *ModelsWorkflowCertificateRequestDetailsModel`

NewModelsWorkflowCertificateRequestDetailsModel instantiates a new ModelsWorkflowCertificateRequestDetailsModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWorkflowCertificateRequestDetailsModelWithDefaults

`func NewModelsWorkflowCertificateRequestDetailsModelWithDefaults() *ModelsWorkflowCertificateRequestDetailsModel`

NewModelsWorkflowCertificateRequestDetailsModelWithDefaults instantiates a new ModelsWorkflowCertificateRequestDetailsModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDenialComment

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetDenialComment() string`

GetDenialComment returns the DenialComment field if non-nil, zero value otherwise.

### GetDenialCommentOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetDenialCommentOk() (*string, bool)`

GetDenialCommentOk returns a tuple with the DenialComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDenialComment

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetDenialComment(v string)`

SetDenialComment sets DenialComment field to given value.

### HasDenialComment

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasDenialComment() bool`

HasDenialComment returns a boolean if a field has been set.

### GetKeyLength

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetKeyLength() string`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetKeyLengthOk() (*string, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetKeyLength(v string)`

SetKeyLength sets KeyLength field to given value.

### HasKeyLength

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasKeyLength() bool`

HasKeyLength returns a boolean if a field has been set.

### GetSANs

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetSANs() []string`

GetSANs returns the SANs field if non-nil, zero value otherwise.

### GetSANsOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetSANsOk() (*[]string, bool)`

GetSANsOk returns a tuple with the SANs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSANs

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetSANs(v []string)`

SetSANs sets SANs field to given value.

### HasSANs

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasSANs() bool`

HasSANs returns a boolean if a field has been set.

### GetCertStores

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCertStores() []ModelsWorkflowCertificateRequestCertStoreModel`

GetCertStores returns the CertStores field if non-nil, zero value otherwise.

### GetCertStoresOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCertStoresOk() (*[]ModelsWorkflowCertificateRequestCertStoreModel, bool)`

GetCertStoresOk returns a tuple with the CertStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStores

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetCertStores(v []ModelsWorkflowCertificateRequestCertStoreModel)`

SetCertStores sets CertStores field to given value.

### HasCertStores

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasCertStores() bool`

HasCertStores returns a boolean if a field has been set.

### GetCurve

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCurve() string`

GetCurve returns the Curve field if non-nil, zero value otherwise.

### GetCurveOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCurveOk() (*string, bool)`

GetCurveOk returns a tuple with the Curve field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurve

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetCurve(v string)`

SetCurve sets Curve field to given value.

### HasCurve

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasCurve() bool`

HasCurve returns a boolean if a field has been set.

### SetCurveNil

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetCurveNil(b bool)`

 SetCurveNil sets the value for Curve to be an explicit nil

### UnsetCurve
`func (o *ModelsWorkflowCertificateRequestDetailsModel) UnsetCurve()`

UnsetCurve ensures that no value is present for Curve, not even an explicit nil
### GetId

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCARequestId

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCARequestId() string`

GetCARequestId returns the CARequestId field if non-nil, zero value otherwise.

### GetCARequestIdOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCARequestIdOk() (*string, bool)`

GetCARequestIdOk returns a tuple with the CARequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARequestId

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetCARequestId(v string)`

SetCARequestId sets CARequestId field to given value.

### HasCARequestId

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasCARequestId() bool`

HasCARequestId returns a boolean if a field has been set.

### GetCommonName

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetCertificateAuthority

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetTemplate

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRequester

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetState

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetState() int64`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetStateOk() (*int64, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetState(v int64)`

SetState sets State field to given value.

### HasState

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateString

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetStateString() string`

GetStateString returns the StateString field if non-nil, zero value otherwise.

### GetStateStringOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetStateStringOk() (*string, bool)`

GetStateStringOk returns a tuple with the StateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateString

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetStateString(v string)`

SetStateString sets StateString field to given value.

### HasStateString

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasStateString() bool`

HasStateString returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsWorkflowCertificateRequestDetailsModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsWorkflowCertificateRequestDetailsModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsWorkflowCertificateRequestDetailsModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


