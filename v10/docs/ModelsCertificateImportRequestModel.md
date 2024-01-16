# ModelsCertificateImportRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | Base64-encoded certificate contents | 
**Password** | Pointer to **string** | Optional password associated if required for a PFX | [optional] 
**Metadata** | Pointer to **map[string]string** | Colleciton of metadata to be associated with the imported certificate | [optional] 
**StoreIds** | Pointer to **[]string** | List of the Keyfactor certificate store identifiers (GUID) with which the imported certificate should be associated | [optional] 
**StoreTypes** | Pointer to [**[]ModelsEnrollmentManagementStoreType**](ModelsEnrollmentManagementStoreType.md) | List of the certificate store types with with the imported certificate should be associated | [optional] 
**Schedule** | Pointer to **time.Time** | Schedule on which the certificate should be imported | [optional] 
**ImportMetadata** | Pointer to **bool** | Whether or not we should validate and import the certificate&#39;s metadata. | [optional] 

## Methods

### NewModelsCertificateImportRequestModel

`func NewModelsCertificateImportRequestModel(certificate string, ) *ModelsCertificateImportRequestModel`

NewModelsCertificateImportRequestModel instantiates a new ModelsCertificateImportRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelsCertificateImportRequestModelWithDefaults

`func NewModelsCertificateImportRequestModelWithDefaults() *ModelsCertificateImportRequestModel`

NewModelsCertificateImportRequestModelWithDefaults instantiates a new ModelsCertificateImportRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *ModelsCertificateImportRequestModel) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *ModelsCertificateImportRequestModel) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *ModelsCertificateImportRequestModel) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPassword

`func (o *ModelsCertificateImportRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ModelsCertificateImportRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ModelsCertificateImportRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ModelsCertificateImportRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetMetadata

`func (o *ModelsCertificateImportRequestModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ModelsCertificateImportRequestModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ModelsCertificateImportRequestModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ModelsCertificateImportRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetStoreIds

`func (o *ModelsCertificateImportRequestModel) GetStoreIds() []string`

GetStoreIds returns the StoreIds field if non-nil, zero value otherwise.

### GetStoreIdsOk

`func (o *ModelsCertificateImportRequestModel) GetStoreIdsOk() (*[]string, bool)`

GetStoreIdsOk returns a tuple with the StoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIds

`func (o *ModelsCertificateImportRequestModel) SetStoreIds(v []string)`

SetStoreIds sets StoreIds field to given value.

### HasStoreIds

`func (o *ModelsCertificateImportRequestModel) HasStoreIds() bool`

HasStoreIds returns a boolean if a field has been set.

### GetStoreTypes

`func (o *ModelsCertificateImportRequestModel) GetStoreTypes() []ModelsEnrollmentManagementStoreType`

GetStoreTypes returns the StoreTypes field if non-nil, zero value otherwise.

### GetStoreTypesOk

`func (o *ModelsCertificateImportRequestModel) GetStoreTypesOk() (*[]ModelsEnrollmentManagementStoreType, bool)`

GetStoreTypesOk returns a tuple with the StoreTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypes

`func (o *ModelsCertificateImportRequestModel) SetStoreTypes(v []ModelsEnrollmentManagementStoreType)`

SetStoreTypes sets StoreTypes field to given value.

### HasStoreTypes

`func (o *ModelsCertificateImportRequestModel) HasStoreTypes() bool`

HasStoreTypes returns a boolean if a field has been set.

### GetSchedule

`func (o *ModelsCertificateImportRequestModel) GetSchedule() time.Time`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ModelsCertificateImportRequestModel) GetScheduleOk() (*time.Time, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ModelsCertificateImportRequestModel) SetSchedule(v time.Time)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ModelsCertificateImportRequestModel) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetImportMetadata

`func (o *ModelsCertificateImportRequestModel) GetImportMetadata() bool`

GetImportMetadata returns the ImportMetadata field if non-nil, zero value otherwise.

### GetImportMetadataOk

`func (o *ModelsCertificateImportRequestModel) GetImportMetadataOk() (*bool, bool)`

GetImportMetadataOk returns a tuple with the ImportMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMetadata

`func (o *ModelsCertificateImportRequestModel) SetImportMetadata(v bool)`

SetImportMetadata sets ImportMetadata field to given value.

### HasImportMetadata

`func (o *ModelsCertificateImportRequestModel) HasImportMetadata() bool`

HasImportMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


