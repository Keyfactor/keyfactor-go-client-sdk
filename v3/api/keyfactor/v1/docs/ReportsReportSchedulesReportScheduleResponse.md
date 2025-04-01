# ReportsReportSchedulesReportScheduleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SendReport** | Pointer to **bool** |  | [optional] 
**SaveReport** | Pointer to **bool** |  | [optional] 
**SaveReportPath** | Pointer to **NullableString** |  | [optional] 
**ReportFormat** | Pointer to **NullableString** |  | [optional] 
**KeyfactorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateCollectionId** | Pointer to **int32** |  | [optional] 
**EmailRecipients** | Pointer to **[]string** |  | [optional] 
**RuntimeParameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewReportsReportSchedulesReportScheduleResponse

`func NewReportsReportSchedulesReportScheduleResponse() *ReportsReportSchedulesReportScheduleResponse`

NewReportsReportSchedulesReportScheduleResponse instantiates a new ReportsReportSchedulesReportScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReportsReportSchedulesReportScheduleResponseWithDefaults

`func NewReportsReportSchedulesReportScheduleResponseWithDefaults() *ReportsReportSchedulesReportScheduleResponse`

NewReportsReportSchedulesReportScheduleResponseWithDefaults instantiates a new ReportsReportSchedulesReportScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReportsReportSchedulesReportScheduleResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReportsReportSchedulesReportScheduleResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ReportsReportSchedulesReportScheduleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSendReport

`func (o *ReportsReportSchedulesReportScheduleResponse) GetSendReport() bool`

GetSendReport returns the SendReport field if non-nil, zero value otherwise.

### GetSendReportOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetSendReportOk() (*bool, bool)`

GetSendReportOk returns a tuple with the SendReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReport

`func (o *ReportsReportSchedulesReportScheduleResponse) SetSendReport(v bool)`

SetSendReport sets SendReport field to given value.

### HasSendReport

`func (o *ReportsReportSchedulesReportScheduleResponse) HasSendReport() bool`

HasSendReport returns a boolean if a field has been set.

### GetSaveReport

`func (o *ReportsReportSchedulesReportScheduleResponse) GetSaveReport() bool`

GetSaveReport returns the SaveReport field if non-nil, zero value otherwise.

### GetSaveReportOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetSaveReportOk() (*bool, bool)`

GetSaveReportOk returns a tuple with the SaveReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReport

`func (o *ReportsReportSchedulesReportScheduleResponse) SetSaveReport(v bool)`

SetSaveReport sets SaveReport field to given value.

### HasSaveReport

`func (o *ReportsReportSchedulesReportScheduleResponse) HasSaveReport() bool`

HasSaveReport returns a boolean if a field has been set.

### GetSaveReportPath

`func (o *ReportsReportSchedulesReportScheduleResponse) GetSaveReportPath() string`

GetSaveReportPath returns the SaveReportPath field if non-nil, zero value otherwise.

### GetSaveReportPathOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetSaveReportPathOk() (*string, bool)`

GetSaveReportPathOk returns a tuple with the SaveReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReportPath

`func (o *ReportsReportSchedulesReportScheduleResponse) SetSaveReportPath(v string)`

SetSaveReportPath sets SaveReportPath field to given value.

### HasSaveReportPath

`func (o *ReportsReportSchedulesReportScheduleResponse) HasSaveReportPath() bool`

HasSaveReportPath returns a boolean if a field has been set.

### SetSaveReportPathNil

`func (o *ReportsReportSchedulesReportScheduleResponse) SetSaveReportPathNil(b bool)`

 SetSaveReportPathNil sets the value for SaveReportPath to be an explicit nil

### UnsetSaveReportPath
`func (o *ReportsReportSchedulesReportScheduleResponse) UnsetSaveReportPath()`

UnsetSaveReportPath ensures that no value is present for SaveReportPath, not even an explicit nil
### GetReportFormat

