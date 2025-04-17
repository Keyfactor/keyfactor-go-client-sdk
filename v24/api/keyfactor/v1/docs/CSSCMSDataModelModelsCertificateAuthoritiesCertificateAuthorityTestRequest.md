# CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CAType** | Pointer to [**CSSCMSCoreEnumsCertificateAuthorityType**](CSSCMSCoreEnumsCertificateAuthorityType.md) |  | [optional] 
**ExplicitCredentials** | Pointer to **bool** |  | [optional] 
**ExplicitPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ExplicitUser** | Pointer to **NullableString** |  | [optional] 
**AuthCertificatePassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**AuthCertificate** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**LogicalName** | Pointer to **NullableString** |  | [optional] 
**HostName** | Pointer to **NullableString** |  | [optional] 
**ForestRoot** | Pointer to **NullableString** |  | [optional] 
**ConfigurationTenant** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ClientId** | Pointer to **NullableString** |  | [optional] 
**TokenURL** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**UseCAConnector** | Pointer to **bool** |  | [optional] 
**ConnectorPool** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**Remote** | Pointer to **bool** |  | [optional] 

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

### GetCAType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetCAType() CSSCMSCoreEnumsCertificateAuthorityType`

GetCAType returns the CAType field if non-nil, zero value otherwise.

### GetCATypeOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetCATypeOk() (*CSSCMSCoreEnumsCertificateAuthorityType, bool)`

GetCATypeOk returns a tuple with the CAType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetCAType(v CSSCMSCoreEnumsCertificateAuthorityType)`

SetCAType sets CAType field to given value.

### HasCAType

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasCAType() bool`

HasCAType returns a boolean if a field has been set.

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
### GetClientSecret

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### SetClientIdNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetClientIdNil(b bool)`

 SetClientIdNil sets the value for ClientId to be an explicit nil

### UnsetClientId
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetClientId()`

UnsetClientId ensures that no value is present for ClientId, not even an explicit nil
### GetTokenURL

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetTokenURL() string`

GetTokenURL returns the TokenURL field if non-nil, zero value otherwise.

### GetTokenURLOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetTokenURLOk() (*string, bool)`

GetTokenURLOk returns a tuple with the TokenURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURL

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetTokenURL(v string)`

SetTokenURL sets TokenURL field to given value.

### HasTokenURL

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasTokenURL() bool`

HasTokenURL returns a boolean if a field has been set.

### SetTokenURLNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetTokenURLNil(b bool)`

 SetTokenURLNil sets the value for TokenURL to be an explicit nil

### UnsetTokenURL
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetTokenURL()`

UnsetTokenURL ensures that no value is present for TokenURL, not even an explicit nil
### GetScope

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetUseCAConnector

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetUseCAConnector() bool`

GetUseCAConnector returns the UseCAConnector field if non-nil, zero value otherwise.

### GetUseCAConnectorOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetUseCAConnectorOk() (*bool, bool)`

GetUseCAConnectorOk returns a tuple with the UseCAConnector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCAConnector

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetUseCAConnector(v bool)`

SetUseCAConnector sets UseCAConnector field to given value.

### HasUseCAConnector

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasUseCAConnector() bool`

HasUseCAConnector returns a boolean if a field has been set.

### GetConnectorPool

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetConnectorPool() string`

GetConnectorPool returns the ConnectorPool field if non-nil, zero value otherwise.

### GetConnectorPoolOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetConnectorPoolOk() (*string, bool)`

GetConnectorPoolOk returns a tuple with the ConnectorPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorPool

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetConnectorPool(v string)`

SetConnectorPool sets ConnectorPool field to given value.

### HasConnectorPool

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasConnectorPool() bool`

HasConnectorPool returns a boolean if a field has been set.

### SetConnectorPoolNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetConnectorPoolNil(b bool)`

 SetConnectorPoolNil sets the value for ConnectorPool to be an explicit nil

### UnsetConnectorPool
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetConnectorPool()`

UnsetConnectorPool ensures that no value is present for ConnectorPool, not even an explicit nil
### GetAudience

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetRemote

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *CSSCMSDataModelModelsCertificateAuthoritiesCertificateAuthorityTestRequest) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


