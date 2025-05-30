# RabbitMQJobQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskQueueURL** | Pointer to **NullableString** |  | [optional] 
**TokenURL** | Pointer to **NullableString** |  | [optional] 
**ClientID** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**AuthType** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewRabbitMQJobQueueResponse

`func NewRabbitMQJobQueueResponse() *RabbitMQJobQueueResponse`

NewRabbitMQJobQueueResponse instantiates a new RabbitMQJobQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRabbitMQJobQueueResponseWithDefaults

`func NewRabbitMQJobQueueResponseWithDefaults() *RabbitMQJobQueueResponse`

NewRabbitMQJobQueueResponseWithDefaults instantiates a new RabbitMQJobQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskQueueURL

`func (o *RabbitMQJobQueueResponse) GetTaskQueueURL() string`

GetTaskQueueURL returns the TaskQueueURL field if non-nil, zero value otherwise.

### GetTaskQueueURLOk

`func (o *RabbitMQJobQueueResponse) GetTaskQueueURLOk() (*string, bool)`

GetTaskQueueURLOk returns a tuple with the TaskQueueURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueURL

`func (o *RabbitMQJobQueueResponse) SetTaskQueueURL(v string)`

SetTaskQueueURL sets TaskQueueURL field to given value.

### HasTaskQueueURL

`func (o *RabbitMQJobQueueResponse) HasTaskQueueURL() bool`

HasTaskQueueURL returns a boolean if a field has been set.

### SetTaskQueueURLNil

`func (o *RabbitMQJobQueueResponse) SetTaskQueueURLNil(b bool)`

 SetTaskQueueURLNil sets the value for TaskQueueURL to be an explicit nil

### UnsetTaskQueueURL
`func (o *RabbitMQJobQueueResponse) UnsetTaskQueueURL()`

UnsetTaskQueueURL ensures that no value is present for TaskQueueURL, not even an explicit nil
### GetTokenURL

`func (o *RabbitMQJobQueueResponse) GetTokenURL() string`

GetTokenURL returns the TokenURL field if non-nil, zero value otherwise.

### GetTokenURLOk

`func (o *RabbitMQJobQueueResponse) GetTokenURLOk() (*string, bool)`

GetTokenURLOk returns a tuple with the TokenURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURL

`func (o *RabbitMQJobQueueResponse) SetTokenURL(v string)`

SetTokenURL sets TokenURL field to given value.

### HasTokenURL

`func (o *RabbitMQJobQueueResponse) HasTokenURL() bool`

HasTokenURL returns a boolean if a field has been set.

### SetTokenURLNil

`func (o *RabbitMQJobQueueResponse) SetTokenURLNil(b bool)`

 SetTokenURLNil sets the value for TokenURL to be an explicit nil

### UnsetTokenURL
`func (o *RabbitMQJobQueueResponse) UnsetTokenURL()`

UnsetTokenURL ensures that no value is present for TokenURL, not even an explicit nil
### GetClientID

`func (o *RabbitMQJobQueueResponse) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *RabbitMQJobQueueResponse) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *RabbitMQJobQueueResponse) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *RabbitMQJobQueueResponse) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### SetClientIDNil

`func (o *RabbitMQJobQueueResponse) SetClientIDNil(b bool)`

 SetClientIDNil sets the value for ClientID to be an explicit nil

### UnsetClientID
`func (o *RabbitMQJobQueueResponse) UnsetClientID()`

UnsetClientID ensures that no value is present for ClientID, not even an explicit nil
### GetClientSecret

`func (o *RabbitMQJobQueueResponse) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *RabbitMQJobQueueResponse) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *RabbitMQJobQueueResponse) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *RabbitMQJobQueueResponse) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScope

`func (o *RabbitMQJobQueueResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RabbitMQJobQueueResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RabbitMQJobQueueResponse) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RabbitMQJobQueueResponse) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RabbitMQJobQueueResponse) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RabbitMQJobQueueResponse) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetUsername

`func (o *RabbitMQJobQueueResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RabbitMQJobQueueResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RabbitMQJobQueueResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RabbitMQJobQueueResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RabbitMQJobQueueResponse) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RabbitMQJobQueueResponse) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetAudience

`func (o *RabbitMQJobQueueResponse) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *RabbitMQJobQueueResponse) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *RabbitMQJobQueueResponse) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *RabbitMQJobQueueResponse) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *RabbitMQJobQueueResponse) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *RabbitMQJobQueueResponse) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetPassword

`func (o *RabbitMQJobQueueResponse) GetPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RabbitMQJobQueueResponse) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RabbitMQJobQueueResponse) SetPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RabbitMQJobQueueResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetAuthType

`func (o *RabbitMQJobQueueResponse) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RabbitMQJobQueueResponse) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RabbitMQJobQueueResponse) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *RabbitMQJobQueueResponse) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *RabbitMQJobQueueResponse) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *RabbitMQJobQueueResponse) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


