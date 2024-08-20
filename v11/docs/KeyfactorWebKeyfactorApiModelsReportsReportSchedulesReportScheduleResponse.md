# KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse

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

### NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse

`func NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse() *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse`

NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse instantiates a new KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse`

NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSendReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetSendReport() bool`

GetSendReport returns the SendReport field if non-nil, zero value otherwise.

### GetSendReportOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetSendReportOk() (*bool, bool)`

GetSendReportOk returns a tuple with the SendReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetSendReport(v bool)`

SetSendReport sets SendReport field to given value.

### HasSendReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasSendReport() bool`

HasSendReport returns a boolean if a field has been set.

### GetSaveReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetSaveReport() bool`

GetSaveReport returns the SaveReport field if non-nil, zero value otherwise.

### GetSaveReportOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetSaveReportOk() (*bool, bool)`

GetSaveReportOk returns a tuple with the SaveReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetSaveReport(v bool)`

SetSaveReport sets SaveReport field to given value.

### HasSaveReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasSaveReport() bool`

HasSaveReport returns a boolean if a field has been set.

### GetSaveReportPath

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetSaveReportPath() string`

GetSaveReportPath returns the SaveReportPath field if non-nil, zero value otherwise.

### GetSaveReportPathOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetSaveReportPathOk() (*string, bool)`

GetSaveReportPathOk returns a tuple with the SaveReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReportPath

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetSaveReportPath(v string)`

SetSaveReportPath sets SaveReportPath field to given value.

### HasSaveReportPath

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasSaveReportPath() bool`

HasSaveReportPath returns a boolean if a field has been set.

### SetSaveReportPathNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetSaveReportPathNil(b bool)`

 SetSaveReportPathNil sets the value for SaveReportPath to be an explicit nil

### UnsetSaveReportPath
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) UnsetSaveReportPath()`

UnsetSaveReportPath ensures that no value is present for SaveReportPath, not even an explicit nil
### GetReportFormat

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.

### SetReportFormatNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetReportFormatNil(b bool)`

 SetReportFormatNil sets the value for ReportFormat to be an explicit nil

### UnsetReportFormat
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) UnsetReportFormat()`

UnsetReportFormat ensures that no value is present for ReportFormat, not even an explicit nil
### GetKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetCertificateCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetCertificateCollectionId() int32`

GetCertificateCollectionId returns the CertificateCollectionId field if non-nil, zero value otherwise.

### GetCertificateCollectionIdOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetCertificateCollectionIdOk() (*int32, bool)`

GetCertificateCollectionIdOk returns a tuple with the CertificateCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetCertificateCollectionId(v int32)`

SetCertificateCollectionId sets CertificateCollectionId field to given value.

### HasCertificateCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasCertificateCollectionId() bool`

HasCertificateCollectionId returns a boolean if a field has been set.

### GetEmailRecipients

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetEmailRecipients() []string`

GetEmailRecipients returns the EmailRecipients field if non-nil, zero value otherwise.

### GetEmailRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetEmailRecipientsOk() (*[]string, bool)`

GetEmailRecipientsOk returns a tuple with the EmailRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipients

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetEmailRecipients(v []string)`

SetEmailRecipients sets EmailRecipients field to given value.

### HasEmailRecipients

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasEmailRecipients() bool`

HasEmailRecipients returns a boolean if a field has been set.

### SetEmailRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetEmailRecipientsNil(b bool)`

 SetEmailRecipientsNil sets the value for EmailRecipients to be an explicit nil

### UnsetEmailRecipients
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) UnsetEmailRecipients()`

UnsetEmailRecipients ensures that no value is present for EmailRecipients, not even an explicit nil
### GetRuntimeParameters

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetRuntimeParameters() map[string]string`

GetRuntimeParameters returns the RuntimeParameters field if non-nil, zero value otherwise.

### GetRuntimeParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) GetRuntimeParametersOk() (*map[string]string, bool)`

GetRuntimeParametersOk returns a tuple with the RuntimeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeParameters

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetRuntimeParameters(v map[string]string)`

SetRuntimeParameters sets RuntimeParameters field to given value.

### HasRuntimeParameters

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) HasRuntimeParameters() bool`

HasRuntimeParameters returns a boolean if a field has been set.

### SetRuntimeParametersNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) SetRuntimeParametersNil(b bool)`

 SetRuntimeParametersNil sets the value for RuntimeParameters to be an explicit nil

### UnsetRuntimeParameters
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleResponse) UnsetRuntimeParameters()`

UnsetRuntimeParameters ensures that no value is present for RuntimeParameters, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


