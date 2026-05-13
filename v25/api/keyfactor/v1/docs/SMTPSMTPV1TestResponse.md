# SMTPSMTPV1TestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayUsername** | Pointer to **NullableString** |  | [optional] 
**SenderAccount** | Pointer to **NullableString** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**TestRecipient** | Pointer to **NullableString** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewSMTPSMTPV1TestResponse

`func NewSMTPSMTPV1TestResponse() *SMTPSMTPV1TestResponse`

NewSMTPSMTPV1TestResponse instantiates a new SMTPSMTPV1TestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMTPSMTPV1TestResponseWithDefaults

`func NewSMTPSMTPV1TestResponseWithDefaults() *SMTPSMTPV1TestResponse`

NewSMTPSMTPV1TestResponseWithDefaults instantiates a new SMTPSMTPV1TestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *SMTPSMTPV1TestResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *SMTPSMTPV1TestResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *SMTPSMTPV1TestResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *SMTPSMTPV1TestResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *SMTPSMTPV1TestResponse) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *SMTPSMTPV1TestResponse) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetId

`func (o *SMTPSMTPV1TestResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SMTPSMTPV1TestResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SMTPSMTPV1TestResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SMTPSMTPV1TestResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *SMTPSMTPV1TestResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *SMTPSMTPV1TestResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *SMTPSMTPV1TestResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *SMTPSMTPV1TestResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *SMTPSMTPV1TestResponse) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *SMTPSMTPV1TestResponse) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *SMTPSMTPV1TestResponse) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *SMTPSMTPV1TestResponse) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayUsername

`func (o *SMTPSMTPV1TestResponse) GetRelayUsername() string`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *SMTPSMTPV1TestResponse) GetRelayUsernameOk() (*string, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *SMTPSMTPV1TestResponse) SetRelayUsername(v string)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *SMTPSMTPV1TestResponse) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### SetRelayUsernameNil

`func (o *SMTPSMTPV1TestResponse) SetRelayUsernameNil(b bool)`

 SetRelayUsernameNil sets the value for RelayUsername to be an explicit nil

### UnsetRelayUsername
`func (o *SMTPSMTPV1TestResponse) UnsetRelayUsername()`

UnsetRelayUsername ensures that no value is present for RelayUsername, not even an explicit nil
### GetSenderAccount

`func (o *SMTPSMTPV1TestResponse) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *SMTPSMTPV1TestResponse) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *SMTPSMTPV1TestResponse) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *SMTPSMTPV1TestResponse) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### SetSenderAccountNil

`func (o *SMTPSMTPV1TestResponse) SetSenderAccountNil(b bool)`

 SetSenderAccountNil sets the value for SenderAccount to be an explicit nil

### UnsetSenderAccount
`func (o *SMTPSMTPV1TestResponse) UnsetSenderAccount()`

UnsetSenderAccount ensures that no value is present for SenderAccount, not even an explicit nil
### GetSenderName

`func (o *SMTPSMTPV1TestResponse) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *SMTPSMTPV1TestResponse) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *SMTPSMTPV1TestResponse) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *SMTPSMTPV1TestResponse) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *SMTPSMTPV1TestResponse) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *SMTPSMTPV1TestResponse) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetTestRecipient

`func (o *SMTPSMTPV1TestResponse) GetTestRecipient() string`

GetTestRecipient returns the TestRecipient field if non-nil, zero value otherwise.

### GetTestRecipientOk

`func (o *SMTPSMTPV1TestResponse) GetTestRecipientOk() (*string, bool)`

GetTestRecipientOk returns a tuple with the TestRecipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestRecipient

`func (o *SMTPSMTPV1TestResponse) SetTestRecipient(v string)`

SetTestRecipient sets TestRecipient field to given value.

### HasTestRecipient

`func (o *SMTPSMTPV1TestResponse) HasTestRecipient() bool`

HasTestRecipient returns a boolean if a field has been set.

### SetTestRecipientNil

`func (o *SMTPSMTPV1TestResponse) SetTestRecipientNil(b bool)`

 SetTestRecipientNil sets the value for TestRecipient to be an explicit nil

### UnsetTestRecipient
`func (o *SMTPSMTPV1TestResponse) UnsetTestRecipient()`

UnsetTestRecipient ensures that no value is present for TestRecipient, not even an explicit nil
### GetUseSSL

`func (o *SMTPSMTPV1TestResponse) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *SMTPSMTPV1TestResponse) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *SMTPSMTPV1TestResponse) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *SMTPSMTPV1TestResponse) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


