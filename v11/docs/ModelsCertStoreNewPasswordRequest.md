# ModelsCertStoreNewPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertStoreId** | **string** |  | 
**NewPassword** | **interface{}** |  | 

## Methods

### NewModelsCertStoreNewPasswordRequest

`func NewModelsCertStoreNewPasswordRequest(certStoreId string, newPassword interface{}, ) *ModelsCertStoreNewPasswordRequest`

NewModelsCertStoreNewPasswordRequest instantiates a new ModelsCertStoreNewPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertStoreNewPasswordRequestWithDefaults

`func NewModelsCertStoreNewPasswordRequestWithDefaults() *ModelsCertStoreNewPasswordRequest`

NewModelsCertStoreNewPasswordRequestWithDefaults instantiates a new ModelsCertStoreNewPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertStoreId

`func (o *ModelsCertStoreNewPasswordRequest) GetCertStoreId() string`

GetCertStoreId returns the CertStoreId field if non-nil, zero value otherwise.

### GetCertStoreIdOk

`func (o *ModelsCertStoreNewPasswordRequest) GetCertStoreIdOk() (*string, bool)`

GetCertStoreIdOk returns a tuple with the CertStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreId

`func (o *ModelsCertStoreNewPasswordRequest) SetCertStoreId(v string)`

SetCertStoreId sets CertStoreId field to given value.


### GetNewPassword

`func (o *ModelsCertStoreNewPasswordRequest) GetNewPassword() interface{}`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ModelsCertStoreNewPasswordRequest) GetNewPasswordOk() (*interface{}, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ModelsCertStoreNewPasswordRequest) SetNewPassword(v interface{})`

SetNewPassword sets NewPassword field to given value.


### SetNewPasswordNil

`func (o *ModelsCertStoreNewPasswordRequest) SetNewPasswordNil(b bool)`

 SetNewPasswordNil sets the value for NewPassword to be an explicit nil

### UnsetNewPassword
`func (o *ModelsCertStoreNewPasswordRequest) UnsetNewPassword()`

UnsetNewPassword ensures that no value is present for NewPassword, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


