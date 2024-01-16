# KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectPart** | **string** | The subject part to apply the regular expression to. | 
**Regex** | Pointer to **string** | The regular expression to apply to the subject part. | [optional] 
**Error** | Pointer to **string** | The error message to show when the regex validation fails. | [optional] 

## Methods

### NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequest

`func NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequest(subjectPart string, ) *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest`

NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequest instantiates a new KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequestWithDefaults

`func NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequestWithDefaults() *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest`

NewKeyfactorApiModelsTemplatesGlobalTemplateRegexRequestWithDefaults instantiates a new KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectPart

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) GetSubjectPart() string`

GetSubjectPart returns the SubjectPart field if non-nil, zero value otherwise.

### GetSubjectPartOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) GetSubjectPartOk() (*string, bool)`

GetSubjectPartOk returns a tuple with the SubjectPart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectPart

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) SetSubjectPart(v string)`

SetSubjectPart sets SubjectPart field to given value.


### GetRegex

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetError

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