`func (o *ReportsReportSchedulesReportScheduleResponse) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *ReportsReportSchedulesReportScheduleResponse) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *ReportsReportSchedulesReportScheduleResponse) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.

### SetReportFormatNil

`func (o *ReportsReportSchedulesReportScheduleResponse) SetReportFormatNil(b bool)`

 SetReportFormatNil sets the value for ReportFormat to be an explicit nil

### UnsetReportFormat
`func (o *ReportsReportSchedulesReportScheduleResponse) UnsetReportFormat()`

UnsetReportFormat ensures that no value is present for ReportFormat, not even an explicit nil
### GetKeyfactorSchedule

`func (o *ReportsReportSchedulesReportScheduleResponse) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *ReportsReportSchedulesReportScheduleResponse) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *ReportsReportSchedulesReportScheduleResponse) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetCertificateCollectionId

`func (o *ReportsReportSchedulesReportScheduleResponse) GetCertificateCollectionId() int32`

GetCertificateCollectionId returns the CertificateCollectionId field if non-nil, zero value otherwise.

### GetCertificateCollectionIdOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetCertificateCollectionIdOk() (*int32, bool)`

GetCertificateCollectionIdOk returns a tuple with the CertificateCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionId

`func (o *ReportsReportSchedulesReportScheduleResponse) SetCertificateCollectionId(v int32)`

SetCertificateCollectionId sets CertificateCollectionId field to given value.

### HasCertificateCollectionId

`func (o *ReportsReportSchedulesReportScheduleResponse) HasCertificateCollectionId() bool`

HasCertificateCollectionId returns a boolean if a field has been set.

### GetEmailRecipients

`func (o *ReportsReportSchedulesReportScheduleResponse) GetEmailRecipients() []string`

GetEmailRecipients returns the EmailRecipients field if non-nil, zero value otherwise.

### GetEmailRecipientsOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetEmailRecipientsOk() (*[]string, bool)`

GetEmailRecipientsOk returns a tuple with the EmailRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipients

`func (o *ReportsReportSchedulesReportScheduleResponse) SetEmailRecipients(v []string)`

SetEmailRecipients sets EmailRecipients field to given value.

### HasEmailRecipients

`func (o *ReportsReportSchedulesReportScheduleResponse) HasEmailRecipients() bool`

HasEmailRecipients returns a boolean if a field has been set.

### SetEmailRecipientsNil

`func (o *ReportsReportSchedulesReportScheduleResponse) SetEmailRecipientsNil(b bool)`

 SetEmailRecipientsNil sets the value for EmailRecipients to be an explicit nil

### UnsetEmailRecipients
`func (o *ReportsReportSchedulesReportScheduleResponse) UnsetEmailRecipients()`

UnsetEmailRecipients ensures that no value is present for EmailRecipients, not even an explicit nil
### GetRuntimeParameters

`func (o *ReportsReportSchedulesReportScheduleResponse) GetRuntimeParameters() map[string]string`

GetRuntimeParameters returns the RuntimeParameters field if non-nil, zero value otherwise.

### GetRuntimeParametersOk

`func (o *ReportsReportSchedulesReportScheduleResponse) GetRuntimeParametersOk() (*map[string]string, bool)`

GetRuntimeParametersOk returns a tuple with the RuntimeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeParameters

`func (o *ReportsReportSchedulesReportScheduleResponse) SetRuntimeParameters(v map[string]string)`

SetRuntimeParameters sets RuntimeParameters field to given value.

### HasRuntimeParameters

`func (o *ReportsReportSchedulesReportScheduleResponse) HasRuntimeParameters() bool`

HasRuntimeParameters returns a boolean if a field has been set.

### SetRuntimeParametersNil

`func (o *ReportsReportSchedulesReportScheduleResponse) SetRuntimeParametersNil(b bool)`

 SetRuntimeParametersNil sets the value for RuntimeParameters to be an explicit nil

### UnsetRuntimeParameters
`func (o *ReportsReportSchedulesReportScheduleResponse) UnsetRuntimeParameters()`

UnsetRuntimeParameters ensures that no value is present for RuntimeParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


