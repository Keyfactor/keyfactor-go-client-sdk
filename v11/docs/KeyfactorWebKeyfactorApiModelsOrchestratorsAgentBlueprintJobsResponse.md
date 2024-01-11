# KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse

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
**AgentBlueprintStores** | Pointer to [**KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse**](KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse() *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponseWithDefaults() *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse`

NewKeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBlueprintJobId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintJobId() string`

GetAgentBlueprintJobId returns the AgentBlueprintJobId field if non-nil, zero value otherwise.

### GetAgentBlueprintJobIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintJobIdOk() (*string, bool)`

GetAgentBlueprintJobIdOk returns a tuple with the AgentBlueprintJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintJobId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintJobId(v string)`

SetAgentBlueprintJobId sets AgentBlueprintJobId field to given value.

### HasAgentBlueprintJobId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintJobId() bool`

HasAgentBlueprintJobId returns a boolean if a field has been set.

### GetAgentBlueprintStoreId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoreId() string`

GetAgentBlueprintStoreId returns the AgentBlueprintStoreId field if non-nil, zero value otherwise.

### GetAgentBlueprintStoreIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoreIdOk() (*string, bool)`

GetAgentBlueprintStoreIdOk returns a tuple with the AgentBlueprintStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStoreId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintStoreId(v string)`

SetAgentBlueprintStoreId sets AgentBlueprintStoreId field to given value.

### HasAgentBlueprintStoreId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintStoreId() bool`

HasAgentBlueprintStoreId returns a boolean if a field has been set.

### GetAgentBlueprintId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintId() string`

GetAgentBlueprintId returns the AgentBlueprintId field if non-nil, zero value otherwise.

### GetAgentBlueprintIdOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintIdOk() (*string, bool)`

GetAgentBlueprintIdOk returns a tuple with the AgentBlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintId(v string)`

SetAgentBlueprintId sets AgentBlueprintId field to given value.

### HasAgentBlueprintId

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintId() bool`

HasAgentBlueprintId returns a boolean if a field has been set.

### GetJobType

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetOperationType

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOperationType() int32`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOperationTypeOk() (*int32, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetOperationType(v int32)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### SetOperationTypeNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetOperationTypeNil(b bool)`

 SetOperationTypeNil sets the value for OperationType to be an explicit nil

### UnsetOperationType
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetOperationType()`

UnsetOperationType ensures that no value is present for OperationType, not even an explicit nil
### GetThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetContents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetAlias

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### SetAliasNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAliasNil(b bool)`

 SetAliasNil sets the value for Alias to be an explicit nil

### UnsetAlias
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetAlias()`

UnsetAlias ensures that no value is present for Alias, not even an explicit nil
### GetPrivateKeyEntry

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetPrivateKeyEntry() bool`

GetPrivateKeyEntry returns the PrivateKeyEntry field if non-nil, zero value otherwise.

### GetPrivateKeyEntryOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetPrivateKeyEntryOk() (*bool, bool)`

GetPrivateKeyEntryOk returns a tuple with the PrivateKeyEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEntry

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetPrivateKeyEntry(v bool)`

SetPrivateKeyEntry sets PrivateKeyEntry field to given value.

### HasPrivateKeyEntry

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasPrivateKeyEntry() bool`

HasPrivateKeyEntry returns a boolean if a field has been set.

### SetPrivateKeyEntryNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetPrivateKeyEntryNil(b bool)`

 SetPrivateKeyEntryNil sets the value for PrivateKeyEntry to be an explicit nil

### UnsetPrivateKeyEntry
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetPrivateKeyEntry()`

UnsetPrivateKeyEntry ensures that no value is present for PrivateKeyEntry, not even an explicit nil
### GetOverwrite

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### SetOverwriteNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetOverwriteNil(b bool)`

 SetOverwriteNil sets the value for Overwrite to be an explicit nil

### UnsetOverwrite
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetOverwrite()`

UnsetOverwrite ensures that no value is present for Overwrite, not even an explicit nil
### GetHasEntryPassword

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasEntryPassword() bool`

GetHasEntryPassword returns the HasEntryPassword field if non-nil, zero value otherwise.

### GetHasEntryPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasEntryPasswordOk() (*bool, bool)`

GetHasEntryPasswordOk returns a tuple with the HasEntryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntryPassword

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetHasEntryPassword(v bool)`

SetHasEntryPassword sets HasEntryPassword field to given value.

### HasHasEntryPassword

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasHasEntryPassword() bool`

HasHasEntryPassword returns a boolean if a field has been set.

### SetHasEntryPasswordNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetHasEntryPasswordNil(b bool)`

 SetHasEntryPasswordNil sets the value for HasEntryPassword to be an explicit nil

### UnsetHasEntryPassword
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetHasEntryPassword()`

UnsetHasEntryPassword ensures that no value is present for HasEntryPassword, not even an explicit nil
### GetHasPfxPassword

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasPfxPassword() bool`

GetHasPfxPassword returns the HasPfxPassword field if non-nil, zero value otherwise.

### GetHasPfxPasswordOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasPfxPasswordOk() (*bool, bool)`

GetHasPfxPasswordOk returns a tuple with the HasPfxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPfxPassword

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetHasPfxPassword(v bool)`

SetHasPfxPassword sets HasPfxPassword field to given value.

### HasHasPfxPassword

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasHasPfxPassword() bool`

HasHasPfxPassword returns a boolean if a field has been set.

### SetHasPfxPasswordNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetHasPfxPasswordNil(b bool)`

 SetHasPfxPasswordNil sets the value for HasPfxPassword to be an explicit nil

### UnsetHasPfxPassword
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetHasPfxPassword()`

UnsetHasPfxPassword ensures that no value is present for HasPfxPassword, not even an explicit nil
### GetRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### SetRequestTimestampNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetRequestTimestampNil(b bool)`

 SetRequestTimestampNil sets the value for RequestTimestamp to be an explicit nil

### UnsetRequestTimestamp
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetRequestTimestamp()`

UnsetRequestTimestamp ensures that no value is present for RequestTimestamp, not even an explicit nil
### GetKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetDirectories

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetDirectories() string`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetDirectoriesOk() (*string, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetDirectories(v string)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### SetDirectoriesNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetDirectoriesNil(b bool)`

 SetDirectoriesNil sets the value for Directories to be an explicit nil

### UnsetDirectories
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetDirectories()`

UnsetDirectories ensures that no value is present for Directories, not even an explicit nil
### GetIgnoredDirectories

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetIgnoredDirectories() string`

GetIgnoredDirectories returns the IgnoredDirectories field if non-nil, zero value otherwise.

### GetIgnoredDirectoriesOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetIgnoredDirectoriesOk() (*string, bool)`

GetIgnoredDirectoriesOk returns a tuple with the IgnoredDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredDirectories

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetIgnoredDirectories(v string)`

SetIgnoredDirectories sets IgnoredDirectories field to given value.

### HasIgnoredDirectories

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasIgnoredDirectories() bool`

HasIgnoredDirectories returns a boolean if a field has been set.

### SetIgnoredDirectoriesNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetIgnoredDirectoriesNil(b bool)`

 SetIgnoredDirectoriesNil sets the value for IgnoredDirectories to be an explicit nil

### UnsetIgnoredDirectories
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetIgnoredDirectories()`

UnsetIgnoredDirectories ensures that no value is present for IgnoredDirectories, not even an explicit nil
### GetSymLinks

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSymLinks() bool`

GetSymLinks returns the SymLinks field if non-nil, zero value otherwise.

### GetSymLinksOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSymLinksOk() (*bool, bool)`

GetSymLinksOk returns a tuple with the SymLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymLinks

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetSymLinks(v bool)`

SetSymLinks sets SymLinks field to given value.

### HasSymLinks

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasSymLinks() bool`

HasSymLinks returns a boolean if a field has been set.

### SetSymLinksNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetSymLinksNil(b bool)`

 SetSymLinksNil sets the value for SymLinks to be an explicit nil

### UnsetSymLinks
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetSymLinks()`

UnsetSymLinks ensures that no value is present for SymLinks, not even an explicit nil
### GetCompatibility

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### SetCompatibilityNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetCompatibilityNil(b bool)`

 SetCompatibilityNil sets the value for Compatibility to be an explicit nil

### UnsetCompatibility
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetCompatibility()`

UnsetCompatibility ensures that no value is present for Compatibility, not even an explicit nil
### GetFileExtensions

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileExtensions() string`

GetFileExtensions returns the FileExtensions field if non-nil, zero value otherwise.

### GetFileExtensionsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileExtensionsOk() (*string, bool)`

GetFileExtensionsOk returns a tuple with the FileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensions

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetFileExtensions(v string)`

SetFileExtensions sets FileExtensions field to given value.

### HasFileExtensions

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasFileExtensions() bool`

HasFileExtensions returns a boolean if a field has been set.

### SetFileExtensionsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetFileExtensionsNil(b bool)`

 SetFileExtensionsNil sets the value for FileExtensions to be an explicit nil

### UnsetFileExtensions
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetFileExtensions()`

UnsetFileExtensions ensures that no value is present for FileExtensions, not even an explicit nil
### GetFileNamePatterns

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileNamePatterns() string`

GetFileNamePatterns returns the FileNamePatterns field if non-nil, zero value otherwise.

### GetFileNamePatternsOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileNamePatternsOk() (*string, bool)`

GetFileNamePatternsOk returns a tuple with the FileNamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePatterns

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetFileNamePatterns(v string)`

SetFileNamePatterns sets FileNamePatterns field to given value.

### HasFileNamePatterns

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasFileNamePatterns() bool`

HasFileNamePatterns returns a boolean if a field has been set.

### SetFileNamePatternsNil

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetFileNamePatternsNil(b bool)`

 SetFileNamePatternsNil sets the value for FileNamePatterns to be an explicit nil

### UnsetFileNamePatterns
`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) UnsetFileNamePatterns()`

UnsetFileNamePatterns ensures that no value is present for FileNamePatterns, not even an explicit nil
### GetAgentBlueprintStores

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStores() KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse`

GetAgentBlueprintStores returns the AgentBlueprintStores field if non-nil, zero value otherwise.

### GetAgentBlueprintStoresOk

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoresOk() (*KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse, bool)`

GetAgentBlueprintStoresOk returns a tuple with the AgentBlueprintStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStores

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintStores(v KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse)`

SetAgentBlueprintStores sets AgentBlueprintStores field to given value.

### HasAgentBlueprintStores

`func (o *KeyfactorWebKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintStores() bool`

HasAgentBlueprintStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


