# KeyfactorApiModelsCertificateStoresAddCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | **int32** |  | 
**CertificateStores** | [**[]ModelsCertificateStoreEntry**](ModelsCertificateStoreEntry.md) |  | 
**Schedule** | [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | 
**CollectionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresAddCertificateRequest

`func NewKeyfactorApiModelsCertificateStoresAddCertificateRequest(certificateId int32, certificateStores []ModelsCertificateStoreEntry, schedule KeyfactorCommonSchedulingKeyfactorSchedule, ) *KeyfactorApiModelsCertificateStoresAddCertificateRequest`

NewKeyfactorApiModelsCertificateStoresAddCertificateRequest instantiates a new KeyfactorApiModelsCertificateStoresAddCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresAddCertificateRequestWithDefaults

`func NewKeyfactorApiModelsCertificateStoresAddCertificateRequestWithDefaults() *KeyfactorApiModelsCertificateStoresAddCertificateRequest`

NewKeyfactorApiModelsCertificateStoresAddCertificateRequestWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresAddCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.


### GetCertificateStores

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateStores() []ModelsCertificateStoreEntry`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateStoresOk() (*[]ModelsCertificateStoreEntry, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCertificateStores(v []ModelsCertificateStoreEntry)`

SetCertificateStores sets CertificateStores field to given value.


### GetSchedule

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.


### GetCollectionId

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *KeyfactorApiModelsCertificateStoresAddCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


