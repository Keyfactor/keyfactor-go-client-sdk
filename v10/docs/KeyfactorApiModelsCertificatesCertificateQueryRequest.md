# KeyfactorApiModelsCertificatesCertificateQueryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueryString** | Pointer to **string** | Contents of the query (ex: field1 -eq value1 AND field2 -gt value2) | [optional] 
**PageReturned** | Pointer to **int64** | The current page within the result set to be returned | [optional] 
**ReturnLimit** | Pointer to **int64** | Maximum number of records to be returned in a single call | [optional] 
**SortField** | Pointer to **string** | Field by which the results should be sorted (view results via Management Portal for sortable columns) | [optional] 
**SortAscending** | Pointer to **int64** | Field sort direction [0&#x3D;ascending, 1&#x3D;descending] | [optional] 
**IncludeRevoked** | Pointer to **bool** | Select &#39;true&#39; to include revoked certificates in the results | [optional] 
**IncludeExpired** | Pointer to **bool** | Select &#39;true&#39; to include expired certificates in the results | [optional] 

## Methods

### NewKeyfactorApiModelsCertificatesCertificateQueryRequest

`func NewKeyfactorApiModelsCertificatesCertificateQueryRequest() *KeyfactorApiModelsCertificatesCertificateQueryRequest`

NewKeyfactorApiModelsCertificatesCertificateQueryRequest instantiates a new KeyfactorApiModelsCertificatesCertificateQueryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsCertificatesCertificateQueryRequestWithDefaults

`func NewKeyfactorApiModelsCertificatesCertificateQueryRequestWithDefaults() *KeyfactorApiModelsCertificatesCertificateQueryRequest`

NewKeyfactorApiModelsCertificatesCertificateQueryRequestWithDefaults instantiates a new KeyfactorApiModelsCertificatesCertificateQueryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueryString

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetQueryString() string`

GetQueryString returns the QueryString field if non-nil, zero value otherwise.

### GetQueryStringOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetQueryStringOk() (*string, bool)`

GetQueryStringOk returns a tuple with the QueryString field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryString

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetQueryString(v string)`

SetQueryString sets QueryString field to given value.

### HasQueryString

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasQueryString() bool`

HasQueryString returns a boolean if a field has been set.

### GetPageReturned

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetPageReturned() int64`

GetPageReturned returns the PageReturned field if non-nil, zero value otherwise.

### GetPageReturnedOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetPageReturnedOk() (*int64, bool)`

GetPageReturnedOk returns a tuple with the PageReturned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageReturned

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetPageReturned(v int64)`

SetPageReturned sets PageReturned field to given value.

### HasPageReturned

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasPageReturned() bool`

HasPageReturned returns a boolean if a field has been set.

### GetReturnLimit

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetReturnLimit() int64`

GetReturnLimit returns the ReturnLimit field if non-nil, zero value otherwise.

### GetReturnLimitOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetReturnLimitOk() (*int64, bool)`

GetReturnLimitOk returns a tuple with the ReturnLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnLimit

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetReturnLimit(v int64)`

SetReturnLimit sets ReturnLimit field to given value.

### HasReturnLimit

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasReturnLimit() bool`

HasReturnLimit returns a boolean if a field has been set.

### GetSortField

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetSortField() string`

GetSortField returns the SortField field if non-nil, zero value otherwise.

### GetSortFieldOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetSortFieldOk() (*string, bool)`

GetSortFieldOk returns a tuple with the SortField field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortField

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetSortField(v string)`

SetSortField sets SortField field to given value.

### HasSortField

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasSortField() bool`

HasSortField returns a boolean if a field has been set.

### GetSortAscending

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetSortAscending() int64`

GetSortAscending returns the SortAscending field if non-nil, zero value otherwise.

### GetSortAscendingOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetSortAscendingOk() (*int64, bool)`

GetSortAscendingOk returns a tuple with the SortAscending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortAscending

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetSortAscending(v int64)`

SetSortAscending sets SortAscending field to given value.

### HasSortAscending

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasSortAscending() bool`

HasSortAscending returns a boolean if a field has been set.

### GetIncludeRevoked

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetIncludeRevoked() bool`

GetIncludeRevoked returns the IncludeRevoked field if non-nil, zero value otherwise.

### GetIncludeRevokedOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetIncludeRevokedOk() (*bool, bool)`

GetIncludeRevokedOk returns a tuple with the IncludeRevoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRevoked

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetIncludeRevoked(v bool)`

SetIncludeRevoked sets IncludeRevoked field to given value.

### HasIncludeRevoked

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasIncludeRevoked() bool`

HasIncludeRevoked returns a boolean if a field has been set.

### GetIncludeExpired

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetIncludeExpired() bool`

GetIncludeExpired returns the IncludeExpired field if non-nil, zero value otherwise.

### GetIncludeExpiredOk

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) GetIncludeExpiredOk() (*bool, bool)`

GetIncludeExpiredOk returns a tuple with the IncludeExpired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeExpired

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) SetIncludeExpired(v bool)`

SetIncludeExpired sets IncludeExpired field to given value.

### HasIncludeExpired

`func (o *KeyfactorApiModelsCertificatesCertificateQueryRequest) HasIncludeExpired() bool`

HasIncludeExpired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


