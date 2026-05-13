# WorkflowsDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**WorkflowType** | Pointer to **NullableString** |  | [optional] 
**Keys** | Pointer to [**[]WorkflowsKeyResponse**](WorkflowsKeyResponse.md) |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**Steps** | Pointer to [**[]WorkflowsDefinitionStepResponse**](WorkflowsDefinitionStepResponse.md) |  | [optional] 
**DraftVersion** | Pointer to **int32** |  | [optional] 
**PublishedVersion** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewWorkflowsDefinitionResponse

`func NewWorkflowsDefinitionResponse() *WorkflowsDefinitionResponse`

NewWorkflowsDefinitionResponse instantiates a new WorkflowsDefinitionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionResponseWithDefaults

`func NewWorkflowsDefinitionResponseWithDefaults() *WorkflowsDefinitionResponse`

NewWorkflowsDefinitionResponseWithDefaults instantiates a new WorkflowsDefinitionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsDefinitionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsDefinitionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsDefinitionResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsDefinitionResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *WorkflowsDefinitionResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsDefinitionResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsDefinitionResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsDefinitionResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsDefinitionResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsDefinitionResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *WorkflowsDefinitionResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowsDefinitionResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowsDefinitionResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowsDefinitionResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowsDefinitionResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowsDefinitionResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetWorkflowType

`func (o *WorkflowsDefinitionResponse) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowsDefinitionResponse) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowsDefinitionResponse) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowsDefinitionResponse) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### SetWorkflowTypeNil

`func (o *WorkflowsDefinitionResponse) SetWorkflowTypeNil(b bool)`

 SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil

### UnsetWorkflowType
`func (o *WorkflowsDefinitionResponse) UnsetWorkflowType()`

UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil
### GetKeys

`func (o *WorkflowsDefinitionResponse) GetKeys() []WorkflowsKeyResponse`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *WorkflowsDefinitionResponse) GetKeysOk() (*[]WorkflowsKeyResponse, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *WorkflowsDefinitionResponse) SetKeys(v []WorkflowsKeyResponse)`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *WorkflowsDefinitionResponse) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### SetKeysNil

`func (o *WorkflowsDefinitionResponse) SetKeysNil(b bool)`

 SetKeysNil sets the value for Keys to be an explicit nil

### UnsetKeys
`func (o *WorkflowsDefinitionResponse) UnsetKeys()`

UnsetKeys ensures that no value is present for Keys, not even an explicit nil
### GetEnabled

`func (o *WorkflowsDefinitionResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsDefinitionResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsDefinitionResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsDefinitionResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetIsPublished

`func (o *WorkflowsDefinitionResponse) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *WorkflowsDefinitionResponse) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *WorkflowsDefinitionResponse) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *WorkflowsDefinitionResponse) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetSteps

`func (o *WorkflowsDefinitionResponse) GetSteps() []WorkflowsDefinitionStepResponse`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowsDefinitionResponse) GetStepsOk() (*[]WorkflowsDefinitionStepResponse, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowsDefinitionResponse) SetSteps(v []WorkflowsDefinitionStepResponse)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkflowsDefinitionResponse) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *WorkflowsDefinitionResponse) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *WorkflowsDefinitionResponse) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetDraftVersion

`func (o *WorkflowsDefinitionResponse) GetDraftVersion() int32`

GetDraftVersion returns the DraftVersion field if non-nil, zero value otherwise.

### GetDraftVersionOk

`func (o *WorkflowsDefinitionResponse) GetDraftVersionOk() (*int32, bool)`

GetDraftVersionOk returns a tuple with the DraftVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftVersion

`func (o *WorkflowsDefinitionResponse) SetDraftVersion(v int32)`

SetDraftVersion sets DraftVersion field to given value.

### HasDraftVersion

`func (o *WorkflowsDefinitionResponse) HasDraftVersion() bool`

HasDraftVersion returns a boolean if a field has been set.

### GetPublishedVersion

`func (o *WorkflowsDefinitionResponse) GetPublishedVersion() int32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *WorkflowsDefinitionResponse) GetPublishedVersionOk() (*int32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *WorkflowsDefinitionResponse) SetPublishedVersion(v int32)`

SetPublishedVersion sets PublishedVersion field to given value.

### HasPublishedVersion

`func (o *WorkflowsDefinitionResponse) HasPublishedVersion() bool`

HasPublishedVersion returns a boolean if a field has been set.

### SetPublishedVersionNil

`func (o *WorkflowsDefinitionResponse) SetPublishedVersionNil(b bool)`

 SetPublishedVersionNil sets the value for PublishedVersion to be an explicit nil

### UnsetPublishedVersion
`func (o *WorkflowsDefinitionResponse) UnsetPublishedVersion()`

UnsetPublishedVersion ensures that no value is present for PublishedVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


