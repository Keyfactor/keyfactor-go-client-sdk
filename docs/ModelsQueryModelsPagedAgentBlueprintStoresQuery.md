# ModelsQueryModelsPagedAgentBlueprintStoresQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageReturned** | Pointer to **int32** | The current page within the result set to be returned | [optional] 
**SkipCount** | Pointer to **int32** | Number of records to be skipped before inclusion in the result set | [optional] [readonly] 
**ReturnLimit** | Pointer to **int32** | Maximum number of records to be returned in a single call | [optional] 
**SortField** | Pointer to **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | [optional] 
**SortAscending** | Pointer to **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | [optional] 

## Methods

### NewModelsQueryModelsPagedAgentBlueprintStoresQuery

`func NewModelsQueryModelsPagedAgentBlueprintStoresQuery() *ModelsQueryModelsPagedAgentBlueprintStoresQuery`

NewModelsQueryModelsPagedAgentBlueprintStoresQuery instantiates a new ModelsQueryModelsPagedAgentBlueprintStoresQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsQueryModelsPagedAgentBlueprintStoresQueryWithDefaults

`func NewModelsQueryModelsPagedAgentBlueprintStoresQueryWithDefaults() *ModelsQueryModelsPagedAgentBlueprintStoresQuery`

NewModelsQueryModelsPagedAgentBlueprintStoresQueryWithDefaults instantiates a new ModelsQueryModelsPagedAgentBlueprintStoresQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageReturned

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetPageReturned() int32`

GetPageReturned returns the PageReturned field if non-nil, zero value otherwise.

### GetPageReturnedOk

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetPageReturnedOk() (*int32, bool)`

GetPageReturnedOk returns a tuple with the PageReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReturned

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) SetPageReturned(v int32)`

SetPageReturned sets PageReturned field to given value.

### HasPageReturned

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) HasPageReturned() bool`

HasPageReturned returns a boolean if a field has been set.

### GetSkipCount

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetSkipCount() int32`

GetSkipCount returns the SkipCount field if non-nil, zero value otherwise.

### GetSkipCountOk

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetSkipCountOk() (*int32, bool)`

GetSkipCountOk returns a tuple with the SkipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCount

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) SetSkipCount(v int32)`

SetSkipCount sets SkipCount field to given value.

### HasSkipCount

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) HasSkipCount() bool`

HasSkipCount returns a boolean if a field has been set.

### GetReturnLimit

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetReturnLimit() int32`

GetReturnLimit returns the ReturnLimit field if non-nil, zero value otherwise.

### GetReturnLimitOk

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetReturnLimitOk() (*int32, bool)`

GetReturnLimitOk returns a tuple with the ReturnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLimit

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) SetReturnLimit(v int32)`

SetReturnLimit sets ReturnLimit field to given value.

### HasReturnLimit

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) HasReturnLimit() bool`

HasReturnLimit returns a boolean if a field has been set.

### GetSortField

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortAscending

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetSortAscending() int32`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) GetSortAscendingOk() (*int32, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) SetSortAscending(v int32)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *ModelsQueryModelsPagedAgentBlueprintStoresQuery) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


