# ModelsCertificateValidationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** | States whether the certificate is valid or not | [optional] 
**Results** | Pointer to **map[string]string** | Lists any reasons why the certificate is not valid | [optional] 

## Methods

### NewModelsCertificateValidationResponse

`func NewModelsCertificateValidationResponse() *ModelsCertificateValidationResponse`

NewModelsCertificateValidationResponse instantiates a new ModelsCertificateValidationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateValidationResponseWithDefaults

`func NewModelsCertificateValidationResponseWithDefaults() *ModelsCertificateValidationResponse`

NewModelsCertificateValidationResponseWithDefaults instantiates a new ModelsCertificateValidationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ModelsCertificateValidationResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ModelsCertificateValidationResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ModelsCertificateValidationResponse) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ModelsCertificateValidationResponse) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetResults

`func (o *ModelsCertificateValidationResponse) GetResults() map[string]string`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ModelsCertificateValidationResponse) GetResultsOk() (*map[string]string, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ModelsCertificateValidationResponse) SetResults(v map[string]string)`

SetResults sets Results field to given value.

### HasResults

`func (o *ModelsCertificateValidationResponse) HasResults() bool`

HasResults returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


