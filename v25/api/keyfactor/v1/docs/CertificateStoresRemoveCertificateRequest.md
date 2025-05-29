# CertificateStoresRemoveCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStores** | [**[]CSSCMSDataModelModelsCertificateLocationSpecifier**](CSSCMSDataModelModelsCertificateLocationSpecifier.md) |  | 
**Schedule** | [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCertificateStoresRemoveCertificateRequest

`func NewCertificateStoresRemoveCertificateRequest(certificateStores []CSSCMSDataModelModelsCertificateLocationSpecifier, schedule KeyfactorCommonSchedulingKeyfactorSchedule, ) *CertificateStoresRemoveCertificateRequest`

NewCertificateStoresRemoveCertificateRequest instantiates a new CertificateStoresRemoveCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresRemoveCertificateRequestWithDefaults

`func NewCertificateStoresRemoveCertificateRequestWithDefaults() *CertificateStoresRemoveCertificateRequest`

NewCertificateStoresRemoveCertificateRequestWithDefaults instantiates a new CertificateStoresRemoveCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStores

`func (o *CertificateStoresRemoveCertificateRequest) GetCertificateStores() []CSSCMSDataModelModelsCertificateLocationSpecifier`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *CertificateStoresRemoveCertificateRequest) GetCertificateStoresOk() (*[]CSSCMSDataModelModelsCertificateLocationSpecifier, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *CertificateStoresRemoveCertificateRequest) SetCertificateStores(v []CSSCMSDataModelModelsCertificateLocationSpecifier)`

SetCertificateStores sets CertificateStores field to given value.


### GetSchedule

`func (o *CertificateStoresRemoveCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CertificateStoresRemoveCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CertificateStoresRemoveCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.


### GetCollectionId

`func (o *CertificateStoresRemoveCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CertificateStoresRemoveCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CertificateStoresRemoveCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CertificateStoresRemoveCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *CertificateStoresRemoveCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *CertificateStoresRemoveCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


