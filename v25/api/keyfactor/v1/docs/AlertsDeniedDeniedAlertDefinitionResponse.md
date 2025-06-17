# AlertsDeniedDeniedAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**Template** | Pointer to [**AlertsAlertTemplateAlertTemplateResponse**](AlertsAlertTemplateAlertTemplateResponse.md) |  | [optional] 
**RegisteredEventHandler** | Pointer to [**EventHandlerRegisteredEventHandlerResponse**](EventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]EventHandlerEventHandlerParameterResponse**](EventHandlerEventHandlerParameterResponse.md) |  | [optional] 

## Methods

### NewAlertsDeniedDeniedAlertDefinitionResponse

`func NewAlertsDeniedDeniedAlertDefinitionResponse() *AlertsDeniedDeniedAlertDefinitionResponse`

NewAlertsDeniedDeniedAlertDefinitionResponse instantiates a new AlertsDeniedDeniedAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsDeniedDeniedAlertDefinitionResponseWithDefaults

`func NewAlertsDeniedDeniedAlertDefinitionResponseWithDefaults() *AlertsDeniedDeniedAlertDefinitionResponse`

NewAlertsDeniedDeniedAlertDefinitionResponseWithDefaults instantiates a new AlertsDeniedDeniedAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *AlertsDeniedDeniedAlertDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSubject

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsDeniedDeniedAlertDefinitionResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsDeniedDeniedAlertDefinitionResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipients

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsDeniedDeniedAlertDefinitionResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetTemplate

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetTemplate() AlertsAlertTemplateAlertTemplateResponse`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetTemplateOk() (*AlertsAlertTemplateAlertTemplateResponse, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetTemplate(v AlertsAlertTemplateAlertTemplateResponse)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetRegisteredEventHandler() EventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*EventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetRegisteredEventHandler(v EventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetEventHandlerParameters() []EventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]EventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetEventHandlerParameters(v []EventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *AlertsDeniedDeniedAlertDefinitionResponse) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *AlertsDeniedDeniedAlertDefinitionResponse) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


