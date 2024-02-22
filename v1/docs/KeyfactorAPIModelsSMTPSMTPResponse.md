# KeyfactorAPIModelsSMTPSMTPResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int32** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**RelayAuthenticationType** | Pointer to **int32** |  | [optional] 
**RelayUsername** | Pointer to **string** |  | [optional] 
**SenderAccount** | Pointer to **string** |  | [optional] 
**SenderName** | Pointer to **string** |  | [optional] 
**UseSSL** | Pointer to **bool** |  | [optional] 

## Methods

### NewKeyfactorAPIModelsSMTPSMTPResponse

`func NewKeyfactorAPIModelsSMTPSMTPResponse() *KeyfactorAPIModelsSMTPSMTPResponse`

NewKeyfactorAPIModelsSMTPSMTPResponse instantiates a new KeyfactorAPIModelsSMTPSMTPResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorAPIModelsSMTPSMTPResponseWithDefaults

`func NewKeyfactorAPIModelsSMTPSMTPResponseWithDefaults() *KeyfactorAPIModelsSMTPSMTPResponse`

NewKeyfactorAPIModelsSMTPSMTPResponseWithDefaults instantiates a new KeyfactorAPIModelsSMTPSMTPResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetId

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPort

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRelayAuthenticationType

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetRelayAuthenticationType() int32`

GetRelayAuthenticationType returns the RelayAuthenticationType field if non-nil, zero value otherwise.

### GetRelayAuthenticationTypeOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetRelayAuthenticationTypeOk() (*int32, bool)`

GetRelayAuthenticationTypeOk returns a tuple with the RelayAuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayAuthenticationType

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetRelayAuthenticationType(v int32)`

SetRelayAuthenticationType sets RelayAuthenticationType field to given value.

### HasRelayAuthenticationType

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasRelayAuthenticationType() bool`

HasRelayAuthenticationType returns a boolean if a field has been set.

### GetRelayUsername

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetRelayUsername() string`

GetRelayUsername returns the RelayUsername field if non-nil, zero value otherwise.

### GetRelayUsernameOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetRelayUsernameOk() (*string, bool)`

GetRelayUsernameOk returns a tuple with the RelayUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelayUsername

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetRelayUsername(v string)`

SetRelayUsername sets RelayUsername field to given value.

### HasRelayUsername

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasRelayUsername() bool`

HasRelayUsername returns a boolean if a field has been set.

### GetSenderAccount

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetSenderAccount() string`

GetSenderAccount returns the SenderAccount field if non-nil, zero value otherwise.

### GetSenderAccountOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetSenderAccountOk() (*string, bool)`

GetSenderAccountOk returns a tuple with the SenderAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderAccount

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetSenderAccount(v string)`

SetSenderAccount sets SenderAccount field to given value.

### HasSenderAccount

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasSenderAccount() bool`

HasSenderAccount returns a boolean if a field has been set.

### GetSenderName

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### GetUseSSL

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetUseSSL() bool`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) GetUseSSLOk() (*bool, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) SetUseSSL(v bool)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *KeyfactorAPIModelsSMTPSMTPResponse) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


