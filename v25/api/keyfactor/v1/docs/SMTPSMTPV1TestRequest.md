# SMTPSMTPV1TestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelayPassword** | Pointer to **NullableString** |  | [optional] 
**RelayUsername** | Pointer to **NullableString** |  | [optional] 
**Host** | **string** |  | 
**Port** | **int32** |  | 
**RelayAuthenticationType** | [**CSSCMSCoreEnumsSMTPRelayAuthenticationType**](CSSCMSCoreEnumsSMTPRelayAuthenticationType.md) |  | 
**SenderAccount** | **string** |  | 
**SenderName** | **string** |  | 
**TestRecipient** | **string** |  | 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewSMTPSMTPV1TestRequest

`func NewSMTPSMTPV1TestRequest(host string, port int32, relayAuthenticationType CSSCMSCoreEnumsSMTPRelayAuthenticationType, senderAccount string, senderName string, testRecipient string, ) *SMTPSMTPV1TestRequest`

NewSMTPSMTPV1TestRequest instantiates a new SMTPSMTPV1TestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSMTPV1TestRequestWithDefaults

`func NewSMTPSMTPV1TestRequestWithDefaults() *SMTPSMTPV1TestRequest`

NewSMTPSMTPV1TestRequestWithDefaults instantiates a new SMTPSMTPV1TestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelayPassword

`func (o *SMTPSMTPV1TestRequest) GetRelayPassword() string`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *SMTPSMTPV1TestRequest) GetRelayPasswordOk() (*string, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *SMTPSMTPV1TestRequest) SetRelayPassword(v string)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *SMTPSMTPV1TestRequest) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### SetRelayPasswordNil

`func (o *SMTPSMTPV1TestRequest) SetRelayPasswordNil(b bool)`

 SetRelayPasswordNil sets the value for RelayPassword to be an explicit nil

### UnsetRelayPassword
`func (o *SMTPSMTPV1TestRequest) UnsetRelayPassword()`

UnsetRelayPassword ensures that no value is present for RelayPassword, not even an explicit nil
### GetRelayUsername

`func (o *SMTPSMTPV1TestRequest) GetRelayUsername() string`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *SMTPSMTPV1TestRequest) GetRelayUsernameOk() (*string, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *SMTPSMTPV1TestRequest) SetRelayUsername(v string)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *SMTPSMTPV1TestRequest) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### SetRelayUsernameNil

`func (o *SMTPSMTPV1TestRequest) SetRelayUsernameNil(b bool)`

 SetRelayUsernameNil sets the value for RelayUsername to be an explicit nil

### UnsetRelayUsername
`func (o *SMTPSMTPV1TestRequest) UnsetRelayUsername()`

UnsetRelayUsername ensures that no value is present for RelayUsername, not even an explicit nil
### GetHost

`func (o *SMTPSMTPV1TestRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSMTPV1TestRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSMTPV1TestRequest) SetHost(v string)`

SetHost sets Host field to given value.


### GetPort

`func (o *SMTPSMTPV1TestRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSMTPV1TestRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSMTPV1TestRequest) SetPort(v int32)`

SetPort sets Port field to given value.


### GetRelayAuthenticationType

`func (o *SMTPSMTPV1TestRequest) GetRelayAuthenticationType() CSSCMSCoreEnumsSMTPRelayAuthenticationType`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *SMTPSMTPV1TestRequest) GetRelayAuthenticationTypeOk() (*CSSCMSCoreEnumsSMTPRelayAuthenticationType, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *SMTPSMTPV1TestRequest) SetRelayAuthenticationType(v CSSCMSCoreEnumsSMTPRelayAuthenticationType)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.


### GetSenderAccount

`func (o *SMTPSMTPV1TestRequest) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *SMTPSMTPV1TestRequest) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *SMTPSMTPV1TestRequest) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.


### GetSenderName

`func (o *SMTPSMTPV1TestRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SMTPSMTPV1TestRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SMTPSMTPV1TestRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.


### GetTestRecipient

`func (o *SMTPSMTPV1TestRequest) GetTestRecipient() string`

GetTestRecipient returns the TestRecipient field if non-nil, zero value otherwise.

### GetTestRecipientOk

`func (o *SMTPSMTPV1TestRequest) GetTestRecipientOk() (*string, bool)`

GetTestRecipientOk returns a tuple with the TestRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRecipient

`func (o *SMTPSMTPV1TestRequest) SetTestRecipient(v string)`

SetTestRecipient sets TestRecipient field to given value.


### GetUseSSL

`func (o *SMTPSMTPV1TestRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SMTPSMTPV1TestRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SMTPSMTPV1TestRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SMTPSMTPV1TestRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


