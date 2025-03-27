# WorkflowsSignalConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalName** | Pointer to **NullableString** | The name of the signal. | [optional] 
**RoleIds** | Pointer to **[]int32** | The roles that are allowed to send the signal. | [optional] 

## Methods

### NewWorkflowsSignalConfigurationRequest

`func NewWorkflowsSignalConfigurationRequest() *WorkflowsSignalConfigurationRequest`

NewWorkflowsSignalConfigurationRequest instantiates a new WorkflowsSignalConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsSignalConfigurationRequestWithDefaults

`func NewWorkflowsSignalConfigurationRequestWithDefaults() *WorkflowsSignalConfigurationRequest`

NewWorkflowsSignalConfigurationRequestWithDefaults instantiates a new WorkflowsSignalConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalName

`func (o *WorkflowsSignalConfigurationRequest) GetSignalName() string`

GetSignalName returns the SignalName field if non-nil, zero value otherwise.

### GetSignalNameOk

`func (o *WorkflowsSignalConfigurationRequest) GetSignalNameOk() (*string, bool)`

GetSignalNameOk returns a tuple with the SignalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalName

`func (o *WorkflowsSignalConfigurationRequest) SetSignalName(v string)`

SetSignalName sets SignalName field to given value.

### HasSignalName

`func (o *WorkflowsSignalConfigurationRequest) HasSignalName() bool`

HasSignalName returns a boolean if a field has been set.

### SetSignalNameNil

`func (o *WorkflowsSignalConfigurationRequest) SetSignalNameNil(b bool)`

 SetSignalNameNil sets the value for SignalName to be an explicit nil

### UnsetSignalName
`func (o *WorkflowsSignalConfigurationRequest) UnsetSignalName()`

UnsetSignalName ensures that no value is present for SignalName, not even an explicit nil
### GetRoleIds

`func (o *WorkflowsSignalConfigurationRequest) GetRoleIds() []int32`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *WorkflowsSignalConfigurationRequest) GetRoleIdsOk() (*[]int32, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *WorkflowsSignalConfigurationRequest) SetRoleIds(v []int32)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *WorkflowsSignalConfigurationRequest) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.

### SetRoleIdsNil

`func (o *WorkflowsSignalConfigurationRequest) SetRoleIdsNil(b bool)`

 SetRoleIdsNil sets the value for RoleIds to be an explicit nil

### UnsetRoleIds
`func (o *WorkflowsSignalConfigurationRequest) UnsetRoleIds()`

UnsetRoleIds ensures that no value is present for RoleIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


