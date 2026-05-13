# CertificateStoreContainersCertificateStoreContainerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateStores** | Pointer to [**[]CertificateStoresCertificateStoreResponse**](CertificateStoresCertificateStoreResponse.md) |  | [optional] 

## Methods

### NewCertificateStoreContainersCertificateStoreContainerResponse

`func NewCertificateStoreContainersCertificateStoreContainerResponse() *CertificateStoreContainersCertificateStoreContainerResponse`

NewCertificateStoreContainersCertificateStoreContainerResponse instantiates a new CertificateStoreContainersCertificateStoreContainerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoreContainersCertificateStoreContainerResponseWithDefaults

`func NewCertificateStoreContainersCertificateStoreContainerResponseWithDefaults() *CertificateStoreContainersCertificateStoreContainerResponse`

NewCertificateStoreContainersCertificateStoreContainerResponseWithDefaults instantiates a new CertificateStoreContainersCertificateStoreContainerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CertificateStoreContainersCertificateStoreContainerResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetSchedule

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetCertificateStores

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetCertificateStores() []CertificateStoresCertificateStoreResponse`

GetCertificateStores returns the CertificateStores field if non-nil, zero value otherwise.

### GetCertificateStoresOk

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) GetCertificateStoresOk() (*[]CertificateStoresCertificateStoreResponse, bool)`

GetCertificateStoresOk returns a tuple with the CertificateStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateStores

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) SetCertificateStores(v []CertificateStoresCertificateStoreResponse)`

SetCertificateStores sets CertificateStores field to given value.

### HasCertificateStores

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) HasCertificateStores() bool`

HasCertificateStores returns a boolean if a field has been set.

### SetCertificateStoresNil

`func (o *CertificateStoreContainersCertificateStoreContainerResponse) SetCertificateStoresNil(b bool)`

 SetCertificateStoresNil sets the value for CertificateStores to be an explicit nil

### UnsetCertificateStores
`func (o *CertificateStoreContainersCertificateStoreContainerResponse) UnsetCertificateStores()`

UnsetCertificateStores ensures that no value is present for CertificateStores, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


