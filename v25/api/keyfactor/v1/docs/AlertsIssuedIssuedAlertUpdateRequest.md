# AlertsIssuedIssuedAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**FriendlyName** | Pointer to **NullableString** |  | [optional] 
**TemplateId** | Pointer to **NullableInt32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerRequest**](EventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterRequest**](EventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewAlertsIssuedIssuedAlertUpdateRequest

`func NewAlertsIssuedIssuedAlertUpdateRequest(displayName string, subject string, message string, ) *AlertsIssuedIssuedAlertUpdateRequest`

NewAlertsIssuedIssuedAlertUpdateRequest instantiates a new AlertsIssuedIssuedAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsIssuedIssuedAlertUpdateRequestWithDefaults

`func NewAlertsIssuedIssuedAlertUpdateRequestWithDefaults() *AlertsIssuedIssuedAlertUpdateRequest`

NewAlertsIssuedIssuedAlertUpdateRequestWithDefaults instantiates a new AlertsIssuedIssuedAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsIssuedIssuedAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetFriendlyName

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetFriendlyName() string`

GetFriendlyName returns the FriendlyName field if non-nil, zero value otherwise.

### GetFriendlyNameOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetFriendlyNameOk() (*string, bool)`

GetFriendlyNameOk returns a tuple with the FriendlyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendlyName

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetFriendlyName(v string)`

SetFriendlyName sets FriendlyName field to given value.

### HasFriendlyName

`func (o *AlertsIssuedIssuedAlertUpdateRequest) HasFriendlyName() bool`

HasFriendlyName returns a boolean if a field has been set.

### SetFriendlyNameNil

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetFriendlyNameNil(b bool)`

 SetFriendlyNameNil sets the value for FriendlyName to be an explicit nil

### UnsetFriendlyName
`func (o *AlertsIssuedIssuedAlertUpdateRequest) UnsetFriendlyName()`

UnsetFriendlyName ensures that no value is present for FriendlyName, not even an explicit nil
### GetTemplateId

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetTemplateId() int32`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetTemplateIdOk() (*int32, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetTemplateId(v int32)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *AlertsIssuedIssuedAlertUpdateRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### SetTemplateIdNil

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetTemplateIdNil(b bool)`

 SetTemplateIdNil sets the value for TemplateId to be an explicit nil

### UnsetTemplateId
`func (o *AlertsIssuedIssuedAlertUpdateRequest) UnsetTemplateId()`

UnsetTemplateId ensures that no value is present for TemplateId, not even an explicit nil
### GetRegisteredEventHandler

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsIssuedIssuedAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsIssuedIssuedAlertUpdateRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsIssuedIssuedAlertUpdateRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetEventHandlerParameters

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetEventHandlerParameters() []EventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsIssuedIssuedAlertUpdateRequest) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsIssuedIssuedAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsIssuedIssuedAlertUpdateRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsIssuedIssuedAlertUpdateRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


