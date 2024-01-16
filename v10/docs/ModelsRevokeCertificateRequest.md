# ModelsRevokeCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateIds** | Pointer to **[]int64** |  | [optional] 
**Reason** | Pointer to **int64** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**CollectionId** | Pointer to **int64** |  | [optional] 

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

`func (o *ModelsRevokeCertificateRequest) GetCertificateIds() []int64`

GetCertificateIds returns the CertificateIds field if non-nil, zero value otherwise.

### GetCertificateIdsOk

`func (o *ModelsRevokeCertificateRequest) GetCertificateIdsOk() (*[]int64, bool)`

GetCertificateIdsOk returns a tuple with the CertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIds

`func (o *ModelsRevokeCertificateRequest) SetCertificateIds(v []int64)`

SetCertificateIds sets CertificateIds field to given value.

### HasCertificateIds

`func (o *ModelsRevokeCertificateRequest) HasCertificateIds() bool`

HasCertificateIds returns a boolean if a field has been set.

### GetReason

`func (o *ModelsRevokeCertificateRequest) GetReason() int64`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ModelsRevokeCertificateRequest) GetReasonOk() (*int64, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ModelsRevokeCertificateRequest) SetReason(v int64)`

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

`func (o *ModelsRevokeCertificateRequest) GetCollectionId() int64`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *ModelsRevokeCertificateRequest) GetCollectionIdOk() (*int64, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *ModelsRevokeCertificateRequest) SetCollectionId(v int64)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *ModelsRevokeCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


