# WorkflowsDefinitionStepSignalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SignalName** | Pointer to **NullableString** | The name of the signal. | [optional] 
**Roles** | Pointer to [**[]WorkflowsDefinitionStepSignalRoleResponse**](WorkflowsDefinitionStepSignalRoleResponse.md) | The information for the roles that are allowed to send this signal. | [optional] 
**RoleTokens** | Pointer to **[]string** | The role tokens that are allowed to send the signal (resolved on run-time of the specific step). | [optional] 

## Methods

### NewWorkflowsDefinitionStepSignalResponse

`func NewWorkflowsDefinitionStepSignalResponse() *WorkflowsDefinitionStepSignalResponse`

NewWorkflowsDefinitionStepSignalResponse instantiates a new WorkflowsDefinitionStepSignalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionStepSignalResponseWithDefaults

`func NewWorkflowsDefinitionStepSignalResponseWithDefaults() *WorkflowsDefinitionStepSignalResponse`

NewWorkflowsDefinitionStepSignalResponseWithDefaults instantiates a new WorkflowsDefinitionStepSignalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignalName

`func (o *WorkflowsDefinitionStepSignalResponse) GetSignalName() string`

GetSignalName returns the SignalName field if non-nil, zero value otherwise.

### GetSignalNameOk

`func (o *WorkflowsDefinitionStepSignalResponse) GetSignalNameOk() (*string, bool)`

GetSignalNameOk returns a tuple with the SignalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignalName

`func (o *WorkflowsDefinitionStepSignalResponse) SetSignalName(v string)`

SetSignalName sets SignalName field to given value.

### HasSignalName

`func (o *WorkflowsDefinitionStepSignalResponse) HasSignalName() bool`

HasSignalName returns a boolean if a field has been set.

### SetSignalNameNil

`func (o *WorkflowsDefinitionStepSignalResponse) SetSignalNameNil(b bool)`

 SetSignalNameNil sets the value for SignalName to be an explicit nil

### UnsetSignalName
`func (o *WorkflowsDefinitionStepSignalResponse) UnsetSignalName()`

UnsetSignalName ensures that no value is present for SignalName, not even an explicit nil
### GetRoles

`func (o *WorkflowsDefinitionStepSignalResponse) GetRoles() []WorkflowsDefinitionStepSignalRoleResponse`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *WorkflowsDefinitionStepSignalResponse) GetRolesOk() (*[]WorkflowsDefinitionStepSignalRoleResponse, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *WorkflowsDefinitionStepSignalResponse) SetRoles(v []WorkflowsDefinitionStepSignalRoleResponse)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *WorkflowsDefinitionStepSignalResponse) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### SetRolesNil

`func (o *WorkflowsDefinitionStepSignalResponse) SetRolesNil(b bool)`

 SetRolesNil sets the value for Roles to be an explicit nil

### UnsetRoles
`func (o *WorkflowsDefinitionStepSignalResponse) UnsetRoles()`

UnsetRoles ensures that no value is present for Roles, not even an explicit nil
### GetRoleTokens

`func (o *WorkflowsDefinitionStepSignalResponse) GetRoleTokens() []string`

GetRoleTokens returns the RoleTokens field if non-nil, zero value otherwise.

### GetRoleTokensOk

`func (o *WorkflowsDefinitionStepSignalResponse) GetRoleTokensOk() (*[]string, bool)`

GetRoleTokensOk returns a tuple with the RoleTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleTokens

`func (o *WorkflowsDefinitionStepSignalResponse) SetRoleTokens(v []string)`

SetRoleTokens sets RoleTokens field to given value.

### HasRoleTokens

`func (o *WorkflowsDefinitionStepSignalResponse) HasRoleTokens() bool`

HasRoleTokens returns a boolean if a field has been set.

### SetRoleTokensNil

`func (o *WorkflowsDefinitionStepSignalResponse) SetRoleTokensNil(b bool)`

 SetRoleTokensNil sets the value for RoleTokens to be an explicit nil

### UnsetRoleTokens
`func (o *WorkflowsDefinitionStepSignalResponse) UnsetRoleTokens()`

UnsetRoleTokens ensures that no value is present for RoleTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


