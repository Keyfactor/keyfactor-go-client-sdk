# KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**CertificateStores** | Pointer to [**[]CSSCMSDataModelModelsCertificateStore**](CSSCMSDataModelModelsCertificateStore.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse() *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse`

NewKeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCertStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetCertificateStores() []CSSCMSDataModelModelsCertificateStore`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) GetCertificateStoresOk() (*[]CSSCMSDataModelModelsCertificateStore, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetCertificateStores(v []CSSCMSDataModelModelsCertificateStore)`

SetCertificateStores sets CertificateStores field to given value.

### HasCertificateStores

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) HasCertificateStores() bool`

HasCertificateStores returns a boolean if a field has been set.

### SetCertificateStoresNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) SetCertificateStoresNil(b bool)`

 SetCertificateStoresNil sets the value for CertificateStores to be an explicit nil

### UnsetCertificateStores
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoreContainersCertificateStoreContainerResponse) UnsetCertificateStores()`

UnsetCertificateStores ensures that no value is present for CertificateStores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


