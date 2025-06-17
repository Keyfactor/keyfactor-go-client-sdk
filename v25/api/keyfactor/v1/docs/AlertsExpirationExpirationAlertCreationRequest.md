# AlertsExpirationExpirationAlertCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExpirationWarningDays** | **int32** |  | 
**CertificateQueryId** | Pointer to **int32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerRequest**](EventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterRequest**](EventHandlerEventHandlerParameterRequest.md) |  | [optional] 
**UseWorkflows** | Pointer to **bool** |  | [optional] [default to false]

## Methods

### NewAlertsExpirationExpirationAlertCreationRequest

`func NewAlertsExpirationExpirationAlertCreationRequest(displayName string, expirationWarningDays int32, ) *AlertsExpirationExpirationAlertCreationRequest`

NewAlertsExpirationExpirationAlertCreationRequest instantiates a new AlertsExpirationExpirationAlertCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsExpirationExpirationAlertCreationRequestWithDefaults

`func NewAlertsExpirationExpirationAlertCreationRequestWithDefaults() *AlertsExpirationExpirationAlertCreationRequest`

NewAlertsExpirationExpirationAlertCreationRequestWithDefaults instantiates a new AlertsExpirationExpirationAlertCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsExpirationExpirationAlertCreationRequest) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsExpirationExpirationAlertCreationRequest) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetExpirationWarningDays() int32`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetExpirationWarningDaysOk() (*int32, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetExpirationWarningDays(v int32)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.


### GetCertificateQueryId

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetCertificateQueryId() int32`

GetCertificateQueryId returns the CertificateQueryId field if non-nil, zero value otherwise.

### GetCertificateQueryIdOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetCertificateQueryIdOk() (*int32, bool)`

GetCertificateQueryIdOk returns a tuple with the CertificateQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQueryId

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetCertificateQueryId(v int32)`

SetCertificateQueryId sets CertificateQueryId field to given value.

### HasCertificateQueryId

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasCertificateQueryId() bool`

HasCertificateQueryId returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsExpirationExpirationAlertCreationRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetEventHandlerParameters() []EventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsExpirationExpirationAlertCreationRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil
### GetUseWorkflows

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetUseWorkflows() bool`

GetUseWorkflows returns the UseWorkflows field if non-nil, zero value otherwise.

### GetUseWorkflowsOk

`func (o *AlertsExpirationExpirationAlertCreationRequest) GetUseWorkflowsOk() (*bool, bool)`

GetUseWorkflowsOk returns a tuple with the UseWorkflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseWorkflows

`func (o *AlertsExpirationExpirationAlertCreationRequest) SetUseWorkflows(v bool)`

SetUseWorkflows sets UseWorkflows field to given value.

### HasUseWorkflows

`func (o *AlertsExpirationExpirationAlertCreationRequest) HasUseWorkflows() bool`

HasUseWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


