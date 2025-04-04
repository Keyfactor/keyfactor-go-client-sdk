# ScriptsScriptCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Contents** | **string** |  | 
**Categories** | Pointer to [**[]CSSCMSDataModelEnumsScriptCategories**](CSSCMSDataModelEnumsScriptCategories.md) |  | [optional] 

## Methods

### NewScriptsScriptCreateRequest

`func NewScriptsScriptCreateRequest(name string, contents string, ) *ScriptsScriptCreateRequest`

NewScriptsScriptCreateRequest instantiates a new ScriptsScriptCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptsScriptCreateRequestWithDefaults

`func NewScriptsScriptCreateRequestWithDefaults() *ScriptsScriptCreateRequest`

NewScriptsScriptCreateRequestWithDefaults instantiates a new ScriptsScriptCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ScriptsScriptCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptsScriptCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptsScriptCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetContents

`func (o *ScriptsScriptCreateRequest) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ScriptsScriptCreateRequest) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ScriptsScriptCreateRequest) SetContents(v string)`

SetContents sets Contents field to given value.


### GetCategories

`func (o *ScriptsScriptCreateRequest) GetCategories() []CSSCMSDataModelEnumsScriptCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ScriptsScriptCreateRequest) GetCategoriesOk() (*[]CSSCMSDataModelEnumsScriptCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ScriptsScriptCreateRequest) SetCategories(v []CSSCMSDataModelEnumsScriptCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ScriptsScriptCreateRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ScriptsScriptCreateRequest) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ScriptsScriptCreateRequest) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


