# SMTPSMTPTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**RelayUsername** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ClientId** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**TokenEndpoint** | Pointer to **NullableString** |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**RequestHeaders** | Pointer to [**[]SharedRequestHeaderRequest**](SharedRequestHeaderRequest.md) |  | [optional] 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**RelayAuthenticationType** | [**CSSCMSCoreEnumsSMTPRelayAuthenticationType**](CSSCMSCoreEnumsSMTPRelayAuthenticationType.md) |  | 
**SenderAccount** | **string** |  | 
**SenderName** | **string** |  | 
**TestRecipient** | **string** |  | 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewSMTPSMTPTestRequest

`func NewSMTPSMTPTestRequest(host string, port int32, relayAuthenticationType CSSCMSCoreEnumsSMTPRelayAuthenticationType, senderAccount string, senderName string, testRecipient string, ) *SMTPSMTPTestRequest`

NewSMTPSMTPTestRequest instantiates a new SMTPSMTPTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSMTPTestRequestWithDefaults

`func NewSMTPSMTPTestRequestWithDefaults() *SMTPSMTPTestRequest`

NewSMTPSMTPTestRequestWithDefaults instantiates a new SMTPSMTPTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayPassword

`func (o *SMTPSMTPTestRequest) GetRelayPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *SMTPSMTPTestRequest) GetRelayPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *SMTPSMTPTestRequest) SetRelayPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *SMTPSMTPTestRequest) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### GetRelayUsername

`func (o *SMTPSMTPTestRequest) GetRelayUsername() CSSCMSDataModelModelsKeyfactorAPISecret`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *SMTPSMTPTestRequest) GetRelayUsernameOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *SMTPSMTPTestRequest) SetRelayUsername(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *SMTPSMTPTestRequest) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### GetClientId

`func (o *SMTPSMTPTestRequest) GetClientId() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *SMTPSMTPTestRequest) GetClientIdOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *SMTPSMTPTestRequest) SetClientId(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *SMTPSMTPTestRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *SMTPSMTPTestRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *SMTPSMTPTestRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *SMTPSMTPTestRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *SMTPSMTPTestRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetTokenEndpoint

`func (o *SMTPSMTPTestRequest) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *SMTPSMTPTestRequest) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *SMTPSMTPTestRequest) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.

### HasTokenEndpoint

`func (o *SMTPSMTPTestRequest) HasTokenEndpoint() bool`

HasTokenEndpoint returns a boolean if a field has been set.

### SetTokenEndpointNil

`func (o *SMTPSMTPTestRequest) SetTokenEndpointNil(b bool)`

 SetTokenEndpointNil sets the value for TokenEndpoint to be an explicit nil

### UnsetTokenEndpoint
`func (o *SMTPSMTPTestRequest) UnsetTokenEndpoint()`

UnsetTokenEndpoint ensures that no value is present for TokenEndpoint, not even an explicit nil
### GetScope

`func (o *SMTPSMTPTestRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SMTPSMTPTestRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SMTPSMTPTestRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *SMTPSMTPTestRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *SMTPSMTPTestRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *SMTPSMTPTestRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAudience

`func (o *SMTPSMTPTestRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *SMTPSMTPTestRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *SMTPSMTPTestRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *SMTPSMTPTestRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *SMTPSMTPTestRequest) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *SMTPSMTPTestRequest) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetRequestHeaders

`func (o *SMTPSMTPTestRequest) GetRequestHeaders() []SharedRequestHeaderRequest`

GetRequestHeaders returns the RequestHeaders field if non-nil, zero value otherwise.

### GetRequestHeadersOk

`func (o *SMTPSMTPTestRequest) GetRequestHeadersOk() (*[]SharedRequestHeaderRequest, bool)`

GetRequestHeadersOk returns a tuple with the RequestHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestHeaders

`func (o *SMTPSMTPTestRequest) SetRequestHeaders(v []SharedRequestHeaderRequest)`

SetRequestHeaders sets RequestHeaders field to given value.

### HasRequestHeaders

`func (o *SMTPSMTPTestRequest) HasRequestHeaders() bool`

HasRequestHeaders returns a boolean if a field has been set.

### SetRequestHeadersNil

`func (o *SMTPSMTPTestRequest) SetRequestHeadersNil(b bool)`

 SetRequestHeadersNil sets the value for RequestHeaders to be an explicit nil

### UnsetRequestHeaders
`func (o *SMTPSMTPTestRequest) UnsetRequestHeaders()`

UnsetRequestHeaders ensures that no value is present for RequestHeaders, not even an explicit nil
### GetHost

`func (o *SMTPSMTPTestRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSMTPTestRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSMTPTestRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SMTPSMTPTestRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSMTPTestRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSMTPTestRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRelayAuthenticationType

`func (o *SMTPSMTPTestRequest) GetRelayAuthenticationType() CSSCMSCoreEnumsSMTPRelayAuthenticationType`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *SMTPSMTPTestRequest) GetRelayAuthenticationTypeOk() (*CSSCMSCoreEnumsSMTPRelayAuthenticationType, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *SMTPSMTPTestRequest) SetRelayAuthenticationType(v CSSCMSCoreEnumsSMTPRelayAuthenticationType)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.


### GetSenderAccount

`func (o *SMTPSMTPTestRequest) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *SMTPSMTPTestRequest) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *SMTPSMTPTestRequest) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.


### GetSenderName

`func (o *SMTPSMTPTestRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SMTPSMTPTestRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SMTPSMTPTestRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.


### GetTestRecipient

`func (o *SMTPSMTPTestRequest) GetTestRecipient() string`

GetTestRecipient returns the TestRecipient field if non-nil, zero value otherwise.

### GetTestRecipientOk

`func (o *SMTPSMTPTestRequest) GetTestRecipientOk() (*string, bool)`

GetTestRecipientOk returns a tuple with the TestRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRecipient

`func (o *SMTPSMTPTestRequest) SetTestRecipient(v string)`

SetTestRecipient sets TestRecipient field to given value.


### GetUseSSL

`func (o *SMTPSMTPTestRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SMTPSMTPTestRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SMTPSMTPTestRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SMTPSMTPTestRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


