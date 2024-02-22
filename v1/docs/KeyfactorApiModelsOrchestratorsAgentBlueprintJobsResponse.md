# KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentBlueprintJobId** | Pointer to **string** |  | [optional] 
**AgentBlueprintStoreId** | Pointer to **string** |  | [optional] 
**AgentBlueprintId** | Pointer to **string** |  | [optional] 
**JobType** | Pointer to **string** |  | [optional] 
**JobTypeName** | Pointer to **string** |  | [optional] 
**OperationType** | Pointer to **int32** |  | [optional] 
**Thumbprint** | Pointer to **string** |  | [optional] 
**Contents** | Pointer to **string** |  | [optional] 
**Alias** | Pointer to **string** |  | [optional] 
**PrivateKeyEntry** | Pointer to **bool** |  | [optional] 
**Overwrite** | Pointer to **bool** |  | [optional] 
**HasEntryPassword** | Pointer to **bool** |  | [optional] 
**HasPfxPassword** | Pointer to **bool** |  | [optional] 
**RequestTimestamp** | Pointer to **time.Time** |  | [optional] 
**KeyfactorSchedule** | Pointer to [**KeyfactorCommonSchedulingKeyfactorSchedule**](KeyfactorCommonSchedulingKeyfactorSchedule.md) |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Directories** | Pointer to **string** |  | [optional] 
**IgnoredDirectories** | Pointer to **string** |  | [optional] 
**SymLinks** | Pointer to **bool** |  | [optional] 
**Compatibility** | Pointer to **bool** |  | [optional] 
**FileExtensions** | Pointer to **string** |  | [optional] 
**FileNamePatterns** | Pointer to **string** |  | [optional] 
**AgentBlueprintStores** | Pointer to [**KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse**](KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse

`func NewKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse() *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse`

NewKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse instantiates a new KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponseWithDefaults

`func NewKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponseWithDefaults() *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse`

NewKeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponseWithDefaults instantiates a new KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentBlueprintJobId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintJobId() string`

GetAgentBlueprintJobId returns the AgentBlueprintJobId field if non-nil, zero value otherwise.

### GetAgentBlueprintJobIdOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintJobIdOk() (*string, bool)`

GetAgentBlueprintJobIdOk returns a tuple with the AgentBlueprintJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintJobId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintJobId(v string)`

SetAgentBlueprintJobId sets AgentBlueprintJobId field to given value.

### HasAgentBlueprintJobId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintJobId() bool`

HasAgentBlueprintJobId returns a boolean if a field has been set.

### GetAgentBlueprintStoreId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoreId() string`

GetAgentBlueprintStoreId returns the AgentBlueprintStoreId field if non-nil, zero value otherwise.

### GetAgentBlueprintStoreIdOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoreIdOk() (*string, bool)`

GetAgentBlueprintStoreIdOk returns a tuple with the AgentBlueprintStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStoreId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintStoreId(v string)`

SetAgentBlueprintStoreId sets AgentBlueprintStoreId field to given value.

### HasAgentBlueprintStoreId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintStoreId() bool`

HasAgentBlueprintStoreId returns a boolean if a field has been set.

### GetAgentBlueprintId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintId() string`

GetAgentBlueprintId returns the AgentBlueprintId field if non-nil, zero value otherwise.

### GetAgentBlueprintIdOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintIdOk() (*string, bool)`

GetAgentBlueprintIdOk returns a tuple with the AgentBlueprintId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintId(v string)`

SetAgentBlueprintId sets AgentBlueprintId field to given value.

### HasAgentBlueprintId

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintId() bool`

HasAgentBlueprintId returns a boolean if a field has been set.

### GetJobType

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobType() string`

GetJobType returns the JobType field if non-nil, zero value otherwise.

### GetJobTypeOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobTypeOk() (*string, bool)`

GetJobTypeOk returns a tuple with the JobType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobType

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetJobType(v string)`

SetJobType sets JobType field to given value.

### HasJobType

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasJobType() bool`

HasJobType returns a boolean if a field has been set.

### GetJobTypeName

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### GetOperationType

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOperationType() int32`

GetOperationType returns the OperationType field if non-nil, zero value otherwise.

### GetOperationTypeOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOperationTypeOk() (*int32, bool)`

GetOperationTypeOk returns a tuple with the OperationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationType

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetOperationType(v int32)`

SetOperationType sets OperationType field to given value.

### HasOperationType

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasOperationType() bool`

HasOperationType returns a boolean if a field has been set.

### GetThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### GetContents

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### GetAlias

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPrivateKeyEntry

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetPrivateKeyEntry() bool`

GetPrivateKeyEntry returns the PrivateKeyEntry field if non-nil, zero value otherwise.

### GetPrivateKeyEntryOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetPrivateKeyEntryOk() (*bool, bool)`

GetPrivateKeyEntryOk returns a tuple with the PrivateKeyEntry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyEntry

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetPrivateKeyEntry(v bool)`

SetPrivateKeyEntry sets PrivateKeyEntry field to given value.

### HasPrivateKeyEntry

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasPrivateKeyEntry() bool`

HasPrivateKeyEntry returns a boolean if a field has been set.

### GetOverwrite

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOverwrite() bool`

GetOverwrite returns the Overwrite field if non-nil, zero value otherwise.

### GetOverwriteOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetOverwriteOk() (*bool, bool)`

GetOverwriteOk returns a tuple with the Overwrite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverwrite

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetOverwrite(v bool)`

SetOverwrite sets Overwrite field to given value.

### HasOverwrite

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasOverwrite() bool`

HasOverwrite returns a boolean if a field has been set.

### GetHasEntryPassword

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasEntryPassword() bool`

GetHasEntryPassword returns the HasEntryPassword field if non-nil, zero value otherwise.

### GetHasEntryPasswordOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasEntryPasswordOk() (*bool, bool)`

GetHasEntryPasswordOk returns a tuple with the HasEntryPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEntryPassword

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetHasEntryPassword(v bool)`

SetHasEntryPassword sets HasEntryPassword field to given value.

### HasHasEntryPassword

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasHasEntryPassword() bool`

HasHasEntryPassword returns a boolean if a field has been set.

### GetHasPfxPassword

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasPfxPassword() bool`

GetHasPfxPassword returns the HasPfxPassword field if non-nil, zero value otherwise.

### GetHasPfxPasswordOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetHasPfxPasswordOk() (*bool, bool)`

GetHasPfxPasswordOk returns a tuple with the HasPfxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPfxPassword

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetHasPfxPassword(v bool)`

SetHasPfxPassword sets HasPfxPassword field to given value.

### HasHasPfxPassword

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasHasPfxPassword() bool`

HasHasPfxPassword returns a boolean if a field has been set.

### GetRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetRequestTimestamp() time.Time`

GetRequestTimestamp returns the RequestTimestamp field if non-nil, zero value otherwise.

### GetRequestTimestampOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetRequestTimestampOk() (*time.Time, bool)`

GetRequestTimestampOk returns a tuple with the RequestTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetRequestTimestamp(v time.Time)`

SetRequestTimestamp sets RequestTimestamp field to given value.

### HasRequestTimestamp

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasRequestTimestamp() bool`

HasRequestTimestamp returns a boolean if a field has been set.

### GetKeyfactorSchedule

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetKeyfactorSchedule() KeyfactorCommonSchedulingKeyfactorSchedule`

GetKeyfactorSchedule returns the KeyfactorSchedule field if non-nil, zero value otherwise.

### GetKeyfactorScheduleOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetKeyfactorScheduleOk() (*KeyfactorCommonSchedulingKeyfactorSchedule, bool)`

GetKeyfactorScheduleOk returns a tuple with the KeyfactorSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorSchedule

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetKeyfactorSchedule(v KeyfactorCommonSchedulingKeyfactorSchedule)`

SetKeyfactorSchedule sets KeyfactorSchedule field to given value.

### HasKeyfactorSchedule

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasKeyfactorSchedule() bool`

HasKeyfactorSchedule returns a boolean if a field has been set.

### GetSubject

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetDirectories

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetDirectories() string`

GetDirectories returns the Directories field if non-nil, zero value otherwise.

### GetDirectoriesOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetDirectoriesOk() (*string, bool)`

GetDirectoriesOk returns a tuple with the Directories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectories

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetDirectories(v string)`

SetDirectories sets Directories field to given value.

### HasDirectories

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasDirectories() bool`

HasDirectories returns a boolean if a field has been set.

### GetIgnoredDirectories

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetIgnoredDirectories() string`

GetIgnoredDirectories returns the IgnoredDirectories field if non-nil, zero value otherwise.

### GetIgnoredDirectoriesOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetIgnoredDirectoriesOk() (*string, bool)`

GetIgnoredDirectoriesOk returns a tuple with the IgnoredDirectories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredDirectories

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetIgnoredDirectories(v string)`

SetIgnoredDirectories sets IgnoredDirectories field to given value.

### HasIgnoredDirectories

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasIgnoredDirectories() bool`

HasIgnoredDirectories returns a boolean if a field has been set.

### GetSymLinks

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSymLinks() bool`

GetSymLinks returns the SymLinks field if non-nil, zero value otherwise.

### GetSymLinksOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetSymLinksOk() (*bool, bool)`

GetSymLinksOk returns a tuple with the SymLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymLinks

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetSymLinks(v bool)`

SetSymLinks sets SymLinks field to given value.

### HasSymLinks

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasSymLinks() bool`

HasSymLinks returns a boolean if a field has been set.

### GetCompatibility

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### GetFileExtensions

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileExtensions() string`

GetFileExtensions returns the FileExtensions field if non-nil, zero value otherwise.

### GetFileExtensionsOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileExtensionsOk() (*string, bool)`

GetFileExtensionsOk returns a tuple with the FileExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileExtensions

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetFileExtensions(v string)`

SetFileExtensions sets FileExtensions field to given value.

### HasFileExtensions

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasFileExtensions() bool`

HasFileExtensions returns a boolean if a field has been set.

### GetFileNamePatterns

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileNamePatterns() string`

GetFileNamePatterns returns the FileNamePatterns field if non-nil, zero value otherwise.

### GetFileNamePatternsOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetFileNamePatternsOk() (*string, bool)`

GetFileNamePatternsOk returns a tuple with the FileNamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileNamePatterns

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetFileNamePatterns(v string)`

SetFileNamePatterns sets FileNamePatterns field to given value.

### HasFileNamePatterns

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasFileNamePatterns() bool`

HasFileNamePatterns returns a boolean if a field has been set.

### GetAgentBlueprintStores

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStores() KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse`

GetAgentBlueprintStores returns the AgentBlueprintStores field if non-nil, zero value otherwise.

### GetAgentBlueprintStoresOk

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) GetAgentBlueprintStoresOk() (*KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse, bool)`

GetAgentBlueprintStoresOk returns a tuple with the AgentBlueprintStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentBlueprintStores

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) SetAgentBlueprintStores(v KeyfactorApiModelsOrchestratorsAgentBlueprintStoresResponse)`

SetAgentBlueprintStores sets AgentBlueprintStores field to given value.

### HasAgentBlueprintStores

`func (o *KeyfactorApiModelsOrchestratorsAgentBlueprintJobsResponse) HasAgentBlueprintStores() bool`

HasAgentBlueprintStores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


