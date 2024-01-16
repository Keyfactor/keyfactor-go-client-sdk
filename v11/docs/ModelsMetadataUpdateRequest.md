# ModelsMetadataUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertificateId** | Pointer to **int32** |  | [optional] 
**Metadata** | **map[string]string** |  | 

## Methods

### NewModelsMetadataUpdateRequest

`func NewModelsMetadataUpdateRequest(metadata map[string]string, ) *ModelsMetadataUpdateRequest`

NewModelsMetadataUpdateRequest instantiates a new ModelsMetadataUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsMetadataUpdateRequestWithDefaults

`func NewModelsMetadataUpdateRequestWithDefaults() *ModelsMetadataUpdateRequest`

NewModelsMetadataUpdateRequestWithDefaults instantiates a new ModelsMetadataUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificateId

`func (o *ModelsMetadataUpdateRequest) GetCertificateId() int32`

GetCertificateId returns the CertificateId field if non-nil, zero value otherwise.

### GetCertificateIdOk

`func (o *ModelsMetadataUpdateRequest) GetCertificateIdOk() (*int32, bool)`

GetCertificateIdOk returns a tuple with the CertificateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateId

`func (o *ModelsMetadataUpdateRequest) SetCertificateId(v int32)`

SetCertificateId sets CertificateId field to given value.

### HasCertificateId

`func (o *ModelsMetadataUpdateRequest) HasCertificateId() bool`

HasCertificateId returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsMetadataUpdateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsMetadataUpdateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsMetadataUpdateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


