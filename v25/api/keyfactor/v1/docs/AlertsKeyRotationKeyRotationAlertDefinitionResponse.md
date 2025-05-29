# AlertsKeyRotationKeyRotationAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipient** | Pointer to **NullableString** |  | [optional] 
**RotationWarningDays** | Pointer to **int32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerResponse**](EventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterResponse**](EventHandlerEventHandlerParameterResponse.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] 
**WorkflowId** | Pointer to **NullableString** |  | [optional] 
**WorkflowName** | Pointer to **NullableString** |  | [optional] 
**WorkflowPublishedVersion** | Pointer to **NullableInt32** |  | [optional] 
**WorkflowEnabled** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewAlertsKeyRotationKeyRotationAlertDefinitionResponse

`func NewAlertsKeyRotationKeyRotationAlertDefinitionResponse() *AlertsKeyRotationKeyRotationAlertDefinitionResponse`

NewAlertsKeyRotationKeyRotationAlertDefinitionResponse instantiates a new AlertsKeyRotationKeyRotationAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsKeyRotationKeyRotationAlertDefinitionResponseWithDefaults

`func NewAlertsKeyRotationKeyRotationAlertDefinitionResponseWithDefaults() *AlertsKeyRotationKeyRotationAlertDefinitionResponse`

NewAlertsKeyRotationKeyRotationAlertDefinitionResponseWithDefaults instantiates a new AlertsKeyRotationKeyRotationAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSubject

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipient

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### SetRecipientNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRecipientNil(b bool)`

 SetRecipientNil sets the value for Recipient to be an explicit nil

### UnsetRecipient
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetRecipient()`

UnsetRecipient ensures that no value is present for Recipient, not even an explicit nil
### GetRotationWarningDays

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRotationWarningDays() int32`

GetRotationWarningDays returns the RotationWarningDays field if non-nil, zero value otherwise.

### GetRotationWarningDaysOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRotationWarningDaysOk() (*int32, bool)`

GetRotationWarningDaysOk returns a tuple with the RotationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationWarningDays

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRotationWarningDays(v int32)`

SetRotationWarningDays sets RotationWarningDays field to given value.

### HasRotationWarningDays

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasRotationWarningDays() bool`

HasRotationWarningDays returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetEventHandlerParameters() []EventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil
### GetUseWorkflows

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.

### GetWorkflowId

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetWorkflowName

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### SetWorkflowNameNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowNameNil(b bool)`

 SetWorkflowNameNil sets the value for WorkflowName to be an explicit nil

### UnsetWorkflowName
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetWorkflowName()`

UnsetWorkflowName ensures that no value is present for WorkflowName, not even an explicit nil
### GetWorkflowPublishedVersion

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowPublishedVersion() int32`

GetWorkflowPublishedVersion returns the WorkflowPublishedVersion field if non-nil, zero value otherwise.

### GetWorkflowPublishedVersionOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowPublishedVersionOk() (*int32, bool)`

GetWorkflowPublishedVersionOk returns a tuple with the WorkflowPublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPublishedVersion

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowPublishedVersion(v int32)`

SetWorkflowPublishedVersion sets WorkflowPublishedVersion field to given value.

### HasWorkflowPublishedVersion

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasWorkflowPublishedVersion() bool`

HasWorkflowPublishedVersion returns a boolean if a field has been set.

### SetWorkflowPublishedVersionNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowPublishedVersionNil(b bool)`

 SetWorkflowPublishedVersionNil sets the value for WorkflowPublishedVersion to be an explicit nil

### UnsetWorkflowPublishedVersion
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetWorkflowPublishedVersion()`

UnsetWorkflowPublishedVersion ensures that no value is present for WorkflowPublishedVersion, not even an explicit nil
### GetWorkflowEnabled

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowEnabled() bool`

GetWorkflowEnabled returns the WorkflowEnabled field if non-nil, zero value otherwise.

### GetWorkflowEnabledOk

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) GetWorkflowEnabledOk() (*bool, bool)`

GetWorkflowEnabledOk returns a tuple with the WorkflowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEnabled

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowEnabled(v bool)`

SetWorkflowEnabled sets WorkflowEnabled field to given value.

### HasWorkflowEnabled

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) HasWorkflowEnabled() bool`

HasWorkflowEnabled returns a boolean if a field has been set.

### SetWorkflowEnabledNil

`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) SetWorkflowEnabledNil(b bool)`

 SetWorkflowEnabledNil sets the value for WorkflowEnabled to be an explicit nil

### UnsetWorkflowEnabled
`func (o *AlertsKeyRotationKeyRotationAlertDefinitionResponse) UnsetWorkflowEnabled()`

UnsetWorkflowEnabled ensures that no value is present for WorkflowEnabled, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


