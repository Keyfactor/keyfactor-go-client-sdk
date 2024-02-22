# KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**ExpirationWarningDays** | Pointer to **int32** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**CertificateQuery** | Pointer to [**KeyfactorWebKeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse**](KeyfactorWebKeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse.md) |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse**](KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse**](KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse

`func NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse() *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse`

NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse instantiates a new KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse`

NewKeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetExpirationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetExpirationWarningDays() int32`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetExpirationWarningDaysOk() (*int32, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetExpirationWarningDays(v int32)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.

### HasExpirationWarningDays

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasExpirationWarningDays() bool`

HasExpirationWarningDays returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetCertificateQuery

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetCertificateQuery() KeyfactorWebKeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse`

GetCertificateQuery returns the CertificateQuery field if non-nil, zero value otherwise.

### GetCertificateQueryOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetCertificateQueryOk() (*KeyfactorWebKeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse, bool)`

GetCertificateQueryOk returns a tuple with the CertificateQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQuery

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetCertificateQuery(v KeyfactorWebKeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse)`

SetCertificateQuery sets CertificateQuery field to given value.

### HasCertificateQuery

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasCertificateQuery() bool`

HasCertificateQuery returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorWebKeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorWebKeyfactorApiModelsEventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.

### SetEventHandlerParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetEventHandlerParametersNil(b bool)`

 SetEventHandlerParametersNil sets the value for EventHandlerParameters to be an explicit nil

### UnsetEventHandlerParameters
`func (o *KeyfactorWebKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) UnsetEventHandlerParameters()`

UnsetEventHandlerParameters ensures that no value is present for EventHandlerParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


