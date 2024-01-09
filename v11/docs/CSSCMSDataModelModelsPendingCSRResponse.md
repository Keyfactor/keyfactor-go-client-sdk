# CSSCMSDataModelModelsPendingCSRResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Csr** | Pointer to **NullableString** |  | [optional] 
**RequestTime** | Pointer to **time.Time** |  | [optional] 
**Subject** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsPendingCSRResponse

`func NewCSSCMSDataModelModelsPendingCSRResponse() *CSSCMSDataModelModelsPendingCSRResponse`

NewCSSCMSDataModelModelsPendingCSRResponse instantiates a new CSSCMSDataModelModelsPendingCSRResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsPendingCSRResponseWithDefaults

`func NewCSSCMSDataModelModelsPendingCSRResponseWithDefaults() *CSSCMSDataModelModelsPendingCSRResponse`

NewCSSCMSDataModelModelsPendingCSRResponseWithDefaults instantiates a new CSSCMSDataModelModelsPendingCSRResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CSSCMSDataModelModelsPendingCSRResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CSSCMSDataModelModelsPendingCSRResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCsr

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetCsr() string`

GetCsr returns the Csr field if non-nil, zero value otherwise.

### GetCsrOk

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetCsrOk() (*string, bool)`

GetCsrOk returns a tuple with the Csr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsr

`func (o *CSSCMSDataModelModelsPendingCSRResponse) SetCsr(v string)`

SetCsr sets Csr field to given value.

### HasCsr

`func (o *CSSCMSDataModelModelsPendingCSRResponse) HasCsr() bool`

HasCsr returns a boolean if a field has been set.

### SetCsrNil

`func (o *CSSCMSDataModelModelsPendingCSRResponse) SetCsrNil(b bool)`

 SetCsrNil sets the value for Csr to be an explicit nil

### UnsetCsr
`func (o *CSSCMSDataModelModelsPendingCSRResponse) UnsetCsr()`

UnsetCsr ensures that no value is present for Csr, not even an explicit nil
### GetRequestTime

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetRequestTime() time.Time`

GetRequestTime returns the RequestTime field if non-nil, zero value otherwise.

### GetRequestTimeOk

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetRequestTimeOk() (*time.Time, bool)`

GetRequestTimeOk returns a tuple with the RequestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTime

`func (o *CSSCMSDataModelModelsPendingCSRResponse) SetRequestTime(v time.Time)`

SetRequestTime sets RequestTime field to given value.

### HasRequestTime

`func (o *CSSCMSDataModelModelsPendingCSRResponse) HasRequestTime() bool`

HasRequestTime returns a boolean if a field has been set.

### GetSubject

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetSubject() []string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CSSCMSDataModelModelsPendingCSRResponse) GetSubjectOk() (*[]string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CSSCMSDataModelModelsPendingCSRResponse) SetSubject(v []string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *CSSCMSDataModelModelsPendingCSRResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *CSSCMSDataModelModelsPendingCSRResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *CSSCMSDataModelModelsPendingCSRResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


