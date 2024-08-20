# KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse

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

### NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse

`func NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse() *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse instantiates a new KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponseWithDefaults

`func NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponseWithDefaults() *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetTemplate

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetTemplate() KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetTemplateOk() (*KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetTemplate(v KeyfactorApiModelsAlertsAlertTemplateAlertTemplateResponse)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsIssuedIssuedAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


