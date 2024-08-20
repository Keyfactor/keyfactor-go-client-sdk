# KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Recipient** | Pointer to **string** |  | [optional] 
**RotationWarningDays** | Pointer to **int64** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse**](KeyfactorApiModelsEventHandlerEventHandlerParameterResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse

`func NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse() *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse instantiates a new KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponseWithDefaults

`func NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponseWithDefaults() *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecipient

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRecipient() string`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRecipientOk() (*string, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRecipient(v string)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRotationWarningDays() int64`

GetRotationWarningDays returns the RotationWarningDays field if non-nil, zero value otherwise.

### GetRotationWarningDaysOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRotationWarningDaysOk() (*int64, bool)`

GetRotationWarningDaysOk returns a tuple with the RotationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRotationWarningDays(v int64)`

SetRotationWarningDays sets RotationWarningDays field to given value.

### HasRotationWarningDays

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasRotationWarningDays() bool`

HasRotationWarningDays returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsKeyRotationKeyRotationAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


