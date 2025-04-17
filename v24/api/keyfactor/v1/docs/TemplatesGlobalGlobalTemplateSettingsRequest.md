# TemplatesGlobalGlobalTemplateSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateRegexes** | [**[]TemplatesGlobalGlobalTemplateRegexRequest**](TemplatesGlobalGlobalTemplateRegexRequest.md) | The regular expressions to use for validation during enrollment. | 
**TemplateDefaults** | [**[]TemplatesGlobalGlobalTemplateDefaultRequest**](TemplatesGlobalGlobalTemplateDefaultRequest.md) | The default values to use during enrollment. | 
**TemplatePolicy** | [**TemplatesGlobalGlobalTemplatePolicyRequest**](TemplatesGlobalGlobalTemplatePolicyRequest.md) |  | 

## Methods

### NewTemplatesGlobalGlobalTemplateSettingsRequest

`func NewTemplatesGlobalGlobalTemplateSettingsRequest(templateRegexes []TemplatesGlobalGlobalTemplateRegexRequest, templateDefaults []TemplatesGlobalGlobalTemplateDefaultRequest, templatePolicy TemplatesGlobalGlobalTemplatePolicyRequest, ) *TemplatesGlobalGlobalTemplateSettingsRequest`

NewTemplatesGlobalGlobalTemplateSettingsRequest instantiates a new TemplatesGlobalGlobalTemplateSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesGlobalGlobalTemplateSettingsRequestWithDefaults

`func NewTemplatesGlobalGlobalTemplateSettingsRequestWithDefaults() *TemplatesGlobalGlobalTemplateSettingsRequest`

NewTemplatesGlobalGlobalTemplateSettingsRequestWithDefaults instantiates a new TemplatesGlobalGlobalTemplateSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateRegexes

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) GetTemplateRegexes() []TemplatesGlobalGlobalTemplateRegexRequest`

GetTemplateRegexes returns the TemplateRegexes field if non-nil, zero value otherwise.

### GetTemplateRegexesOk

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) GetTemplateRegexesOk() (*[]TemplatesGlobalGlobalTemplateRegexRequest, bool)`

GetTemplateRegexesOk returns a tuple with the TemplateRegexes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateRegexes

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) SetTemplateRegexes(v []TemplatesGlobalGlobalTemplateRegexRequest)`

SetTemplateRegexes sets TemplateRegexes field to given value.


### GetTemplateDefaults

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) GetTemplateDefaults() []TemplatesGlobalGlobalTemplateDefaultRequest`

GetTemplateDefaults returns the TemplateDefaults field if non-nil, zero value otherwise.

### GetTemplateDefaultsOk

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) GetTemplateDefaultsOk() (*[]TemplatesGlobalGlobalTemplateDefaultRequest, bool)`

GetTemplateDefaultsOk returns a tuple with the TemplateDefaults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateDefaults

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) SetTemplateDefaults(v []TemplatesGlobalGlobalTemplateDefaultRequest)`

SetTemplateDefaults sets TemplateDefaults field to given value.


### GetTemplatePolicy

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) GetTemplatePolicy() TemplatesGlobalGlobalTemplatePolicyRequest`

GetTemplatePolicy returns the TemplatePolicy field if non-nil, zero value otherwise.

### GetTemplatePolicyOk

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) GetTemplatePolicyOk() (*TemplatesGlobalGlobalTemplatePolicyRequest, bool)`

GetTemplatePolicyOk returns a tuple with the TemplatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplatePolicy

`func (o *TemplatesGlobalGlobalTemplateSettingsRequest) SetTemplatePolicy(v TemplatesGlobalGlobalTemplatePolicyRequest)`

SetTemplatePolicy sets TemplatePolicy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


