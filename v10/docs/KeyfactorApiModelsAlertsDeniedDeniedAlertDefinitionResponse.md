# KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to [**KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse**](KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse.md) |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse**](KeyfactorApiModelsEventHandlerEventHandlerParameterResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse

`func NewKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse() *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse instantiates a new KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponseWithDefaults

`func NewKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponseWithDefaults() *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetTemplate

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetTemplate() KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetTemplateOk() (*KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetTemplate(v KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsDeniedDeniedAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


