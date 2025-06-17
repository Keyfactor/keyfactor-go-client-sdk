# EnrollmentEnrollmentManagementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stores** | Pointer to [**[]EnrollmentManagementStoreRequest**](EnrollmentManagementStoreRequest.md) | The stores to add the certificate to. Values in this collection will take precedence over values in CSS.CMS.Data.Model.Models.Enrollment.SpecificEnrollmentManagementRequest.StoreTypes. | [optional] 
**StoreIds** | Pointer to **[]string** |  | [optional] 
**StoreTypes** | Pointer to [**[]EnrollmentManagementStoreTypeRequest**](EnrollmentManagementStoreTypeRequest.md) |  | [optional] 
**CertificateId** | Pointer to **int32** |  | [optional] 
**RequestId** | Pointer to **int32** |  | [optional] 
**Password** | **string** |  | 
**JobTime** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewEnrollmentEnrollmentManagementRequest

`func NewEnrollmentEnrollmentManagementRequest(password string, ) *EnrollmentEnrollmentManagementRequest`

NewEnrollmentEnrollmentManagementRequest instantiates a new EnrollmentEnrollmentManagementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollmentEnrollmentManagementRequestWithDefaults

`func NewEnrollmentEnrollmentManagementRequestWithDefaults() *EnrollmentEnrollmentManagementRequest`

NewEnrollmentEnrollmentManagementRequestWithDefaults instantiates a new EnrollmentEnrollmentManagementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStores

`func (o *EnrollmentEnrollmentManagementRequest) GetStores() []EnrollmentManagementStoreRequest`

GetStores returns the Stores field if non-nil, zero value otherwise.

### GetStoresOk

`func (o *EnrollmentEnrollmentManagementRequest) GetStoresOk() (*[]EnrollmentManagementStoreRequest, bool)`

GetStoresOk returns a tuple with the Stores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStores

`func (o *EnrollmentEnrollmentManagementRequest) SetStores(v []EnrollmentManagementStoreRequest)`

SetStores sets Stores field to given value.

### HasStores

`func (o *EnrollmentEnrollmentManagementRequest) HasStores() bool`

HasStores returns a boolean if a field has been set.

### SetStoresNil

`func (o *EnrollmentEnrollmentManagementRequest) SetStoresNil(b bool)`

 SetStoresNil sets the value for Stores to be an explicit nil

### UnsetStores
`func (o *EnrollmentEnrollmentManagementRequest) UnsetStores()`

UnsetStores ensures that no value is present for Stores, not even an explicit nil
### GetStoreIds

`func (o *EnrollmentEnrollmentManagementRequest) GetStoreIds() []string`

GetStoreIds returns the StoreIds field if non-nil, zero value otherwise.

### GetStoreIdsOk

`func (o *EnrollmentEnrollmentManagementRequest) GetStoreIdsOk() (*[]string, bool)`

GetStoreIdsOk returns a tuple with the StoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIds

`func (o *EnrollmentEnrollmentManagementRequest) SetStoreIds(v []string)`

SetStoreIds sets StoreIds field to given value.

### HasStoreIds

`func (o *EnrollmentEnrollmentManagementRequest) HasStoreIds() bool`

HasStoreIds returns a boolean if a field has been set.

### SetStoreIdsNil

`func (o *EnrollmentEnrollmentManagementRequest) SetStoreIdsNil(b bool)`

 SetStoreIdsNil sets the value for StoreIds to be an explicit nil

### UnsetStoreIds
`func (o *EnrollmentEnrollmentManagementRequest) UnsetStoreIds()`

UnsetStoreIds ensures that no value is present for StoreIds, not even an explicit nil
### GetStoreTypes

`func (o *EnrollmentEnrollmentManagementRequest) GetStoreTypes() []EnrollmentManagementStoreTypeRequest`

GetStoreTypes returns the StoreTypes field if non-nil, zero value otherwise.

### GetStoreTypesOk

`func (o *EnrollmentEnrollmentManagementRequest) GetStoreTypesOk() (*[]EnrollmentManagementStoreTypeRequest, bool)`

GetStoreTypesOk returns a tuple with the StoreTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypes

`func (o *EnrollmentEnrollmentManagementRequest) SetStoreTypes(v []EnrollmentManagementStoreTypeRequest)`

SetStoreTypes sets StoreTypes field to given value.

### HasStoreTypes

`func (o *EnrollmentEnrollmentManagementRequest) HasStoreTypes() bool`

HasStoreTypes returns a boolean if a field has been set.

### SetStoreTypesNil

`func (o *EnrollmentEnrollmentManagementRequest) SetStoreTypesNil(b bool)`

 SetStoreTypesNil sets the value for StoreTypes to be an explicit nil

### UnsetStoreTypes
`func (o *EnrollmentEnrollmentManagementRequest) UnsetStoreTypes()`

UnsetStoreTypes ensures that no value is present for StoreTypes, not even an explicit nil
### GetCertificateId

`func (o *EnrollmentEnrollmentManagementRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *EnrollmentEnrollmentManagementRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *EnrollmentEnrollmentManagementRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *EnrollmentEnrollmentManagementRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetRequestId

`func (o *EnrollmentEnrollmentManagementRequest) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *EnrollmentEnrollmentManagementRequest) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *EnrollmentEnrollmentManagementRequest) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *EnrollmentEnrollmentManagementRequest) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetPassword

`func (o *EnrollmentEnrollmentManagementRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *EnrollmentEnrollmentManagementRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *EnrollmentEnrollmentManagementRequest) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetJobTime

`func (o *EnrollmentEnrollmentManagementRequest) GetJobTime() time.Time`

GetJobTime returns the JobTime field if non-nil, zero value otherwise.

### GetJobTimeOk

`func (o *EnrollmentEnrollmentManagementRequest) GetJobTimeOk() (*time.Time, bool)`

GetJobTimeOk returns a tuple with the JobTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTime

`func (o *EnrollmentEnrollmentManagementRequest) SetJobTime(v time.Time)`

SetJobTime sets JobTime field to given value.

### HasJobTime

`func (o *EnrollmentEnrollmentManagementRequest) HasJobTime() bool`

HasJobTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


