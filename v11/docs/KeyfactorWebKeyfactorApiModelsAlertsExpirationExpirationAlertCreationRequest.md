# KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**ExpirationWarningDays** | **int32** |  | 
**CertificateQueryId** | Pointer to **int32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest

`func NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest(displayName string, subject string, message string, expirationWarningDays int32, ) *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest instantiates a new KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest`

NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExpirationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetExpirationWarningDays() int32`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetExpirationWarningDaysOk() (*int32, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetExpirationWarningDays(v int32)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.


### GetCertificateQueryId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetCertificateQueryId() int32`

GetCertificateQueryId returns the CertificateQueryId field if non-nil, zero value otherwise.

### GetCertificateQueryIdOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetCertificateQueryIdOk() (*int32, bool)`

GetCertificateQueryIdOk returns a tuple with the CertificateQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQueryId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetCertificateQueryId(v int32)`

SetCertificateQueryId sets CertificateQueryId field to given value.

### HasCertificateQueryId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasCertificateQueryId() bool`

HasCertificateQueryId returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetEventHandlerParametersOk() (*[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


