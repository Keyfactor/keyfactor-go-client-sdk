# KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to [**KeyfactorWebKeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse**](KeyfactorWebKeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse.md) |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse**](KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse**](KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse

`func NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse() *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse`

NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse instantiates a new KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse`

NewKeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetTemplate

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetTemplate() KeyfactorWebKeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetTemplateOk() (*KeyfactorWebKeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetTemplate(v KeyfactorWebKeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *KeyfactorWebKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


