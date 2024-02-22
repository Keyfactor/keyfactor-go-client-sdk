# KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**TemplateId** | Pointer to **int64** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest

`func NewKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest(displayName string, subject string, message string, ) *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest`

NewKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest instantiates a new KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequestWithDefaults

`func NewKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequestWithDefaults() *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest`

NewKeyfactorApiModelsAlertsPendingPendingAlertUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTemplateId

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetTemplateId() int64`

GetTemplateId returns the TemplateId field if non-nil, zero value otherwise.

### GetTemplateIdOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetTemplateIdOk() (*int64, bool)`

GetTemplateIdOk returns a tuple with the TemplateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateId

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetTemplateId(v int64)`

SetTemplateId sets TemplateId field to given value.

### HasTemplateId

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasTemplateId() bool`

HasTemplateId returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsPendingPendingAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


