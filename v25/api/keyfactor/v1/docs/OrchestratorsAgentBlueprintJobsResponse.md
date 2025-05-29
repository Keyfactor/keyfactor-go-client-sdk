# OrchestratorsAgentBlueprintJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBlueprintJobId** | Pointer to **string** |  | [optional] 
**AgentBlueprintStoreId** | Pointer to **string** |  | [optional] 
**AgentBlueprintId** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**JobTypeName** | Pointer to **NullableString** |  | [optional] 
**OperationType** | Pointer to **NullableInt32** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**Contents** | Pointer to **NullableString** |  | [optional] 
**Alias** | Pointer to **NullableString** |  | [optional] 
**PrivateKeyEntry** | Pointer to **NullableBool** |  | [optional] 
**Overwrite** | Pointer to **NullableBool** |  | [optional] 
**HasEntryPassword** | Pointer to **NullableBool** |  | [optional] 
**HasPfxPassword** | Pointer to **NullableBool** |  | [optional] 
**RequestTimestamp** | Pointer to **NullableTime** |  | [optional] 
**KeyfactorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**Directories** | Pointer to **NullableString** |  | [optional] 
**IgnoredDirectories** | Pointer to **NullableString** |  | [optional] 
**SymLinks** | Pointer to **NullableBool** |  | [optional] 
**Compatibility** | Pointer to **NullableBool** |  | [optional] 
**FileExtensions** | Pointer to **NullableString** |  | [optional] 
**FileNamePatterns** | Pointer to **NullableString** |  | [optional] 
**AgentBlueprintStores** | Pointer to [**OrchestratorsAgentBlueprintStoresResponse**](OrchestratorsAgentBlueprintStoresResponse.md) |  | [optional] 

## Methods

### NewOrchestratorsAgentBlueprintJobsResponse

`func NewOrchestratorsAgentBlueprintJobsResponse() *OrchestratorsAgentBlueprintJobsResponse`

NewOrchestratorsAgentBlueprintJobsResponse instantiates a new OrchestratorsAgentBlueprintJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorsAgentBlueprintJobsResponseWithDefaults

`func NewOrchestratorsAgentBlueprintJobsResponseWithDefaults() *OrchestratorsAgentBlueprintJobsResponse`

NewOrchestratorsAgentBlueprintJobsResponseWithDefaults instantiates a new OrchestratorsAgentBlueprintJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBlueprintJobId

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintJobId() string`

GetAgentBlueprintJobId returns the AgentBlueprintJobId field if non-nil, zero value otherwise.

### GetAgentBlueprintJobIdOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintJobIdOk() (*string, bool)`

GetAgentBlueprintJobIdOk returns a tuple with the AgentBlueprintJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintJobId

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintJobId(v string)`

SetAgentBlueprintJobId sets AgentBlueprintJobId field to given value.

### HasAgentBlueprintJobId

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintJobId() bool`

HasAgentBlueprintJobId returns a boolean if a field has been set.

### GetAgentBlueprintStoreId

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoreId() string`

GetAgentBlueprintStoreId returns the AgentBlueprintStoreId field if non-nil, zero value otherwise.

### GetAgentBlueprintStoreIdOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoreIdOk() (*string, bool)`

GetAgentBlueprintStoreIdOk returns a tuple with the AgentBlueprintStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStoreId

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintStoreId(v string)`

SetAgentBlueprintStoreId sets AgentBlueprintStoreId field to given value.

### HasAgentBlueprintStoreId

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintStoreId() bool`

HasAgentBlueprintStoreId returns a boolean if a field has been set.

### GetAgentBlueprintId

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintId() string`

GetAgentBlueprintId returns the AgentBlueprintId field if non-nil, zero value otherwise.

### GetAgentBlueprintIdOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintIdOk() (*string, bool)`

GetAgentBlueprintIdOk returns a tuple with the AgentBlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintId

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintId(v string)`

SetAgentBlueprintId sets AgentBlueprintId field to given value.

### HasAgentBlueprintId

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintId() bool`

HasAgentBlueprintId returns a boolean if a field has been set.

