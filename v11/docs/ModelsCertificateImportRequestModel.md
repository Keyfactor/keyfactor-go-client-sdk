# ModelsCertificateImportRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**StoreIds** | Pointer to **[]string** |  | [optional] 
**StoreTypes** | Pointer to [**[]ModelsEnrollmentManagementStoreType**](ModelsEnrollmentManagementStoreType.md) |  | [optional] 
**Schedule** | Pointer to **NullableTime** |  | [optional] 

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

### SetPasswordNil

`func (o *ModelsCertificateImportRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *ModelsCertificateImportRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
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

### SetMetadataNil

`func (o *ModelsCertificateImportRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ModelsCertificateImportRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
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

### SetStoreIdsNil

`func (o *ModelsCertificateImportRequestModel) SetStoreIdsNil(b bool)`

 SetStoreIdsNil sets the value for StoreIds to be an explicit nil

### UnsetStoreIds
`func (o *ModelsCertificateImportRequestModel) UnsetStoreIds()`

UnsetStoreIds ensures that no value is present for StoreIds, not even an explicit nil
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

### SetStoreTypesNil

`func (o *ModelsCertificateImportRequestModel) SetStoreTypesNil(b bool)`

 SetStoreTypesNil sets the value for StoreTypes to be an explicit nil

### UnsetStoreTypes
`func (o *ModelsCertificateImportRequestModel) UnsetStoreTypes()`

UnsetStoreTypes ensures that no value is present for StoreTypes, not even an explicit nil
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

### SetScheduleNil

`func (o *ModelsCertificateImportRequestModel) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *ModelsCertificateImportRequestModel) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


