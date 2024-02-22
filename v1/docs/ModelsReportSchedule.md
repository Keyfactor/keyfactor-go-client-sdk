# ModelsReportSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SendReport** | Pointer to **bool** |  | [optional] 
**SaveReport** | Pointer to **bool** |  | [optional] 
**SaveReportPath** | Pointer to **string** |  | [optional] 
**ReportFormat** | Pointer to **string** |  | [optional] 
**KeyfactorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateCollectionId** | Pointer to **int32** |  | [optional] 
**EmailRecipients** | Pointer to **[]string** |  | [optional] 
**RuntimeParameters** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewModelsReportSchedule

`func NewModelsReportSchedule() *ModelsReportSchedule`

NewModelsReportSchedule instantiates a new ModelsReportSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsReportScheduleWithDefaults

`func NewModelsReportScheduleWithDefaults() *ModelsReportSchedule`

NewModelsReportScheduleWithDefaults instantiates a new ModelsReportSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsReportSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsReportSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsReportSchedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsReportSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSendReport

`func (o *ModelsReportSchedule) GetSendReport() bool`

GetSendReport returns the SendReport field if non-nil, zero value otherwise.

### GetSendReportOk

`func (o *ModelsReportSchedule) GetSendReportOk() (*bool, bool)`

GetSendReportOk returns a tuple with the SendReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReport

`func (o *ModelsReportSchedule) SetSendReport(v bool)`

SetSendReport sets SendReport field to given value.

### HasSendReport

`func (o *ModelsReportSchedule) HasSendReport() bool`

HasSendReport returns a boolean if a field has been set.

### GetSaveReport

`func (o *ModelsReportSchedule) GetSaveReport() bool`

GetSaveReport returns the SaveReport field if non-nil, zero value otherwise.

### GetSaveReportOk

`func (o *ModelsReportSchedule) GetSaveReportOk() (*bool, bool)`

GetSaveReportOk returns a tuple with the SaveReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReport

`func (o *ModelsReportSchedule) SetSaveReport(v bool)`

SetSaveReport sets SaveReport field to given value.

### HasSaveReport

`func (o *ModelsReportSchedule) HasSaveReport() bool`

HasSaveReport returns a boolean if a field has been set.

### GetSaveReportPath

`func (o *ModelsReportSchedule) GetSaveReportPath() string`

GetSaveReportPath returns the SaveReportPath field if non-nil, zero value otherwise.

### GetSaveReportPathOk

`func (o *ModelsReportSchedule) GetSaveReportPathOk() (*string, bool)`

GetSaveReportPathOk returns a tuple with the SaveReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReportPath

`func (o *ModelsReportSchedule) SetSaveReportPath(v string)`

SetSaveReportPath sets SaveReportPath field to given value.

### HasSaveReportPath

`func (o *ModelsReportSchedule) HasSaveReportPath() bool`

HasSaveReportPath returns a boolean if a field has been set.

### GetReportFormat

`func (o *ModelsReportSchedule) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *ModelsReportSchedule) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *ModelsReportSchedule) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *ModelsReportSchedule) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.

### GetKeyfactorSchedule

`func (o *ModelsReportSchedule) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *ModelsReportSchedule) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *ModelsReportSchedule) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *ModelsReportSchedule) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetCertificateCollectionId

`func (o *ModelsReportSchedule) GetCertificateCollectionId() int32`

GetCertificateCollectionId returns the CertificateCollectionId field if non-nil, zero value otherwise.

### GetCertificateCollectionIdOk

`func (o *ModelsReportSchedule) GetCertificateCollectionIdOk() (*int32, bool)`

GetCertificateCollectionIdOk returns a tuple with the CertificateCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionId

`func (o *ModelsReportSchedule) SetCertificateCollectionId(v int32)`

SetCertificateCollectionId sets CertificateCollectionId field to given value.

### HasCertificateCollectionId

`func (o *ModelsReportSchedule) HasCertificateCollectionId() bool`

HasCertificateCollectionId returns a boolean if a field has been set.

### GetEmailRecipients

`func (o *ModelsReportSchedule) GetEmailRecipients() []string`

GetEmailRecipients returns the EmailRecipients field if non-nil, zero value otherwise.

### GetEmailRecipientsOk

`func (o *ModelsReportSchedule) GetEmailRecipientsOk() (*[]string, bool)`

GetEmailRecipientsOk returns a tuple with the EmailRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipients

`func (o *ModelsReportSchedule) SetEmailRecipients(v []string)`

SetEmailRecipients sets EmailRecipients field to given value.

### HasEmailRecipients

`func (o *ModelsReportSchedule) HasEmailRecipients() bool`

HasEmailRecipients returns a boolean if a field has been set.

### GetRuntimeParameters

`func (o *ModelsReportSchedule) GetRuntimeParameters() map[string]string`

GetRuntimeParameters returns the RuntimeParameters field if non-nil, zero value otherwise.

### GetRuntimeParametersOk

`func (o *ModelsReportSchedule) GetRuntimeParametersOk() (*map[string]string, bool)`

GetRuntimeParametersOk returns a tuple with the RuntimeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeParameters

`func (o *ModelsReportSchedule) SetRuntimeParameters(v map[string]string)`

SetRuntimeParameters sets RuntimeParameters field to given value.

### HasRuntimeParameters

`func (o *ModelsReportSchedule) HasRuntimeParameters() bool`

HasRuntimeParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


