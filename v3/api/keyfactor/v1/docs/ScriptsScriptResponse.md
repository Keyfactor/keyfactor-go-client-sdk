# ScriptsScriptResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Contents** | Pointer to **NullableString** |  | [optional] 
**Categories** | Pointer to [**[]CSSCMSDataModelEnumsScriptCategories**](CSSCMSDataModelEnumsScriptCategories.md) |  | [optional] 

## Methods

### NewScriptsScriptResponse

`func NewScriptsScriptResponse() *ScriptsScriptResponse`

NewScriptsScriptResponse instantiates a new ScriptsScriptResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScriptsScriptResponseWithDefaults

`func NewScriptsScriptResponseWithDefaults() *ScriptsScriptResponse`

NewScriptsScriptResponseWithDefaults instantiates a new ScriptsScriptResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScriptsScriptResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScriptsScriptResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScriptsScriptResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ScriptsScriptResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ScriptsScriptResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScriptsScriptResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScriptsScriptResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ScriptsScriptResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *ScriptsScriptResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *ScriptsScriptResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetContents

`func (o *ScriptsScriptResponse) GetContents() string`

GetContents returns the Contents field if non-nil, zero value otherwise.

### GetContentsOk

`func (o *ScriptsScriptResponse) GetContentsOk() (*string, bool)`

GetContentsOk returns a tuple with the Contents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContents

`func (o *ScriptsScriptResponse) SetContents(v string)`

SetContents sets Contents field to given value.

### HasContents

`func (o *ScriptsScriptResponse) HasContents() bool`

HasContents returns a boolean if a field has been set.

### SetContentsNil

`func (o *ScriptsScriptResponse) SetContentsNil(b bool)`

 SetContentsNil sets the value for Contents to be an explicit nil

### UnsetContents
`func (o *ScriptsScriptResponse) UnsetContents()`

UnsetContents ensures that no value is present for Contents, not even an explicit nil
### GetCategories

`func (o *ScriptsScriptResponse) GetCategories() []CSSCMSDataModelEnumsScriptCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ScriptsScriptResponse) GetCategoriesOk() (*[]CSSCMSDataModelEnumsScriptCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ScriptsScriptResponse) SetCategories(v []CSSCMSDataModelEnumsScriptCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ScriptsScriptResponse) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### SetCategoriesNil

`func (o *ScriptsScriptResponse) SetCategoriesNil(b bool)`

 SetCategoriesNil sets the value for Categories to be an explicit nil

### UnsetCategories
`func (o *ScriptsScriptResponse) UnsetCategories()`

UnsetCategories ensures that no value is present for Categories, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


