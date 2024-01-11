# KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **NullableString** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayPassword** | Pointer to **NullableString** |  | [optional] 
**RelayUsername** | Pointer to **NullableString** |  | [optional] 
**SenderAccount** | Pointer to **NullableString** |  | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequest

`func NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequest() *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest`

NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequest instantiates a new KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest`

NewKeyfactorWebKeyfactorApiModelsSMTPSMTPRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetId

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayPassword

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetRelayPassword() string`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetRelayPasswordOk() (*string, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetRelayPassword(v string)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### SetRelayPasswordNil

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetRelayPasswordNil(b bool)`

 SetRelayPasswordNil sets the value for RelayPassword to be an explicit nil

### UnsetRelayPassword
`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) UnsetRelayPassword()`

UnsetRelayPassword ensures that no value is present for RelayPassword, not even an explicit nil
### GetRelayUsername

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetRelayUsername() string`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetRelayUsernameOk() (*string, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetRelayUsername(v string)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### SetRelayUsernameNil

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetRelayUsernameNil(b bool)`

 SetRelayUsernameNil sets the value for RelayUsername to be an explicit nil

### UnsetRelayUsername
`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) UnsetRelayUsername()`

UnsetRelayUsername ensures that no value is present for RelayUsername, not even an explicit nil
### GetSenderAccount

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### SetSenderAccountNil

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetSenderAccountNil(b bool)`

 SetSenderAccountNil sets the value for SenderAccount to be an explicit nil

### UnsetSenderAccount
`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) UnsetSenderAccount()`

UnsetSenderAccount ensures that no value is present for SenderAccount, not even an explicit nil
### GetSenderName

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetUseSSL

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *KeyfactorWebKeyfactorApiModelsSMTPSMTPRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


