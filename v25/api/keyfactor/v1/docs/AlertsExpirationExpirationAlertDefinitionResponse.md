# AlertsExpirationExpirationAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExpirationWarningDays** | Pointer to **int32** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**CertificateQuery** | Pointer to [**AlertsAlertCertificateQueryAlertCertificateQueryResponse**](AlertsAlertCertificateQueryAlertCertificateQueryResponse.md) |  | [optional] 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerResponse**](EventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterResponse**](EventHandlerEventHandlerParameterResponse.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] 
**WorkflowId** | Pointer to **NullableString** |  | [optional] 
**WorkflowName** | Pointer to **NullableString** |  | [optional] 
**WorkflowPublishedVersion** | Pointer to **NullableInt32** |  | [optional] 
**WorkflowEnabled** | Pointer to **NullableBool** |  | [optional] 
**Schedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 

## Methods

### NewAlertsExpirationExpirationAlertDefinitionResponse

`func NewAlertsExpirationExpirationAlertDefinitionResponse() *AlertsExpirationExpirationAlertDefinitionResponse`

NewAlertsExpirationExpirationAlertDefinitionResponse instantiates a new AlertsExpirationExpirationAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsExpirationExpirationAlertDefinitionResponseWithDefaults

`func NewAlertsExpirationExpirationAlertDefinitionResponseWithDefaults() *AlertsExpirationExpirationAlertDefinitionResponse`

NewAlertsExpirationExpirationAlertDefinitionResponseWithDefaults instantiates a new AlertsExpirationExpirationAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSubject

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetExpirationWarningDays() int32`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetExpirationWarningDaysOk() (*int32, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetExpirationWarningDays(v int32)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.

### HasExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasExpirationWarningDays() bool`

HasExpirationWarningDays returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetCertificateQuery

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetCertificateQuery() AlertsAlertCertificateQueryAlertCertificateQueryResponse`

GetCertificateQuery returns the CertificateQuery field if non-nil, zero value otherwise.

### GetCertificateQueryOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetCertificateQueryOk() (*AlertsAlertCertificateQueryAlertCertificateQueryResponse, bool)`

GetCertificateQueryOk returns a tuple with the CertificateQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQuery

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetCertificateQuery(v AlertsAlertCertificateQueryAlertCertificateQueryResponse)`

SetCertificateQuery sets CertificateQuery field to given value.

### HasCertificateQuery

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasCertificateQuery() bool`

HasCertificateQuery returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetEventHandlerParameters() []EventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil
### GetUseWorkflows

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.

### GetWorkflowId

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### SetWorkflowIdNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowIdNil(b bool)`

 SetWorkflowIdNil sets the value for WorkflowId to be an explicit nil

### UnsetWorkflowId
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetWorkflowId()`

UnsetWorkflowId ensures that no value is present for WorkflowId, not even an explicit nil
### GetWorkflowName

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowName() string`

GetWorkflowName returns the WorkflowName field if non-nil, zero value otherwise.

### GetWorkflowNameOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowNameOk() (*string, bool)`

GetWorkflowNameOk returns a tuple with the WorkflowName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowName

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowName(v string)`

SetWorkflowName sets WorkflowName field to given value.

### HasWorkflowName

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasWorkflowName() bool`

HasWorkflowName returns a boolean if a field has been set.

### SetWorkflowNameNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowNameNil(b bool)`

 SetWorkflowNameNil sets the value for WorkflowName to be an explicit nil

### UnsetWorkflowName
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetWorkflowName()`

UnsetWorkflowName ensures that no value is present for WorkflowName, not even an explicit nil
### GetWorkflowPublishedVersion

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowPublishedVersion() int32`

GetWorkflowPublishedVersion returns the WorkflowPublishedVersion field if non-nil, zero value otherwise.

### GetWorkflowPublishedVersionOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowPublishedVersionOk() (*int32, bool)`

GetWorkflowPublishedVersionOk returns a tuple with the WorkflowPublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowPublishedVersion

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowPublishedVersion(v int32)`

SetWorkflowPublishedVersion sets WorkflowPublishedVersion field to given value.

### HasWorkflowPublishedVersion

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasWorkflowPublishedVersion() bool`

HasWorkflowPublishedVersion returns a boolean if a field has been set.

### SetWorkflowPublishedVersionNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowPublishedVersionNil(b bool)`

 SetWorkflowPublishedVersionNil sets the value for WorkflowPublishedVersion to be an explicit nil

### UnsetWorkflowPublishedVersion
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetWorkflowPublishedVersion()`

UnsetWorkflowPublishedVersion ensures that no value is present for WorkflowPublishedVersion, not even an explicit nil
### GetWorkflowEnabled

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowEnabled() bool`

GetWorkflowEnabled returns the WorkflowEnabled field if non-nil, zero value otherwise.

### GetWorkflowEnabledOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetWorkflowEnabledOk() (*bool, bool)`

GetWorkflowEnabledOk returns a tuple with the WorkflowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEnabled

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowEnabled(v bool)`

SetWorkflowEnabled sets WorkflowEnabled field to given value.

### HasWorkflowEnabled

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasWorkflowEnabled() bool`

HasWorkflowEnabled returns a boolean if a field has been set.

### SetWorkflowEnabledNil

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetWorkflowEnabledNil(b bool)`

 SetWorkflowEnabledNil sets the value for WorkflowEnabled to be an explicit nil

### UnsetWorkflowEnabled
`func (o *AlertsExpirationExpirationAlertDefinitionResponse) UnsetWorkflowEnabled()`

UnsetWorkflowEnabled ensures that no value is present for WorkflowEnabled, not even an explicit nil
### GetSchedule

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) GetScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) SetSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AlertsExpirationExpirationAlertDefinitionResponse) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


