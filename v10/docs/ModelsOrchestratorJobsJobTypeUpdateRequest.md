# ModelsOrchestratorJobsJobTypeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobTypeName** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**JobTypeFields** | Pointer to [**[]ModelsOrchestratorJobsJobTypeFieldRequest**](ModelsOrchestratorJobsJobTypeFieldRequest.md) |  | [optional] 

## Methods

### NewModelsOrchestratorJobsJobTypeUpdateRequest

`func NewModelsOrchestratorJobsJobTypeUpdateRequest(id string, jobTypeName string, ) *ModelsOrchestratorJobsJobTypeUpdateRequest`

NewModelsOrchestratorJobsJobTypeUpdateRequest instantiates a new ModelsOrchestratorJobsJobTypeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults

`func NewModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults() *ModelsOrchestratorJobsJobTypeUpdateRequest`

NewModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults instantiates a new ModelsOrchestratorJobsJobTypeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetJobTypeName

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetDescription

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJobTypeFields

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeFields() []ModelsOrchestratorJobsJobTypeFieldRequest`

GetJobTypeFields returns the JobTypeFields field if non-nil, zero value otherwise.

### GetJobTypeFieldsOk

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeFieldsOk() (*[]ModelsOrchestratorJobsJobTypeFieldRequest, bool)`

GetJobTypeFieldsOk returns a tuple with the JobTypeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeFields

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeFields(v []ModelsOrchestratorJobsJobTypeFieldRequest)`

SetJobTypeFields sets JobTypeFields field to given value.

### HasJobTypeFields

`func (o *ModelsOrchestratorJobsJobTypeUpdateRequest) HasJobTypeFields() bool`

HasJobTypeFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


