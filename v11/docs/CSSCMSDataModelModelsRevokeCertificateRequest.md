# CSSCMSDataModelModelsRevokeCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateIds** | Pointer to **[]int32** |  | [optional] 
**Reason** | Pointer to [**KeyfactorPKIEnumsRevokeCode**](KeyfactorPKIEnumsRevokeCode.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsRevokeCertificateRequest

`func NewCSSCMSDataModelModelsRevokeCertificateRequest() *CSSCMSDataModelModelsRevokeCertificateRequest`

NewCSSCMSDataModelModelsRevokeCertificateRequest instantiates a new CSSCMSDataModelModelsRevokeCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsRevokeCertificateRequestWithDefaults

`func NewCSSCMSDataModelModelsRevokeCertificateRequestWithDefaults() *CSSCMSDataModelModelsRevokeCertificateRequest`

NewCSSCMSDataModelModelsRevokeCertificateRequestWithDefaults instantiates a new CSSCMSDataModelModelsRevokeCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateIds

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetCertificateIds() []int32`

GetCertificateIds returns the CertificateIds field if non-nil, zero value otherwise.

### GetCertificateIdsOk

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetCertificateIdsOk() (*[]int32, bool)`

GetCertificateIdsOk returns a tuple with the CertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIds

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetCertificateIds(v []int32)`

SetCertificateIds sets CertificateIds field to given value.

### HasCertificateIds

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) HasCertificateIds() bool`

HasCertificateIds returns a boolean if a field has been set.

### SetCertificateIdsNil

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetCertificateIdsNil(b bool)`

 SetCertificateIdsNil sets the value for CertificateIds to be an explicit nil

### UnsetCertificateIds
`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) UnsetCertificateIds()`

UnsetCertificateIds ensures that no value is present for CertificateIds, not even an explicit nil
### GetReason

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetReason() KeyfactorPKIEnumsRevokeCode`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetReasonOk() (*KeyfactorPKIEnumsRevokeCode, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetReason(v KeyfactorPKIEnumsRevokeCode)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEffectiveDate

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetCollectionId

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *CSSCMSDataModelModelsRevokeCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


