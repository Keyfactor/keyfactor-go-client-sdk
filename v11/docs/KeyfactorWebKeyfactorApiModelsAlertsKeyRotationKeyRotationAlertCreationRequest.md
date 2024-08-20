# KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**RotationWarningDays** | **int32** |  | 
**RegisteredEventHandler** | Pointer to [**KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest

`func NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest(displayName string, subject string, message string, rotationWarningDays int32, ) *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRotationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRotationWarningDays() int32`

GetRotationWarningDays returns the RotationWarningDays field if non-nil, zero value otherwise.

### GetRotationWarningDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRotationWarningDaysOk() (*int32, bool)`

GetRotationWarningDaysOk returns a tuple with the RotationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetRotationWarningDays(v int32)`

SetRotationWarningDays sets RotationWarningDays field to given value.


### GetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetEventHandlerParametersOk() (*[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *KeyfactorWebKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


