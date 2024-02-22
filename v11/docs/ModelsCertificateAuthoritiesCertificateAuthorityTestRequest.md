# ModelsCertificateAuthoritiesCertificateAuthorityTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CaType** | Pointer to **int32** |  | [optional] 
**ExplicitCredentials** | Pointer to **bool** |  | [optional] 
**ExplicitPassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**ExplicitUser** | Pointer to **NullableString** |  | [optional] 
**AuthCertificatePassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**AuthCertificate** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **NullableString** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsCertificateAuthoritiesCertificateAuthorityTestRequest

`func NewModelsCertificateAuthoritiesCertificateAuthorityTestRequest() *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest`

NewModelsCertificateAuthoritiesCertificateAuthorityTestRequest instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateAuthoritiesCertificateAuthorityTestRequestWithDefaults

`func NewModelsCertificateAuthoritiesCertificateAuthorityTestRequestWithDefaults() *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest`

NewModelsCertificateAuthoritiesCertificateAuthorityTestRequestWithDefaults instantiates a new ModelsCertificateAuthoritiesCertificateAuthorityTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetCaType() int32`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetCaTypeOk() (*int32, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetCaType(v int32)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitPassword() ModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitPassword(v ModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### SetExplicitUserNil

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitUserNil(b bool)`

 SetExplicitUserNil sets the value for ExplicitUser to be an explicit nil

### UnsetExplicitUser
`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetExplicitUser()`

UnsetExplicitUser ensures that no value is present for ExplicitUser, not even an explicit nil
### GetAuthCertificatePassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificatePassword() ModelsKeyfactorAPISecret`

GetAuthCertificatePassword returns the AuthCertificatePassword field if non-nil, zero value otherwise.

### GetAuthCertificatePasswordOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificatePasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetAuthCertificatePasswordOk returns a tuple with the AuthCertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificatePassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetAuthCertificatePassword(v ModelsKeyfactorAPISecret)`

SetAuthCertificatePassword sets AuthCertificatePassword field to given value.

### HasAuthCertificatePassword

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasAuthCertificatePassword() bool`

HasAuthCertificatePassword returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificate() ModelsKeyfactorAPISecret`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificateOk() (*ModelsKeyfactorAPISecret, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetAuthCertificate(v ModelsKeyfactorAPISecret)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil
### GetHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *ModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


