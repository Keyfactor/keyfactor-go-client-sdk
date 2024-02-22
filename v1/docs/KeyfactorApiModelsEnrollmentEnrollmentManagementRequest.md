# KeyfactorApiModelsEnrollmentEnrollmentManagementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stores** | Pointer to [**[]KeyfactorApiModelsEnrollmentManagementStoreRequest**](KeyfactorApiModelsEnrollmentManagementStoreRequest.md) | The stores to add the certificate to. Values in this collection will take precedence over values in {Models.Enrollment.SpecificEnrollmentManagementRequest.StoreTypes}. | [optional] 
**StoreIds** | Pointer to **[]string** |  | [optional] 
**StoreTypes** | Pointer to [**[]KeyfactorApiModelsEnrollmentManagementStoreTypeRequest**](KeyfactorApiModelsEnrollmentManagementStoreTypeRequest.md) |  | [optional] 
**CertificateId** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **int32** |  | [optional] 
**Password** | **string** |  | 
**JobTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequest

`func NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequest(password string, ) *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest`

NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequest instantiates a new KeyfactorApiModelsEnrollmentEnrollmentManagementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequestWithDefaults

`func NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequestWithDefaults() *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest`

NewKeyfactorApiModelsEnrollmentEnrollmentManagementRequestWithDefaults instantiates a new KeyfactorApiModelsEnrollmentEnrollmentManagementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStores

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStores() []KeyfactorApiModelsEnrollmentManagementStoreRequest`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoresOk() (*[]KeyfactorApiModelsEnrollmentManagementStoreRequest, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetStores(v []KeyfactorApiModelsEnrollmentManagementStoreRequest)`

SetStores sets Stores field to given value.

### HasStores

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasStores() bool`

HasStores returns a boolean if a field has been set.

### GetStoreIds

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreIds() []string`

GetStoreIds returns the StoreIds field if non-nil, zero value otherwise.

### GetStoreIdsOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreIdsOk() (*[]string, bool)`

GetStoreIdsOk returns a tuple with the StoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIds

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetStoreIds(v []string)`

SetStoreIds sets StoreIds field to given value.

### HasStoreIds

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasStoreIds() bool`

HasStoreIds returns a boolean if a field has been set.

### GetStoreTypes

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreTypes() []KeyfactorApiModelsEnrollmentManagementStoreTypeRequest`

GetStoreTypes returns the StoreTypes field if non-nil, zero value otherwise.

### GetStoreTypesOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetStoreTypesOk() (*[]KeyfactorApiModelsEnrollmentManagementStoreTypeRequest, bool)`

GetStoreTypesOk returns a tuple with the StoreTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypes

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetStoreTypes(v []KeyfactorApiModelsEnrollmentManagementStoreTypeRequest)`

SetStoreTypes sets StoreTypes field to given value.

### HasStoreTypes

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasStoreTypes() bool`

HasStoreTypes returns a boolean if a field has been set.

### GetCertificateId

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetRequestId

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetPassword

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetJobTime

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetJobTime() time.Time`

GetJobTime returns the JobTime field if non-nil, zero value otherwise.

### GetJobTimeOk

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) GetJobTimeOk() (*time.Time, bool)`

GetJobTimeOk returns a tuple with the JobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTime

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) SetJobTime(v time.Time)`

SetJobTime sets JobTime field to given value.

### HasJobTime

`func (o *KeyfactorApiModelsEnrollmentEnrollmentManagementRequest) HasJobTime() bool`

HasJobTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


