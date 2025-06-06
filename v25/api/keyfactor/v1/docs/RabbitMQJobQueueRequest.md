# RabbitMQJobQueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthType** | Pointer to **NullableString** |  | [optional] 
**TaskQueueURL** | Pointer to **NullableString** |  | [optional] 
**TokenURL** | Pointer to **NullableString** |  | [optional] 
**ClientID** | Pointer to **NullableString** |  | [optional] 
**ClientSecret** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**Scope** | Pointer to **NullableString** |  | [optional] 
**Audience** | Pointer to **NullableString** |  | [optional] 
**Username** | Pointer to **NullableString** |  | [optional] 
**Password** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 

## Methods

### NewRabbitMQJobQueueRequest

`func NewRabbitMQJobQueueRequest() *RabbitMQJobQueueRequest`

NewRabbitMQJobQueueRequest instantiates a new RabbitMQJobQueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRabbitMQJobQueueRequestWithDefaults

`func NewRabbitMQJobQueueRequestWithDefaults() *RabbitMQJobQueueRequest`

NewRabbitMQJobQueueRequestWithDefaults instantiates a new RabbitMQJobQueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthType

`func (o *RabbitMQJobQueueRequest) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *RabbitMQJobQueueRequest) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *RabbitMQJobQueueRequest) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *RabbitMQJobQueueRequest) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### SetAuthTypeNil

`func (o *RabbitMQJobQueueRequest) SetAuthTypeNil(b bool)`

 SetAuthTypeNil sets the value for AuthType to be an explicit nil

### UnsetAuthType
`func (o *RabbitMQJobQueueRequest) UnsetAuthType()`

UnsetAuthType ensures that no value is present for AuthType, not even an explicit nil
### GetTaskQueueURL

`func (o *RabbitMQJobQueueRequest) GetTaskQueueURL() string`

GetTaskQueueURL returns the TaskQueueURL field if non-nil, zero value otherwise.

### GetTaskQueueURLOk

`func (o *RabbitMQJobQueueRequest) GetTaskQueueURLOk() (*string, bool)`

GetTaskQueueURLOk returns a tuple with the TaskQueueURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskQueueURL

`func (o *RabbitMQJobQueueRequest) SetTaskQueueURL(v string)`

SetTaskQueueURL sets TaskQueueURL field to given value.

### HasTaskQueueURL

`func (o *RabbitMQJobQueueRequest) HasTaskQueueURL() bool`

HasTaskQueueURL returns a boolean if a field has been set.

### SetTaskQueueURLNil

`func (o *RabbitMQJobQueueRequest) SetTaskQueueURLNil(b bool)`

 SetTaskQueueURLNil sets the value for TaskQueueURL to be an explicit nil

### UnsetTaskQueueURL
`func (o *RabbitMQJobQueueRequest) UnsetTaskQueueURL()`

UnsetTaskQueueURL ensures that no value is present for TaskQueueURL, not even an explicit nil
### GetTokenURL

`func (o *RabbitMQJobQueueRequest) GetTokenURL() string`

GetTokenURL returns the TokenURL field if non-nil, zero value otherwise.

### GetTokenURLOk

`func (o *RabbitMQJobQueueRequest) GetTokenURLOk() (*string, bool)`

GetTokenURLOk returns a tuple with the TokenURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenURL

`func (o *RabbitMQJobQueueRequest) SetTokenURL(v string)`

SetTokenURL sets TokenURL field to given value.

### HasTokenURL

`func (o *RabbitMQJobQueueRequest) HasTokenURL() bool`

HasTokenURL returns a boolean if a field has been set.

### SetTokenURLNil

`func (o *RabbitMQJobQueueRequest) SetTokenURLNil(b bool)`

 SetTokenURLNil sets the value for TokenURL to be an explicit nil

### UnsetTokenURL
`func (o *RabbitMQJobQueueRequest) UnsetTokenURL()`

UnsetTokenURL ensures that no value is present for TokenURL, not even an explicit nil
### GetClientID

`func (o *RabbitMQJobQueueRequest) GetClientID() string`

GetClientID returns the ClientID field if non-nil, zero value otherwise.

### GetClientIDOk

`func (o *RabbitMQJobQueueRequest) GetClientIDOk() (*string, bool)`

GetClientIDOk returns a tuple with the ClientID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientID

`func (o *RabbitMQJobQueueRequest) SetClientID(v string)`

SetClientID sets ClientID field to given value.

### HasClientID

`func (o *RabbitMQJobQueueRequest) HasClientID() bool`

HasClientID returns a boolean if a field has been set.

### SetClientIDNil

`func (o *RabbitMQJobQueueRequest) SetClientIDNil(b bool)`

 SetClientIDNil sets the value for ClientID to be an explicit nil

### UnsetClientID
`func (o *RabbitMQJobQueueRequest) UnsetClientID()`

UnsetClientID ensures that no value is present for ClientID, not even an explicit nil
### GetClientSecret

`func (o *RabbitMQJobQueueRequest) GetClientSecret() CSSCMSDataModelModelsKeyfactorAPISecret`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *RabbitMQJobQueueRequest) GetClientSecretOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *RabbitMQJobQueueRequest) SetClientSecret(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *RabbitMQJobQueueRequest) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScope

`func (o *RabbitMQJobQueueRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *RabbitMQJobQueueRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *RabbitMQJobQueueRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *RabbitMQJobQueueRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### SetScopeNil

`func (o *RabbitMQJobQueueRequest) SetScopeNil(b bool)`

 SetScopeNil sets the value for Scope to be an explicit nil

### UnsetScope
`func (o *RabbitMQJobQueueRequest) UnsetScope()`

UnsetScope ensures that no value is present for Scope, not even an explicit nil
### GetAudience

`func (o *RabbitMQJobQueueRequest) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *RabbitMQJobQueueRequest) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *RabbitMQJobQueueRequest) SetAudience(v string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *RabbitMQJobQueueRequest) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *RabbitMQJobQueueRequest) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *RabbitMQJobQueueRequest) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetUsername

`func (o *RabbitMQJobQueueRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RabbitMQJobQueueRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RabbitMQJobQueueRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *RabbitMQJobQueueRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### SetUsernameNil

`func (o *RabbitMQJobQueueRequest) SetUsernameNil(b bool)`

 SetUsernameNil sets the value for Username to be an explicit nil

### UnsetUsername
`func (o *RabbitMQJobQueueRequest) UnsetUsername()`

UnsetUsername ensures that no value is present for Username, not even an explicit nil
### GetPassword

`func (o *RabbitMQJobQueueRequest) GetPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RabbitMQJobQueueRequest) GetPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RabbitMQJobQueueRequest) SetPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *RabbitMQJobQueueRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