### GetJobType

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetJobTypeName

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetOperationType

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetOperationType() int32`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetOperationTypeOk() (*int32, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetOperationType(v int32)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetThumbprint

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetContents

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetAlias

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetPrivateKeyEntry

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetPrivateKeyEntry() bool`

GetPrivateKeyEntry returns the PrivateKeyEntry field if non-nil, zero value otherwise.

### GetPrivateKeyEntryOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetPrivateKeyEntryOk() (*bool, bool)`

GetPrivateKeyEntryOk returns a tuple with the PrivateKeyEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEntry

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetPrivateKeyEntry(v bool)`

SetPrivateKeyEntry sets PrivateKeyEntry field to given value.

### HasPrivateKeyEntry

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasPrivateKeyEntry() bool`

HasPrivateKeyEntry returns a boolean if a field has been set.

### SetPrivateKeyEntryNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetPrivateKeyEntryNil(b bool)`

 SetPrivateKeyEntryNil sets the value for PrivateKeyEntry to be an explicit nil

### UnsetPrivateKeyEntry
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetPrivateKeyEntry()`

UnsetPrivateKeyEntry ensures that no value is present for PrivateKeyEntry, not even an explicit nil
### GetOverwrite

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetHasEntryPassword

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetHasEntryPassword() bool`

GetHasEntryPassword returns the HasEntryPassword field if non-nil, zero value otherwise.

### GetHasEntryPasswordOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetHasEntryPasswordOk() (*bool, bool)`

GetHasEntryPasswordOk returns a tuple with the HasEntryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntryPassword

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetHasEntryPassword(v bool)`

SetHasEntryPassword sets HasEntryPassword field to given value.

### HasHasEntryPassword

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasHasEntryPassword() bool`

HasHasEntryPassword returns a boolean if a field has been set.

### SetHasEntryPasswordNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetHasEntryPasswordNil(b bool)`

 SetHasEntryPasswordNil sets the value for HasEntryPassword to be an explicit nil

### UnsetHasEntryPassword
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetHasEntryPassword()`

UnsetHasEntryPassword ensures that no value is present for HasEntryPassword, not even an explicit nil
### GetHasPfxPassword

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetHasPfxPassword() bool`

GetHasPfxPassword returns the HasPfxPassword field if non-nil, zero value otherwise.

### GetHasPfxPasswordOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetHasPfxPasswordOk() (*bool, bool)`

GetHasPfxPasswordOk returns a tuple with the HasPfxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPfxPassword

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetHasPfxPassword(v bool)`

SetHasPfxPassword sets HasPfxPassword field to given value.

### HasHasPfxPassword

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasHasPfxPassword() bool`

HasHasPfxPassword returns a boolean if a field has been set.

### SetHasPfxPasswordNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetHasPfxPasswordNil(b bool)`

 SetHasPfxPasswordNil sets the value for HasPfxPassword to be an explicit nil

### UnsetHasPfxPassword
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetHasPfxPassword()`

UnsetHasPfxPassword ensures that no value is present for HasPfxPassword, not even an explicit nil
### GetRequestTimestamp

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### SetRequestTimestampNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetRequestTimestampNil(b bool)`

 SetRequestTimestampNil sets the value for RequestTimestamp to be an explicit nil

