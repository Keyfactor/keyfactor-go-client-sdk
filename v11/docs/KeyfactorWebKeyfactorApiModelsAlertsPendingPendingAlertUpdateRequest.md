# KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest(displayName string, subject string, message string, ) *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTemplateId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetEventHandlerParametersOk() (*[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *KeyfactorWebKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


