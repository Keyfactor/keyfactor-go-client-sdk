# ModelsWorkflowCertificateRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewModelsWorkflowCertificateRequestModel

`func NewModelsWorkflowCertificateRequestModel() *ModelsWorkflowCertificateRequestModel`

NewModelsWorkflowCertificateRequestModel instantiates a new ModelsWorkflowCertificateRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWorkflowCertificateRequestModelWithDefaults

`func NewModelsWorkflowCertificateRequestModelWithDefaults() *ModelsWorkflowCertificateRequestModel`

NewModelsWorkflowCertificateRequestModelWithDefaults instantiates a new ModelsWorkflowCertificateRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsWorkflowCertificateRequestModel) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsWorkflowCertificateRequestModel) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsWorkflowCertificateRequestModel) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsWorkflowCertificateRequestModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCARequestId

`func (o *ModelsWorkflowCertificateRequestModel) GetCARequestId() string`

GetCARequestId returns the CARequestId field if non-nil, zero value otherwise.

### GetCARequestIdOk

`func (o *ModelsWorkflowCertificateRequestModel) GetCARequestIdOk() (*string, bool)`

GetCARequestIdOk returns a tuple with the CARequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARequestId

`func (o *ModelsWorkflowCertificateRequestModel) SetCARequestId(v string)`

SetCARequestId sets CARequestId field to given value.

### HasCARequestId

`func (o *ModelsWorkflowCertificateRequestModel) HasCARequestId() bool`

HasCARequestId returns a boolean if a field has been set.

### GetCommonName

`func (o *ModelsWorkflowCertificateRequestModel) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *ModelsWorkflowCertificateRequestModel) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *ModelsWorkflowCertificateRequestModel) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *ModelsWorkflowCertificateRequestModel) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetDistinguishedName

`func (o *ModelsWorkflowCertificateRequestModel) GetDistinguishedName() string`

GetDistinguishedName returns the DistinguishedName field if non-nil, zero value otherwise.

### GetDistinguishedNameOk

`func (o *ModelsWorkflowCertificateRequestModel) GetDistinguishedNameOk() (*string, bool)`

GetDistinguishedNameOk returns a tuple with the DistinguishedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistinguishedName

`func (o *ModelsWorkflowCertificateRequestModel) SetDistinguishedName(v string)`

SetDistinguishedName sets DistinguishedName field to given value.

### HasDistinguishedName

`func (o *ModelsWorkflowCertificateRequestModel) HasDistinguishedName() bool`

HasDistinguishedName returns a boolean if a field has been set.

### GetSubmissionDate

`func (o *ModelsWorkflowCertificateRequestModel) GetSubmissionDate() time.Time`

GetSubmissionDate returns the SubmissionDate field if non-nil, zero value otherwise.

### GetSubmissionDateOk

`func (o *ModelsWorkflowCertificateRequestModel) GetSubmissionDateOk() (*time.Time, bool)`

GetSubmissionDateOk returns a tuple with the SubmissionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionDate

`func (o *ModelsWorkflowCertificateRequestModel) SetSubmissionDate(v time.Time)`

SetSubmissionDate sets SubmissionDate field to given value.

### HasSubmissionDate

`func (o *ModelsWorkflowCertificateRequestModel) HasSubmissionDate() bool`

HasSubmissionDate returns a boolean if a field has been set.

### GetCertificateAuthority

`func (o *ModelsWorkflowCertificateRequestModel) GetCertificateAuthority() string`

GetCertificateAuthority returns the CertificateAuthority field if non-nil, zero value otherwise.

### GetCertificateAuthorityOk

`func (o *ModelsWorkflowCertificateRequestModel) GetCertificateAuthorityOk() (*string, bool)`

GetCertificateAuthorityOk returns a tuple with the CertificateAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateAuthority

`func (o *ModelsWorkflowCertificateRequestModel) SetCertificateAuthority(v string)`

SetCertificateAuthority sets CertificateAuthority field to given value.

### HasCertificateAuthority

`func (o *ModelsWorkflowCertificateRequestModel) HasCertificateAuthority() bool`

HasCertificateAuthority returns a boolean if a field has been set.

### GetTemplate

`func (o *ModelsWorkflowCertificateRequestModel) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *ModelsWorkflowCertificateRequestModel) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *ModelsWorkflowCertificateRequestModel) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *ModelsWorkflowCertificateRequestModel) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRequester

`func (o *ModelsWorkflowCertificateRequestModel) GetRequester() string`

GetRequester returns the Requester field if non-nil, zero value otherwise.

### GetRequesterOk

`func (o *ModelsWorkflowCertificateRequestModel) GetRequesterOk() (*string, bool)`

GetRequesterOk returns a tuple with the Requester field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequester

`func (o *ModelsWorkflowCertificateRequestModel) SetRequester(v string)`

SetRequester sets Requester field to given value.

### HasRequester

`func (o *ModelsWorkflowCertificateRequestModel) HasRequester() bool`

HasRequester returns a boolean if a field has been set.

### GetState

`func (o *ModelsWorkflowCertificateRequestModel) GetState() int64`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ModelsWorkflowCertificateRequestModel) GetStateOk() (*int64, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ModelsWorkflowCertificateRequestModel) SetState(v int64)`

SetState sets State field to given value.

### HasState

`func (o *ModelsWorkflowCertificateRequestModel) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStateString

`func (o *ModelsWorkflowCertificateRequestModel) GetStateString() string`

GetStateString returns the StateString field if non-nil, zero value otherwise.

### GetStateStringOk

`func (o *ModelsWorkflowCertificateRequestModel) GetStateStringOk() (*string, bool)`

GetStateStringOk returns a tuple with the StateString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateString

`func (o *ModelsWorkflowCertificateRequestModel) SetStateString(v string)`

SetStateString sets StateString field to given value.

### HasStateString

`func (o *ModelsWorkflowCertificateRequestModel) HasStateString() bool`

HasStateString returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsWorkflowCertificateRequestModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsWorkflowCertificateRequestModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsWorkflowCertificateRequestModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsWorkflowCertificateRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


