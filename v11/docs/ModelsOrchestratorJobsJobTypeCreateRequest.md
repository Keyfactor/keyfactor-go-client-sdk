# ModelsOrchestratorJobsJobTypeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTypeName** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**JobTypeFields** | Pointer to [**[]ModelsOrchestratorJobsJobTypeFieldRequest**](ModelsOrchestratorJobsJobTypeFieldRequest.md) |  | [optional] 

## Methods

### NewModelsOrchestratorJobsJobTypeCreateRequest

`func NewModelsOrchestratorJobsJobTypeCreateRequest(jobTypeName string, ) *ModelsOrchestratorJobsJobTypeCreateRequest`

NewModelsOrchestratorJobsJobTypeCreateRequest instantiates a new ModelsOrchestratorJobsJobTypeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrchestratorJobsJobTypeCreateRequestWithDefaults

`func NewModelsOrchestratorJobsJobTypeCreateRequestWithDefaults() *ModelsOrchestratorJobsJobTypeCreateRequest`

NewModelsOrchestratorJobsJobTypeCreateRequestWithDefaults instantiates a new ModelsOrchestratorJobsJobTypeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTypeName

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetDescription

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJobTypeFields

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeFields() []ModelsOrchestratorJobsJobTypeFieldRequest`

GetJobTypeFields returns the JobTypeFields field if non-nil, zero value otherwise.

### GetJobTypeFieldsOk

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeFieldsOk() (*[]ModelsOrchestratorJobsJobTypeFieldRequest, bool)`

GetJobTypeFieldsOk returns a tuple with the JobTypeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeFields

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) SetJobTypeFields(v []ModelsOrchestratorJobsJobTypeFieldRequest)`

SetJobTypeFields sets JobTypeFields field to given value.

### HasJobTypeFields

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) HasJobTypeFields() bool`

HasJobTypeFields returns a boolean if a field has been set.

### SetJobTypeFieldsNil

`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) SetJobTypeFieldsNil(b bool)`

 SetJobTypeFieldsNil sets the value for JobTypeFields to be an explicit nil

### UnsetJobTypeFields
`func (o *ModelsOrchestratorJobsJobTypeCreateRequest) UnsetJobTypeFields()`

UnsetJobTypeFields ensures that no value is present for JobTypeFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


