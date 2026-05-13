# CSSCMSDataModelModelsCertStoreContainerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**OverwriteSchedules** | Pointer to **bool** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateStores** | Pointer to [**[]CSSCMSDataModelModelsCertificateStore**](CSSCMSDataModelModelsCertificateStore.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertStoreContainerRequest

`func NewCSSCMSDataModelModelsCertStoreContainerRequest(name string, ) *CSSCMSDataModelModelsCertStoreContainerRequest`

NewCSSCMSDataModelModelsCertStoreContainerRequest instantiates a new CSSCMSDataModelModelsCertStoreContainerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertStoreContainerRequestWithDefaults

`func NewCSSCMSDataModelModelsCertStoreContainerRequestWithDefaults() *CSSCMSDataModelModelsCertStoreContainerRequest`

NewCSSCMSDataModelModelsCertStoreContainerRequestWithDefaults instantiates a new CSSCMSDataModelModelsCertStoreContainerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOverwriteSchedules

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetOverwriteSchedules() bool`

GetOverwriteSchedules returns the OverwriteSchedules field if non-nil, zero value otherwise.

### GetOverwriteSchedulesOk

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetOverwriteSchedulesOk() (*bool, bool)`

GetOverwriteSchedulesOk returns a tuple with the OverwriteSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwriteSchedules

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) SetOverwriteSchedules(v bool)`

SetOverwriteSchedules sets OverwriteSchedules field to given value.

### HasOverwriteSchedules

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) HasOverwriteSchedules() bool`

HasOverwriteSchedules returns a boolean if a field has been set.

### GetSchedule

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCertificateStores

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetCertificateStores() []CSSCMSDataModelModelsCertificateStore`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) GetCertificateStoresOk() (*[]CSSCMSDataModelModelsCertificateStore, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) SetCertificateStores(v []CSSCMSDataModelModelsCertificateStore)`

SetCertificateStores sets CertificateStores field to given value.

### HasCertificateStores

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) HasCertificateStores() bool`

HasCertificateStores returns a boolean if a field has been set.

### SetCertificateStoresNil

`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) SetCertificateStoresNil(b bool)`

 SetCertificateStoresNil sets the value for CertificateStores to be an explicit nil

### UnsetCertificateStores
`func (o *CSSCMSDataModelModelsCertStoreContainerRequest) UnsetCertificateStores()`

UnsetCertificateStores ensures that no value is present for CertificateStores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


