# ModelsReport

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
**ReportParameter** | Pointer to [**[]ModelsReportParameters**](ModelsReportParameters.md) |  | [optional] 
**Schedules** | Pointer to [**[]ModelsReportSchedule**](ModelsReportSchedule.md) |  | [optional] 
**AcceptedScheduleFormats** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelsReport

`func NewModelsReport() *ModelsReport`

NewModelsReport instantiates a new ModelsReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsReportWithDefaults

`func NewModelsReportWithDefaults() *ModelsReport`

NewModelsReportWithDefaults instantiates a new ModelsReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsReport) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsReport) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsReport) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsReport) HasId() bool`

HasId returns a boolean if a field has been set.

### GetScheduled

`func (o *ModelsReport) GetScheduled() int32`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *ModelsReport) GetScheduledOk() (*int32, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *ModelsReport) SetScheduled(v int32)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *ModelsReport) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### SetScheduledNil

`func (o *ModelsReport) SetScheduledNil(b bool)`

 SetScheduledNil sets the value for Scheduled to be an explicit nil

### UnsetScheduled
`func (o *ModelsReport) UnsetScheduled()`

UnsetScheduled ensures that no value is present for Scheduled, not even an explicit nil
### GetDisplayName

`func (o *ModelsReport) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ModelsReport) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ModelsReport) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ModelsReport) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *ModelsReport) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *ModelsReport) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *ModelsReport) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsReport) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsReport) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsReport) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelsReport) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelsReport) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetReportPath

`func (o *ModelsReport) GetReportPath() string`

GetReportPath returns the ReportPath field if non-nil, zero value otherwise.

### GetReportPathOk

`func (o *ModelsReport) GetReportPathOk() (*string, bool)`

GetReportPathOk returns a tuple with the ReportPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportPath

`func (o *ModelsReport) SetReportPath(v string)`

SetReportPath sets ReportPath field to given value.

### HasReportPath

`func (o *ModelsReport) HasReportPath() bool`

HasReportPath returns a boolean if a field has been set.

### SetReportPathNil

`func (o *ModelsReport) SetReportPathNil(b bool)`

 SetReportPathNil sets the value for ReportPath to be an explicit nil

### UnsetReportPath
`func (o *ModelsReport) UnsetReportPath()`

UnsetReportPath ensures that no value is present for ReportPath, not even an explicit nil
### GetVersionNumber

`func (o *ModelsReport) GetVersionNumber() string`

GetVersionNumber returns the VersionNumber field if non-nil, zero value otherwise.

### GetVersionNumberOk

`func (o *ModelsReport) GetVersionNumberOk() (*string, bool)`

GetVersionNumberOk returns a tuple with the VersionNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionNumber

`func (o *ModelsReport) SetVersionNumber(v string)`

SetVersionNumber sets VersionNumber field to given value.

### HasVersionNumber

`func (o *ModelsReport) HasVersionNumber() bool`

HasVersionNumber returns a boolean if a field has been set.

### SetVersionNumberNil

`func (o *ModelsReport) SetVersionNumberNil(b bool)`

 SetVersionNumberNil sets the value for VersionNumber to be an explicit nil

### UnsetVersionNumber
`func (o *ModelsReport) UnsetVersionNumber()`

UnsetVersionNumber ensures that no value is present for VersionNumber, not even an explicit nil
### GetCategories

`func (o *ModelsReport) GetCategories() string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ModelsReport) GetCategoriesOk() (*string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ModelsReport) SetCategories(v string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ModelsReport) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ModelsReport) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ModelsReport) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil
### GetShortName

`func (o *ModelsReport) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ModelsReport) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ModelsReport) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ModelsReport) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### SetShortNameNil

`func (o *ModelsReport) SetShortNameNil(b bool)`

 SetShortNameNil sets the value for ShortName to be an explicit nil

### UnsetShortName
`func (o *ModelsReport) UnsetShortName()`

UnsetShortName ensures that no value is present for ShortName, not even an explicit nil
### GetInNavigator

`func (o *ModelsReport) GetInNavigator() bool`

GetInNavigator returns the InNavigator field if non-nil, zero value otherwise.

### GetInNavigatorOk

`func (o *ModelsReport) GetInNavigatorOk() (*bool, bool)`

GetInNavigatorOk returns a tuple with the InNavigator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInNavigator

`func (o *ModelsReport) SetInNavigator(v bool)`

SetInNavigator sets InNavigator field to given value.

### HasInNavigator

`func (o *ModelsReport) HasInNavigator() bool`

HasInNavigator returns a boolean if a field has been set.

### GetFavorite

`func (o *ModelsReport) GetFavorite() bool`

GetFavorite returns the Favorite field if non-nil, zero value otherwise.

