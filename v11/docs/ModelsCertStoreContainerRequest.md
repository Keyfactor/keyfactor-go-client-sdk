# ModelsCertStoreContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**OverwriteSchedules** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**CertificateStores** | Pointer to [**[]ModelsCertificateStore**](ModelsCertificateStore.md) |  | [optional] 

## Methods

### NewModelsCertStoreContainerRequest

`func NewModelsCertStoreContainerRequest() *ModelsCertStoreContainerRequest`

NewModelsCertStoreContainerRequest instantiates a new ModelsCertStoreContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertStoreContainerRequestWithDefaults

`func NewModelsCertStoreContainerRequestWithDefaults() *ModelsCertStoreContainerRequest`

NewModelsCertStoreContainerRequestWithDefaults instantiates a new ModelsCertStoreContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertStoreContainerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertStoreContainerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertStoreContainerRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertStoreContainerRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ModelsCertStoreContainerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCertStoreContainerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCertStoreContainerRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsCertStoreContainerRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelsCertStoreContainerRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelsCertStoreContainerRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOverwriteSchedules

`func (o *ModelsCertStoreContainerRequest) GetOverwriteSchedules() bool`

GetOverwriteSchedules returns the OverwriteSchedules field if non-nil, zero value otherwise.

### GetOverwriteSchedulesOk

`func (o *ModelsCertStoreContainerRequest) GetOverwriteSchedulesOk() (*bool, bool)`

GetOverwriteSchedulesOk returns a tuple with the OverwriteSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteSchedules

`func (o *ModelsCertStoreContainerRequest) SetOverwriteSchedules(v bool)`

SetOverwriteSchedules sets OverwriteSchedules field to given value.

### HasOverwriteSchedules

`func (o *ModelsCertStoreContainerRequest) HasOverwriteSchedules() bool`

HasOverwriteSchedules returns a boolean if a field has been set.

### GetSchedule

`func (o *ModelsCertStoreContainerRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsCertStoreContainerRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsCertStoreContainerRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsCertStoreContainerRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCertStoreType

`func (o *ModelsCertStoreContainerRequest) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *ModelsCertStoreContainerRequest) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *ModelsCertStoreContainerRequest) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *ModelsCertStoreContainerRequest) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetCertificateStores

`func (o *ModelsCertStoreContainerRequest) GetCertificateStores() []ModelsCertificateStore`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *ModelsCertStoreContainerRequest) GetCertificateStoresOk() (*[]ModelsCertificateStore, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *ModelsCertStoreContainerRequest) SetCertificateStores(v []ModelsCertificateStore)`

SetCertificateStores sets CertificateStores field to given value.

### HasCertificateStores

`func (o *ModelsCertStoreContainerRequest) HasCertificateStores() bool`

HasCertificateStores returns a boolean if a field has been set.

### SetCertificateStoresNil

`func (o *ModelsCertStoreContainerRequest) SetCertificateStoresNil(b bool)`

 SetCertificateStoresNil sets the value for CertificateStores to be an explicit nil

### UnsetCertificateStores
`func (o *ModelsCertStoreContainerRequest) UnsetCertificateStores()`

UnsetCertificateStores ensures that no value is present for CertificateStores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


