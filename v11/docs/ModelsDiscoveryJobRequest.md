# ModelsDiscoveryJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientMachine** | Pointer to **NullableString** |  | [optional] 
**AgentId** | Pointer to **NullableString** |  | [optional] 
**Type** | **int32** |  | 
**JobExecutionTimestamp** | Pointer to **NullableTime** |  | [optional] 
**Dirs** | Pointer to **NullableString** |  | [optional] 
**IgnoredDirs** | Pointer to **NullableString** |  | [optional] 
**Extensions** | Pointer to **NullableString** |  | [optional] 
**NamePatterns** | Pointer to **NullableString** |  | [optional] 
**SymLinks** | Pointer to **bool** |  | [optional] 
**Compatibility** | Pointer to **bool** |  | [optional] 
**ServerUsername** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**ServerPassword** | Pointer to [**ModelsKeyfactorAPISecret**](ModelsKeyfactorAPISecret.md) |  | [optional] 
**ServerUseSsl** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsDiscoveryJobRequest

`func NewModelsDiscoveryJobRequest(type_ int32, ) *ModelsDiscoveryJobRequest`

NewModelsDiscoveryJobRequest instantiates a new ModelsDiscoveryJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsDiscoveryJobRequestWithDefaults

`func NewModelsDiscoveryJobRequestWithDefaults() *ModelsDiscoveryJobRequest`

NewModelsDiscoveryJobRequestWithDefaults instantiates a new ModelsDiscoveryJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMachine

`func (o *ModelsDiscoveryJobRequest) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *ModelsDiscoveryJobRequest) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *ModelsDiscoveryJobRequest) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *ModelsDiscoveryJobRequest) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *ModelsDiscoveryJobRequest) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *ModelsDiscoveryJobRequest) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetAgentId

`func (o *ModelsDiscoveryJobRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *ModelsDiscoveryJobRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *ModelsDiscoveryJobRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *ModelsDiscoveryJobRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *ModelsDiscoveryJobRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *ModelsDiscoveryJobRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetType

`func (o *ModelsDiscoveryJobRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ModelsDiscoveryJobRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ModelsDiscoveryJobRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetJobExecutionTimestamp

`func (o *ModelsDiscoveryJobRequest) GetJobExecutionTimestamp() time.Time`

GetJobExecutionTimestamp returns the JobExecutionTimestamp field if non-nil, zero value otherwise.

### GetJobExecutionTimestampOk

`func (o *ModelsDiscoveryJobRequest) GetJobExecutionTimestampOk() (*time.Time, bool)`

GetJobExecutionTimestampOk returns a tuple with the JobExecutionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecutionTimestamp

`func (o *ModelsDiscoveryJobRequest) SetJobExecutionTimestamp(v time.Time)`

SetJobExecutionTimestamp sets JobExecutionTimestamp field to given value.

### HasJobExecutionTimestamp

`func (o *ModelsDiscoveryJobRequest) HasJobExecutionTimestamp() bool`

HasJobExecutionTimestamp returns a boolean if a field has been set.

### SetJobExecutionTimestampNil

`func (o *ModelsDiscoveryJobRequest) SetJobExecutionTimestampNil(b bool)`

 SetJobExecutionTimestampNil sets the value for JobExecutionTimestamp to be an explicit nil

### UnsetJobExecutionTimestamp
`func (o *ModelsDiscoveryJobRequest) UnsetJobExecutionTimestamp()`

UnsetJobExecutionTimestamp ensures that no value is present for JobExecutionTimestamp, not even an explicit nil
### GetDirs

`func (o *ModelsDiscoveryJobRequest) GetDirs() string`

GetDirs returns the Dirs field if non-nil, zero value otherwise.

### GetDirsOk

`func (o *ModelsDiscoveryJobRequest) GetDirsOk() (*string, bool)`

GetDirsOk returns a tuple with the Dirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirs

`func (o *ModelsDiscoveryJobRequest) SetDirs(v string)`

SetDirs sets Dirs field to given value.

### HasDirs

`func (o *ModelsDiscoveryJobRequest) HasDirs() bool`

HasDirs returns a boolean if a field has been set.

### SetDirsNil

`func (o *ModelsDiscoveryJobRequest) SetDirsNil(b bool)`

 SetDirsNil sets the value for Dirs to be an explicit nil

### UnsetDirs
`func (o *ModelsDiscoveryJobRequest) UnsetDirs()`

UnsetDirs ensures that no value is present for Dirs, not even an explicit nil
### GetIgnoredDirs

`func (o *ModelsDiscoveryJobRequest) GetIgnoredDirs() string`

GetIgnoredDirs returns the IgnoredDirs field if non-nil, zero value otherwise.

### GetIgnoredDirsOk

`func (o *ModelsDiscoveryJobRequest) GetIgnoredDirsOk() (*string, bool)`

GetIgnoredDirsOk returns a tuple with the IgnoredDirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredDirs

`func (o *ModelsDiscoveryJobRequest) SetIgnoredDirs(v string)`

SetIgnoredDirs sets IgnoredDirs field to given value.

### HasIgnoredDirs

`func (o *ModelsDiscoveryJobRequest) HasIgnoredDirs() bool`

HasIgnoredDirs returns a boolean if a field has been set.

### SetIgnoredDirsNil

`func (o *ModelsDiscoveryJobRequest) SetIgnoredDirsNil(b bool)`

 SetIgnoredDirsNil sets the value for IgnoredDirs to be an explicit nil

### UnsetIgnoredDirs
`func (o *ModelsDiscoveryJobRequest) UnsetIgnoredDirs()`

UnsetIgnoredDirs ensures that no value is present for IgnoredDirs, not even an explicit nil
### GetExtensions

`func (o *ModelsDiscoveryJobRequest) GetExtensions() string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ModelsDiscoveryJobRequest) GetExtensionsOk() (*string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ModelsDiscoveryJobRequest) SetExtensions(v string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ModelsDiscoveryJobRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *ModelsDiscoveryJobRequest) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *ModelsDiscoveryJobRequest) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetNamePatterns

