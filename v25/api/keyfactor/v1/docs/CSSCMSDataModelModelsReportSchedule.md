# CSSCMSDataModelModelsReportSchedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**SendReport** | Pointer to **bool** |  | [optional] 
**SaveReport** | Pointer to **bool** |  | [optional] 
**SaveReportPath** | Pointer to **NullableString** |  | [optional] 
**ReportFormat** | Pointer to **NullableString** |  | [optional] 
**KeyfactorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**CertificateCollectionId** | Pointer to **NullableInt32** |  | [optional] 
**CertificateCollectionName** | Pointer to **NullableString** |  | [optional] 
**EmailRecipients** | Pointer to **[]string** |  | [optional] 
**RuntimeParameters** | Pointer to **map[string]string** |  | [optional] 
**ReportDefinitionId** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] [readonly] 
**Enabled** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewCSSCMSDataModelModelsReportSchedule

`func NewCSSCMSDataModelModelsReportSchedule() *CSSCMSDataModelModelsReportSchedule`

NewCSSCMSDataModelModelsReportSchedule instantiates a new CSSCMSDataModelModelsReportSchedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsReportScheduleWithDefaults

`func NewCSSCMSDataModelModelsReportScheduleWithDefaults() *CSSCMSDataModelModelsReportSchedule`

NewCSSCMSDataModelModelsReportScheduleWithDefaults instantiates a new CSSCMSDataModelModelsReportSchedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsReportSchedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsReportSchedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsReportSchedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSendReport

`func (o *CSSCMSDataModelModelsReportSchedule) GetSendReport() bool`

GetSendReport returns the SendReport field if non-nil, zero value otherwise.

### GetSendReportOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetSendReportOk() (*bool, bool)`

GetSendReportOk returns a tuple with the SendReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendReport

`func (o *CSSCMSDataModelModelsReportSchedule) SetSendReport(v bool)`

SetSendReport sets SendReport field to given value.

### HasSendReport

`func (o *CSSCMSDataModelModelsReportSchedule) HasSendReport() bool`

HasSendReport returns a boolean if a field has been set.

### GetSaveReport

`func (o *CSSCMSDataModelModelsReportSchedule) GetSaveReport() bool`

GetSaveReport returns the SaveReport field if non-nil, zero value otherwise.

### GetSaveReportOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetSaveReportOk() (*bool, bool)`

GetSaveReportOk returns a tuple with the SaveReport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReport

`func (o *CSSCMSDataModelModelsReportSchedule) SetSaveReport(v bool)`

SetSaveReport sets SaveReport field to given value.

### HasSaveReport

`func (o *CSSCMSDataModelModelsReportSchedule) HasSaveReport() bool`

HasSaveReport returns a boolean if a field has been set.

### GetSaveReportPath

`func (o *CSSCMSDataModelModelsReportSchedule) GetSaveReportPath() string`

GetSaveReportPath returns the SaveReportPath field if non-nil, zero value otherwise.

### GetSaveReportPathOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetSaveReportPathOk() (*string, bool)`

GetSaveReportPathOk returns a tuple with the SaveReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveReportPath

`func (o *CSSCMSDataModelModelsReportSchedule) SetSaveReportPath(v string)`

SetSaveReportPath sets SaveReportPath field to given value.

### HasSaveReportPath

`func (o *CSSCMSDataModelModelsReportSchedule) HasSaveReportPath() bool`

HasSaveReportPath returns a boolean if a field has been set.

### SetSaveReportPathNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetSaveReportPathNil(b bool)`

 SetSaveReportPathNil sets the value for SaveReportPath to be an explicit nil

### UnsetSaveReportPath
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetSaveReportPath()`

UnsetSaveReportPath ensures that no value is present for SaveReportPath, not even an explicit nil
### GetReportFormat

`func (o *CSSCMSDataModelModelsReportSchedule) GetReportFormat() string`

GetReportFormat returns the ReportFormat field if non-nil, zero value otherwise.

### GetReportFormatOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetReportFormatOk() (*string, bool)`

GetReportFormatOk returns a tuple with the ReportFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportFormat

`func (o *CSSCMSDataModelModelsReportSchedule) SetReportFormat(v string)`

SetReportFormat sets ReportFormat field to given value.

### HasReportFormat

`func (o *CSSCMSDataModelModelsReportSchedule) HasReportFormat() bool`

HasReportFormat returns a boolean if a field has been set.

### SetReportFormatNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetReportFormatNil(b bool)`

 SetReportFormatNil sets the value for ReportFormat to be an explicit nil

### UnsetReportFormat
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetReportFormat()`

UnsetReportFormat ensures that no value is present for ReportFormat, not even an explicit nil
### GetKeyfactorSchedule

`func (o *CSSCMSDataModelModelsReportSchedule) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *CSSCMSDataModelModelsReportSchedule) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *CSSCMSDataModelModelsReportSchedule) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetCertificateCollectionId

`func (o *CSSCMSDataModelModelsReportSchedule) GetCertificateCollectionId() int32`

GetCertificateCollectionId returns the CertificateCollectionId field if non-nil, zero value otherwise.

### GetCertificateCollectionIdOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetCertificateCollectionIdOk() (*int32, bool)`

GetCertificateCollectionIdOk returns a tuple with the CertificateCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionId

