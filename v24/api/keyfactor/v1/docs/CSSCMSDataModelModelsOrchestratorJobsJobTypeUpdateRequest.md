# CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**JobTypeName** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**JobTypeFields** | Pointer to [**[]CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest**](CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest.md) |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest

`func NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest(id string, jobTypeName string, ) *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest`

NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest instantiates a new CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults

`func NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults() *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest`

NewCSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequestWithDefaults instantiates a new CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.


### GetDescription

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJobTypeFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeFields() []CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest`

GetJobTypeFields returns the JobTypeFields field if non-nil, zero value otherwise.

### GetJobTypeFieldsOk

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) GetJobTypeFieldsOk() (*[]CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest, bool)`

GetJobTypeFieldsOk returns a tuple with the JobTypeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeFields(v []CSSCMSDataModelModelsOrchestratorJobsJobTypeFieldRequest)`

SetJobTypeFields sets JobTypeFields field to given value.

### HasJobTypeFields

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) HasJobTypeFields() bool`

HasJobTypeFields returns a boolean if a field has been set.

### SetJobTypeFieldsNil

`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) SetJobTypeFieldsNil(b bool)`

 SetJobTypeFieldsNil sets the value for JobTypeFields to be an explicit nil

### UnsetJobTypeFields
`func (o *CSSCMSDataModelModelsOrchestratorJobsJobTypeUpdateRequest) UnsetJobTypeFields()`

UnsetJobTypeFields ensures that no value is present for JobTypeFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


