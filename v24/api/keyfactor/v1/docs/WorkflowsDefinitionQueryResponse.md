# WorkflowsDefinitionQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**Key** | Pointer to **NullableString** |  | [optional] 
**KeyDisplayName** | Pointer to **NullableString** |  | [optional] 
**WorkflowType** | Pointer to **NullableString** |  | [optional] 
**DraftVersion** | Pointer to **int32** |  | [optional] 
**PublishedVersion** | Pointer to **NullableInt32** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewWorkflowsDefinitionQueryResponse

`func NewWorkflowsDefinitionQueryResponse() *WorkflowsDefinitionQueryResponse`

NewWorkflowsDefinitionQueryResponse instantiates a new WorkflowsDefinitionQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowsDefinitionQueryResponseWithDefaults

`func NewWorkflowsDefinitionQueryResponseWithDefaults() *WorkflowsDefinitionQueryResponse`

NewWorkflowsDefinitionQueryResponseWithDefaults instantiates a new WorkflowsDefinitionQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowsDefinitionQueryResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowsDefinitionQueryResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowsDefinitionQueryResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowsDefinitionQueryResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayName

`func (o *WorkflowsDefinitionQueryResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *WorkflowsDefinitionQueryResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *WorkflowsDefinitionQueryResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *WorkflowsDefinitionQueryResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *WorkflowsDefinitionQueryResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *WorkflowsDefinitionQueryResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetKey

`func (o *WorkflowsDefinitionQueryResponse) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *WorkflowsDefinitionQueryResponse) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *WorkflowsDefinitionQueryResponse) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *WorkflowsDefinitionQueryResponse) HasKey() bool`

HasKey returns a boolean if a field has been set.

### SetKeyNil

`func (o *WorkflowsDefinitionQueryResponse) SetKeyNil(b bool)`

 SetKeyNil sets the value for Key to be an explicit nil

### UnsetKey
`func (o *WorkflowsDefinitionQueryResponse) UnsetKey()`

UnsetKey ensures that no value is present for Key, not even an explicit nil
### GetKeyDisplayName

`func (o *WorkflowsDefinitionQueryResponse) GetKeyDisplayName() string`

GetKeyDisplayName returns the KeyDisplayName field if non-nil, zero value otherwise.

### GetKeyDisplayNameOk

`func (o *WorkflowsDefinitionQueryResponse) GetKeyDisplayNameOk() (*string, bool)`

GetKeyDisplayNameOk returns a tuple with the KeyDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyDisplayName

`func (o *WorkflowsDefinitionQueryResponse) SetKeyDisplayName(v string)`

SetKeyDisplayName sets KeyDisplayName field to given value.

### HasKeyDisplayName

`func (o *WorkflowsDefinitionQueryResponse) HasKeyDisplayName() bool`

HasKeyDisplayName returns a boolean if a field has been set.

### SetKeyDisplayNameNil

`func (o *WorkflowsDefinitionQueryResponse) SetKeyDisplayNameNil(b bool)`

 SetKeyDisplayNameNil sets the value for KeyDisplayName to be an explicit nil

### UnsetKeyDisplayName
`func (o *WorkflowsDefinitionQueryResponse) UnsetKeyDisplayName()`

UnsetKeyDisplayName ensures that no value is present for KeyDisplayName, not even an explicit nil
### GetWorkflowType

`func (o *WorkflowsDefinitionQueryResponse) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *WorkflowsDefinitionQueryResponse) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *WorkflowsDefinitionQueryResponse) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *WorkflowsDefinitionQueryResponse) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.

### SetWorkflowTypeNil

`func (o *WorkflowsDefinitionQueryResponse) SetWorkflowTypeNil(b bool)`

 SetWorkflowTypeNil sets the value for WorkflowType to be an explicit nil

### UnsetWorkflowType
`func (o *WorkflowsDefinitionQueryResponse) UnsetWorkflowType()`

UnsetWorkflowType ensures that no value is present for WorkflowType, not even an explicit nil
### GetDraftVersion

`func (o *WorkflowsDefinitionQueryResponse) GetDraftVersion() int32`

GetDraftVersion returns the DraftVersion field if non-nil, zero value otherwise.

### GetDraftVersionOk

`func (o *WorkflowsDefinitionQueryResponse) GetDraftVersionOk() (*int32, bool)`

GetDraftVersionOk returns a tuple with the DraftVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraftVersion

`func (o *WorkflowsDefinitionQueryResponse) SetDraftVersion(v int32)`

SetDraftVersion sets DraftVersion field to given value.

### HasDraftVersion

`func (o *WorkflowsDefinitionQueryResponse) HasDraftVersion() bool`

HasDraftVersion returns a boolean if a field has been set.

### GetPublishedVersion

`func (o *WorkflowsDefinitionQueryResponse) GetPublishedVersion() int32`

GetPublishedVersion returns the PublishedVersion field if non-nil, zero value otherwise.

### GetPublishedVersionOk

`func (o *WorkflowsDefinitionQueryResponse) GetPublishedVersionOk() (*int32, bool)`

GetPublishedVersionOk returns a tuple with the PublishedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedVersion

`func (o *WorkflowsDefinitionQueryResponse) SetPublishedVersion(v int32)`

SetPublishedVersion sets PublishedVersion field to given value.

### HasPublishedVersion

`func (o *WorkflowsDefinitionQueryResponse) HasPublishedVersion() bool`

HasPublishedVersion returns a boolean if a field has been set.

### SetPublishedVersionNil

`func (o *WorkflowsDefinitionQueryResponse) SetPublishedVersionNil(b bool)`

 SetPublishedVersionNil sets the value for PublishedVersion to be an explicit nil

### UnsetPublishedVersion
`func (o *WorkflowsDefinitionQueryResponse) UnsetPublishedVersion()`

UnsetPublishedVersion ensures that no value is present for PublishedVersion, not even an explicit nil
### GetEnabled

`func (o *WorkflowsDefinitionQueryResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *WorkflowsDefinitionQueryResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *WorkflowsDefinitionQueryResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *WorkflowsDefinitionQueryResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


