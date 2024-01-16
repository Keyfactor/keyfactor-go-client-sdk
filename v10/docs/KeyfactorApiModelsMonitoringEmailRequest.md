# KeyfactorApiModelsMonitoringEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableReminder** | Pointer to **bool** |  | [optional] 
**WarningDays** | Pointer to **int64** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 

## Methods

### NewKeyfactorApiModelsMonitoringEmailRequest

`func NewKeyfactorApiModelsMonitoringEmailRequest() *KeyfactorApiModelsMonitoringEmailRequest`

NewKeyfactorApiModelsMonitoringEmailRequest instantiates a new KeyfactorApiModelsMonitoringEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsMonitoringEmailRequestWithDefaults

`func NewKeyfactorApiModelsMonitoringEmailRequestWithDefaults() *KeyfactorApiModelsMonitoringEmailRequest`

NewKeyfactorApiModelsMonitoringEmailRequestWithDefaults instantiates a new KeyfactorApiModelsMonitoringEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableReminder

`func (o *KeyfactorApiModelsMonitoringEmailRequest) GetEnableReminder() bool`

GetEnableReminder returns the EnableReminder field if non-nil, zero value otherwise.

### GetEnableReminderOk

`func (o *KeyfactorApiModelsMonitoringEmailRequest) GetEnableReminderOk() (*bool, bool)`

GetEnableReminderOk returns a tuple with the EnableReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReminder

`func (o *KeyfactorApiModelsMonitoringEmailRequest) SetEnableReminder(v bool)`

SetEnableReminder sets EnableReminder field to given value.

### HasEnableReminder

`func (o *KeyfactorApiModelsMonitoringEmailRequest) HasEnableReminder() bool`

HasEnableReminder returns a boolean if a field has been set.

### GetWarningDays

`func (o *KeyfactorApiModelsMonitoringEmailRequest) GetWarningDays() int64`

GetWarningDays returns the WarningDays field if non-nil, zero value otherwise.

### GetWarningDaysOk

`func (o *KeyfactorApiModelsMonitoringEmailRequest) GetWarningDaysOk() (*int64, bool)`

GetWarningDaysOk returns a tuple with the WarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDays

`func (o *KeyfactorApiModelsMonitoringEmailRequest) SetWarningDays(v int64)`

SetWarningDays sets WarningDays field to given value.

### HasWarningDays

`func (o *KeyfactorApiModelsMonitoringEmailRequest) HasWarningDays() bool`

HasWarningDays returns a boolean if a field has been set.

### GetRecipients

`func (o *KeyfactorApiModelsMonitoringEmailRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *KeyfactorApiModelsMonitoringEmailRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *KeyfactorApiModelsMonitoringEmailRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *KeyfactorApiModelsMonitoringEmailRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


