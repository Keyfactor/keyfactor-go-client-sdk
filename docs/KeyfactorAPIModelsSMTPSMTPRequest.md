# KeyfactorAPIModelsSMTPSMTPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayPassword** | Pointer to **string** |  | [optional] 
**RelayUsername** | Pointer to **string** |  | [optional] 
**SenderAccount** | Pointer to **string** |  | [optional] 
**SenderName** | Pointer to **string** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorAPIModelsSMTPSMTPRequest

`func NewKeyfactorAPIModelsSMTPSMTPRequest() *KeyfactorAPIModelsSMTPSMTPRequest`

NewKeyfactorAPIModelsSMTPSMTPRequest instantiates a new KeyfactorAPIModelsSMTPSMTPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorAPIModelsSMTPSMTPRequestWithDefaults

`func NewKeyfactorAPIModelsSMTPSMTPRequestWithDefaults() *KeyfactorAPIModelsSMTPSMTPRequest`

NewKeyfactorAPIModelsSMTPSMTPRequestWithDefaults instantiates a new KeyfactorAPIModelsSMTPSMTPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayPassword

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetRelayPassword() string`

GetRelayPassword returns the RelayPassword field if non-nil, zero value otherwise.

### GetRelayPasswordOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetRelayPasswordOk() (*string, bool)`

GetRelayPasswordOk returns a tuple with the RelayPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayPassword

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetRelayPassword(v string)`

SetRelayPassword sets RelayPassword field to given value.

### HasRelayPassword

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasRelayPassword() bool`

HasRelayPassword returns a boolean if a field has been set.

### GetRelayUsername

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetRelayUsername() string`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetRelayUsernameOk() (*string, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetRelayUsername(v string)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### GetSenderAccount

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### GetSenderName

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetUseSSL

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *KeyfactorAPIModelsSMTPSMTPRequest) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


