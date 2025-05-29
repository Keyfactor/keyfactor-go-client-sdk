# AlertsExpirationExpirationAlertResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CAName** | Pointer to **NullableString** |  | [optional] 
**CARow** | Pointer to **int64** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**Expiry** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Message** | Pointer to **NullableString** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**SendDate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAlertsExpirationExpirationAlertResponse

`func NewAlertsExpirationExpirationAlertResponse() *AlertsExpirationExpirationAlertResponse`

NewAlertsExpirationExpirationAlertResponse instantiates a new AlertsExpirationExpirationAlertResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertsExpirationExpirationAlertResponseWithDefaults

`func NewAlertsExpirationExpirationAlertResponseWithDefaults() *AlertsExpirationExpirationAlertResponse`

NewAlertsExpirationExpirationAlertResponseWithDefaults instantiates a new AlertsExpirationExpirationAlertResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCAName

`func (o *AlertsExpirationExpirationAlertResponse) GetCAName() string`

GetCAName returns the CAName field if non-nil, zero value otherwise.

### GetCANameOk

`func (o *AlertsExpirationExpirationAlertResponse) GetCANameOk() (*string, bool)`

GetCANameOk returns a tuple with the CAName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCAName

`func (o *AlertsExpirationExpirationAlertResponse) SetCAName(v string)`

SetCAName sets CAName field to given value.

### HasCAName

`func (o *AlertsExpirationExpirationAlertResponse) HasCAName() bool`

HasCAName returns a boolean if a field has been set.

### SetCANameNil

`func (o *AlertsExpirationExpirationAlertResponse) SetCANameNil(b bool)`

 SetCANameNil sets the value for CAName to be an explicit nil

### UnsetCAName
`func (o *AlertsExpirationExpirationAlertResponse) UnsetCAName()`

UnsetCAName ensures that no value is present for CAName, not even an explicit nil
### GetCARow

`func (o *AlertsExpirationExpirationAlertResponse) GetCARow() int64`

GetCARow returns the CARow field if non-nil, zero value otherwise.

### GetCARowOk

`func (o *AlertsExpirationExpirationAlertResponse) GetCARowOk() (*int64, bool)`

GetCARowOk returns a tuple with the CARow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCARow

`func (o *AlertsExpirationExpirationAlertResponse) SetCARow(v int64)`

SetCARow sets CARow field to given value.

### HasCARow

`func (o *AlertsExpirationExpirationAlertResponse) HasCARow() bool`

HasCARow returns a boolean if a field has been set.

### GetIssuedCN

`func (o *AlertsExpirationExpirationAlertResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *AlertsExpirationExpirationAlertResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *AlertsExpirationExpirationAlertResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *AlertsExpirationExpirationAlertResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *AlertsExpirationExpirationAlertResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *AlertsExpirationExpirationAlertResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetExpiry

`func (o *AlertsExpirationExpirationAlertResponse) GetExpiry() string`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *AlertsExpirationExpirationAlertResponse) GetExpiryOk() (*string, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *AlertsExpirationExpirationAlertResponse) SetExpiry(v string)`

SetExpiry sets Expiry field to given value.

### HasExpiry

`func (o *AlertsExpirationExpirationAlertResponse) HasExpiry() bool`

HasExpiry returns a boolean if a field has been set.

### SetExpiryNil

`func (o *AlertsExpirationExpirationAlertResponse) SetExpiryNil(b bool)`

 SetExpiryNil sets the value for Expiry to be an explicit nil

### UnsetExpiry
`func (o *AlertsExpirationExpirationAlertResponse) UnsetExpiry()`

UnsetExpiry ensures that no value is present for Expiry, not even an explicit nil
### GetSubject

`func (o *AlertsExpirationExpirationAlertResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *AlertsExpirationExpirationAlertResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *AlertsExpirationExpirationAlertResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *AlertsExpirationExpirationAlertResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *AlertsExpirationExpirationAlertResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *AlertsExpirationExpirationAlertResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetMessage

`func (o *AlertsExpirationExpirationAlertResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AlertsExpirationExpirationAlertResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AlertsExpirationExpirationAlertResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AlertsExpirationExpirationAlertResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *AlertsExpirationExpirationAlertResponse) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *AlertsExpirationExpirationAlertResponse) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetRecipients

`func (o *AlertsExpirationExpirationAlertResponse) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *AlertsExpirationExpirationAlertResponse) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *AlertsExpirationExpirationAlertResponse) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *AlertsExpirationExpirationAlertResponse) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *AlertsExpirationExpirationAlertResponse) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *AlertsExpirationExpirationAlertResponse) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil
### GetSendDate

`func (o *AlertsExpirationExpirationAlertResponse) GetSendDate() string`

GetSendDate returns the SendDate field if non-nil, zero value otherwise.

### GetSendDateOk

`func (o *AlertsExpirationExpirationAlertResponse) GetSendDateOk() (*string, bool)`

GetSendDateOk returns a tuple with the SendDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDate

`func (o *AlertsExpirationExpirationAlertResponse) SetSendDate(v string)`

SetSendDate sets SendDate field to given value.

### HasSendDate

`func (o *AlertsExpirationExpirationAlertResponse) HasSendDate() bool`

HasSendDate returns a boolean if a field has been set.

### SetSendDateNil

`func (o *AlertsExpirationExpirationAlertResponse) SetSendDateNil(b bool)`

 SetSendDateNil sets the value for SendDate to be an explicit nil

### UnsetSendDate
`func (o *AlertsExpirationExpirationAlertResponse) UnsetSendDate()`

UnsetSendDate ensures that no value is present for SendDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


