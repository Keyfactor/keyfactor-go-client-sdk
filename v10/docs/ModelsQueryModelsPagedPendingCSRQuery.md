# ModelsQueryModelsPagedPendingCSRQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | [optional] 
**PageReturned** | Pointer to **int32** | The current page within the result set to be returned | [optional] 
**ReturnLimit** | Pointer to **int32** | Maximum number of records to be returned in a single call | [optional] 
**SortField** | Pointer to **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | [optional] 
**SortAscending** | Pointer to **int32** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | [optional] 

## Methods

### NewModelsQueryModelsPagedPendingCSRQuery

`func NewModelsQueryModelsPagedPendingCSRQuery() *ModelsQueryModelsPagedPendingCSRQuery`

NewModelsQueryModelsPagedPendingCSRQuery instantiates a new ModelsQueryModelsPagedPendingCSRQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsQueryModelsPagedPendingCSRQueryWithDefaults

`func NewModelsQueryModelsPagedPendingCSRQueryWithDefaults() *ModelsQueryModelsPagedPendingCSRQuery`

NewModelsQueryModelsPagedPendingCSRQueryWithDefaults instantiates a new ModelsQueryModelsPagedPendingCSRQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ModelsQueryModelsPagedPendingCSRQuery) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *ModelsQueryModelsPagedPendingCSRQuery) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetPageReturned

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetPageReturned() int32`

GetPageReturned returns the PageReturned field if non-nil, zero value otherwise.

### GetPageReturnedOk

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetPageReturnedOk() (*int32, bool)`

GetPageReturnedOk returns a tuple with the PageReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReturned

`func (o *ModelsQueryModelsPagedPendingCSRQuery) SetPageReturned(v int32)`

SetPageReturned sets PageReturned field to given value.

### HasPageReturned

`func (o *ModelsQueryModelsPagedPendingCSRQuery) HasPageReturned() bool`

HasPageReturned returns a boolean if a field has been set.

### GetReturnLimit

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetReturnLimit() int32`

GetReturnLimit returns the ReturnLimit field if non-nil, zero value otherwise.

### GetReturnLimitOk

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetReturnLimitOk() (*int32, bool)`

GetReturnLimitOk returns a tuple with the ReturnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLimit

`func (o *ModelsQueryModelsPagedPendingCSRQuery) SetReturnLimit(v int32)`

SetReturnLimit sets ReturnLimit field to given value.

### HasReturnLimit

`func (o *ModelsQueryModelsPagedPendingCSRQuery) HasReturnLimit() bool`

HasReturnLimit returns a boolean if a field has been set.

### GetSortField

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *ModelsQueryModelsPagedPendingCSRQuery) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *ModelsQueryModelsPagedPendingCSRQuery) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortAscending

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetSortAscending() int32`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *ModelsQueryModelsPagedPendingCSRQuery) GetSortAscendingOk() (*int32, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *ModelsQueryModelsPagedPendingCSRQuery) SetSortAscending(v int32)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *ModelsQueryModelsPagedPendingCSRQuery) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


