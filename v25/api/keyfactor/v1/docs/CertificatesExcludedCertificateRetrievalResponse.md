# CertificatesExcludedCertificateRetrievalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**CertificateThumbprint** | Pointer to **NullableString** |  | [optional] 
**IssuedCN** | Pointer to **NullableString** |  | [optional] 
**ExcludingUser** | Pointer to **NullableString** |  | [optional] 
**DateExcluded** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCertificatesExcludedCertificateRetrievalResponse

`func NewCertificatesExcludedCertificateRetrievalResponse() *CertificatesExcludedCertificateRetrievalResponse`

NewCertificatesExcludedCertificateRetrievalResponse instantiates a new CertificatesExcludedCertificateRetrievalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesExcludedCertificateRetrievalResponseWithDefaults

`func NewCertificatesExcludedCertificateRetrievalResponseWithDefaults() *CertificatesExcludedCertificateRetrievalResponse`

NewCertificatesExcludedCertificateRetrievalResponseWithDefaults instantiates a new CertificatesExcludedCertificateRetrievalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CertificatesExcludedCertificateRetrievalResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCertificateThumbprint

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetCertificateThumbprint() string`

GetCertificateThumbprint returns the CertificateThumbprint field if non-nil, zero value otherwise.

### GetCertificateThumbprintOk

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetCertificateThumbprintOk() (*string, bool)`

GetCertificateThumbprintOk returns a tuple with the CertificateThumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateThumbprint

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetCertificateThumbprint(v string)`

SetCertificateThumbprint sets CertificateThumbprint field to given value.

### HasCertificateThumbprint

`func (o *CertificatesExcludedCertificateRetrievalResponse) HasCertificateThumbprint() bool`

HasCertificateThumbprint returns a boolean if a field has been set.

### SetCertificateThumbprintNil

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetCertificateThumbprintNil(b bool)`

 SetCertificateThumbprintNil sets the value for CertificateThumbprint to be an explicit nil

### UnsetCertificateThumbprint
`func (o *CertificatesExcludedCertificateRetrievalResponse) UnsetCertificateThumbprint()`

UnsetCertificateThumbprint ensures that no value is present for CertificateThumbprint, not even an explicit nil
### GetIssuedCN

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetIssuedCN() string`

GetIssuedCN returns the IssuedCN field if non-nil, zero value otherwise.

### GetIssuedCNOk

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetIssuedCNOk() (*string, bool)`

GetIssuedCNOk returns a tuple with the IssuedCN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedCN

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetIssuedCN(v string)`

SetIssuedCN sets IssuedCN field to given value.

### HasIssuedCN

`func (o *CertificatesExcludedCertificateRetrievalResponse) HasIssuedCN() bool`

HasIssuedCN returns a boolean if a field has been set.

### SetIssuedCNNil

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetIssuedCNNil(b bool)`

 SetIssuedCNNil sets the value for IssuedCN to be an explicit nil

### UnsetIssuedCN
`func (o *CertificatesExcludedCertificateRetrievalResponse) UnsetIssuedCN()`

UnsetIssuedCN ensures that no value is present for IssuedCN, not even an explicit nil
### GetExcludingUser

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetExcludingUser() string`

GetExcludingUser returns the ExcludingUser field if non-nil, zero value otherwise.

### GetExcludingUserOk

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetExcludingUserOk() (*string, bool)`

GetExcludingUserOk returns a tuple with the ExcludingUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludingUser

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetExcludingUser(v string)`

SetExcludingUser sets ExcludingUser field to given value.

### HasExcludingUser

`func (o *CertificatesExcludedCertificateRetrievalResponse) HasExcludingUser() bool`

HasExcludingUser returns a boolean if a field has been set.

### SetExcludingUserNil

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetExcludingUserNil(b bool)`

 SetExcludingUserNil sets the value for ExcludingUser to be an explicit nil

### UnsetExcludingUser
`func (o *CertificatesExcludedCertificateRetrievalResponse) UnsetExcludingUser()`

UnsetExcludingUser ensures that no value is present for ExcludingUser, not even an explicit nil
### GetDateExcluded

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetDateExcluded() time.Time`

GetDateExcluded returns the DateExcluded field if non-nil, zero value otherwise.

### GetDateExcludedOk

`func (o *CertificatesExcludedCertificateRetrievalResponse) GetDateExcludedOk() (*time.Time, bool)`

GetDateExcludedOk returns a tuple with the DateExcluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateExcluded

`func (o *CertificatesExcludedCertificateRetrievalResponse) SetDateExcluded(v time.Time)`

SetDateExcluded sets DateExcluded field to given value.

### HasDateExcluded

`func (o *CertificatesExcludedCertificateRetrievalResponse) HasDateExcluded() bool`

HasDateExcluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


