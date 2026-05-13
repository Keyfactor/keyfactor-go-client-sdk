# WorkflowsDefinitionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Description** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Keys** | Pointer to **[]string** | The Keys to set on the existing workflow definition. | [optional] 

## Methods

### NewWorkflowsDefinitionUpdateRequest

`func NewWorkflowsDefinitionUpdateRequest(displayName string, description string, ) *WorkflowsDefinitionUpdateRequest`

NewWorkflowsDefinitionUpdateRequest instantiates a new WorkflowsDefinitionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionUpdateRequestWithDefaults

`func NewWorkflowsDefinitionUpdateRequestWithDefaults() *WorkflowsDefinitionUpdateRequest`

NewWorkflowsDefinitionUpdateRequestWithDefaults instantiates a new WorkflowsDefinitionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowsDefinitionUpdateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsDefinitionUpdateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsDefinitionUpdateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *WorkflowsDefinitionUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowsDefinitionUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowsDefinitionUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetEnabled

`func (o *WorkflowsDefinitionUpdateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsDefinitionUpdateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsDefinitionUpdateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsDefinitionUpdateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeys

`func (o *WorkflowsDefinitionUpdateRequest) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *WorkflowsDefinitionUpdateRequest) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *WorkflowsDefinitionUpdateRequest) SetKeys(v []string)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *WorkflowsDefinitionUpdateRequest) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *WorkflowsDefinitionUpdateRequest) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *WorkflowsDefinitionUpdateRequest) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


