# ModelsCertificateStoreUpdateServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Username** | [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | 
**Password** | [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | 
**UseSSL** | **bool** |  | 
**Name** | **string** |  | 
**Container** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewModelsCertificateStoreUpdateServerRequest

`func NewModelsCertificateStoreUpdateServerRequest(id int32, username ModelsKeyfactorAPISecret, password ModelsKeyfactorAPISecret, useSSL bool, name string, ) *ModelsCertificateStoreUpdateServerRequest`

NewModelsCertificateStoreUpdateServerRequest instantiates a new ModelsCertificateStoreUpdateServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreUpdateServerRequestWithDefaults

`func NewModelsCertificateStoreUpdateServerRequestWithDefaults() *ModelsCertificateStoreUpdateServerRequest`

NewModelsCertificateStoreUpdateServerRequestWithDefaults instantiates a new ModelsCertificateStoreUpdateServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateStoreUpdateServerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateStoreUpdateServerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateStoreUpdateServerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *ModelsCertificateStoreUpdateServerRequest) GetUsername() ModelsKeyfactorAPISecret`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsCertificateStoreUpdateServerRequest) GetUsernameOk() (*ModelsKeyfactorAPISecret, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsCertificateStoreUpdateServerRequest) SetUsername(v ModelsKeyfactorAPISecret)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ModelsCertificateStoreUpdateServerRequest) GetPassword() ModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateStoreUpdateServerRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateStoreUpdateServerRequest) SetPassword(v ModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.


### GetUseSSL

`func (o *ModelsCertificateStoreUpdateServerRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *ModelsCertificateStoreUpdateServerRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *ModelsCertificateStoreUpdateServerRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.


### GetName

`func (o *ModelsCertificateStoreUpdateServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCertificateStoreUpdateServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCertificateStoreUpdateServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContainer

`func (o *ModelsCertificateStoreUpdateServerRequest) GetContainer() int32`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelsCertificateStoreUpdateServerRequest) GetContainerOk() (*int32, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelsCertificateStoreUpdateServerRequest) SetContainer(v int32)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelsCertificateStoreUpdateServerRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *ModelsCertificateStoreUpdateServerRequest) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *ModelsCertificateStoreUpdateServerRequest) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


