# CSSCMSDataModelModelsCertificateStoreCreateServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | 
**Password** | [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | 
**UseSSL** | **bool** |  | 
**ServerType** | **int32** |  | 
**Name** | **string** |  | 
**Container** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreCreateServerRequest

`func NewCSSCMSDataModelModelsCertificateStoreCreateServerRequest(username CSSCMSDataModelModelsKeyfactorAPISecret, password CSSCMSDataModelModelsKeyfactorAPISecret, useSSL bool, serverType int32, name string, ) *CSSCMSDataModelModelsCertificateStoreCreateServerRequest`

NewCSSCMSDataModelModelsCertificateStoreCreateServerRequest instantiates a new CSSCMSDataModelModelsCertificateStoreCreateServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreCreateServerRequestWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreCreateServerRequestWithDefaults() *CSSCMSDataModelModelsCertificateStoreCreateServerRequest`

NewCSSCMSDataModelModelsCertificateStoreCreateServerRequestWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreCreateServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetUsername() CSSCMSDataModelModelsKeyfactorAPISecret`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetUsernameOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetUsername(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.


### GetUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.


### GetServerType

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetServerType(v int32)`

SetServerType sets ServerType field to given value.


### GetName

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContainer

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetContainer() int32`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) GetContainerOk() (*int32, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetContainer(v int32)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *CSSCMSDataModelModelsCertificateStoreCreateServerRequest) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


