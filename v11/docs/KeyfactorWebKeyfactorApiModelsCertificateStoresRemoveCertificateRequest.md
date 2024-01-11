# KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateStores** | [**[]ModelsCertificateLocationSpecifier**](ModelsCertificateLocationSpecifier.md) |  | 
**Schedule** | [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | 
**CollectionId** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest(certificateStores []ModelsCertificateLocationSpecifier, schedule KeyfactorCommonSchedulingKeyfactorSchedule, ) *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCertificateStores() []ModelsCertificateLocationSpecifier`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCertificateStoresOk() (*[]ModelsCertificateLocationSpecifier, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCertificateStores(v []ModelsCertificateLocationSpecifier)`

SetCertificateStores sets CertificateStores field to given value.


### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.


### GetCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCollectionId() int32`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) GetCollectionIdOk() (*int32, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCollectionId(v int32)`

SetCollectionId sets CollectionId field to given value.

### HasCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) HasCollectionId() bool`

HasCollectionId returns a boolean if a field has been set.

### SetCollectionIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) SetCollectionIdNil(b bool)`

 SetCollectionIdNil sets the value for CollectionId to be an explicit nil

### UnsetCollectionId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresRemoveCertificateRequest) UnsetCollectionId()`

UnsetCollectionId ensures that no value is present for CollectionId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


