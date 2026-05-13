# CertificatesCertificateLocationsGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreType** | Pointer to **NullableString** |  | [optional] [readonly] 
**StoreTypeId** | Pointer to **int32** |  | [optional] [readonly] 
**StoreCount** | Pointer to **int32** |  | [optional] [readonly] 
**Locations** | Pointer to [**[]CertificatesCertificateStoreLocationsDetailResponse**](CertificatesCertificateStoreLocationsDetailResponse.md) |  | [optional] 

## Methods

### NewCertificatesCertificateLocationsGroupResponse

`func NewCertificatesCertificateLocationsGroupResponse() *CertificatesCertificateLocationsGroupResponse`

NewCertificatesCertificateLocationsGroupResponse instantiates a new CertificatesCertificateLocationsGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificatesCertificateLocationsGroupResponseWithDefaults

`func NewCertificatesCertificateLocationsGroupResponseWithDefaults() *CertificatesCertificateLocationsGroupResponse`

NewCertificatesCertificateLocationsGroupResponseWithDefaults instantiates a new CertificatesCertificateLocationsGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreType

`func (o *CertificatesCertificateLocationsGroupResponse) GetStoreType() string`

GetStoreType returns the StoreType field if non-nil, zero value otherwise.

### GetStoreTypeOk

`func (o *CertificatesCertificateLocationsGroupResponse) GetStoreTypeOk() (*string, bool)`

GetStoreTypeOk returns a tuple with the StoreType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreType

`func (o *CertificatesCertificateLocationsGroupResponse) SetStoreType(v string)`

SetStoreType sets StoreType field to given value.

### HasStoreType

`func (o *CertificatesCertificateLocationsGroupResponse) HasStoreType() bool`

HasStoreType returns a boolean if a field has been set.

### SetStoreTypeNil

`func (o *CertificatesCertificateLocationsGroupResponse) SetStoreTypeNil(b bool)`

 SetStoreTypeNil sets the value for StoreType to be an explicit nil

### UnsetStoreType
`func (o *CertificatesCertificateLocationsGroupResponse) UnsetStoreType()`

UnsetStoreType ensures that no value is present for StoreType, not even an explicit nil
### GetStoreTypeId

`func (o *CertificatesCertificateLocationsGroupResponse) GetStoreTypeId() int32`

GetStoreTypeId returns the StoreTypeId field if non-nil, zero value otherwise.

### GetStoreTypeIdOk

`func (o *CertificatesCertificateLocationsGroupResponse) GetStoreTypeIdOk() (*int32, bool)`

GetStoreTypeIdOk returns a tuple with the StoreTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypeId

`func (o *CertificatesCertificateLocationsGroupResponse) SetStoreTypeId(v int32)`

SetStoreTypeId sets StoreTypeId field to given value.

### HasStoreTypeId

`func (o *CertificatesCertificateLocationsGroupResponse) HasStoreTypeId() bool`

HasStoreTypeId returns a boolean if a field has been set.

### GetStoreCount

`func (o *CertificatesCertificateLocationsGroupResponse) GetStoreCount() int32`

GetStoreCount returns the StoreCount field if non-nil, zero value otherwise.

### GetStoreCountOk

`func (o *CertificatesCertificateLocationsGroupResponse) GetStoreCountOk() (*int32, bool)`

GetStoreCountOk returns a tuple with the StoreCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreCount

`func (o *CertificatesCertificateLocationsGroupResponse) SetStoreCount(v int32)`

SetStoreCount sets StoreCount field to given value.

### HasStoreCount

`func (o *CertificatesCertificateLocationsGroupResponse) HasStoreCount() bool`

HasStoreCount returns a boolean if a field has been set.

### GetLocations

`func (o *CertificatesCertificateLocationsGroupResponse) GetLocations() []CertificatesCertificateStoreLocationsDetailResponse`

GetLocations returns the Locations field if non-nil, zero value otherwise.

### GetLocationsOk

`func (o *CertificatesCertificateLocationsGroupResponse) GetLocationsOk() (*[]CertificatesCertificateStoreLocationsDetailResponse, bool)`

GetLocationsOk returns a tuple with the Locations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocations

`func (o *CertificatesCertificateLocationsGroupResponse) SetLocations(v []CertificatesCertificateStoreLocationsDetailResponse)`

SetLocations sets Locations field to given value.

### HasLocations

`func (o *CertificatesCertificateLocationsGroupResponse) HasLocations() bool`

HasLocations returns a boolean if a field has been set.

### SetLocationsNil

`func (o *CertificatesCertificateLocationsGroupResponse) SetLocationsNil(b bool)`

 SetLocationsNil sets the value for Locations to be an explicit nil

### UnsetLocations
`func (o *CertificatesCertificateLocationsGroupResponse) UnsetLocations()`

UnsetLocations ensures that no value is present for Locations, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


