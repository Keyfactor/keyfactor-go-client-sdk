# ModelsCertificateDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssuedDN** | Pointer to **NullableString** |  | [optional] 
**IssuerDN** | Pointer to **NullableString** |  | [optional] 
**Thumbprint** | Pointer to **NullableString** |  | [optional] 
**NotAfter** | Pointer to **time.Time** |  | [optional] 
**NotBefore** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**IsEndEntity** | Pointer to **bool** |  | [optional] 

## Methods

### NewModelsCertificateDetails

`func NewModelsCertificateDetails() *ModelsCertificateDetails`

NewModelsCertificateDetails instantiates a new ModelsCertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateDetailsWithDefaults

`func NewModelsCertificateDetailsWithDefaults() *ModelsCertificateDetails`

NewModelsCertificateDetailsWithDefaults instantiates a new ModelsCertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuedDN

`func (o *ModelsCertificateDetails) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *ModelsCertificateDetails) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *ModelsCertificateDetails) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *ModelsCertificateDetails) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *ModelsCertificateDetails) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *ModelsCertificateDetails) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetIssuerDN

`func (o *ModelsCertificateDetails) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *ModelsCertificateDetails) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *ModelsCertificateDetails) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *ModelsCertificateDetails) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *ModelsCertificateDetails) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *ModelsCertificateDetails) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *ModelsCertificateDetails) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *ModelsCertificateDetails) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *ModelsCertificateDetails) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *ModelsCertificateDetails) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *ModelsCertificateDetails) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *ModelsCertificateDetails) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetNotAfter

`func (o *ModelsCertificateDetails) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *ModelsCertificateDetails) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *ModelsCertificateDetails) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *ModelsCertificateDetails) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *ModelsCertificateDetails) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *ModelsCertificateDetails) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *ModelsCertificateDetails) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *ModelsCertificateDetails) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsCertificateDetails) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsCertificateDetails) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsCertificateDetails) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsCertificateDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ModelsCertificateDetails) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ModelsCertificateDetails) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsEndEntity

`func (o *ModelsCertificateDetails) GetIsEndEntity() bool`

GetIsEndEntity returns the IsEndEntity field if non-nil, zero value otherwise.

### GetIsEndEntityOk

`func (o *ModelsCertificateDetails) GetIsEndEntityOk() (*bool, bool)`

GetIsEndEntityOk returns a tuple with the IsEndEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEndEntity

`func (o *ModelsCertificateDetails) SetIsEndEntity(v bool)`

SetIsEndEntity sets IsEndEntity field to given value.

### HasIsEndEntity

`func (o *ModelsCertificateDetails) HasIsEndEntity() bool`

HasIsEndEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