### UnsetRequestTimestamp
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetRequestTimestamp()`

UnsetRequestTimestamp ensures that no value is present for RequestTimestamp, not even an explicit nil
### GetKeyfactorSchedule

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetSubject

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetDirectories

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetDirectories() string`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetDirectoriesOk() (*string, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetDirectories(v string)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### SetDirectoriesNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetDirectoriesNil(b bool)`

 SetDirectoriesNil sets the value for Directories to be an explicit nil

### UnsetDirectories
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetDirectories()`

UnsetDirectories ensures that no value is present for Directories, not even an explicit nil
### GetIgnoredDirectories

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetIgnoredDirectories() string`

GetIgnoredDirectories returns the IgnoredDirectories field if non-nil, zero value otherwise.

### GetIgnoredDirectoriesOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetIgnoredDirectoriesOk() (*string, bool)`

GetIgnoredDirectoriesOk returns a tuple with the IgnoredDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredDirectories

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetIgnoredDirectories(v string)`

SetIgnoredDirectories sets IgnoredDirectories field to given value.

### HasIgnoredDirectories

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasIgnoredDirectories() bool`

HasIgnoredDirectories returns a boolean if a field has been set.

### SetIgnoredDirectoriesNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetIgnoredDirectoriesNil(b bool)`

 SetIgnoredDirectoriesNil sets the value for IgnoredDirectories to be an explicit nil

### UnsetIgnoredDirectories
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetIgnoredDirectories()`

UnsetIgnoredDirectories ensures that no value is present for IgnoredDirectories, not even an explicit nil
### GetSymLinks

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetSymLinks() bool`

GetSymLinks returns the SymLinks field if non-nil, zero value otherwise.

### GetSymLinksOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetSymLinksOk() (*bool, bool)`

GetSymLinksOk returns a tuple with the SymLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymLinks

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetSymLinks(v bool)`

SetSymLinks sets SymLinks field to given value.

### HasSymLinks

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasSymLinks() bool`

HasSymLinks returns a boolean if a field has been set.

### SetSymLinksNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetSymLinksNil(b bool)`

 SetSymLinksNil sets the value for SymLinks to be an explicit nil

### UnsetSymLinks
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetSymLinks()`

UnsetSymLinks ensures that no value is present for SymLinks, not even an explicit nil
### GetCompatibility

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### SetCompatibilityNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetCompatibilityNil(b bool)`

 SetCompatibilityNil sets the value for Compatibility to be an explicit nil

### UnsetCompatibility
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetCompatibility()`

UnsetCompatibility ensures that no value is present for Compatibility, not even an explicit nil
### GetFileExtensions

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetFileExtensions() string`

GetFileExtensions returns the FileExtensions field if non-nil, zero value otherwise.

### GetFileExtensionsOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetFileExtensionsOk() (*string, bool)`

GetFileExtensionsOk returns a tuple with the FileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensions

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetFileExtensions(v string)`

SetFileExtensions sets FileExtensions field to given value.

### HasFileExtensions

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasFileExtensions() bool`

HasFileExtensions returns a boolean if a field has been set.

### SetFileExtensionsNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetFileExtensionsNil(b bool)`

 SetFileExtensionsNil sets the value for FileExtensions to be an explicit nil

### UnsetFileExtensions
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetFileExtensions()`

UnsetFileExtensions ensures that no value is present for FileExtensions, not even an explicit nil
### GetFileNamePatterns

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetFileNamePatterns() string`

GetFileNamePatterns returns the FileNamePatterns field if non-nil, zero value otherwise.

### GetFileNamePatternsOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetFileNamePatternsOk() (*string, bool)`

GetFileNamePatternsOk returns a tuple with the FileNamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePatterns

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetFileNamePatterns(v string)`

SetFileNamePatterns sets FileNamePatterns field to given value.

### HasFileNamePatterns

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasFileNamePatterns() bool`

HasFileNamePatterns returns a boolean if a field has been set.

### SetFileNamePatternsNil

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetFileNamePatternsNil(b bool)`

 SetFileNamePatternsNil sets the value for FileNamePatterns to be an explicit nil

### UnsetFileNamePatterns
`func (o *OrchestratorsAgentBlueprintJobsResponse) UnsetFileNamePatterns()`

UnsetFileNamePatterns ensures that no value is present for FileNamePatterns, not even an explicit nil
### GetAgentBlueprintStores

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStores() OrchestratorsAgentBlueprintStoresResponse`

GetAgentBlueprintStores returns the AgentBlueprintStores field if non-nil, zero value otherwise.

### GetAgentBlueprintStoresOk

`func (o *OrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoresOk() (*OrchestratorsAgentBlueprintStoresResponse, bool)`

GetAgentBlueprintStoresOk returns a tuple with the AgentBlueprintStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStores

`func (o *OrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintStores(v OrchestratorsAgentBlueprintStoresResponse)`

SetAgentBlueprintStores sets AgentBlueprintStores field to given value.

### HasAgentBlueprintStores

`func (o *OrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintStores() bool`

HasAgentBlueprintStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


