# OrchestratorJobsJobTypeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** |  | [optional] 
**JobTypeName** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**JobTypeFields** | Pointer to [**[]OrchestratorJobsJobTypeFieldResponse**](OrchestratorJobsJobTypeFieldResponse.md) |  | [optional] 

## Methods

### NewOrchestratorJobsJobTypeResponse

`func NewOrchestratorJobsJobTypeResponse() *OrchestratorJobsJobTypeResponse`

NewOrchestratorJobsJobTypeResponse instantiates a new OrchestratorJobsJobTypeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrchestratorJobsJobTypeResponseWithDefaults

`func NewOrchestratorJobsJobTypeResponseWithDefaults() *OrchestratorJobsJobTypeResponse`

NewOrchestratorJobsJobTypeResponseWithDefaults instantiates a new OrchestratorJobsJobTypeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrchestratorJobsJobTypeResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrchestratorJobsJobTypeResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrchestratorJobsJobTypeResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrchestratorJobsJobTypeResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *OrchestratorJobsJobTypeResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *OrchestratorJobsJobTypeResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetJobTypeName

`func (o *OrchestratorJobsJobTypeResponse) GetJobTypeName() string`

GetJobTypeName returns the JobTypeName field if non-nil, zero value otherwise.

### GetJobTypeNameOk

`func (o *OrchestratorJobsJobTypeResponse) GetJobTypeNameOk() (*string, bool)`

GetJobTypeNameOk returns a tuple with the JobTypeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeName

`func (o *OrchestratorJobsJobTypeResponse) SetJobTypeName(v string)`

SetJobTypeName sets JobTypeName field to given value.

### HasJobTypeName

`func (o *OrchestratorJobsJobTypeResponse) HasJobTypeName() bool`

HasJobTypeName returns a boolean if a field has been set.

### SetJobTypeNameNil

`func (o *OrchestratorJobsJobTypeResponse) SetJobTypeNameNil(b bool)`

 SetJobTypeNameNil sets the value for JobTypeName to be an explicit nil

### UnsetJobTypeName
`func (o *OrchestratorJobsJobTypeResponse) UnsetJobTypeName()`

UnsetJobTypeName ensures that no value is present for JobTypeName, not even an explicit nil
### GetDescription

`func (o *OrchestratorJobsJobTypeResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OrchestratorJobsJobTypeResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OrchestratorJobsJobTypeResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OrchestratorJobsJobTypeResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OrchestratorJobsJobTypeResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OrchestratorJobsJobTypeResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetJobTypeFields

`func (o *OrchestratorJobsJobTypeResponse) GetJobTypeFields() []OrchestratorJobsJobTypeFieldResponse`

GetJobTypeFields returns the JobTypeFields field if non-nil, zero value otherwise.

### GetJobTypeFieldsOk

`func (o *OrchestratorJobsJobTypeResponse) GetJobTypeFieldsOk() (*[]OrchestratorJobsJobTypeFieldResponse, bool)`

GetJobTypeFieldsOk returns a tuple with the JobTypeFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobTypeFields

`func (o *OrchestratorJobsJobTypeResponse) SetJobTypeFields(v []OrchestratorJobsJobTypeFieldResponse)`

SetJobTypeFields sets JobTypeFields field to given value.

### HasJobTypeFields

`func (o *OrchestratorJobsJobTypeResponse) HasJobTypeFields() bool`

HasJobTypeFields returns a boolean if a field has been set.

### SetJobTypeFieldsNil

`func (o *OrchestratorJobsJobTypeResponse) SetJobTypeFieldsNil(b bool)`

 SetJobTypeFieldsNil sets the value for JobTypeFields to be an explicit nil

### UnsetJobTypeFields
`func (o *OrchestratorJobsJobTypeResponse) UnsetJobTypeFields()`

UnsetJobTypeFields ensures that no value is present for JobTypeFields, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


