# CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CaType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 
**ExplicitCredentials** | Pointer to **bool** |  | [optional] 
**ExplicitPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ExplicitUser** | Pointer to **NullableString** |  | [optional] 
**AuthCertificatePassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**AuthCertificate** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **NullableString** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest

`func NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest() *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest`

NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest instantiates a new CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequestWithDefaults

`func NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequestWithDefaults() *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest`

NewCSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequestWithDefaults instantiates a new CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCaType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetCaType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCaType returns the CaType field if non-nil, zero value otherwise.

### GetCaTypeOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetCaTypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCaTypeOk returns a tuple with the CaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetCaType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCaType sets CaType field to given value.

### HasCaType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasCaType() bool`

HasCaType returns a boolean if a field has been set.

### GetExplicitCredentials

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitCredentials() bool`

GetExplicitCredentials returns the ExplicitCredentials field if non-nil, zero value otherwise.

### GetExplicitCredentialsOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitCredentialsOk() (*bool, bool)`

GetExplicitCredentialsOk returns a tuple with the ExplicitCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitCredentials

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitCredentials(v bool)`

SetExplicitCredentials sets ExplicitCredentials field to given value.

### HasExplicitCredentials

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasExplicitCredentials() bool`

HasExplicitCredentials returns a boolean if a field has been set.

### GetExplicitPassword

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetExplicitPassword returns the ExplicitPassword field if non-nil, zero value otherwise.

### GetExplicitPasswordOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetExplicitPasswordOk returns a tuple with the ExplicitPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitPassword

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetExplicitPassword sets ExplicitPassword field to given value.

### HasExplicitPassword

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasExplicitPassword() bool`

HasExplicitPassword returns a boolean if a field has been set.

### GetExplicitUser

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitUser() string`

GetExplicitUser returns the ExplicitUser field if non-nil, zero value otherwise.

### GetExplicitUserOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetExplicitUserOk() (*string, bool)`

GetExplicitUserOk returns a tuple with the ExplicitUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitUser

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitUser(v string)`

SetExplicitUser sets ExplicitUser field to given value.

### HasExplicitUser

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasExplicitUser() bool`

HasExplicitUser returns a boolean if a field has been set.

### SetExplicitUserNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetExplicitUserNil(b bool)`

 SetExplicitUserNil sets the value for ExplicitUser to be an explicit nil

### UnsetExplicitUser
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetExplicitUser()`

UnsetExplicitUser ensures that no value is present for ExplicitUser, not even an explicit nil
### GetAuthCertificatePassword

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificatePassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetAuthCertificatePassword returns the AuthCertificatePassword field if non-nil, zero value otherwise.

### GetAuthCertificatePasswordOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificatePasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetAuthCertificatePasswordOk returns a tuple with the AuthCertificatePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificatePassword

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetAuthCertificatePassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetAuthCertificatePassword sets AuthCertificatePassword field to given value.

### HasAuthCertificatePassword

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasAuthCertificatePassword() bool`

HasAuthCertificatePassword returns a boolean if a field has been set.

### GetAuthCertificate

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificate() CSSCMSDataModelModelsKeyfactorAPISecret`

GetAuthCertificate returns the AuthCertificate field if non-nil, zero value otherwise.

### GetAuthCertificateOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAuthCertificateOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetAuthCertificateOk returns a tuple with the AuthCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthCertificate

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetAuthCertificate(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetAuthCertificate sets AuthCertificate field to given value.

### HasAuthCertificate

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasAuthCertificate() bool`

HasAuthCertificate returns a boolean if a field has been set.

### GetLogicalName

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetLogicalName() string`

GetLogicalName returns the LogicalName field if non-nil, zero value otherwise.

### GetLogicalNameOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetLogicalNameOk() (*string, bool)`

GetLogicalNameOk returns a tuple with the LogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogicalName

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetLogicalName(v string)`

SetLogicalName sets LogicalName field to given value.

### HasLogicalName

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasLogicalName() bool`

HasLogicalName returns a boolean if a field has been set.

### SetLogicalNameNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetLogicalNameNil(b bool)`

 SetLogicalNameNil sets the value for LogicalName to be an explicit nil

### UnsetLogicalName
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetLogicalName()`

UnsetLogicalName ensures that no value is present for LogicalName, not even an explicit nil
### GetHostName

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### SetHostNameNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetHostNameNil(b bool)`

 SetHostNameNil sets the value for HostName to be an explicit nil

### UnsetHostName
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetHostName()`

UnsetHostName ensures that no value is present for HostName, not even an explicit nil
### GetForestRoot

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetForestRoot() string`

GetForestRoot returns the ForestRoot field if non-nil, zero value otherwise.

### GetForestRootOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetForestRootOk() (*string, bool)`

GetForestRootOk returns a tuple with the ForestRoot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForestRoot

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetForestRoot(v string)`

SetForestRoot sets ForestRoot field to given value.

### HasForestRoot

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasForestRoot() bool`

HasForestRoot returns a boolean if a field has been set.

### SetForestRootNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetForestRootNil(b bool)`

 SetForestRootNil sets the value for ForestRoot to be an explicit nil

### UnsetForestRoot
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetForestRoot()`

UnsetForestRoot ensures that no value is present for ForestRoot, not even an explicit nil
### GetConfigurationTenant

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetConfigurationTenant() string`

GetConfigurationTenant returns the ConfigurationTenant field if non-nil, zero value otherwise.

### GetConfigurationTenantOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetConfigurationTenantOk() (*string, bool)`

GetConfigurationTenantOk returns a tuple with the ConfigurationTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigurationTenant

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetConfigurationTenant(v string)`

SetConfigurationTenant sets ConfigurationTenant field to given value.

### HasConfigurationTenant

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasConfigurationTenant() bool`

HasConfigurationTenant returns a boolean if a field has been set.

### SetConfigurationTenantNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetConfigurationTenantNil(b bool)`

 SetConfigurationTenantNil sets the value for ConfigurationTenant to be an explicit nil

### UnsetConfigurationTenant
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetConfigurationTenant()`

UnsetConfigurationTenant ensures that no value is present for ConfigurationTenant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


