# KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DisplayName** | **string** |  | 
**Subject** | **string** |  | 
**Message** | **string** |  | 
**ExpirationWarningDays** | **int64** |  | 
**CertificateQueryId** | Pointer to **int64** |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest.md) |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest**](KeyfactorApiModelsEventHandlerEventHandlerParameterRequest.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest(displayName string, subject string, message string, expirationWarningDays int64, ) *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequestWithDefaults

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequestWithDefaults() *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequestWithDefaults instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetExpirationWarningDays() int64`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetExpirationWarningDaysOk() (*int64, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetExpirationWarningDays(v int64)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.


### GetCertificateQueryId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetCertificateQueryId() int64`

GetCertificateQueryId returns the CertificateQueryId field if non-nil, zero value otherwise.

### GetCertificateQueryIdOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetCertificateQueryIdOk() (*int64, bool)`

GetCertificateQueryIdOk returns a tuple with the CertificateQueryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQueryId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetCertificateQueryId(v int64)`

SetCertificateQueryId sets CertificateQueryId field to given value.

### HasCertificateQueryId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasCertificateQueryId() bool`

HasCertificateQueryId returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerRequest)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterRequest, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterRequest)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertUpdateRequest) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


