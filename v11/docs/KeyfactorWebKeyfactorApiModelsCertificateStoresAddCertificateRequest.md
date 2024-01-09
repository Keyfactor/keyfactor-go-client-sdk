# KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | **int32** |  | 
**CertificateStores** | [**[]CSSCMSDataModelModelsCertificateStoreEntry**](CSSCMSDataModelModelsCertificateStoreEntry.md) |  | 
**Schedule** | [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest(certificateId int32, certificateStores []CSSCMSDataModelModelsCertificateStoreEntry, schedule KeyfactorCommonSchedulingKeyfactorSchedule, ) *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.


### GetCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateStores() []CSSCMSDataModelModelsCertificateStoreEntry`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCertificateStoresOk() (*[]CSSCMSDataModelModelsCertificateStoreEntry, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCertificateStores(v []CSSCMSDataModelModelsCertificateStoreEntry)`

SetCertificateStores sets CertificateStores field to given value.


### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.


### GetCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresAddCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


