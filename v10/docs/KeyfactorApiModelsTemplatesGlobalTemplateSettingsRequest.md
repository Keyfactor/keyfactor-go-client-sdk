# KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateRegexes** | [**[]KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest**](KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest.md) | The regular expressions to use for validation during enrollment. | 
**TemplateDefaults** | [**[]KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest**](KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest.md) | The default values to use during enrollment. | 
**TemplatePolicy** | [**KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest**](KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest.md) |  | 

## Methods

### NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest

`func NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest(templateRegexes []KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest, templateDefaults []KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest, templatePolicy KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest, ) *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest`

NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest instantiates a new KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequestWithDefaults

`func NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequestWithDefaults() *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest`

NewKeyfactorApiModelsTemplatesGlobalTemplateSettingsRequestWithDefaults instantiates a new KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateRegexes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) GetTemplateRegexes() []KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) GetTemplateRegexesOk() (*[]KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) SetTemplateRegexes(v []KeyfactorApiModelsTemplatesGlobalTemplateRegexRequest)`

SetTemplateRegexes sets TemplateRegexes field to given value.


### GetTemplateDefaults

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) GetTemplateDefaults() []KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) GetTemplateDefaultsOk() (*[]KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) SetTemplateDefaults(v []KeyfactorApiModelsTemplatesGlobalTemplateDefaultRequest)`

SetTemplateDefaults sets TemplateDefaults field to given value.


### GetTemplatePolicy

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) GetTemplatePolicy() KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) GetTemplatePolicyOk() (*KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *KeyfactorApiModelsTemplatesGlobalTemplateSettingsRequest) SetTemplatePolicy(v KeyfactorApiModelsTemplatesGlobalTemplatePolicyRequest)`

SetTemplatePolicy sets TemplatePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


