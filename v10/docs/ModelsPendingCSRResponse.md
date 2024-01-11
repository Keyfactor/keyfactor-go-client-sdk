# ModelsPendingCSRResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CSR** | Pointer to **string** |  | [optional] 
**RequestTime** | Pointer to **time.Time** |  | [optional] 
**Subject** | Pointer to **[]string** |  | [optional] 

## Methods

### NewModelsPendingCSRResponse

`func NewModelsPendingCSRResponse() *ModelsPendingCSRResponse`

NewModelsPendingCSRResponse instantiates a new ModelsPendingCSRResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsPendingCSRResponseWithDefaults

`func NewModelsPendingCSRResponseWithDefaults() *ModelsPendingCSRResponse`

NewModelsPendingCSRResponseWithDefaults instantiates a new ModelsPendingCSRResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelsPendingCSRResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelsPendingCSRResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelsPendingCSRResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ModelsPendingCSRResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCSR

`func (o *ModelsPendingCSRResponse) GetCSR() string`

GetCSR returns the CSR field if non-nil, zero value otherwise.

### GetCSROk

`func (o *ModelsPendingCSRResponse) GetCSROk() (*string, bool)`

GetCSROk returns a tuple with the CSR field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCSR

`func (o *ModelsPendingCSRResponse) SetCSR(v string)`

SetCSR sets CSR field to given value.

### HasCSR

`func (o *ModelsPendingCSRResponse) HasCSR() bool`

HasCSR returns a boolean if a field has been set.

### GetRequestTime

`func (o *ModelsPendingCSRResponse) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *ModelsPendingCSRResponse) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *ModelsPendingCSRResponse) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *ModelsPendingCSRResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetSubject

`func (o *ModelsPendingCSRResponse) GetSubject() []string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ModelsPendingCSRResponse) GetSubjectOk() (*[]string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ModelsPendingCSRResponse) SetSubject(v []string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ModelsPendingCSRResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


