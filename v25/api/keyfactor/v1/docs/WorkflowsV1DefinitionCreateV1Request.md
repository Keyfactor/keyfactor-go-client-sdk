# WorkflowsV1DefinitionCreateV1Request

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **NullableString** | Display name of the Definition | [optional] 
**Description** | Pointer to **NullableString** | Description of the Definition | [optional] 
**Key** | Pointer to **NullableString** | Key to be used to look up definition when starting a new workflow. For enrollment workflowTypes, this should be a template | [optional] 
**WorkflowType** | Pointer to **NullableString** | The Type of Workflow | [optional] 
**Enabled** | Pointer to **bool** | If the Definition is enabled | [optional] 

## Methods

### NewWorkflowsV1DefinitionCreateV1Request

`func NewWorkflowsV1DefinitionCreateV1Request() *WorkflowsV1DefinitionCreateV1Request`

NewWorkflowsV1DefinitionCreateV1Request instantiates a new WorkflowsV1DefinitionCreateV1Request object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsV1DefinitionCreateV1RequestWithDefaults

`func NewWorkflowsV1DefinitionCreateV1RequestWithDefaults() *WorkflowsV1DefinitionCreateV1Request`

NewWorkflowsV1DefinitionCreateV1RequestWithDefaults instantiates a new WorkflowsV1DefinitionCreateV1Request object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *WorkflowsV1DefinitionCreateV1Request) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsV1DefinitionCreateV1Request) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsV1DefinitionCreateV1Request) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsV1DefinitionCreateV1Request) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsV1DefinitionCreateV1Request) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsV1DefinitionCreateV1Request) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *WorkflowsV1DefinitionCreateV1Request) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowsV1DefinitionCreateV1Request) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowsV1DefinitionCreateV1Request) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowsV1DefinitionCreateV1Request) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowsV1DefinitionCreateV1Request) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowsV1DefinitionCreateV1Request) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKey

`func (o *WorkflowsV1DefinitionCreateV1Request) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowsV1DefinitionCreateV1Request) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowsV1DefinitionCreateV1Request) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkflowsV1DefinitionCreateV1Request) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *WorkflowsV1DefinitionCreateV1Request) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *WorkflowsV1DefinitionCreateV1Request) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetWorkflowType

`func (o *WorkflowsV1DefinitionCreateV1Request) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowsV1DefinitionCreateV1Request) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowsV1DefinitionCreateV1Request) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowsV1DefinitionCreateV1Request) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### SetWorkflowTypeNil

`func (o *WorkflowsV1DefinitionCreateV1Request) SetWorkflowTypeNil(b bool)`

 SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil

### UnsetWorkflowType
`func (o *WorkflowsV1DefinitionCreateV1Request) UnsetWorkflowType()`

UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil
### GetEnabled

`func (o *WorkflowsV1DefinitionCreateV1Request) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsV1DefinitionCreateV1Request) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsV1DefinitionCreateV1Request) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsV1DefinitionCreateV1Request) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


