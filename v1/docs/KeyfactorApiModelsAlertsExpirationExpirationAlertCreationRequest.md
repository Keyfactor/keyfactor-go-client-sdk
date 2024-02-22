# KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**ExpirationWarningDays** | **int32** |  | 
**CertificateQueryId** | Pointer to **int32** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest(displayName string, subject string, message string, expirationWarningDays int32, ) *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequestWithDefaults

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequestWithDefaults() *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetExpirationWarningDays() int32`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetExpirationWarningDaysOk() (*int32, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetExpirationWarningDays(v int32)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.


### GetCertificateQueryId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetCertificateQueryId() int32`

GetCertificateQueryId returns the CertificateQueryId field if non-nil, zero value otherwise.

### GetCertificateQueryIdOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetCertificateQueryIdOk() (*int32, bool)`

GetCertificateQueryIdOk returns a tuple with the CertificateQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQueryId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetCertificateQueryId(v int32)`

SetCertificateQueryId sets CertificateQueryId field to given value.

### HasCertificateQueryId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasCertificateQueryId() bool`

HasCertificateQueryId returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertCreationRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


