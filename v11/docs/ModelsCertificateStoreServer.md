# ModelsCertificateStoreServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Username** | Pointer to [**ModelsKeyfactorSecret**](ModelsKeyfactorSecret.md) |  | [optional] 
**Password** | Pointer to [**ModelsKeyfactorSecret**](ModelsKeyfactorSecret.md) |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 
**ServerType** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsCertificateStoreServer

`func NewModelsCertificateStoreServer() *ModelsCertificateStoreServer`

NewModelsCertificateStoreServer instantiates a new ModelsCertificateStoreServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateStoreServerWithDefaults

`func NewModelsCertificateStoreServerWithDefaults() *ModelsCertificateStoreServer`

NewModelsCertificateStoreServerWithDefaults instantiates a new ModelsCertificateStoreServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateStoreServer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateStoreServer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateStoreServer) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateStoreServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *ModelsCertificateStoreServer) GetUsername() ModelsKeyfactorSecret`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ModelsCertificateStoreServer) GetUsernameOk() (*ModelsKeyfactorSecret, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ModelsCertificateStoreServer) SetUsername(v ModelsKeyfactorSecret)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ModelsCertificateStoreServer) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ModelsCertificateStoreServer) GetPassword() ModelsKeyfactorSecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateStoreServer) GetPasswordOk() (*ModelsKeyfactorSecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateStoreServer) SetPassword(v ModelsKeyfactorSecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsCertificateStoreServer) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetUseSSL

`func (o *ModelsCertificateStoreServer) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *ModelsCertificateStoreServer) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *ModelsCertificateStoreServer) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *ModelsCertificateStoreServer) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetServerType

`func (o *ModelsCertificateStoreServer) GetServerType() int32`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ModelsCertificateStoreServer) GetServerTypeOk() (*int32, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ModelsCertificateStoreServer) SetServerType(v int32)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ModelsCertificateStoreServer) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetName

`func (o *ModelsCertificateStoreServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ModelsCertificateStoreServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ModelsCertificateStoreServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ModelsCertificateStoreServer) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ModelsCertificateStoreServer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ModelsCertificateStoreServer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


