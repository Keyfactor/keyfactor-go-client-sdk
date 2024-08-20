# ModelsQueryModelsPagedMetadataFieldQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | [optional] 
**PageReturned** | Pointer to **int64** | The current page within the result set to be returned | [optional] 
**ReturnLimit** | Pointer to **int64** | Maximum number of records to be returned in a single call | [optional] 
**SortField** | Pointer to **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | [optional] 
**SortAscending** | Pointer to **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | [optional] 

## Methods

### NewModelsQueryModelsPagedMetadataFieldQuery

`func NewModelsQueryModelsPagedMetadataFieldQuery() *ModelsQueryModelsPagedMetadataFieldQuery`

NewModelsQueryModelsPagedMetadataFieldQuery instantiates a new ModelsQueryModelsPagedMetadataFieldQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsQueryModelsPagedMetadataFieldQueryWithDefaults

`func NewModelsQueryModelsPagedMetadataFieldQueryWithDefaults() *ModelsQueryModelsPagedMetadataFieldQuery`

NewModelsQueryModelsPagedMetadataFieldQueryWithDefaults instantiates a new ModelsQueryModelsPagedMetadataFieldQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetPageReturned

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetPageReturned() int64`

GetPageReturned returns the PageReturned field if non-nil, zero value otherwise.

### GetPageReturnedOk

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetPageReturnedOk() (*int64, bool)`

GetPageReturnedOk returns a tuple with the PageReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReturned

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) SetPageReturned(v int64)`

SetPageReturned sets PageReturned field to given value.

### HasPageReturned

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) HasPageReturned() bool`

HasPageReturned returns a boolean if a field has been set.

### GetReturnLimit

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetReturnLimit() int64`

GetReturnLimit returns the ReturnLimit field if non-nil, zero value otherwise.

### GetReturnLimitOk

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetReturnLimitOk() (*int64, bool)`

GetReturnLimitOk returns a tuple with the ReturnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLimit

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) SetReturnLimit(v int64)`

SetReturnLimit sets ReturnLimit field to given value.

### HasReturnLimit

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) HasReturnLimit() bool`

HasReturnLimit returns a boolean if a field has been set.

### GetSortField

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortAscending

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetSortAscending() int64`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) GetSortAscendingOk() (*int64, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) SetSortAscending(v int64)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *ModelsQueryModelsPagedMetadataFieldQuery) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


