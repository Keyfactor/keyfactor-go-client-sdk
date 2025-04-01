# MonitoringEmailRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableReminder** | Pointer to **bool** |  | [optional] 
**WarningDays** | Pointer to **int32** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 

## Methods

### NewMonitoringEmailRequest

`func NewMonitoringEmailRequest() *MonitoringEmailRequest`

NewMonitoringEmailRequest instantiates a new MonitoringEmailRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonitoringEmailRequestWithDefaults

`func NewMonitoringEmailRequestWithDefaults() *MonitoringEmailRequest`

NewMonitoringEmailRequestWithDefaults instantiates a new MonitoringEmailRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableReminder

`func (o *MonitoringEmailRequest) GetEnableReminder() bool`

GetEnableReminder returns the EnableReminder field if non-nil, zero value otherwise.

### GetEnableReminderOk

`func (o *MonitoringEmailRequest) GetEnableReminderOk() (*bool, bool)`

GetEnableReminderOk returns a tuple with the EnableReminder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableReminder

`func (o *MonitoringEmailRequest) SetEnableReminder(v bool)`

SetEnableReminder sets EnableReminder field to given value.

### HasEnableReminder

`func (o *MonitoringEmailRequest) HasEnableReminder() bool`

HasEnableReminder returns a boolean if a field has been set.

### GetWarningDays

`func (o *MonitoringEmailRequest) GetWarningDays() int32`

GetWarningDays returns the WarningDays field if non-nil, zero value otherwise.

### GetWarningDaysOk

`func (o *MonitoringEmailRequest) GetWarningDaysOk() (*int32, bool)`

GetWarningDaysOk returns a tuple with the WarningDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningDays

`func (o *MonitoringEmailRequest) SetWarningDays(v int32)`

SetWarningDays sets WarningDays field to given value.

### HasWarningDays

`func (o *MonitoringEmailRequest) HasWarningDays() bool`

HasWarningDays returns a boolean if a field has been set.

### GetRecipients

`func (o *MonitoringEmailRequest) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *MonitoringEmailRequest) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *MonitoringEmailRequest) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *MonitoringEmailRequest) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### SetRecipientsNil

`func (o *MonitoringEmailRequest) SetRecipientsNil(b bool)`

 SetRecipientsNil sets the value for Recipients to be an explicit nil

### UnsetRecipients
`func (o *MonitoringEmailRequest) UnsetRecipients()`

UnsetRecipients ensures that no value is present for Recipients, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


