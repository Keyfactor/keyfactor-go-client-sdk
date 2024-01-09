# CSSCMSDataModelModelsDiscoveryJobRequest

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
**ServerUsername** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ServerPassword** | Pointer to [**CSSCMSDataModelModelsKeyfactorAPISecret**](CSSCMSDataModelModelsKeyfactorAPISecret.md) |  | [optional] 
**ServerUseSsl** | Pointer to **bool** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsDiscoveryJobRequest

`func NewCSSCMSDataModelModelsDiscoveryJobRequest(type_ int32, ) *CSSCMSDataModelModelsDiscoveryJobRequest`

NewCSSCMSDataModelModelsDiscoveryJobRequest instantiates a new CSSCMSDataModelModelsDiscoveryJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsDiscoveryJobRequestWithDefaults

`func NewCSSCMSDataModelModelsDiscoveryJobRequestWithDefaults() *CSSCMSDataModelModelsDiscoveryJobRequest`

NewCSSCMSDataModelModelsDiscoveryJobRequestWithDefaults instantiates a new CSSCMSDataModelModelsDiscoveryJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientMachine

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetClientMachine() string`

GetClientMachine returns the ClientMachine field if non-nil, zero value otherwise.

### GetClientMachineOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetClientMachineOk() (*string, bool)`

GetClientMachineOk returns a tuple with the ClientMachine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientMachine

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetClientMachine(v string)`

SetClientMachine sets ClientMachine field to given value.

### HasClientMachine

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasClientMachine() bool`

HasClientMachine returns a boolean if a field has been set.

### SetClientMachineNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetClientMachineNil(b bool)`

 SetClientMachineNil sets the value for ClientMachine to be an explicit nil

