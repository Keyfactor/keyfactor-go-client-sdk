# WorkflowsV1DefinitionV1Response

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
**Steps** | Pointer to [**[]WorkflowsV1DefinitionStepV1Response**](WorkflowsV1DefinitionStepV1Response.md) |  | [optional] 
**DraftVersion** | Pointer to **int32** |  | [optional] 
**PublishedVersion** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkflowsV1DefinitionV1Response

`func NewWorkflowsV1DefinitionV1Response() *WorkflowsV1DefinitionV1Response`

NewWorkflowsV1DefinitionV1Response instantiates a new WorkflowsV1DefinitionV1Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsV1DefinitionV1ResponseWithDefaults

`func NewWorkflowsV1DefinitionV1ResponseWithDefaults() *WorkflowsV1DefinitionV1Response`

NewWorkflowsV1DefinitionV1ResponseWithDefaults instantiates a new WorkflowsV1DefinitionV1Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsV1DefinitionV1Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsV1DefinitionV1Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsV1DefinitionV1Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsV1DefinitionV1Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *WorkflowsV1DefinitionV1Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsV1DefinitionV1Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsV1DefinitionV1Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsV1DefinitionV1Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsV1DefinitionV1Response) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsV1DefinitionV1Response) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetDescription

`func (o *WorkflowsV1DefinitionV1Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowsV1DefinitionV1Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowsV1DefinitionV1Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowsV1DefinitionV1Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowsV1DefinitionV1Response) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowsV1DefinitionV1Response) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetKey

`func (o *WorkflowsV1DefinitionV1Response) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowsV1DefinitionV1Response) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowsV1DefinitionV1Response) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkflowsV1DefinitionV1Response) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *WorkflowsV1DefinitionV1Response) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *WorkflowsV1DefinitionV1Response) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyDisplayName

`func (o *WorkflowsV1DefinitionV1Response) GetKeyDisplayName() string`

GetKeyDisplayName returns the KeyDisplayName field if non-nil, zero value otherwise.

### GetKeyDisplayNameOk

`func (o *WorkflowsV1DefinitionV1Response) GetKeyDisplayNameOk() (*string, bool)`

GetKeyDisplayNameOk returns a tuple with the KeyDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDisplayName

`func (o *WorkflowsV1DefinitionV1Response) SetKeyDisplayName(v string)`

SetKeyDisplayName sets KeyDisplayName field to given value.

### HasKeyDisplayName

`func (o *WorkflowsV1DefinitionV1Response) HasKeyDisplayName() bool`

HasKeyDisplayName returns a boolean if a field has been set.

### SetKeyDisplayNameNil

`func (o *WorkflowsV1DefinitionV1Response) SetKeyDisplayNameNil(b bool)`

 SetKeyDisplayNameNil sets the value for KeyDisplayName to be an explicit nil

### UnsetKeyDisplayName
`func (o *WorkflowsV1DefinitionV1Response) UnsetKeyDisplayName()`

UnsetKeyDisplayName ensures that no value is present for KeyDisplayName, not even an explicit nil
### GetIsPublished

`func (o *WorkflowsV1DefinitionV1Response) GetIsPublished() bool`

GetIsPublished returns the IsPublished field if non-nil, zero value otherwise.

### GetIsPublishedOk

`func (o *WorkflowsV1DefinitionV1Response) GetIsPublishedOk() (*bool, bool)`

GetIsPublishedOk returns a tuple with the IsPublished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublished

`func (o *WorkflowsV1DefinitionV1Response) SetIsPublished(v bool)`

SetIsPublished sets IsPublished field to given value.

### HasIsPublished

`func (o *WorkflowsV1DefinitionV1Response) HasIsPublished() bool`

HasIsPublished returns a boolean if a field has been set.

### GetWorkflowType

`func (o *WorkflowsV1DefinitionV1Response) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowsV1DefinitionV1Response) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowsV1DefinitionV1Response) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowsV1DefinitionV1Response) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### SetWorkflowTypeNil

`func (o *WorkflowsV1DefinitionV1Response) SetWorkflowTypeNil(b bool)`

 SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil

### UnsetWorkflowType
`func (o *WorkflowsV1DefinitionV1Response) UnsetWorkflowType()`

UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil
### GetSteps

`func (o *WorkflowsV1DefinitionV1Response) GetSteps() []WorkflowsV1DefinitionStepV1Response`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *WorkflowsV1DefinitionV1Response) GetStepsOk() (*[]WorkflowsV1DefinitionStepV1Response, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *WorkflowsV1DefinitionV1Response) SetSteps(v []WorkflowsV1DefinitionStepV1Response)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *WorkflowsV1DefinitionV1Response) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### SetStepsNil

`func (o *WorkflowsV1DefinitionV1Response) SetStepsNil(b bool)`

 SetStepsNil sets the value for Steps to be an explicit nil

### UnsetSteps
`func (o *WorkflowsV1DefinitionV1Response) UnsetSteps()`

UnsetSteps ensures that no value is present for Steps, not even an explicit nil
### GetDraftVersion

`func (o *WorkflowsV1DefinitionV1Response) GetDraftVersion() int32`

GetDraftVersion returns the DraftVersion field if non-nil, zero value otherwise.

### GetDraftVersionOk

`func (o *WorkflowsV1DefinitionV1Response) GetDraftVersionOk() (*int32, bool)`

GetDraftVersionOk returns a tuple with the DraftVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftVersion

`func (o *WorkflowsV1DefinitionV1Response) SetDraftVersion(v int32)`

SetDraftVersion sets DraftVersion field to given value.

### HasDraftVersion

`func (o *WorkflowsV1DefinitionV1Response) HasDraftVersion() bool`

HasDraftVersion returns a boolean if a field has been set.

### GetPublishedVersion

`func (o *WorkflowsV1DefinitionV1Response) GetPublishedVersion() int32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *WorkflowsV1DefinitionV1Response) GetPublishedVersionOk() (*int32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *WorkflowsV1DefinitionV1Response) SetPublishedVersion(v int32)`

SetPublishedVersion sets PublishedVersion field to given value.

### HasPublishedVersion

`func (o *WorkflowsV1DefinitionV1Response) HasPublishedVersion() bool`

HasPublishedVersion returns a boolean if a field has been set.

### SetPublishedVersionNil

`func (o *WorkflowsV1DefinitionV1Response) SetPublishedVersionNil(b bool)`

 SetPublishedVersionNil sets the value for PublishedVersion to be an explicit nil

### UnsetPublishedVersion
`func (o *WorkflowsV1DefinitionV1Response) UnsetPublishedVersion()`

UnsetPublishedVersion ensures that no value is present for PublishedVersion, not even an explicit nil
### GetEnabled

`func (o *WorkflowsV1DefinitionV1Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsV1DefinitionV1Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsV1DefinitionV1Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsV1DefinitionV1Response) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


