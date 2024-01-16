# KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**RotationWarningDays** | **int64** |  | 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest

`func NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest(displayName string, subject string, message string, rotationWarningDays int64, ) *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest`

NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest instantiates a new KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequestWithDefaults

`func NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequestWithDefaults() *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest`

NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetRotationWarningDays() int64`

GetRotationWarningDays returns the RotationWarningDays field if non-nil, zero value otherwise.

### GetRotationWarningDaysOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetRotationWarningDaysOk() (*int64, bool)`

GetRotationWarningDaysOk returns a tuple with the RotationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetRotationWarningDays(v int64)`

SetRotationWarningDays sets RotationWarningDays field to given value.


### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


