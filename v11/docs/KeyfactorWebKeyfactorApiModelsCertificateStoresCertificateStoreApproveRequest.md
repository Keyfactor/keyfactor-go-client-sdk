# KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**ContainerId** | Pointer to **NullableInt32** |  | [optional] 
**CertStoreType** | Pointer to **int32** |  | [optional] 
**Properties** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest() *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest`

NewKeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContainerId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetContainerId() int32`

GetContainerId returns the ContainerId field if non-nil, zero value otherwise.

### GetContainerIdOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetContainerIdOk() (*int32, bool)`

GetContainerIdOk returns a tuple with the ContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetContainerId(v int32)`

SetContainerId sets ContainerId field to given value.

### HasContainerId

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) HasContainerId() bool`

HasContainerId returns a boolean if a field has been set.

### SetContainerIdNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetContainerIdNil(b bool)`

 SetContainerIdNil sets the value for ContainerId to be an explicit nil

### UnsetContainerId
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) UnsetContainerId()`

UnsetContainerId ensures that no value is present for ContainerId, not even an explicit nil
### GetCertStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetCertStoreType() int32`

GetCertStoreType returns the CertStoreType field if non-nil, zero value otherwise.

### GetCertStoreTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetCertStoreTypeOk() (*int32, bool)`

GetCertStoreTypeOk returns a tuple with the CertStoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetCertStoreType(v int32)`

SetCertStoreType sets CertStoreType field to given value.

### HasCertStoreType

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) HasCertStoreType() bool`

HasCertStoreType returns a boolean if a field has been set.

### GetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetProperties() string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetPropertiesOk() (*string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetProperties(v string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### SetPropertiesNil

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetPropertiesNil(b bool)`

 SetPropertiesNil sets the value for Properties to be an explicit nil

### UnsetProperties
`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) UnsetProperties()`

UnsetProperties ensures that no value is present for Properties, not even an explicit nil
### GetPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetPassword() ModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) SetPassword(v ModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *KeyfactorWebKeyfactorApiModelsCertificateStoresCertificateStoreApproveRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


