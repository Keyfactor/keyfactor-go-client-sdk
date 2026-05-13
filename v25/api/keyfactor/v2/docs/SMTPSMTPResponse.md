# SMTPSMTPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayUsername** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**RelayPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**SenderAccount** | Pointer to **NullableString** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 
**ClientId** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**TokenEndpoint** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**RequestHeaders** | Pointer to [**[]SharedRequestHeaderResponse**](SharedRequestHeaderResponse.md) |  | [optional] 

## Methods

### NewSMTPSMTPResponse

`func NewSMTPSMTPResponse() *SMTPSMTPResponse`

NewSMTPSMTPResponse instantiates a new SMTPSMTPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSMTPResponseWithDefaults

`func NewSMTPSMTPResponseWithDefaults() *SMTPSMTPResponse`

NewSMTPSMTPResponseWithDefaults instantiates a new SMTPSMTPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SMTPSMTPResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSMTPResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSMTPResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPSMTPResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SMTPSMTPResponse) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SMTPSMTPResponse) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetId

`func (o *SMTPSMTPResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMTPSMTPResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMTPSMTPResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SMTPSMTPResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *SMTPSMTPResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSMTPResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSMTPResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPSMTPResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *SMTPSMTPResponse) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *SMTPSMTPResponse) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *SMTPSMTPResponse) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *SMTPSMTPResponse) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayUsername

`func (o *SMTPSMTPResponse) GetRelayUsername() CSSCMSDataModelModelsKeyfactorAPISecret`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *SMTPSMTPResponse) GetRelayUsernameOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *SMTPSMTPResponse) SetRelayUsername(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *SMTPSMTPResponse) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### GetRelayPassword

`func (o *SMTPSMTPResponse) GetRelayPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *SMTPSMTPResponse) GetRelayPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *SMTPSMTPResponse) SetRelayPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *SMTPSMTPResponse) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### GetSenderAccount

`func (o *SMTPSMTPResponse) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *SMTPSMTPResponse) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *SMTPSMTPResponse) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *SMTPSMTPResponse) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### SetSenderAccountNil

`func (o *SMTPSMTPResponse) SetSenderAccountNil(b bool)`

 SetSenderAccountNil sets the value for SenderAccount to be an explicit nil

### UnsetSenderAccount
`func (o *SMTPSMTPResponse) UnsetSenderAccount()`

UnsetSenderAccount ensures that no value is present for SenderAccount, not even an explicit nil
### GetSenderName

`func (o *SMTPSMTPResponse) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SMTPSMTPResponse) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SMTPSMTPResponse) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *SMTPSMTPResponse) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *SMTPSMTPResponse) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *SMTPSMTPResponse) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetUseSSL

`func (o *SMTPSMTPResponse) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SMTPSMTPResponse) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SMTPSMTPResponse) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SMTPSMTPResponse) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetClientId

`func (o *SMTPSMTPResponse) GetClientId() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SMTPSMTPResponse) GetClientIdOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SMTPSMTPResponse) SetClientId(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SMTPSMTPResponse) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *SMTPSMTPResponse) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SMTPSMTPResponse) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SMTPSMTPResponse) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *SMTPSMTPResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *SMTPSMTPResponse) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *SMTPSMTPResponse) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *SMTPSMTPResponse) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *SMTPSMTPResponse) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### SetTokenEndpointNil

`func (o *SMTPSMTPResponse) SetTokenEndpointNil(b bool)`

 SetTokenEndpointNil sets the value for TokenEndpoint to be an explicit nil

### UnsetTokenEndpoint
`func (o *SMTPSMTPResponse) UnsetTokenEndpoint()`

UnsetTokenEndpoint ensures that no value is present for TokenEndpoint, not even an explicit nil
### GetScope

`func (o *SMTPSMTPResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SMTPSMTPResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SMTPSMTPResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SMTPSMTPResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *SMTPSMTPResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *SMTPSMTPResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAudience

`func (o *SMTPSMTPResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SMTPSMTPResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SMTPSMTPResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SMTPSMTPResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *SMTPSMTPResponse) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *SMTPSMTPResponse) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetRequestHeaders

`func (o *SMTPSMTPResponse) GetRequestHeaders() []SharedRequestHeaderResponse`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SMTPSMTPResponse) GetRequestHeadersOk() (*[]SharedRequestHeaderResponse, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *SMTPSMTPResponse) SetRequestHeaders(v []SharedRequestHeaderResponse)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *SMTPSMTPResponse) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *SMTPSMTPResponse) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *SMTPSMTPResponse) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


