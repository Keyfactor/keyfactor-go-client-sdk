# KeyfactorApiModelsAlertsExpirationExpirationAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CAName** | Pointer to **string** |  | [optional] 
**CARow** | Pointer to **int64** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**Expiry** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**SendDate** | Pointer to **string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponse

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponse() *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponse instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponseWithDefaults

`func NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponseWithDefaults() *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse`

NewKeyfactorApiModelsAlertsExpirationExpirationAlertResponseWithDefaults instantiates a new KeyfactorApiModelsAlertsExpirationExpirationAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCAName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCAName() string`

GetCAName returns the CAName field if non-nil, zero value otherwise.

### GetCANameOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCANameOk() (*string, bool)`

GetCANameOk returns a tuple with the CAName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetCAName(v string)`

SetCAName sets CAName field to given value.

### HasCAName

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasCAName() bool`

HasCAName returns a boolean if a field has been set.

### GetCARow

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCARow() int64`

GetCARow returns the CARow field if non-nil, zero value otherwise.

### GetCARowOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetCARowOk() (*int64, bool)`

GetCARowOk returns a tuple with the CARow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARow

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetCARow(v int64)`

SetCARow sets CARow field to given value.

### HasCARow

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasCARow() bool`

HasCARow returns a boolean if a field has been set.

### GetIssuedCN

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetExpiry

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSendDate

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSendDate() string`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) GetSendDateOk() (*string, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) SetSendDate(v string)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *KeyfactorApiModelsAlertsExpirationExpirationAlertResponse) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


