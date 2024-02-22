# KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**RotationWarningDays** | **int64** |  | 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest

`func NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest(displayName string, subject string, message string, rotationWarningDays int64, ) *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest`

NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest instantiates a new KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequestWithDefaults

`func NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequestWithDefaults() *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest`

NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRotationWarningDays() int64`

GetRotationWarningDays returns the RotationWarningDays field if non-nil, zero value otherwise.

### GetRotationWarningDaysOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRotationWarningDaysOk() (*int64, bool)`

GetRotationWarningDaysOk returns a tuple with the RotationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetRotationWarningDays(v int64)`

SetRotationWarningDays sets RotationWarningDays field to given value.


### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertCreationRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


