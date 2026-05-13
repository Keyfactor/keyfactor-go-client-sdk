# WorkflowsDefinitionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | **string** |  | 
**Description** | **string** |  | 
**WorkflowType** | **string** |  | 
**Enabled** | Pointer to **bool** |  | [optional] 
**Keys** | **[]string** |  | 

## Methods

### NewWorkflowsDefinitionCreateRequest

`func NewWorkflowsDefinitionCreateRequest(displayName string, description string, workflowType string, keys []string, ) *WorkflowsDefinitionCreateRequest`

NewWorkflowsDefinitionCreateRequest instantiates a new WorkflowsDefinitionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionCreateRequestWithDefaults

`func NewWorkflowsDefinitionCreateRequestWithDefaults() *WorkflowsDefinitionCreateRequest`

NewWorkflowsDefinitionCreateRequestWithDefaults instantiates a new WorkflowsDefinitionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowsDefinitionCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsDefinitionCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsDefinitionCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetDescription

`func (o *WorkflowsDefinitionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowsDefinitionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowsDefinitionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetWorkflowType

`func (o *WorkflowsDefinitionCreateRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowsDefinitionCreateRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowsDefinitionCreateRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.


### GetEnabled

`func (o *WorkflowsDefinitionCreateRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsDefinitionCreateRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsDefinitionCreateRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsDefinitionCreateRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetKeys

`func (o *WorkflowsDefinitionCreateRequest) GetKeys() []string`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *WorkflowsDefinitionCreateRequest) GetKeysOk() (*[]string, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *WorkflowsDefinitionCreateRequest) SetKeys(v []string)`

SetKeys sets Keys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


