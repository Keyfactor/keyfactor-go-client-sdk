# KeyfactorApiModelsCertificateStoresRemoveCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStores** | [**[]ModelsCertificateLocationSpecifier**](ModelsCertificateLocationSpecifier.md) |  | 
**Schedule** | [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | 
**CollectionId** | Pointer to **int32** |  | [optional] 

## Methods

### NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequest

`func NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequest(certificateStores []ModelsCertificateLocationSpecifier, schedule KeyfactorCommonSchedulingKeyfactorSchedule, ) *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest`

NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequest instantiates a new KeyfactorApiModelsCertificateStoresRemoveCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults

`func NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults() *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest`

NewKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults instantiates a new KeyfactorApiModelsCertificateStoresRemoveCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStores

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCertificateStores() []ModelsCertificateLocationSpecifier`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCertificateStoresOk() (*[]ModelsCertificateLocationSpecifier, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCertificateStores(v []ModelsCertificateLocationSpecifier)`

SetCertificateStores sets CertificateStores field to given value.


### GetSchedule

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.


### GetCollectionId

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *KeyfactorApiModelsCertificateStoresRemoveCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


