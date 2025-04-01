# CSSCMSDataModelModelsMetadataAllUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** |  | [optional] 
**CertificateIds** | Pointer to **[]int32** |  | [optional] 
**Metadata** | [**[]CSSCMSDataModelModelsMetadataSingleUpdateRequest**](CSSCMSDataModelModelsMetadataSingleUpdateRequest.md) |  | 

## Methods

### NewCSSCMSDataModelModelsMetadataAllUpdateRequest

`func NewCSSCMSDataModelModelsMetadataAllUpdateRequest(metadata []CSSCMSDataModelModelsMetadataSingleUpdateRequest, ) *CSSCMSDataModelModelsMetadataAllUpdateRequest`

NewCSSCMSDataModelModelsMetadataAllUpdateRequest instantiates a new CSSCMSDataModelModelsMetadataAllUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsMetadataAllUpdateRequestWithDefaults

`func NewCSSCMSDataModelModelsMetadataAllUpdateRequestWithDefaults() *CSSCMSDataModelModelsMetadataAllUpdateRequest`

NewCSSCMSDataModelModelsMetadataAllUpdateRequestWithDefaults instantiates a new CSSCMSDataModelModelsMetadataAllUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetCertificateIds

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) GetCertificateIds() []int32`

GetCertificateIds returns the CertificateIds field if non-nil, zero value otherwise.

### GetCertificateIdsOk

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) GetCertificateIdsOk() (*[]int32, bool)`

GetCertificateIdsOk returns a tuple with the CertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIds

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) SetCertificateIds(v []int32)`

SetCertificateIds sets CertificateIds field to given value.

### HasCertificateIds

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) HasCertificateIds() bool`

HasCertificateIds returns a boolean if a field has been set.

### SetCertificateIdsNil

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) SetCertificateIdsNil(b bool)`

 SetCertificateIdsNil sets the value for CertificateIds to be an explicit nil

### UnsetCertificateIds
`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) UnsetCertificateIds()`

UnsetCertificateIds ensures that no value is present for CertificateIds, not even an explicit nil
### GetMetadata

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) GetMetadata() []CSSCMSDataModelModelsMetadataSingleUpdateRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) GetMetadataOk() (*[]CSSCMSDataModelModelsMetadataSingleUpdateRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsMetadataAllUpdateRequest) SetMetadata(v []CSSCMSDataModelModelsMetadataSingleUpdateRequest)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


