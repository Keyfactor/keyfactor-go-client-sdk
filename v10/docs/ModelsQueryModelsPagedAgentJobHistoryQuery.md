# ModelsQueryModelsPagedAgentJobHistoryQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | [optional] 
**PageReturned** | Pointer to **int64** | The current page within the result set to be returned | [optional] 
**ReturnLimit** | Pointer to **int64** | Maximum number of records to be returned in a single call | [optional] 
**SortField** | Pointer to **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | [optional] 
**SortAscending** | Pointer to **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | [optional] 

## Methods

### NewModelsQueryModelsPagedAgentJobHistoryQuery

`func NewModelsQueryModelsPagedAgentJobHistoryQuery() *ModelsQueryModelsPagedAgentJobHistoryQuery`

NewModelsQueryModelsPagedAgentJobHistoryQuery instantiates a new ModelsQueryModelsPagedAgentJobHistoryQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsQueryModelsPagedAgentJobHistoryQueryWithDefaults

`func NewModelsQueryModelsPagedAgentJobHistoryQueryWithDefaults() *ModelsQueryModelsPagedAgentJobHistoryQuery`

NewModelsQueryModelsPagedAgentJobHistoryQueryWithDefaults instantiates a new ModelsQueryModelsPagedAgentJobHistoryQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetPageReturned

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetPageReturned() int64`

GetPageReturned returns the PageReturned field if non-nil, zero value otherwise.

### GetPageReturnedOk

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetPageReturnedOk() (*int64, bool)`

GetPageReturnedOk returns a tuple with the PageReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReturned

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) SetPageReturned(v int64)`

SetPageReturned sets PageReturned field to given value.

### HasPageReturned

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) HasPageReturned() bool`

HasPageReturned returns a boolean if a field has been set.

### GetReturnLimit

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetReturnLimit() int64`

GetReturnLimit returns the ReturnLimit field if non-nil, zero value otherwise.

### GetReturnLimitOk

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetReturnLimitOk() (*int64, bool)`

GetReturnLimitOk returns a tuple with the ReturnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLimit

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) SetReturnLimit(v int64)`

SetReturnLimit sets ReturnLimit field to given value.

### HasReturnLimit

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) HasReturnLimit() bool`

HasReturnLimit returns a boolean if a field has been set.

### GetSortField

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortAscending

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetSortAscending() int64`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) GetSortAscendingOk() (*int64, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) SetSortAscending(v int64)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *ModelsQueryModelsPagedAgentJobHistoryQuery) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


