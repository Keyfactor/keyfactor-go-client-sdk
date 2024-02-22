# ModelsPendingCSRResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 
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

### GetCsr

`func (o *ModelsPendingCSRResponse) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *ModelsPendingCSRResponse) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *ModelsPendingCSRResponse) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *ModelsPendingCSRResponse) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *ModelsPendingCSRResponse) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *ModelsPendingCSRResponse) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
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

### SetSubjectNil

`func (o *ModelsPendingCSRResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *ModelsPendingCSRResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


