# AlertsExpirationExpirationAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExpirationWarningDays** | **int32** |  | 
**CertificateQueryId** | Pointer to **int32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerRequest**](EventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterRequest**](EventHandlerEventHandlerParameterRequest.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] [default to false]
**WorkflowId** | Pointer to **NullableString** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 

## Methods

### NewAlertsExpirationExpirationAlertUpdateRequest

`func NewAlertsExpirationExpirationAlertUpdateRequest(displayName string, expirationWarningDays int32, ) *AlertsExpirationExpirationAlertUpdateRequest`

NewAlertsExpirationExpirationAlertUpdateRequest instantiates a new AlertsExpirationExpirationAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsExpirationExpirationAlertUpdateRequestWithDefaults

`func NewAlertsExpirationExpirationAlertUpdateRequestWithDefaults() *AlertsExpirationExpirationAlertUpdateRequest`

NewAlertsExpirationExpirationAlertUpdateRequestWithDefaults instantiates a new AlertsExpirationExpirationAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsExpirationExpirationAlertUpdateRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsExpirationExpirationAlertUpdateRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetExpirationWarningDays() int32`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetExpirationWarningDaysOk() (*int32, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetExpirationWarningDays(v int32)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.


### GetCertificateQueryId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetCertificateQueryId() int32`

GetCertificateQueryId returns the CertificateQueryId field if non-nil, zero value otherwise.

### GetCertificateQueryIdOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetCertificateQueryIdOk() (*int32, bool)`

GetCertificateQueryIdOk returns a tuple with the CertificateQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQueryId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetCertificateQueryId(v int32)`

SetCertificateQueryId sets CertificateQueryId field to given value.

### HasCertificateQueryId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasCertificateQueryId() bool`

HasCertificateQueryId returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsExpirationExpirationAlertUpdateRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetEventHandlerParameters() []EventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsExpirationExpirationAlertUpdateRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil
### GetUseWorkflows

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.

### GetWorkflowId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *AlertsExpirationExpirationAlertUpdateRequest) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetSchedule

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AlertsExpirationExpirationAlertUpdateRequest) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AlertsExpirationExpirationAlertUpdateRequest) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AlertsExpirationExpirationAlertUpdateRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


