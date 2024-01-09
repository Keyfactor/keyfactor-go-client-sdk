# CSSCMSDataModelModelsCertificateDetails

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

### NewCSSCMSDataModelModelsCertificateDetails

`func NewCSSCMSDataModelModelsCertificateDetails() *CSSCMSDataModelModelsCertificateDetails`

NewCSSCMSDataModelModelsCertificateDetails instantiates a new CSSCMSDataModelModelsCertificateDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateDetailsWithDefaults

`func NewCSSCMSDataModelModelsCertificateDetailsWithDefaults() *CSSCMSDataModelModelsCertificateDetails`

NewCSSCMSDataModelModelsCertificateDetailsWithDefaults instantiates a new CSSCMSDataModelModelsCertificateDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuedDN

`func (o *CSSCMSDataModelModelsCertificateDetails) GetIssuedDN() string`

GetIssuedDN returns the IssuedDN field if non-nil, zero value otherwise.

### GetIssuedDNOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetIssuedDNOk() (*string, bool)`

GetIssuedDNOk returns a tuple with the IssuedDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDN

`func (o *CSSCMSDataModelModelsCertificateDetails) SetIssuedDN(v string)`

SetIssuedDN sets IssuedDN field to given value.

### HasIssuedDN

`func (o *CSSCMSDataModelModelsCertificateDetails) HasIssuedDN() bool`

HasIssuedDN returns a boolean if a field has been set.

### SetIssuedDNNil

`func (o *CSSCMSDataModelModelsCertificateDetails) SetIssuedDNNil(b bool)`

 SetIssuedDNNil sets the value for IssuedDN to be an explicit nil

### UnsetIssuedDN
`func (o *CSSCMSDataModelModelsCertificateDetails) UnsetIssuedDN()`

UnsetIssuedDN ensures that no value is present for IssuedDN, not even an explicit nil
### GetIssuerDN

`func (o *CSSCMSDataModelModelsCertificateDetails) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *CSSCMSDataModelModelsCertificateDetails) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *CSSCMSDataModelModelsCertificateDetails) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### SetIssuerDNNil

`func (o *CSSCMSDataModelModelsCertificateDetails) SetIssuerDNNil(b bool)`

 SetIssuerDNNil sets the value for IssuerDN to be an explicit nil

### UnsetIssuerDN
`func (o *CSSCMSDataModelModelsCertificateDetails) UnsetIssuerDN()`

UnsetIssuerDN ensures that no value is present for IssuerDN, not even an explicit nil
### GetThumbprint

`func (o *CSSCMSDataModelModelsCertificateDetails) GetThumbprint() string`

GetThumbprint returns the Thumbprint field if non-nil, zero value otherwise.

### GetThumbprintOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetThumbprintOk() (*string, bool)`

GetThumbprintOk returns a tuple with the Thumbprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbprint

`func (o *CSSCMSDataModelModelsCertificateDetails) SetThumbprint(v string)`

SetThumbprint sets Thumbprint field to given value.

### HasThumbprint

`func (o *CSSCMSDataModelModelsCertificateDetails) HasThumbprint() bool`

HasThumbprint returns a boolean if a field has been set.

### SetThumbprintNil

`func (o *CSSCMSDataModelModelsCertificateDetails) SetThumbprintNil(b bool)`

 SetThumbprintNil sets the value for Thumbprint to be an explicit nil

### UnsetThumbprint
`func (o *CSSCMSDataModelModelsCertificateDetails) UnsetThumbprint()`

UnsetThumbprint ensures that no value is present for Thumbprint, not even an explicit nil
### GetNotAfter

`func (o *CSSCMSDataModelModelsCertificateDetails) GetNotAfter() time.Time`

GetNotAfter returns the NotAfter field if non-nil, zero value otherwise.

### GetNotAfterOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetNotAfterOk() (*time.Time, bool)`

GetNotAfterOk returns a tuple with the NotAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAfter

`func (o *CSSCMSDataModelModelsCertificateDetails) SetNotAfter(v time.Time)`

SetNotAfter sets NotAfter field to given value.

### HasNotAfter

`func (o *CSSCMSDataModelModelsCertificateDetails) HasNotAfter() bool`

HasNotAfter returns a boolean if a field has been set.

### GetNotBefore

`func (o *CSSCMSDataModelModelsCertificateDetails) GetNotBefore() time.Time`

GetNotBefore returns the NotBefore field if non-nil, zero value otherwise.

### GetNotBeforeOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetNotBeforeOk() (*time.Time, bool)`

GetNotBeforeOk returns a tuple with the NotBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotBefore

`func (o *CSSCMSDataModelModelsCertificateDetails) SetNotBefore(v time.Time)`

SetNotBefore sets NotBefore field to given value.

### HasNotBefore

`func (o *CSSCMSDataModelModelsCertificateDetails) HasNotBefore() bool`

HasNotBefore returns a boolean if a field has been set.

### GetMetadata

`func (o *CSSCMSDataModelModelsCertificateDetails) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsCertificateDetails) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsCertificateDetails) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsCertificateDetails) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsCertificateDetails) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsEndEntity

`func (o *CSSCMSDataModelModelsCertificateDetails) GetIsEndEntity() bool`

GetIsEndEntity returns the IsEndEntity field if non-nil, zero value otherwise.

### GetIsEndEntityOk

`func (o *CSSCMSDataModelModelsCertificateDetails) GetIsEndEntityOk() (*bool, bool)`

GetIsEndEntityOk returns a tuple with the IsEndEntity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEndEntity

`func (o *CSSCMSDataModelModelsCertificateDetails) SetIsEndEntity(v bool)`

SetIsEndEntity sets IsEndEntity field to given value.

### HasIsEndEntity

`func (o *CSSCMSDataModelModelsCertificateDetails) HasIsEndEntity() bool`

HasIsEndEntity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


