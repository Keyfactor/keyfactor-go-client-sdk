# KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ExpirationWarningDays** | Pointer to **int64** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**CertificateQuery** | Pointer to [**KeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse**](KeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse.md) |  | [optional] 
**RegisteredEventHandler** | Pointer to [**KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse**](KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse.md) |  | [optional] 
**EventHandlerParameters** | Pointer to [**[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse**](KeyfactorApiModelsEventHandlerEventHandlerParameterResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse() *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponseWithDefaults

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponseWithDefaults() *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetExpirationWarningDays() int64`

GetExpirationWarningDays returns the ExpirationWarningDays field if non-nil, zero value otherwise.

### GetExpirationWarningDaysOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetExpirationWarningDaysOk() (*int64, bool)`

GetExpirationWarningDaysOk returns a tuple with the ExpirationWarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetExpirationWarningDays(v int64)`

SetExpirationWarningDays sets ExpirationWarningDays field to given value.

### HasExpirationWarningDays

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasExpirationWarningDays() bool`

HasExpirationWarningDays returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetCertificateQuery

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetCertificateQuery() KeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse`

GetCertificateQuery returns the CertificateQuery field if non-nil, zero value otherwise.

### GetCertificateQueryOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetCertificateQueryOk() (*KeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse, bool)`

GetCertificateQueryOk returns a tuple with the CertificateQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateQuery

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetCertificateQuery(v KeyfactorApiModelsAlertsAlertCertificateQueryAlertCertificateQueryResponse)`

SetCertificateQuery sets CertificateQuery field to given value.

### HasCertificateQuery

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasCertificateQuery() bool`

HasCertificateQuery returns a boolean if a field has been set.

### GetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRegisteredEventHandler() KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse`

GetRegisteredEventHandler returns the RegisteredEventHandler field if non-nil, zero value otherwise.

### GetRegisteredEventHandlerOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetRegisteredEventHandlerOk() (*KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse, bool)`

GetRegisteredEventHandlerOk returns a tuple with the RegisteredEventHandler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetRegisteredEventHandler(v KeyfactorApiModelsEventHandlerRegisteredEventHandlerResponse)`

SetRegisteredEventHandler sets RegisteredEventHandler field to given value.

### HasRegisteredEventHandler

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasRegisteredEventHandler() bool`

HasRegisteredEventHandler returns a boolean if a field has been set.

### GetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetEventHandlerParameters() []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse`

GetEventHandlerParameters returns the EventHandlerParameters field if non-nil, zero value otherwise.

### GetEventHandlerParametersOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) GetEventHandlerParametersOk() (*[]KeyfactorApiModelsEventHandlerEventHandlerParameterResponse, bool)`

GetEventHandlerParametersOk returns a tuple with the EventHandlerParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) SetEventHandlerParameters(v []KeyfactorApiModelsEventHandlerEventHandlerParameterResponse)`

SetEventHandlerParameters sets EventHandlerParameters field to given value.

### HasEventHandlerParameters

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertDefinitionResponse) HasEventHandlerParameters() bool`

HasEventHandlerParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


