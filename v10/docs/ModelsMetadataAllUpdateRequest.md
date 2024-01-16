# ModelsMetadataAllUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **string** |  | [optional] 
**CertificateIds** | Pointer to **[]int64** |  | [optional] 
**Metadata** | [**[]ModelsMetadataSingleUpdateRequest**](ModelsMetadataSingleUpdateRequest.md) |  | 

## Methods

### NewModelsMetadataAllUpdateRequest

`func NewModelsMetadataAllUpdateRequest(metadata []ModelsMetadataSingleUpdateRequest, ) *ModelsMetadataAllUpdateRequest`

NewModelsMetadataAllUpdateRequest instantiates a new ModelsMetadataAllUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMetadataAllUpdateRequestWithDefaults

`func NewModelsMetadataAllUpdateRequestWithDefaults() *ModelsMetadataAllUpdateRequest`

NewModelsMetadataAllUpdateRequestWithDefaults instantiates a new ModelsMetadataAllUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *ModelsMetadataAllUpdateRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *ModelsMetadataAllUpdateRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *ModelsMetadataAllUpdateRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *ModelsMetadataAllUpdateRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### GetCertificateIds

`func (o *ModelsMetadataAllUpdateRequest) GetCertificateIds() []int64`

GetCertificateIds returns the CertificateIds field if non-nil, zero value otherwise.

### GetCertificateIdsOk

`func (o *ModelsMetadataAllUpdateRequest) GetCertificateIdsOk() (*[]int64, bool)`

GetCertificateIdsOk returns a tuple with the CertificateIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateIds

`func (o *ModelsMetadataAllUpdateRequest) SetCertificateIds(v []int64)`

SetCertificateIds sets CertificateIds field to given value.

### HasCertificateIds

`func (o *ModelsMetadataAllUpdateRequest) HasCertificateIds() bool`

HasCertificateIds returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsMetadataAllUpdateRequest) GetMetadata() []ModelsMetadataSingleUpdateRequest`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsMetadataAllUpdateRequest) GetMetadataOk() (*[]ModelsMetadataSingleUpdateRequest, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsMetadataAllUpdateRequest) SetMetadata(v []ModelsMetadataSingleUpdateRequest)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


