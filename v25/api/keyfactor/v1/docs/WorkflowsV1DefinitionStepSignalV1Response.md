# WorkflowsV1DefinitionStepSignalV1Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalName** | Pointer to **NullableString** | The name of the signal. | [optional] 
**RoleIds** | Pointer to **[]int32** | The roles that are allowed to send this signal. | [optional] 
**RoleTokens** | Pointer to **[]string** | The role tokens that are allowed to send the signal (resolved on run-time of the specific step). | [optional] 

## Methods

### NewWorkflowsV1DefinitionStepSignalV1Response

`func NewWorkflowsV1DefinitionStepSignalV1Response() *WorkflowsV1DefinitionStepSignalV1Response`

NewWorkflowsV1DefinitionStepSignalV1Response instantiates a new WorkflowsV1DefinitionStepSignalV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsV1DefinitionStepSignalV1ResponseWithDefaults

`func NewWorkflowsV1DefinitionStepSignalV1ResponseWithDefaults() *WorkflowsV1DefinitionStepSignalV1Response`

NewWorkflowsV1DefinitionStepSignalV1ResponseWithDefaults instantiates a new WorkflowsV1DefinitionStepSignalV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalName

`func (o *WorkflowsV1DefinitionStepSignalV1Response) GetSignalName() string`

GetSignalName returns the SignalName field if non-nil, zero value otherwise.

### GetSignalNameOk

`func (o *WorkflowsV1DefinitionStepSignalV1Response) GetSignalNameOk() (*string, bool)`

GetSignalNameOk returns a tuple with the SignalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalName

`func (o *WorkflowsV1DefinitionStepSignalV1Response) SetSignalName(v string)`

SetSignalName sets SignalName field to given value.

### HasSignalName

`func (o *WorkflowsV1DefinitionStepSignalV1Response) HasSignalName() bool`

HasSignalName returns a boolean if a field has been set.

### SetSignalNameNil

`func (o *WorkflowsV1DefinitionStepSignalV1Response) SetSignalNameNil(b bool)`

 SetSignalNameNil sets the value for SignalName to be an explicit nil

### UnsetSignalName
`func (o *WorkflowsV1DefinitionStepSignalV1Response) UnsetSignalName()`

UnsetSignalName ensures that no value is present for SignalName, not even an explicit nil
### GetRoleIds

`func (o *WorkflowsV1DefinitionStepSignalV1Response) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *WorkflowsV1DefinitionStepSignalV1Response) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *WorkflowsV1DefinitionStepSignalV1Response) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *WorkflowsV1DefinitionStepSignalV1Response) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### SetRoleIdsNil

`func (o *WorkflowsV1DefinitionStepSignalV1Response) SetRoleIdsNil(b bool)`

 SetRoleIdsNil sets the value for RoleIds to be an explicit nil

### UnsetRoleIds
`func (o *WorkflowsV1DefinitionStepSignalV1Response) UnsetRoleIds()`

UnsetRoleIds ensures that no value is present for RoleIds, not even an explicit nil
### GetRoleTokens

`func (o *WorkflowsV1DefinitionStepSignalV1Response) GetRoleTokens() []string`

GetRoleTokens returns the RoleTokens field if non-nil, zero value otherwise.

### GetRoleTokensOk

`func (o *WorkflowsV1DefinitionStepSignalV1Response) GetRoleTokensOk() (*[]string, bool)`

GetRoleTokensOk returns a tuple with the RoleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTokens

`func (o *WorkflowsV1DefinitionStepSignalV1Response) SetRoleTokens(v []string)`

SetRoleTokens sets RoleTokens field to given value.

### HasRoleTokens

`func (o *WorkflowsV1DefinitionStepSignalV1Response) HasRoleTokens() bool`

HasRoleTokens returns a boolean if a field has been set.

### SetRoleTokensNil

`func (o *WorkflowsV1DefinitionStepSignalV1Response) SetRoleTokensNil(b bool)`

 SetRoleTokensNil sets the value for RoleTokens to be an explicit nil

### UnsetRoleTokens
`func (o *WorkflowsV1DefinitionStepSignalV1Response) UnsetRoleTokens()`

UnsetRoleTokens ensures that no value is present for RoleTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


