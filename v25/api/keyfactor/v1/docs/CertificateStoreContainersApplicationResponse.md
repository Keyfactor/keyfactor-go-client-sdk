# CertificateStoreContainersApplicationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateStores** | Pointer to [**[]CertificateStoresCertificateStoreApplicationResponse**](CertificateStoresCertificateStoreApplicationResponse.md) |  | [optional] 

## Methods

### NewCertificateStoreContainersApplicationResponse

`func NewCertificateStoreContainersApplicationResponse() *CertificateStoreContainersApplicationResponse`

NewCertificateStoreContainersApplicationResponse instantiates a new CertificateStoreContainersApplicationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoreContainersApplicationResponseWithDefaults

`func NewCertificateStoreContainersApplicationResponseWithDefaults() *CertificateStoreContainersApplicationResponse`

NewCertificateStoreContainersApplicationResponseWithDefaults instantiates a new CertificateStoreContainersApplicationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateStoreContainersApplicationResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateStoreContainersApplicationResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateStoreContainersApplicationResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateStoreContainersApplicationResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CertificateStoreContainersApplicationResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoreContainersApplicationResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoreContainersApplicationResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStoreContainersApplicationResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateStoreContainersApplicationResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateStoreContainersApplicationResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchedule

`func (o *CertificateStoreContainersApplicationResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CertificateStoreContainersApplicationResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CertificateStoreContainersApplicationResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CertificateStoreContainersApplicationResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCertificateStores

`func (o *CertificateStoreContainersApplicationResponse) GetCertificateStores() []CertificateStoresCertificateStoreApplicationResponse`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *CertificateStoreContainersApplicationResponse) GetCertificateStoresOk() (*[]CertificateStoresCertificateStoreApplicationResponse, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *CertificateStoreContainersApplicationResponse) SetCertificateStores(v []CertificateStoresCertificateStoreApplicationResponse)`

SetCertificateStores sets CertificateStores field to given value.

### HasCertificateStores

`func (o *CertificateStoreContainersApplicationResponse) HasCertificateStores() bool`

HasCertificateStores returns a boolean if a field has been set.

### SetCertificateStoresNil

`func (o *CertificateStoreContainersApplicationResponse) SetCertificateStoresNil(b bool)`

 SetCertificateStoresNil sets the value for CertificateStores to be an explicit nil

### UnsetCertificateStores
`func (o *CertificateStoreContainersApplicationResponse) UnsetCertificateStores()`

UnsetCertificateStores ensures that no value is present for CertificateStores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


