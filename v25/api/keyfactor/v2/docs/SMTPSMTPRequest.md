# SMTPSMTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**RelayUsername** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**SenderAccount** | Pointer to **NullableString** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**TokenEndpoint** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**RequestHeaders** | Pointer to [**[]SharedRequestHeaderRequest**](SharedRequestHeaderRequest.md) |  | [optional] 

## Methods

### NewSMTPSMTPRequest

`func NewSMTPSMTPRequest() *SMTPSMTPRequest`

NewSMTPSMTPRequest instantiates a new SMTPSMTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSMTPRequestWithDefaults

`func NewSMTPSMTPRequestWithDefaults() *SMTPSMTPRequest`

NewSMTPSMTPRequestWithDefaults instantiates a new SMTPSMTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SMTPSMTPRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSMTPRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSMTPRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPSMTPRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SMTPSMTPRequest) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SMTPSMTPRequest) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SMTPSMTPRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSMTPRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSMTPRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPSMTPRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *SMTPSMTPRequest) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *SMTPSMTPRequest) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *SMTPSMTPRequest) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *SMTPSMTPRequest) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayPassword

`func (o *SMTPSMTPRequest) GetRelayPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *SMTPSMTPRequest) GetRelayPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *SMTPSMTPRequest) SetRelayPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *SMTPSMTPRequest) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### GetRelayUsername

`func (o *SMTPSMTPRequest) GetRelayUsername() CSSCMSDataModelModelsKeyfactorAPISecret`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *SMTPSMTPRequest) GetRelayUsernameOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *SMTPSMTPRequest) SetRelayUsername(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *SMTPSMTPRequest) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### GetSenderAccount

`func (o *SMTPSMTPRequest) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *SMTPSMTPRequest) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *SMTPSMTPRequest) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *SMTPSMTPRequest) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### SetSenderAccountNil

`func (o *SMTPSMTPRequest) SetSenderAccountNil(b bool)`

 SetSenderAccountNil sets the value for SenderAccount to be an explicit nil

### UnsetSenderAccount
`func (o *SMTPSMTPRequest) UnsetSenderAccount()`

UnsetSenderAccount ensures that no value is present for SenderAccount, not even an explicit nil
### GetSenderName

`func (o *SMTPSMTPRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SMTPSMTPRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SMTPSMTPRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *SMTPSMTPRequest) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *SMTPSMTPRequest) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *SMTPSMTPRequest) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetUseSSL

`func (o *SMTPSMTPRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SMTPSMTPRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SMTPSMTPRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SMTPSMTPRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetClientId

`func (o *SMTPSMTPRequest) GetClientId() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SMTPSMTPRequest) GetClientIdOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SMTPSMTPRequest) SetClientId(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SMTPSMTPRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *SMTPSMTPRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SMTPSMTPRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SMTPSMTPRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *SMTPSMTPRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *SMTPSMTPRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *SMTPSMTPRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *SMTPSMTPRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *SMTPSMTPRequest) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### SetTokenEndpointNil

`func (o *SMTPSMTPRequest) SetTokenEndpointNil(b bool)`

 SetTokenEndpointNil sets the value for TokenEndpoint to be an explicit nil

### UnsetTokenEndpoint
`func (o *SMTPSMTPRequest) UnsetTokenEndpoint()`

UnsetTokenEndpoint ensures that no value is present for TokenEndpoint, not even an explicit nil
### GetScope

`func (o *SMTPSMTPRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SMTPSMTPRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SMTPSMTPRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SMTPSMTPRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *SMTPSMTPRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *SMTPSMTPRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAudience

`func (o *SMTPSMTPRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SMTPSMTPRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SMTPSMTPRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SMTPSMTPRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *SMTPSMTPRequest) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *SMTPSMTPRequest) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetRequestHeaders

`func (o *SMTPSMTPRequest) GetRequestHeaders() []SharedRequestHeaderRequest`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SMTPSMTPRequest) GetRequestHeadersOk() (*[]SharedRequestHeaderRequest, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *SMTPSMTPRequest) SetRequestHeaders(v []SharedRequestHeaderRequest)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *SMTPSMTPRequest) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *SMTPSMTPRequest) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *SMTPSMTPRequest) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