### GetFavoriteOk

`func (o *ModelsReport) GetFavoriteOk() (*bool, bool)`

GetFavoriteOk returns a tuple with the Favorite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorite

`func (o *ModelsReport) SetFavorite(v bool)`

SetFavorite sets Favorite field to given value.

### HasFavorite

`func (o *ModelsReport) HasFavorite() bool`

HasFavorite returns a boolean if a field has been set.

### GetRemoveDuplicates

`func (o *ModelsReport) GetRemoveDuplicates() bool`

GetRemoveDuplicates returns the RemoveDuplicates field if non-nil, zero value otherwise.

### GetRemoveDuplicatesOk

`func (o *ModelsReport) GetRemoveDuplicatesOk() (*bool, bool)`

GetRemoveDuplicatesOk returns a tuple with the RemoveDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoveDuplicates

`func (o *ModelsReport) SetRemoveDuplicates(v bool)`

SetRemoveDuplicates sets RemoveDuplicates field to given value.

### HasRemoveDuplicates

`func (o *ModelsReport) HasRemoveDuplicates() bool`

HasRemoveDuplicates returns a boolean if a field has been set.

### GetUsesCollection

`func (o *ModelsReport) GetUsesCollection() bool`

GetUsesCollection returns the UsesCollection field if non-nil, zero value otherwise.

### GetUsesCollectionOk

`func (o *ModelsReport) GetUsesCollectionOk() (*bool, bool)`

GetUsesCollectionOk returns a tuple with the UsesCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsesCollection

`func (o *ModelsReport) SetUsesCollection(v bool)`

SetUsesCollection sets UsesCollection field to given value.

### HasUsesCollection

`func (o *ModelsReport) HasUsesCollection() bool`

HasUsesCollection returns a boolean if a field has been set.

### GetReportParameter

`func (o *ModelsReport) GetReportParameter() []ModelsReportParameters`

GetReportParameter returns the ReportParameter field if non-nil, zero value otherwise.

### GetReportParameterOk

`func (o *ModelsReport) GetReportParameterOk() (*[]ModelsReportParameters, bool)`

GetReportParameterOk returns a tuple with the ReportParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportParameter

`func (o *ModelsReport) SetReportParameter(v []ModelsReportParameters)`

SetReportParameter sets ReportParameter field to given value.

### HasReportParameter

`func (o *ModelsReport) HasReportParameter() bool`

HasReportParameter returns a boolean if a field has been set.

### SetReportParameterNil

`func (o *ModelsReport) SetReportParameterNil(b bool)`

 SetReportParameterNil sets the value for ReportParameter to be an explicit nil

### UnsetReportParameter
`func (o *ModelsReport) UnsetReportParameter()`

UnsetReportParameter ensures that no value is present for ReportParameter, not even an explicit nil
### GetSchedules

`func (o *ModelsReport) GetSchedules() []ModelsReportSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *ModelsReport) GetSchedulesOk() (*[]ModelsReportSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *ModelsReport) SetSchedules(v []ModelsReportSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *ModelsReport) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### SetSchedulesNil

`func (o *ModelsReport) SetSchedulesNil(b bool)`

 SetSchedulesNil sets the value for Schedules to be an explicit nil

### UnsetSchedules
`func (o *ModelsReport) UnsetSchedules()`

UnsetSchedules ensures that no value is present for Schedules, not even an explicit nil
### GetAcceptedScheduleFormats

`func (o *ModelsReport) GetAcceptedScheduleFormats() []string`

GetAcceptedScheduleFormats returns the AcceptedScheduleFormats field if non-nil, zero value otherwise.

### GetAcceptedScheduleFormatsOk

`func (o *ModelsReport) GetAcceptedScheduleFormatsOk() (*[]string, bool)`

GetAcceptedScheduleFormatsOk returns a tuple with the AcceptedScheduleFormats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedScheduleFormats

`func (o *ModelsReport) SetAcceptedScheduleFormats(v []string)`

SetAcceptedScheduleFormats sets AcceptedScheduleFormats field to given value.

### HasAcceptedScheduleFormats

`func (o *ModelsReport) HasAcceptedScheduleFormats() bool`

HasAcceptedScheduleFormats returns a boolean if a field has been set.

### SetAcceptedScheduleFormatsNil

`func (o *ModelsReport) SetAcceptedScheduleFormatsNil(b bool)`

 SetAcceptedScheduleFormatsNil sets the value for AcceptedScheduleFormats to be an explicit nil

### UnsetAcceptedScheduleFormats
`func (o *ModelsReport) UnsetAcceptedScheduleFormats()`

UnsetAcceptedScheduleFormats ensures that no value is present for AcceptedScheduleFormats, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


