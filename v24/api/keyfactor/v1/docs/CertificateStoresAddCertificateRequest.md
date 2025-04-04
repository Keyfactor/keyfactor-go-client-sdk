# CertificateStoresAddCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | **int32** |  | 
**CertificateStores** | [**[]CSSCMSDataModelModelsCertificateStoreEntry**](CSSCMSDataModelModelsCertificateStoreEntry.md) |  | 
**Schedule** | [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCertificateStoresAddCertificateRequest

`func NewCertificateStoresAddCertificateRequest(certificateId int32, certificateStores []CSSCMSDataModelModelsCertificateStoreEntry, schedule KeyfactorCommonSchedulingKeyfactorSchedule, ) *CertificateStoresAddCertificateRequest`

NewCertificateStoresAddCertificateRequest instantiates a new CertificateStoresAddCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresAddCertificateRequestWithDefaults

`func NewCertificateStoresAddCertificateRequestWithDefaults() *CertificateStoresAddCertificateRequest`

NewCertificateStoresAddCertificateRequestWithDefaults instantiates a new CertificateStoresAddCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *CertificateStoresAddCertificateRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *CertificateStoresAddCertificateRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *CertificateStoresAddCertificateRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.


### GetCertificateStores

`func (o *CertificateStoresAddCertificateRequest) GetCertificateStores() []CSSCMSDataModelModelsCertificateStoreEntry`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *CertificateStoresAddCertificateRequest) GetCertificateStoresOk() (*[]CSSCMSDataModelModelsCertificateStoreEntry, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *CertificateStoresAddCertificateRequest) SetCertificateStores(v []CSSCMSDataModelModelsCertificateStoreEntry)`

SetCertificateStores sets CertificateStores field to given value.


### GetSchedule

`func (o *CertificateStoresAddCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CertificateStoresAddCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CertificateStoresAddCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.


### GetCollectionId

`func (o *CertificateStoresAddCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *CertificateStoresAddCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *CertificateStoresAddCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *CertificateStoresAddCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *CertificateStoresAddCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *CertificateStoresAddCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


