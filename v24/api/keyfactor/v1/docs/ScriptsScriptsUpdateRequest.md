# ScriptsScriptsUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Contents** | Pointer to **NullableString** |  | [optional] 
**Categories** | Pointer to [**[]CSSCMSDataModelEnumsScriptCategories**](CSSCMSDataModelEnumsScriptCategories.md) |  | [optional] 

## Methods

### NewScriptsScriptsUpdateRequest

`func NewScriptsScriptsUpdateRequest() *ScriptsScriptsUpdateRequest`

NewScriptsScriptsUpdateRequest instantiates a new ScriptsScriptsUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptsScriptsUpdateRequestWithDefaults

`func NewScriptsScriptsUpdateRequestWithDefaults() *ScriptsScriptsUpdateRequest`

NewScriptsScriptsUpdateRequestWithDefaults instantiates a new ScriptsScriptsUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScriptsScriptsUpdateRequest) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptsScriptsUpdateRequest) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptsScriptsUpdateRequest) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScriptsScriptsUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetContents

`func (o *ScriptsScriptsUpdateRequest) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ScriptsScriptsUpdateRequest) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ScriptsScriptsUpdateRequest) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ScriptsScriptsUpdateRequest) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *ScriptsScriptsUpdateRequest) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *ScriptsScriptsUpdateRequest) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCategories

`func (o *ScriptsScriptsUpdateRequest) GetCategories() []CSSCMSDataModelEnumsScriptCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ScriptsScriptsUpdateRequest) GetCategoriesOk() (*[]CSSCMSDataModelEnumsScriptCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ScriptsScriptsUpdateRequest) SetCategories(v []CSSCMSDataModelEnumsScriptCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ScriptsScriptsUpdateRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ScriptsScriptsUpdateRequest) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ScriptsScriptsUpdateRequest) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


