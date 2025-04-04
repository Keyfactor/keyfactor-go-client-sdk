# AlertsKeyRotationKeyRotationAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**RotationWarningDays** | **int32** |  | 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerRequest**](EventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterRequest**](EventHandlerEventHandlerParameterRequest.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] 

## Methods

### NewAlertsKeyRotationKeyRotationAlertUpdateRequest

`func NewAlertsKeyRotationKeyRotationAlertUpdateRequest(displayName string, rotationWarningDays int32, ) *AlertsKeyRotationKeyRotationAlertUpdateRequest`

NewAlertsKeyRotationKeyRotationAlertUpdateRequest instantiates a new AlertsKeyRotationKeyRotationAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsKeyRotationKeyRotationAlertUpdateRequestWithDefaults

`func NewAlertsKeyRotationKeyRotationAlertUpdateRequestWithDefaults() *AlertsKeyRotationKeyRotationAlertUpdateRequest`

NewAlertsKeyRotationKeyRotationAlertUpdateRequestWithDefaults instantiates a new AlertsKeyRotationKeyRotationAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRotationWarningDays

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetRotationWarningDays() int32`

GetRotationWarningDays returns the RotationWarningDays field if non-nil, zero value otherwise.

### GetRotationWarningDaysOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetRotationWarningDaysOk() (*int32, bool)`

GetRotationWarningDaysOk returns a tuple with the RotationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationWarningDays

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetRotationWarningDays(v int32)`

SetRotationWarningDays sets RotationWarningDays field to given value.


### GetRegisteredEventHandler

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetEventHandlerParameters() []EventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil
### GetUseWorkflows

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *AlertsKeyRotationKeyRotationAlertUpdateRequest) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


