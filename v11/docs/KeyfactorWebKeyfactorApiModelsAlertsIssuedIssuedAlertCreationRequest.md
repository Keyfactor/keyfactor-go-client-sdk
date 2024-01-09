# KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest

`func NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest(displayName string, subject string, message string, ) *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTemplateId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) GetEventHandlerParametersOk() (*[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *KeyfactorWebKeyfactorApiModelsAlertsIssuedIssuedAlertCreationRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


