# ModelsRevokeCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateIds** | Pointer to **[]int32** |  | [optional] 
**Reason** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewModelsRevokeCertificateRequest

`func NewModelsRevokeCertificateRequest() *ModelsRevokeCertificateRequest`

NewModelsRevokeCertificateRequest instantiates a new ModelsRevokeCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsRevokeCertificateRequestWithDefaults

`func NewModelsRevokeCertificateRequestWithDefaults() *ModelsRevokeCertificateRequest`

NewModelsRevokeCertificateRequestWithDefaults instantiates a new ModelsRevokeCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateIds

`func (o *ModelsRevokeCertificateRequest) GetCertificateIds() []int32`

GetCertificateIds returns the CertificateIds field if non-nil, zero value otherwise.

### GetCertificateIdsOk

`func (o *ModelsRevokeCertificateRequest) GetCertificateIdsOk() (*[]int32, bool)`

GetCertificateIdsOk returns a tuple with the CertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIds

`func (o *ModelsRevokeCertificateRequest) SetCertificateIds(v []int32)`

SetCertificateIds sets CertificateIds field to given value.

### HasCertificateIds

`func (o *ModelsRevokeCertificateRequest) HasCertificateIds() bool`

HasCertificateIds returns a boolean if a field has been set.

### SetCertificateIdsNil

`func (o *ModelsRevokeCertificateRequest) SetCertificateIdsNil(b bool)`

 SetCertificateIdsNil sets the value for CertificateIds to be an explicit nil

### UnsetCertificateIds
`func (o *ModelsRevokeCertificateRequest) UnsetCertificateIds()`

UnsetCertificateIds ensures that no value is present for CertificateIds, not even an explicit nil
### GetReason

`func (o *ModelsRevokeCertificateRequest) GetReason() int32`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelsRevokeCertificateRequest) GetReasonOk() (*int32, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelsRevokeCertificateRequest) SetReason(v int32)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ModelsRevokeCertificateRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *ModelsRevokeCertificateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelsRevokeCertificateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelsRevokeCertificateRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelsRevokeCertificateRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelsRevokeCertificateRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelsRevokeCertificateRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEffectiveDate

`func (o *ModelsRevokeCertificateRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *ModelsRevokeCertificateRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *ModelsRevokeCertificateRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *ModelsRevokeCertificateRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetCollectionId

`func (o *ModelsRevokeCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ModelsRevokeCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ModelsRevokeCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *ModelsRevokeCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *ModelsRevokeCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *ModelsRevokeCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


