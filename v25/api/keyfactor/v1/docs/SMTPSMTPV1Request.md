# SMTPSMTPV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayPassword** | Pointer to **NullableString** |  | [optional] 
**RelayUsername** | Pointer to **NullableString** |  | [optional] 
**SenderAccount** | Pointer to **NullableString** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewSMTPSMTPV1Request

`func NewSMTPSMTPV1Request() *SMTPSMTPV1Request`

NewSMTPSMTPV1Request instantiates a new SMTPSMTPV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSMTPV1RequestWithDefaults

`func NewSMTPSMTPV1RequestWithDefaults() *SMTPSMTPV1Request`

NewSMTPSMTPV1RequestWithDefaults instantiates a new SMTPSMTPV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SMTPSMTPV1Request) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSMTPV1Request) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSMTPV1Request) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPSMTPV1Request) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SMTPSMTPV1Request) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SMTPSMTPV1Request) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *SMTPSMTPV1Request) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSMTPV1Request) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSMTPV1Request) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPSMTPV1Request) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *SMTPSMTPV1Request) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *SMTPSMTPV1Request) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *SMTPSMTPV1Request) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *SMTPSMTPV1Request) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayPassword

`func (o *SMTPSMTPV1Request) GetRelayPassword() string`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *SMTPSMTPV1Request) GetRelayPasswordOk() (*string, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *SMTPSMTPV1Request) SetRelayPassword(v string)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *SMTPSMTPV1Request) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### SetRelayPasswordNil

`func (o *SMTPSMTPV1Request) SetRelayPasswordNil(b bool)`

 SetRelayPasswordNil sets the value for RelayPassword to be an explicit nil

### UnsetRelayPassword
`func (o *SMTPSMTPV1Request) UnsetRelayPassword()`

UnsetRelayPassword ensures that no value is present for RelayPassword, not even an explicit nil
### GetRelayUsername

`func (o *SMTPSMTPV1Request) GetRelayUsername() string`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *SMTPSMTPV1Request) GetRelayUsernameOk() (*string, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *SMTPSMTPV1Request) SetRelayUsername(v string)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *SMTPSMTPV1Request) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### SetRelayUsernameNil

`func (o *SMTPSMTPV1Request) SetRelayUsernameNil(b bool)`

 SetRelayUsernameNil sets the value for RelayUsername to be an explicit nil

### UnsetRelayUsername
`func (o *SMTPSMTPV1Request) UnsetRelayUsername()`

UnsetRelayUsername ensures that no value is present for RelayUsername, not even an explicit nil
### GetSenderAccount

`func (o *SMTPSMTPV1Request) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *SMTPSMTPV1Request) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *SMTPSMTPV1Request) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *SMTPSMTPV1Request) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### SetSenderAccountNil

`func (o *SMTPSMTPV1Request) SetSenderAccountNil(b bool)`

 SetSenderAccountNil sets the value for SenderAccount to be an explicit nil

### UnsetSenderAccount
`func (o *SMTPSMTPV1Request) UnsetSenderAccount()`

UnsetSenderAccount ensures that no value is present for SenderAccount, not even an explicit nil
### GetSenderName

`func (o *SMTPSMTPV1Request) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SMTPSMTPV1Request) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SMTPSMTPV1Request) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *SMTPSMTPV1Request) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *SMTPSMTPV1Request) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *SMTPSMTPV1Request) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetUseSSL

`func (o *SMTPSMTPV1Request) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SMTPSMTPV1Request) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SMTPSMTPV1Request) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SMTPSMTPV1Request) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


