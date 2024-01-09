# CSSCMSDataModelModelsReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Scheduled** | Pointer to **NullableInt32** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**ReportPath** | Pointer to **NullableString** |  | [optional] 
**VersionNumber** | Pointer to **NullableString** |  | [optional] 
**Categories** | Pointer to **NullableString** |  | [optional] 
**ShortName** | Pointer to **NullableString** |  | [optional] 
**InNavigator** | Pointer to **bool** |  | [optional] 
**Favorite** | Pointer to **bool** |  | [optional] 
**RemoveDuplicates** | Pointer to **bool** |  | [optional] 
**UsesCollection** | Pointer to **bool** |  | [optional] 
**ReportParameter** | Pointer to [**[]CSSCMSDataModelModelsReportParameters**](CSSCMSDataModelModelsReportParameters.md) |  | [optional] 
**Schedules** | Pointer to [**[]CSSCMSDataModelModelsReportSchedule**](CSSCMSDataModelModelsReportSchedule.md) |  | [optional] 
**AcceptedScheduleFormats** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsReport

`func NewCSSCMSDataModelModelsReport() *CSSCMSDataModelModelsReport`

NewCSSCMSDataModelModelsReport instantiates a new CSSCMSDataModelModelsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsReportWithDefaults

`func NewCSSCMSDataModelModelsReportWithDefaults() *CSSCMSDataModelModelsReport`

NewCSSCMSDataModelModelsReportWithDefaults instantiates a new CSSCMSDataModelModelsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsReport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduled

`func (o *CSSCMSDataModelModelsReport) GetScheduled() int32`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *CSSCMSDataModelModelsReport) GetScheduledOk() (*int32, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *CSSCMSDataModelModelsReport) SetScheduled(v int32)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *CSSCMSDataModelModelsReport) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### SetScheduledNil

`func (o *CSSCMSDataModelModelsReport) SetScheduledNil(b bool)`

 SetScheduledNil sets the value for Scheduled to be an explicit nil

### UnsetScheduled
`func (o *CSSCMSDataModelModelsReport) UnsetScheduled()`

UnsetScheduled ensures that no value is present for Scheduled, not even an explicit nil
### GetDisplayName

`func (o *CSSCMSDataModelModelsReport) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CSSCMSDataModelModelsReport) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CSSCMSDataModelModelsReport) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CSSCMSDataModelModelsReport) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *CSSCMSDataModelModelsReport) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *CSSCMSDataModelModelsReport) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *CSSCMSDataModelModelsReport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CSSCMSDataModelModelsReport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CSSCMSDataModelModelsReport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CSSCMSDataModelModelsReport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CSSCMSDataModelModelsReport) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CSSCMSDataModelModelsReport) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReportPath

`func (o *CSSCMSDataModelModelsReport) GetReportPath() string`

GetReportPath returns the ReportPath field if non-nil, zero value otherwise.

### GetReportPathOk

`func (o *CSSCMSDataModelModelsReport) GetReportPathOk() (*string, bool)`

GetReportPathOk returns a tuple with the ReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPath

`func (o *CSSCMSDataModelModelsReport) SetReportPath(v string)`

SetReportPath sets ReportPath field to given value.

### HasReportPath

`func (o *CSSCMSDataModelModelsReport) HasReportPath() bool`

HasReportPath returns a boolean if a field has been set.

### SetReportPathNil

`func (o *CSSCMSDataModelModelsReport) SetReportPathNil(b bool)`

 SetReportPathNil sets the value for ReportPath to be an explicit nil

### UnsetReportPath
`func (o *CSSCMSDataModelModelsReport) UnsetReportPath()`

UnsetReportPath ensures that no value is present for ReportPath, not even an explicit nil
### GetVersionNumber

`func (o *CSSCMSDataModelModelsReport) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *CSSCMSDataModelModelsReport) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *CSSCMSDataModelModelsReport) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *CSSCMSDataModelModelsReport) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### SetVersionNumberNil

`func (o *CSSCMSDataModelModelsReport) SetVersionNumberNil(b bool)`

 SetVersionNumberNil sets the value for VersionNumber to be an explicit nil

### UnsetVersionNumber
`func (o *CSSCMSDataModelModelsReport) UnsetVersionNumber()`

UnsetVersionNumber ensures that no value is present for VersionNumber, not even an explicit nil
### GetCategories

`func (o *CSSCMSDataModelModelsReport) GetCategories() string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CSSCMSDataModelModelsReport) GetCategoriesOk() (*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CSSCMSDataModelModelsReport) SetCategories(v string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *CSSCMSDataModelModelsReport) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *CSSCMSDataModelModelsReport) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *CSSCMSDataModelModelsReport) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetShortName

`func (o *CSSCMSDataModelModelsReport) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *CSSCMSDataModelModelsReport) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *CSSCMSDataModelModelsReport) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *CSSCMSDataModelModelsReport) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *CSSCMSDataModelModelsReport) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *CSSCMSDataModelModelsReport) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetInNavigator

`func (o *CSSCMSDataModelModelsReport) GetInNavigator() bool`

GetInNavigator returns the InNavigator field if non-nil, zero value otherwise.

### GetInNavigatorOk

`func (o *CSSCMSDataModelModelsReport) GetInNavigatorOk() (*bool, bool)`

GetInNavigatorOk returns a tuple with the InNavigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInNavigator

