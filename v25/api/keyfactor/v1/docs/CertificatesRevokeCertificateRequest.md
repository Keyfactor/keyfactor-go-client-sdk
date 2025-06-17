# CertificatesRevokeCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateIds** | Pointer to **[]int32** |  | [optional] 
**Reason** | Pointer to [**KeyfactorPKIEnumsRevokeCode**](KeyfactorPKIEnumsRevokeCode.md) |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 
**EffectiveDate** | Pointer to **time.Time** |  | [optional] 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 
**PublishCRL** | Pointer to **bool** |  | [optional] 

## Methods

### NewCertificatesRevokeCertificateRequest

`func NewCertificatesRevokeCertificateRequest() *CertificatesRevokeCertificateRequest`

NewCertificatesRevokeCertificateRequest instantiates a new CertificatesRevokeCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesRevokeCertificateRequestWithDefaults

`func NewCertificatesRevokeCertificateRequestWithDefaults() *CertificatesRevokeCertificateRequest`

NewCertificatesRevokeCertificateRequestWithDefaults instantiates a new CertificatesRevokeCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateIds

`func (o *CertificatesRevokeCertificateRequest) GetCertificateIds() []int32`

GetCertificateIds returns the CertificateIds field if non-nil, zero value otherwise.

### GetCertificateIdsOk

`func (o *CertificatesRevokeCertificateRequest) GetCertificateIdsOk() (*[]int32, bool)`

GetCertificateIdsOk returns a tuple with the CertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIds

`func (o *CertificatesRevokeCertificateRequest) SetCertificateIds(v []int32)`

SetCertificateIds sets CertificateIds field to given value.

### HasCertificateIds

`func (o *CertificatesRevokeCertificateRequest) HasCertificateIds() bool`

HasCertificateIds returns a boolean if a field has been set.

### SetCertificateIdsNil

`func (o *CertificatesRevokeCertificateRequest) SetCertificateIdsNil(b bool)`

 SetCertificateIdsNil sets the value for CertificateIds to be an explicit nil

### UnsetCertificateIds
`func (o *CertificatesRevokeCertificateRequest) UnsetCertificateIds()`

UnsetCertificateIds ensures that no value is present for CertificateIds, not even an explicit nil
### GetReason

`func (o *CertificatesRevokeCertificateRequest) GetReason() KeyfactorPKIEnumsRevokeCode`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CertificatesRevokeCertificateRequest) GetReasonOk() (*KeyfactorPKIEnumsRevokeCode, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CertificatesRevokeCertificateRequest) SetReason(v KeyfactorPKIEnumsRevokeCode)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CertificatesRevokeCertificateRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetComment

`func (o *CertificatesRevokeCertificateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CertificatesRevokeCertificateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CertificatesRevokeCertificateRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CertificatesRevokeCertificateRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *CertificatesRevokeCertificateRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *CertificatesRevokeCertificateRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil
### GetEffectiveDate

`func (o *CertificatesRevokeCertificateRequest) GetEffectiveDate() time.Time`

GetEffectiveDate returns the EffectiveDate field if non-nil, zero value otherwise.

### GetEffectiveDateOk

`func (o *CertificatesRevokeCertificateRequest) GetEffectiveDateOk() (*time.Time, bool)`

GetEffectiveDateOk returns a tuple with the EffectiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDate

`func (o *CertificatesRevokeCertificateRequest) SetEffectiveDate(v time.Time)`

SetEffectiveDate sets EffectiveDate field to given value.

### HasEffectiveDate

`func (o *CertificatesRevokeCertificateRequest) HasEffectiveDate() bool`

HasEffectiveDate returns a boolean if a field has been set.

### GetCollectionId

`func (o *CertificatesRevokeCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CertificatesRevokeCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CertificatesRevokeCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CertificatesRevokeCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *CertificatesRevokeCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *CertificatesRevokeCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil
### GetPublishCRL

`func (o *CertificatesRevokeCertificateRequest) GetPublishCRL() bool`

GetPublishCRL returns the PublishCRL field if non-nil, zero value otherwise.

### GetPublishCRLOk

`func (o *CertificatesRevokeCertificateRequest) GetPublishCRLOk() (*bool, bool)`

GetPublishCRLOk returns a tuple with the PublishCRL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishCRL

`func (o *CertificatesRevokeCertificateRequest) SetPublishCRL(v bool)`

SetPublishCRL sets PublishCRL field to given value.

### HasPublishCRL

`func (o *CertificatesRevokeCertificateRequest) HasPublishCRL() bool`

HasPublishCRL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


