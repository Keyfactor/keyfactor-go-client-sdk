# WorkflowsDefinitionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**KeyDisplayName** | Pointer to **NullableString** |  | [optional] 
**IsPublished** | Pointer to **bool** |  | [optional] 
**WorkflowType** | Pointer to **NullableString** |  | [optional] 
**Steps** | Pointer to [**[]WorkflowsDefinitionStepResponse**](WorkflowsDefinitionStepResponse.md) |  | [optional] 
**DraftVersion** | Pointer to **int32** |  | [optional] 
**PublishedVersion** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

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
### GetKey

`func (o *WorkflowsDefinitionResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowsDefinitionResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowsDefinitionResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkflowsDefinitionResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *WorkflowsDefinitionResponse) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *WorkflowsDefinitionResponse) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyDisplayName

`func (o *WorkflowsDefinitionResponse) GetKeyDisplayName() string`

GetKeyDisplayName returns the KeyDisplayName field if non-nil, zero value otherwise.

### GetKeyDisplayNameOk

`func (o *WorkflowsDefinitionResponse) GetKeyDisplayNameOk() (*string, bool)`

GetKeyDisplayNameOk returns a tuple with the KeyDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDisplayName

`func (o *WorkflowsDefinitionResponse) SetKeyDisplayName(v string)`

SetKeyDisplayName sets KeyDisplayName field to given value.

### HasKeyDisplayName

`func (o *WorkflowsDefinitionResponse) HasKeyDisplayName() bool`

HasKeyDisplayName returns a boolean if a field has been set.

### SetKeyDisplayNameNil

`func (o *WorkflowsDefinitionResponse) SetKeyDisplayNameNil(b bool)`

 SetKeyDisplayNameNil sets the value for KeyDisplayName to be an explicit nil

### UnsetKeyDisplayName
`func (o *WorkflowsDefinitionResponse) UnsetKeyDisplayName()`

UnsetKeyDisplayName ensures that no value is present for KeyDisplayName, not even an explicit nil
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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


