# KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**SendReport** | Pointer to **bool** |  | [optional] 
**SaveReport** | Pointer to **bool** |  | [optional] 
**SaveReportPath** | Pointer to **NullableString** |  | [optional] 
**ReportFormat** | **string** |  | 
**KeyfactorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateCollectionId** | Pointer to **NullableInt32** |  | [optional] 
**EmailRecipients** | Pointer to **[]string** |  | [optional] 
**RuntimeParameters** | **map[string]string** |  | 

## Methods

### NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest

`func NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest(id int32, reportFormat string, runtimeParameters map[string]string, ) *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest instantiates a new KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequestWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequestWithDefaults() *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest`

NewKeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequestWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.


### GetSendReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetSendReport() bool`

GetSendReport returns the SendReport field if non-nil, zero value otherwise.

### GetSendReportOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetSendReportOk() (*bool, bool)`

GetSendReportOk returns a tuple with the SendReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetSendReport(v bool)`

SetSendReport sets SendReport field to given value.

### HasSendReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) HasSendReport() bool`

HasSendReport returns a boolean if a field has been set.

### GetSaveReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetSaveReport() bool`

GetSaveReport returns the SaveReport field if non-nil, zero value otherwise.

### GetSaveReportOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetSaveReportOk() (*bool, bool)`

GetSaveReportOk returns a tuple with the SaveReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetSaveReport(v bool)`

SetSaveReport sets SaveReport field to given value.

### HasSaveReport

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) HasSaveReport() bool`

HasSaveReport returns a boolean if a field has been set.

### GetSaveReportPath

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetSaveReportPath() string`

GetSaveReportPath returns the SaveReportPath field if non-nil, zero value otherwise.

### GetSaveReportPathOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetSaveReportPathOk() (*string, bool)`

GetSaveReportPathOk returns a tuple with the SaveReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReportPath

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetSaveReportPath(v string)`

SetSaveReportPath sets SaveReportPath field to given value.

### HasSaveReportPath

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) HasSaveReportPath() bool`

HasSaveReportPath returns a boolean if a field has been set.

### SetSaveReportPathNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetSaveReportPathNil(b bool)`

 SetSaveReportPathNil sets the value for SaveReportPath to be an explicit nil

### UnsetSaveReportPath
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) UnsetSaveReportPath()`

UnsetSaveReportPath ensures that no value is present for SaveReportPath, not even an explicit nil
### GetReportFormat

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.


### GetKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetCertificateCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetCertificateCollectionId() int32`

GetCertificateCollectionId returns the CertificateCollectionId field if non-nil, zero value otherwise.

### GetCertificateCollectionIdOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetCertificateCollectionIdOk() (*int32, bool)`

GetCertificateCollectionIdOk returns a tuple with the CertificateCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetCertificateCollectionId(v int32)`

SetCertificateCollectionId sets CertificateCollectionId field to given value.

### HasCertificateCollectionId

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) HasCertificateCollectionId() bool`

HasCertificateCollectionId returns a boolean if a field has been set.

### SetCertificateCollectionIdNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetCertificateCollectionIdNil(b bool)`

 SetCertificateCollectionIdNil sets the value for CertificateCollectionId to be an explicit nil

### UnsetCertificateCollectionId
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) UnsetCertificateCollectionId()`

UnsetCertificateCollectionId ensures that no value is present for CertificateCollectionId, not even an explicit nil
### GetEmailRecipients

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetEmailRecipients() []string`

GetEmailRecipients returns the EmailRecipients field if non-nil, zero value otherwise.

### GetEmailRecipientsOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetEmailRecipientsOk() (*[]string, bool)`

GetEmailRecipientsOk returns a tuple with the EmailRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipients

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetEmailRecipients(v []string)`

SetEmailRecipients sets EmailRecipients field to given value.

### HasEmailRecipients

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) HasEmailRecipients() bool`

HasEmailRecipients returns a boolean if a field has been set.

### SetEmailRecipientsNil

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetEmailRecipientsNil(b bool)`

 SetEmailRecipientsNil sets the value for EmailRecipients to be an explicit nil

### UnsetEmailRecipients
`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) UnsetEmailRecipients()`

UnsetEmailRecipients ensures that no value is present for EmailRecipients, not even an explicit nil
### GetRuntimeParameters

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetRuntimeParameters() map[string]string`

GetRuntimeParameters returns the RuntimeParameters field if non-nil, zero value otherwise.

### GetRuntimeParametersOk

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) GetRuntimeParametersOk() (*map[string]string, bool)`

GetRuntimeParametersOk returns a tuple with the RuntimeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeParameters

`func (o *KeyfactorWebKeyfactorApiModelsReportsReportSchedulesReportScheduleUpdateRequest) SetRuntimeParameters(v map[string]string)`

SetRuntimeParameters sets RuntimeParameters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