`func (o *CSSCMSDataModelModelsReportSchedule) SetCertificateCollectionId(v int32)`

SetCertificateCollectionId sets CertificateCollectionId field to given value.

### HasCertificateCollectionId

`func (o *CSSCMSDataModelModelsReportSchedule) HasCertificateCollectionId() bool`

HasCertificateCollectionId returns a boolean if a field has been set.

### SetCertificateCollectionIdNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetCertificateCollectionIdNil(b bool)`

 SetCertificateCollectionIdNil sets the value for CertificateCollectionId to be an explicit nil

### UnsetCertificateCollectionId
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetCertificateCollectionId()`

UnsetCertificateCollectionId ensures that no value is present for CertificateCollectionId, not even an explicit nil
### GetCertificateCollectionName

`func (o *CSSCMSDataModelModelsReportSchedule) GetCertificateCollectionName() string`

GetCertificateCollectionName returns the CertificateCollectionName field if non-nil, zero value otherwise.

### GetCertificateCollectionNameOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetCertificateCollectionNameOk() (*string, bool)`

GetCertificateCollectionNameOk returns a tuple with the CertificateCollectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateCollectionName

`func (o *CSSCMSDataModelModelsReportSchedule) SetCertificateCollectionName(v string)`

SetCertificateCollectionName sets CertificateCollectionName field to given value.

### HasCertificateCollectionName

`func (o *CSSCMSDataModelModelsReportSchedule) HasCertificateCollectionName() bool`

HasCertificateCollectionName returns a boolean if a field has been set.

### SetCertificateCollectionNameNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetCertificateCollectionNameNil(b bool)`

 SetCertificateCollectionNameNil sets the value for CertificateCollectionName to be an explicit nil

### UnsetCertificateCollectionName
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetCertificateCollectionName()`

UnsetCertificateCollectionName ensures that no value is present for CertificateCollectionName, not even an explicit nil
### GetEmailRecipients

`func (o *CSSCMSDataModelModelsReportSchedule) GetEmailRecipients() []string`

GetEmailRecipients returns the EmailRecipients field if non-nil, zero value otherwise.

### GetEmailRecipientsOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetEmailRecipientsOk() (*[]string, bool)`

GetEmailRecipientsOk returns a tuple with the EmailRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailRecipients

`func (o *CSSCMSDataModelModelsReportSchedule) SetEmailRecipients(v []string)`

SetEmailRecipients sets EmailRecipients field to given value.

### HasEmailRecipients

`func (o *CSSCMSDataModelModelsReportSchedule) HasEmailRecipients() bool`

HasEmailRecipients returns a boolean if a field has been set.

### SetEmailRecipientsNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetEmailRecipientsNil(b bool)`

 SetEmailRecipientsNil sets the value for EmailRecipients to be an explicit nil

### UnsetEmailRecipients
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetEmailRecipients()`

UnsetEmailRecipients ensures that no value is present for EmailRecipients, not even an explicit nil
### GetRuntimeParameters

`func (o *CSSCMSDataModelModelsReportSchedule) GetRuntimeParameters() map[string]string`

GetRuntimeParameters returns the RuntimeParameters field if non-nil, zero value otherwise.

### GetRuntimeParametersOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetRuntimeParametersOk() (*map[string]string, bool)`

GetRuntimeParametersOk returns a tuple with the RuntimeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuntimeParameters

`func (o *CSSCMSDataModelModelsReportSchedule) SetRuntimeParameters(v map[string]string)`

SetRuntimeParameters sets RuntimeParameters field to given value.

### HasRuntimeParameters

`func (o *CSSCMSDataModelModelsReportSchedule) HasRuntimeParameters() bool`

HasRuntimeParameters returns a boolean if a field has been set.

### SetRuntimeParametersNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetRuntimeParametersNil(b bool)`

 SetRuntimeParametersNil sets the value for RuntimeParameters to be an explicit nil

### UnsetRuntimeParameters
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetRuntimeParameters()`

UnsetRuntimeParameters ensures that no value is present for RuntimeParameters, not even an explicit nil
### GetReportDefinitionId

`func (o *CSSCMSDataModelModelsReportSchedule) GetReportDefinitionId() int32`

GetReportDefinitionId returns the ReportDefinitionId field if non-nil, zero value otherwise.

### GetReportDefinitionIdOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetReportDefinitionIdOk() (*int32, bool)`

GetReportDefinitionIdOk returns a tuple with the ReportDefinitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportDefinitionId

`func (o *CSSCMSDataModelModelsReportSchedule) SetReportDefinitionId(v int32)`

SetReportDefinitionId sets ReportDefinitionId field to given value.

### HasReportDefinitionId

`func (o *CSSCMSDataModelModelsReportSchedule) HasReportDefinitionId() bool`

HasReportDefinitionId returns a boolean if a field has been set.

### GetName

`func (o *CSSCMSDataModelModelsReportSchedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CSSCMSDataModelModelsReportSchedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CSSCMSDataModelModelsReportSchedule) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CSSCMSDataModelModelsReportSchedule) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CSSCMSDataModelModelsReportSchedule) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetEnabled

`func (o *CSSCMSDataModelModelsReportSchedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CSSCMSDataModelModelsReportSchedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CSSCMSDataModelModelsReportSchedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CSSCMSDataModelModelsReportSchedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


