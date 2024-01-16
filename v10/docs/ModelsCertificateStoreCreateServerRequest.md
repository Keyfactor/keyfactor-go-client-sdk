# ModelsCertificateStoreCreateServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | 
**Password** | [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | 
**UseSSL** | **bool** |  | 
**ServerType** | **int64** |  | 
**Name** | **string** |  | 
**Container** | Pointer to **int64** |  | [optional] 

## Methods

### NewModelsCertificateStoreCreateServerRequest

`func NewModelsCertificateStoreCreateServerRequest(username ModelsKeyfactorAPISecret, password ModelsKeyfactorAPISecret, useSSL bool, serverType int64, name string, ) *ModelsCertificateStoreCreateServerRequest`

NewModelsCertificateStoreCreateServerRequest instantiates a new ModelsCertificateStoreCreateServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreCreateServerRequestWithDefaults

`func NewModelsCertificateStoreCreateServerRequestWithDefaults() *ModelsCertificateStoreCreateServerRequest`

NewModelsCertificateStoreCreateServerRequestWithDefaults instantiates a new ModelsCertificateStoreCreateServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ModelsCertificateStoreCreateServerRequest) GetUsername() ModelsKeyfactorAPISecret`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsCertificateStoreCreateServerRequest) GetUsernameOk() (*ModelsKeyfactorAPISecret, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsCertificateStoreCreateServerRequest) SetUsername(v ModelsKeyfactorAPISecret)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *ModelsCertificateStoreCreateServerRequest) GetPassword() ModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateStoreCreateServerRequest) GetPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateStoreCreateServerRequest) SetPassword(v ModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.


### GetUseSSL

`func (o *ModelsCertificateStoreCreateServerRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *ModelsCertificateStoreCreateServerRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *ModelsCertificateStoreCreateServerRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.


### GetServerType

`func (o *ModelsCertificateStoreCreateServerRequest) GetServerType() int64`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ModelsCertificateStoreCreateServerRequest) GetServerTypeOk() (*int64, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ModelsCertificateStoreCreateServerRequest) SetServerType(v int64)`

SetServerType sets ServerType field to given value.


### GetName

`func (o *ModelsCertificateStoreCreateServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCertificateStoreCreateServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCertificateStoreCreateServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContainer

`func (o *ModelsCertificateStoreCreateServerRequest) GetContainer() int64`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *ModelsCertificateStoreCreateServerRequest) GetContainerOk() (*int64, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *ModelsCertificateStoreCreateServerRequest) SetContainer(v int64)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *ModelsCertificateStoreCreateServerRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


