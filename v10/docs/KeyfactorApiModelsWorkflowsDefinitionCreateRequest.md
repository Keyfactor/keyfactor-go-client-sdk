# KeyfactorApiModelsWorkflowsDefinitionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayName** | Pointer to **string** | Display name of the Definition | [optional] 
**Description** | Pointer to **string** | Description of the Definition | [optional] 
**Key** | Pointer to **string** | Key to be used to look up definition when starting a new workflow.  For enrollment workflowTypes, this should be a template | [optional] 
**WorkflowType** | Pointer to **string** | The Type of Workflow | [optional] 

## Methods

### NewKeyfactorApiModelsWorkflowsDefinitionCreateRequest

`func NewKeyfactorApiModelsWorkflowsDefinitionCreateRequest() *KeyfactorApiModelsWorkflowsDefinitionCreateRequest`

NewKeyfactorApiModelsWorkflowsDefinitionCreateRequest instantiates a new KeyfactorApiModelsWorkflowsDefinitionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsWorkflowsDefinitionCreateRequestWithDefaults

`func NewKeyfactorApiModelsWorkflowsDefinitionCreateRequestWithDefaults() *KeyfactorApiModelsWorkflowsDefinitionCreateRequest`

NewKeyfactorApiModelsWorkflowsDefinitionCreateRequestWithDefaults instantiates a new KeyfactorApiModelsWorkflowsDefinitionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetDescription

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKey

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetWorkflowType

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetWorkflowType() string`

GetWorkflowType returns the WorkflowType field if non-nil, zero value otherwise.

### GetWorkflowTypeOk

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) GetWorkflowTypeOk() (*string, bool)`

GetWorkflowTypeOk returns a tuple with the WorkflowType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowType

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) SetWorkflowType(v string)`

SetWorkflowType sets WorkflowType field to given value.

### HasWorkflowType

`func (o *KeyfactorApiModelsWorkflowsDefinitionCreateRequest) HasWorkflowType() bool`

HasWorkflowType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


