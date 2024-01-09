# CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobTypeName** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**JobTypeFields** | Pointer to [**[]CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest**](CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest

`func NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest(jobTypeName string, ) *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest`

NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest instantiates a new CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequestWithDefaults

`func NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequestWithDefaults() *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest`

NewCSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequestWithDefaults instantiates a new CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetDescription

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJobTypeFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeFields() []CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest`

GetJobTypeFields returns the JobTypeFields field if non-nil, zero value otherwise.

### GetJobTypeFieldsOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) GetJobTypeFieldsOk() (*[]CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest, bool)`

GetJobTypeFieldsOk returns a tuple with the JobTypeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) SetJobTypeFields(v []CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest)`

SetJobTypeFields sets JobTypeFields field to given value.

### HasJobTypeFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) HasJobTypeFields() bool`

HasJobTypeFields returns a boolean if a field has been set.

### SetJobTypeFieldsNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) SetJobTypeFieldsNil(b bool)`

 SetJobTypeFieldsNil sets the value for JobTypeFields to be an explicit nil

### UnsetJobTypeFields
`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeCreateRequest) UnsetJobTypeFields()`

UnsetJobTypeFields ensures that no value is present for JobTypeFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