`func (o *CSSCMSDataModelModelsReport) SetInNavigator(v bool)`

SetInNavigator sets InNavigator field to given value.

### HasInNavigator

`func (o *CSSCMSDataModelModelsReport) HasInNavigator() bool`

HasInNavigator returns a boolean if a field has been set.

### GetFavorite

`func (o *CSSCMSDataModelModelsReport) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *CSSCMSDataModelModelsReport) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *CSSCMSDataModelModelsReport) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *CSSCMSDataModelModelsReport) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetRemoveDuplicates

`func (o *CSSCMSDataModelModelsReport) GetRemoveDuplicates() bool`

GetRemoveDuplicates returns the RemoveDuplicates field if non-nil, zero value otherwise.

### GetRemoveDuplicatesOk

`func (o *CSSCMSDataModelModelsReport) GetRemoveDuplicatesOk() (*bool, bool)`

GetRemoveDuplicatesOk returns a tuple with the RemoveDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDuplicates

`func (o *CSSCMSDataModelModelsReport) SetRemoveDuplicates(v bool)`

SetRemoveDuplicates sets RemoveDuplicates field to given value.

### HasRemoveDuplicates

`func (o *CSSCMSDataModelModelsReport) HasRemoveDuplicates() bool`

HasRemoveDuplicates returns a boolean if a field has been set.

### GetUsesCollection

`func (o *CSSCMSDataModelModelsReport) GetUsesCollection() bool`

GetUsesCollection returns the UsesCollection field if non-nil, zero value otherwise.

### GetUsesCollectionOk

`func (o *CSSCMSDataModelModelsReport) GetUsesCollectionOk() (*bool, bool)`

GetUsesCollectionOk returns a tuple with the UsesCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCollection

`func (o *CSSCMSDataModelModelsReport) SetUsesCollection(v bool)`

SetUsesCollection sets UsesCollection field to given value.

### HasUsesCollection

`func (o *CSSCMSDataModelModelsReport) HasUsesCollection() bool`

HasUsesCollection returns a boolean if a field has been set.

### GetReportParameter

`func (o *CSSCMSDataModelModelsReport) GetReportParameter() []CSSCMSDataModelModelsReportParameters`

GetReportParameter returns the ReportParameter field if non-nil, zero value otherwise.

### GetReportParameterOk

`func (o *CSSCMSDataModelModelsReport) GetReportParameterOk() (*[]CSSCMSDataModelModelsReportParameters, bool)`

GetReportParameterOk returns a tuple with the ReportParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportParameter

`func (o *CSSCMSDataModelModelsReport) SetReportParameter(v []CSSCMSDataModelModelsReportParameters)`

SetReportParameter sets ReportParameter field to given value.

### HasReportParameter

`func (o *CSSCMSDataModelModelsReport) HasReportParameter() bool`

HasReportParameter returns a boolean if a field has been set.

### SetReportParameterNil

`func (o *CSSCMSDataModelModelsReport) SetReportParameterNil(b bool)`

 SetReportParameterNil sets the value for ReportParameter to be an explicit nil

### UnsetReportParameter
`func (o *CSSCMSDataModelModelsReport) UnsetReportParameter()`

UnsetReportParameter ensures that no value is present for ReportParameter, not even an explicit nil
### GetSchedules

`func (o *CSSCMSDataModelModelsReport) GetSchedules() []CSSCMSDataModelModelsReportSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *CSSCMSDataModelModelsReport) GetSchedulesOk() (*[]CSSCMSDataModelModelsReportSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *CSSCMSDataModelModelsReport) SetSchedules(v []CSSCMSDataModelModelsReportSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *CSSCMSDataModelModelsReport) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedulesNil

`func (o *CSSCMSDataModelModelsReport) SetSchedulesNil(b bool)`

 SetSchedulesNil sets the value for Schedules to be an explicit nil

### UnsetSchedules
`func (o *CSSCMSDataModelModelsReport) UnsetSchedules()`

UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
### GetAcceptedScheduleFormats

`func (o *CSSCMSDataModelModelsReport) GetAcceptedScheduleFormats() []string`

GetAcceptedScheduleFormats returns the AcceptedScheduleFormats field if non-nil, zero value otherwise.

### GetAcceptedScheduleFormatsOk

`func (o *CSSCMSDataModelModelsReport) GetAcceptedScheduleFormatsOk() (*[]string, bool)`

GetAcceptedScheduleFormatsOk returns a tuple with the AcceptedScheduleFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedScheduleFormats

`func (o *CSSCMSDataModelModelsReport) SetAcceptedScheduleFormats(v []string)`

SetAcceptedScheduleFormats sets AcceptedScheduleFormats field to given value.

### HasAcceptedScheduleFormats

`func (o *CSSCMSDataModelModelsReport) HasAcceptedScheduleFormats() bool`

HasAcceptedScheduleFormats returns a boolean if a field has been set.

### SetAcceptedScheduleFormatsNil

`func (o *CSSCMSDataModelModelsReport) SetAcceptedScheduleFormatsNil(b bool)`

 SetAcceptedScheduleFormatsNil sets the value for AcceptedScheduleFormats to be an explicit nil

### UnsetAcceptedScheduleFormats
`func (o *CSSCMSDataModelModelsReport) UnsetAcceptedScheduleFormats()`

UnsetAcceptedScheduleFormats ensures that no value is present for AcceptedScheduleFormats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