### UnsetClientMachine
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetClientMachine()`

UnsetClientMachine ensures that no value is present for ClientMachine, not even an explicit nil
### GetAgentId

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetAgentId() string`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetAgentIdOk() (*string, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetAgentId(v string)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### SetAgentIdNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetAgentIdNil(b bool)`

 SetAgentIdNil sets the value for AgentId to be an explicit nil

### UnsetAgentId
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetAgentId()`

UnsetAgentId ensures that no value is present for AgentId, not even an explicit nil
### GetType

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetType(v int32)`

SetType sets Type field to given value.


### GetJobExecutionTimestamp

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetJobExecutionTimestamp() time.Time`

GetJobExecutionTimestamp returns the JobExecutionTimestamp field if non-nil, zero value otherwise.

### GetJobExecutionTimestampOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetJobExecutionTimestampOk() (*time.Time, bool)`

GetJobExecutionTimestampOk returns a tuple with the JobExecutionTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobExecutionTimestamp

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetJobExecutionTimestamp(v time.Time)`

SetJobExecutionTimestamp sets JobExecutionTimestamp field to given value.

### HasJobExecutionTimestamp

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasJobExecutionTimestamp() bool`

HasJobExecutionTimestamp returns a boolean if a field has been set.

### SetJobExecutionTimestampNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetJobExecutionTimestampNil(b bool)`

 SetJobExecutionTimestampNil sets the value for JobExecutionTimestamp to be an explicit nil

### UnsetJobExecutionTimestamp
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetJobExecutionTimestamp()`

UnsetJobExecutionTimestamp ensures that no value is present for JobExecutionTimestamp, not even an explicit nil
### GetDirs

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetDirs() string`

GetDirs returns the Dirs field if non-nil, zero value otherwise.

### GetDirsOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetDirsOk() (*string, bool)`

GetDirsOk returns a tuple with the Dirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirs

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetDirs(v string)`

SetDirs sets Dirs field to given value.

### HasDirs

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasDirs() bool`

HasDirs returns a boolean if a field has been set.

### SetDirsNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetDirsNil(b bool)`

 SetDirsNil sets the value for Dirs to be an explicit nil

### UnsetDirs
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetDirs()`

UnsetDirs ensures that no value is present for Dirs, not even an explicit nil
### GetIgnoredDirs

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetIgnoredDirs() string`

GetIgnoredDirs returns the IgnoredDirs field if non-nil, zero value otherwise.

### GetIgnoredDirsOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetIgnoredDirsOk() (*string, bool)`

GetIgnoredDirsOk returns a tuple with the IgnoredDirs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoredDirs

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetIgnoredDirs(v string)`

SetIgnoredDirs sets IgnoredDirs field to given value.

### HasIgnoredDirs

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasIgnoredDirs() bool`

HasIgnoredDirs returns a boolean if a field has been set.

### SetIgnoredDirsNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetIgnoredDirsNil(b bool)`

 SetIgnoredDirsNil sets the value for IgnoredDirs to be an explicit nil

### UnsetIgnoredDirs
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetIgnoredDirs()`

UnsetIgnoredDirs ensures that no value is present for IgnoredDirs, not even an explicit nil
### GetExtensions

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetExtensions() string`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetExtensionsOk() (*string, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetExtensions(v string)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### SetExtensionsNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetExtensionsNil(b bool)`

 SetExtensionsNil sets the value for Extensions to be an explicit nil

### UnsetExtensions
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetExtensions()`

UnsetExtensions ensures that no value is present for Extensions, not even an explicit nil
### GetNamePatterns

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetNamePatterns() string`

GetNamePatterns returns the NamePatterns field if non-nil, zero value otherwise.

### GetNamePatternsOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetNamePatternsOk() (*string, bool)`

GetNamePatternsOk returns a tuple with the NamePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePatterns

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetNamePatterns(v string)`

SetNamePatterns sets NamePatterns field to given value.

### HasNamePatterns

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasNamePatterns() bool`

HasNamePatterns returns a boolean if a field has been set.

### SetNamePatternsNil

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetNamePatternsNil(b bool)`

 SetNamePatternsNil sets the value for NamePatterns to be an explicit nil

### UnsetNamePatterns
`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) UnsetNamePatterns()`

UnsetNamePatterns ensures that no value is present for NamePatterns, not even an explicit nil
### GetSymLinks

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetSymLinks() bool`

GetSymLinks returns the SymLinks field if non-nil, zero value otherwise.

### GetSymLinksOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetSymLinksOk() (*bool, bool)`

GetSymLinksOk returns a tuple with the SymLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymLinks

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetSymLinks(v bool)`

SetSymLinks sets SymLinks field to given value.

### HasSymLinks

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasSymLinks() bool`

HasSymLinks returns a boolean if a field has been set.

### GetCompatibility

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetCompatibility() bool`

GetCompatibility returns the Compatibility field if non-nil, zero value otherwise.

### GetCompatibilityOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetCompatibilityOk() (*bool, bool)`

GetCompatibilityOk returns a tuple with the Compatibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibility

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetCompatibility(v bool)`

SetCompatibility sets Compatibility field to given value.

### HasCompatibility

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasCompatibility() bool`

HasCompatibility returns a boolean if a field has been set.

### GetServerUsername

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetServerUsername() CSSCMSDataModelModelsKeyfactorAPISecret`

GetServerUsername returns the ServerUsername field if non-nil, zero value otherwise.

### GetServerUsernameOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetServerUsernameOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetServerUsernameOk returns a tuple with the ServerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUsername

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetServerUsername(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetServerUsername sets ServerUsername field to given value.

### HasServerUsername

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasServerUsername() bool`

HasServerUsername returns a boolean if a field has been set.

### GetServerPassword

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetServerPassword() CSSCMSDataModelModelsKeyfactorAPISecret`

GetServerPassword returns the ServerPassword field if non-nil, zero value otherwise.

### GetServerPasswordOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetServerPasswordOk() (*CSSCMSDataModelModelsKeyfactorAPISecret, bool)`

GetServerPasswordOk returns a tuple with the ServerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPassword

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetServerPassword(v CSSCMSDataModelModelsKeyfactorAPISecret)`

SetServerPassword sets ServerPassword field to given value.

### HasServerPassword

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasServerPassword() bool`

HasServerPassword returns a boolean if a field has been set.

### GetServerUseSsl

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetServerUseSsl() bool`

GetServerUseSsl returns the ServerUseSsl field if non-nil, zero value otherwise.

### GetServerUseSslOk

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) GetServerUseSslOk() (*bool, bool)`

GetServerUseSslOk returns a tuple with the ServerUseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUseSsl

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) SetServerUseSsl(v bool)`

SetServerUseSsl sets ServerUseSsl field to given value.

### HasServerUseSsl

`func (o *CSSCMSDataModelModelsDiscoveryJobRequest) HasServerUseSsl() bool`

HasServerUseSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


