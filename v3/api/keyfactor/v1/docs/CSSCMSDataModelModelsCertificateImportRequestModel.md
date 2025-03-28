# CSSCMSDataModelModelsCertificateImportRequestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** |  | 
**Password** | Pointer to **NullableString** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**StoreIds** | Pointer to **[]string** |  | [optional] 
**StoreTypes** | Pointer to [**[]CSSCMSDataModelModelsEnrollmentManagementStoreType**](CSSCMSDataModelModelsEnrollmentManagementStoreType.md) |  | [optional] 
**Schedule** | Pointer to **NullableTime** |  | [optional] 
**OwnerRoleId** | Pointer to **NullableInt32** |  | [optional] 
**OwnerRoleName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCSSCMSDataModelModelsCertificateImportRequestModel

`func NewCSSCMSDataModelModelsCertificateImportRequestModel(certificate string, ) *CSSCMSDataModelModelsCertificateImportRequestModel`

NewCSSCMSDataModelModelsCertificateImportRequestModel instantiates a new CSSCMSDataModelModelsCertificateImportRequestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsCertificateImportRequestModelWithDefaults

`func NewCSSCMSDataModelModelsCertificateImportRequestModelWithDefaults() *CSSCMSDataModelModelsCertificateImportRequestModel`

NewCSSCMSDataModelModelsCertificateImportRequestModelWithDefaults instantiates a new CSSCMSDataModelModelsCertificateImportRequestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetPassword

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### SetPasswordNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetPasswordNil(b bool)`

 SetPasswordNil sets the value for Password to be an explicit nil

### UnsetPassword
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetPassword()`

UnsetPassword ensures that no value is present for Password, not even an explicit nil
### GetMetadata

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetStoreIds

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetStoreIds() []string`

GetStoreIds returns the StoreIds field if non-nil, zero value otherwise.

### GetStoreIdsOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetStoreIdsOk() (*[]string, bool)`

GetStoreIdsOk returns a tuple with the StoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreIds

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetStoreIds(v []string)`

SetStoreIds sets StoreIds field to given value.

### HasStoreIds

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasStoreIds() bool`

HasStoreIds returns a boolean if a field has been set.

### SetStoreIdsNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetStoreIdsNil(b bool)`

 SetStoreIdsNil sets the value for StoreIds to be an explicit nil

### UnsetStoreIds
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetStoreIds()`

UnsetStoreIds ensures that no value is present for StoreIds, not even an explicit nil
### GetStoreTypes

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetStoreTypes() []CSSCMSDataModelModelsEnrollmentManagementStoreType`

GetStoreTypes returns the StoreTypes field if non-nil, zero value otherwise.

### GetStoreTypesOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetStoreTypesOk() (*[]CSSCMSDataModelModelsEnrollmentManagementStoreType, bool)`

GetStoreTypesOk returns a tuple with the StoreTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreTypes

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetStoreTypes(v []CSSCMSDataModelModelsEnrollmentManagementStoreType)`

SetStoreTypes sets StoreTypes field to given value.

### HasStoreTypes

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasStoreTypes() bool`

HasStoreTypes returns a boolean if a field has been set.

### SetStoreTypesNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetStoreTypesNil(b bool)`

 SetStoreTypesNil sets the value for StoreTypes to be an explicit nil

### UnsetStoreTypes
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetStoreTypes()`

UnsetStoreTypes ensures that no value is present for StoreTypes, not even an explicit nil
### GetSchedule

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetSchedule() time.Time`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetScheduleOk() (*time.Time, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetSchedule(v time.Time)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### SetScheduleNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetScheduleNil(b bool)`

 SetScheduleNil sets the value for Schedule to be an explicit nil

### UnsetSchedule
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetSchedule()`

UnsetSchedule ensures that no value is present for Schedule, not even an explicit nil
### GetOwnerRoleId

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetOwnerRoleId() int32`

GetOwnerRoleId returns the OwnerRoleId field if non-nil, zero value otherwise.

### GetOwnerRoleIdOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetOwnerRoleIdOk() (*int32, bool)`

GetOwnerRoleIdOk returns a tuple with the OwnerRoleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleId

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetOwnerRoleId(v int32)`

SetOwnerRoleId sets OwnerRoleId field to given value.

### HasOwnerRoleId

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasOwnerRoleId() bool`

HasOwnerRoleId returns a boolean if a field has been set.

### SetOwnerRoleIdNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetOwnerRoleIdNil(b bool)`

 SetOwnerRoleIdNil sets the value for OwnerRoleId to be an explicit nil

### UnsetOwnerRoleId
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetOwnerRoleId()`

UnsetOwnerRoleId ensures that no value is present for OwnerRoleId, not even an explicit nil
### GetOwnerRoleName

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetOwnerRoleName() string`

GetOwnerRoleName returns the OwnerRoleName field if non-nil, zero value otherwise.

### GetOwnerRoleNameOk

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) GetOwnerRoleNameOk() (*string, bool)`

GetOwnerRoleNameOk returns a tuple with the OwnerRoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerRoleName

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetOwnerRoleName(v string)`

SetOwnerRoleName sets OwnerRoleName field to given value.

### HasOwnerRoleName

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) HasOwnerRoleName() bool`

HasOwnerRoleName returns a boolean if a field has been set.

### SetOwnerRoleNameNil

`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) SetOwnerRoleNameNil(b bool)`

 SetOwnerRoleNameNil sets the value for OwnerRoleName to be an explicit nil

### UnsetOwnerRoleName
`func (o *CSSCMSDataModelModelsCertificateImportRequestModel) UnsetOwnerRoleName()`

UnsetOwnerRoleName ensures that no value is present for OwnerRoleName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