`func (o *ModelsDiscoveryJobRequest) GetNamePatterns() string`

GetNamePatterns returns the NamePatterns field if non-nil, zero value otherwise.

### GetNamePatternsOk

`func (o *ModelsDiscoveryJobRequest) GetNamePatternsOk() (*string, bool)`

GetNamePatternsOk returns a tuple with the NamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePatterns

`func (o *ModelsDiscoveryJobRequest) SetNamePatterns(v string)`

SetNamePatterns sets NamePatterns field to given value.

### HasNamePatterns

`func (o *ModelsDiscoveryJobRequest) HasNamePatterns() bool`

HasNamePatterns returns a boolean if a field has been set.

### SetNamePatternsNil

`func (o *ModelsDiscoveryJobRequest) SetNamePatternsNil(b bool)`

 SetNamePatternsNil sets the value for NamePatterns to be an explicit nil

### UnsetNamePatterns
`func (o *ModelsDiscoveryJobRequest) UnsetNamePatterns()`

UnsetNamePatterns ensures that no value is present for NamePatterns, not even an explicit nil
### GetSymLinks

`func (o *ModelsDiscoveryJobRequest) GetSymLinks() bool`

GetSymLinks returns the SymLinks field if non-nil, zero value otherwise.

### GetSymLinksOk

`func (o *ModelsDiscoveryJobRequest) GetSymLinksOk() (*bool, bool)`

GetSymLinksOk returns a tuple with the SymLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymLinks

`func (o *ModelsDiscoveryJobRequest) SetSymLinks(v bool)`

SetSymLinks sets SymLinks field to given value.

### HasSymLinks

`func (o *ModelsDiscoveryJobRequest) HasSymLinks() bool`

HasSymLinks returns a boolean if a field has been set.

### GetCompatibility

`func (o *ModelsDiscoveryJobRequest) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *ModelsDiscoveryJobRequest) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *ModelsDiscoveryJobRequest) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *ModelsDiscoveryJobRequest) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### GetServerUsername

`func (o *ModelsDiscoveryJobRequest) GetServerUsername() ModelsKeyfactorAPISecret`

GetServerUsername returns the ServerUsername field if non-nil, zero value otherwise.

### GetServerUsernameOk

`func (o *ModelsDiscoveryJobRequest) GetServerUsernameOk() (*ModelsKeyfactorAPISecret, bool)`

GetServerUsernameOk returns a tuple with the ServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUsername

`func (o *ModelsDiscoveryJobRequest) SetServerUsername(v ModelsKeyfactorAPISecret)`

SetServerUsername sets ServerUsername field to given value.

### HasServerUsername

`func (o *ModelsDiscoveryJobRequest) HasServerUsername() bool`

HasServerUsername returns a boolean if a field has been set.

### GetServerPassword

`func (o *ModelsDiscoveryJobRequest) GetServerPassword() ModelsKeyfactorAPISecret`

GetServerPassword returns the ServerPassword field if non-nil, zero value otherwise.

### GetServerPasswordOk

`func (o *ModelsDiscoveryJobRequest) GetServerPasswordOk() (*ModelsKeyfactorAPISecret, bool)`

GetServerPasswordOk returns a tuple with the ServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPassword

`func (o *ModelsDiscoveryJobRequest) SetServerPassword(v ModelsKeyfactorAPISecret)`

SetServerPassword sets ServerPassword field to given value.

### HasServerPassword

`func (o *ModelsDiscoveryJobRequest) HasServerPassword() bool`

HasServerPassword returns a boolean if a field has been set.

### GetServerUseSsl

`func (o *ModelsDiscoveryJobRequest) GetServerUseSsl() bool`

GetServerUseSsl returns the ServerUseSsl field if non-nil, zero value otherwise.

### GetServerUseSslOk

`func (o *ModelsDiscoveryJobRequest) GetServerUseSslOk() (*bool, bool)`

GetServerUseSslOk returns a tuple with the ServerUseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUseSsl

`func (o *ModelsDiscoveryJobRequest) SetServerUseSsl(v bool)`

SetServerUseSsl sets ServerUseSsl field to given value.

### HasServerUseSsl

`func (o *ModelsDiscoveryJobRequest) HasServerUseSsl() bool`

HasServerUseSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


