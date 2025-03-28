# CertificateStoresCertificateStoreApproveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **NullableInt32** |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**InventorySchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewCertificateStoresCertificateStoreApproveRequest

`func NewCertificateStoresCertificateStoreApproveRequest() *CertificateStoresCertificateStoreApproveRequest`

NewCertificateStoresCertificateStoreApproveRequest instantiates a new CertificateStoresCertificateStoreApproveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresCertificateStoreApproveRequestWithDefaults

`func NewCertificateStoresCertificateStoreApproveRequestWithDefaults() *CertificateStoresCertificateStoreApproveRequest`

NewCertificateStoresCertificateStoreApproveRequestWithDefaults instantiates a new CertificateStoresCertificateStoreApproveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificateStoresCertificateStoreApproveRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificateStoresCertificateStoreApproveRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificateStoresCertificateStoreApproveRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CertificateStoresCertificateStoreApproveRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContainerId

`func (o *CertificateStoresCertificateStoreApproveRequest) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *CertificateStoresCertificateStoreApproveRequest) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *CertificateStoresCertificateStoreApproveRequest) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *CertificateStoresCertificateStoreApproveRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *CertificateStoresCertificateStoreApproveRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *CertificateStoresCertificateStoreApproveRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetCertStoreType

`func (o *CertificateStoresCertificateStoreApproveRequest) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *CertificateStoresCertificateStoreApproveRequest) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *CertificateStoresCertificateStoreApproveRequest) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *CertificateStoresCertificateStoreApproveRequest) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetProperties

`func (o *CertificateStoresCertificateStoreApproveRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CertificateStoresCertificateStoreApproveRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CertificateStoresCertificateStoreApproveRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CertificateStoresCertificateStoreApproveRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *CertificateStoresCertificateStoreApproveRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *CertificateStoresCertificateStoreApproveRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetInventorySchedule

`func (o *CertificateStoresCertificateStoreApproveRequest) GetInventorySchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetInventorySchedule returns the InventorySchedule field if non-nil, zero value otherwise.

### GetInventoryScheduleOk

`func (o *CertificateStoresCertificateStoreApproveRequest) GetInventoryScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetInventoryScheduleOk returns a tuple with the InventorySchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventorySchedule

`func (o *CertificateStoresCertificateStoreApproveRequest) SetInventorySchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetInventorySchedule sets InventorySchedule field to given value.

### HasInventorySchedule

`func (o *CertificateStoresCertificateStoreApproveRequest) HasInventorySchedule() bool`

HasInventorySchedule returns a boolean if a field has been set.

### GetPassword

`func (o *CertificateStoresCertificateStoreApproveRequest) GetPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CertificateStoresCertificateStoreApproveRequest) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CertificateStoresCertificateStoreApproveRequest) SetPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CertificateStoresCertificateStoreApproveRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


