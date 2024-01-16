# ModelsQueryModelsPagedAgentBlueprintJobsQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageReturned** | Pointer to **int64** | The current page within the result set to be returned | [optional] 
**SkipCount** | Pointer to **int64** | Number of records to be skipped before inclusion in the result set | [optional] [readonly] 
**ReturnLimit** | Pointer to **int64** | Maximum number of records to be returned in a single call | [optional] 
**SortField** | Pointer to **string** | Field by which the results should be sorted (OperationStart, OperationEnd, UserName) | [optional] 
**SortAscending** | Pointer to **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | [optional] 

## Methods

### NewModelsQueryModelsPagedAgentBlueprintJobsQuery

`func NewModelsQueryModelsPagedAgentBlueprintJobsQuery() *ModelsQueryModelsPagedAgentBlueprintJobsQuery`

NewModelsQueryModelsPagedAgentBlueprintJobsQuery instantiates a new ModelsQueryModelsPagedAgentBlueprintJobsQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsQueryModelsPagedAgentBlueprintJobsQueryWithDefaults

`func NewModelsQueryModelsPagedAgentBlueprintJobsQueryWithDefaults() *ModelsQueryModelsPagedAgentBlueprintJobsQuery`

NewModelsQueryModelsPagedAgentBlueprintJobsQueryWithDefaults instantiates a new ModelsQueryModelsPagedAgentBlueprintJobsQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPageReturned

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetPageReturned() int64`

GetPageReturned returns the PageReturned field if non-nil, zero value otherwise.

### GetPageReturnedOk

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetPageReturnedOk() (*int64, bool)`

GetPageReturnedOk returns a tuple with the PageReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReturned

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetPageReturned(v int64)`

SetPageReturned sets PageReturned field to given value.

### HasPageReturned

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasPageReturned() bool`

HasPageReturned returns a boolean if a field has been set.

### GetSkipCount

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSkipCount() int64`

GetSkipCount returns the SkipCount field if non-nil, zero value otherwise.

### GetSkipCountOk

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSkipCountOk() (*int64, bool)`

GetSkipCountOk returns a tuple with the SkipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipCount

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetSkipCount(v int64)`

SetSkipCount sets SkipCount field to given value.

### HasSkipCount

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasSkipCount() bool`

HasSkipCount returns a boolean if a field has been set.

### GetReturnLimit

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetReturnLimit() int64`

GetReturnLimit returns the ReturnLimit field if non-nil, zero value otherwise.

### GetReturnLimitOk

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetReturnLimitOk() (*int64, bool)`

GetReturnLimitOk returns a tuple with the ReturnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLimit

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetReturnLimit(v int64)`

SetReturnLimit sets ReturnLimit field to given value.

### HasReturnLimit

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasReturnLimit() bool`

HasReturnLimit returns a boolean if a field has been set.

### GetSortField

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortAscending

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortAscending() int64`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) GetSortAscendingOk() (*int64, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) SetSortAscending(v int64)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *ModelsQueryModelsPagedAgentBlueprintJobsQuery) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


