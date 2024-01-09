# CSSCMSDataModelModelsCertificateStoreServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to [**CSSCMSDataModelModelsKeyfactorSecret**](CSSCMSDataModelModelsKeyfactorSecret.md) |  | [optional] 
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorSecret**](CSSCMSDataModelModelsKeyfactorSecret.md) |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 
**ServerType** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateStoreServer

`func NewCSSCMSDataModelModelsCertificateStoreServer() *CSSCMSDataModelModelsCertificateStoreServer`

NewCSSCMSDataModelModelsCertificateStoreServer instantiates a new CSSCMSDataModelModelsCertificateStoreServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateStoreServerWithDefaults

`func NewCSSCMSDataModelModelsCertificateStoreServerWithDefaults() *CSSCMSDataModelModelsCertificateStoreServer`

NewCSSCMSDataModelModelsCertificateStoreServerWithDefaults instantiates a new CSSCMSDataModelModelsCertificateStoreServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsCertificateStoreServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetUsername() CSSCMSDataModelModelsKeyfactorSecret`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetUsernameOk() (*CSSCMSDataModelModelsKeyfactorSecret, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetUsername(v CSSCMSDataModelModelsKeyfactorSecret)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *CSSCMSDataModelModelsCertificateStoreServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetPassword() CSSCMSDataModelModelsKeyfactorSecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorSecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetPassword(v CSSCMSDataModelModelsKeyfactorSecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CSSCMSDataModelModelsCertificateStoreServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *CSSCMSDataModelModelsCertificateStoreServer) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetServerType

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetServerType(v int32)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *CSSCMSDataModelModelsCertificateStoreServer) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsCertificateStoreServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSDataModelModelsCertificateStoreServer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSDataModelModelsCertificateStoreServer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSDataModelModelsCertificateStoreServer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


