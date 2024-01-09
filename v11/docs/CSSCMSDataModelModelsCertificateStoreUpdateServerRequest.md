# CSSCMSDataModelModelsCertificateStoreUpdateServerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Username** | [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | 
**Password** | [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | 
**UseSSL** | **bool** |  | 
**Name** | **string** |  | 
**Container** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreUpdateServerRequest

`func NewCSSCMSDataModelModelsCertificateStoreUpdateServerRequest(id int32, username CSSCMSDataModelModelsKeyfactorAPISecret, password CSSCMSDataModelModelsKeyfactorAPISecret, useSSL bool, name string, ) *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest`

NewCSSCMSDataModelModelsCertificateStoreUpdateServerRequest instantiates a new CSSCMSDataModelModelsCertificateStoreUpdateServerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreUpdateServerRequestWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreUpdateServerRequestWithDefaults() *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest`

NewCSSCMSDataModelModelsCertificateStoreUpdateServerRequestWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreUpdateServerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetUsername

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetUsername() CSSCMSDataModelModelsKeyfactorAPISecret`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetUsernameOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetUsername(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetUsername sets Username field to given value.


### GetPassword

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.


### GetUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.


### GetName

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContainer

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetContainer() int32`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) GetContainerOk() (*int32, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetContainer(v int32)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *CSSCMSDataModelModelsCertificateStoreUpdateServerRequest) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


