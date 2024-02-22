# ModelsWorkflowProcessedCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaRowId** | Pointer to **NullableInt32** |  | [optional] [readonly] 
**CaRequestId** | Pointer to **NullableString** |  | [optional] 
**CaHost** | Pointer to **NullableString** |  | [optional] 
**CaLogicalName** | Pointer to **NullableString** |  | [optional] 
**KeyfactorRequestId** | Pointer to **int32** |  | [optional] 
**Comment** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewModelsWorkflowProcessedCertificateRequest

`func NewModelsWorkflowProcessedCertificateRequest() *ModelsWorkflowProcessedCertificateRequest`

NewModelsWorkflowProcessedCertificateRequest instantiates a new ModelsWorkflowProcessedCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsWorkflowProcessedCertificateRequestWithDefaults

`func NewModelsWorkflowProcessedCertificateRequestWithDefaults() *ModelsWorkflowProcessedCertificateRequest`

NewModelsWorkflowProcessedCertificateRequestWithDefaults instantiates a new ModelsWorkflowProcessedCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaRowId

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaRowId() int32`

GetCaRowId returns the CaRowId field if non-nil, zero value otherwise.

### GetCaRowIdOk

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaRowIdOk() (*int32, bool)`

GetCaRowIdOk returns a tuple with the CaRowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRowId

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaRowId(v int32)`

SetCaRowId sets CaRowId field to given value.

### HasCaRowId

`func (o *ModelsWorkflowProcessedCertificateRequest) HasCaRowId() bool`

HasCaRowId returns a boolean if a field has been set.

### SetCaRowIdNil

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaRowIdNil(b bool)`

 SetCaRowIdNil sets the value for CaRowId to be an explicit nil

### UnsetCaRowId
`func (o *ModelsWorkflowProcessedCertificateRequest) UnsetCaRowId()`

UnsetCaRowId ensures that no value is present for CaRowId, not even an explicit nil
### GetCaRequestId

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaRequestId() string`

GetCaRequestId returns the CaRequestId field if non-nil, zero value otherwise.

### GetCaRequestIdOk

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaRequestIdOk() (*string, bool)`

GetCaRequestIdOk returns a tuple with the CaRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaRequestId

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaRequestId(v string)`

SetCaRequestId sets CaRequestId field to given value.

### HasCaRequestId

`func (o *ModelsWorkflowProcessedCertificateRequest) HasCaRequestId() bool`

HasCaRequestId returns a boolean if a field has been set.

### SetCaRequestIdNil

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaRequestIdNil(b bool)`

 SetCaRequestIdNil sets the value for CaRequestId to be an explicit nil

### UnsetCaRequestId
`func (o *ModelsWorkflowProcessedCertificateRequest) UnsetCaRequestId()`

UnsetCaRequestId ensures that no value is present for CaRequestId, not even an explicit nil
### GetCaHost

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaHost() string`

GetCaHost returns the CaHost field if non-nil, zero value otherwise.

### GetCaHostOk

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaHostOk() (*string, bool)`

GetCaHostOk returns a tuple with the CaHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaHost

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaHost(v string)`

SetCaHost sets CaHost field to given value.

### HasCaHost

`func (o *ModelsWorkflowProcessedCertificateRequest) HasCaHost() bool`

HasCaHost returns a boolean if a field has been set.

### SetCaHostNil

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaHostNil(b bool)`

 SetCaHostNil sets the value for CaHost to be an explicit nil

### UnsetCaHost
`func (o *ModelsWorkflowProcessedCertificateRequest) UnsetCaHost()`

UnsetCaHost ensures that no value is present for CaHost, not even an explicit nil
### GetCaLogicalName

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaLogicalName() string`

GetCaLogicalName returns the CaLogicalName field if non-nil, zero value otherwise.

### GetCaLogicalNameOk

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCaLogicalNameOk() (*string, bool)`

GetCaLogicalNameOk returns a tuple with the CaLogicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaLogicalName

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaLogicalName(v string)`

SetCaLogicalName sets CaLogicalName field to given value.

### HasCaLogicalName

`func (o *ModelsWorkflowProcessedCertificateRequest) HasCaLogicalName() bool`

HasCaLogicalName returns a boolean if a field has been set.

### SetCaLogicalNameNil

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCaLogicalNameNil(b bool)`

 SetCaLogicalNameNil sets the value for CaLogicalName to be an explicit nil

### UnsetCaLogicalName
`func (o *ModelsWorkflowProcessedCertificateRequest) UnsetCaLogicalName()`

UnsetCaLogicalName ensures that no value is present for CaLogicalName, not even an explicit nil
### GetKeyfactorRequestId

`func (o *ModelsWorkflowProcessedCertificateRequest) GetKeyfactorRequestId() int32`

GetKeyfactorRequestId returns the KeyfactorRequestId field if non-nil, zero value otherwise.

### GetKeyfactorRequestIdOk

`func (o *ModelsWorkflowProcessedCertificateRequest) GetKeyfactorRequestIdOk() (*int32, bool)`

GetKeyfactorRequestIdOk returns a tuple with the KeyfactorRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorRequestId

`func (o *ModelsWorkflowProcessedCertificateRequest) SetKeyfactorRequestId(v int32)`

SetKeyfactorRequestId sets KeyfactorRequestId field to given value.

### HasKeyfactorRequestId

`func (o *ModelsWorkflowProcessedCertificateRequest) HasKeyfactorRequestId() bool`

HasKeyfactorRequestId returns a boolean if a field has been set.

### GetComment

`func (o *ModelsWorkflowProcessedCertificateRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *ModelsWorkflowProcessedCertificateRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *ModelsWorkflowProcessedCertificateRequest) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *ModelsWorkflowProcessedCertificateRequest) HasComment() bool`

HasComment returns a boolean if a field has been set.

### SetCommentNil

`func (o *ModelsWorkflowProcessedCertificateRequest) SetCommentNil(b bool)`

 SetCommentNil sets the value for Comment to be an explicit nil

### UnsetComment
`func (o *ModelsWorkflowProcessedCertificateRequest) UnsetComment()`

UnsetComment ensures that no value is present for Comment, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


