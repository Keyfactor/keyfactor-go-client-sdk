# CSSCMSDataModelModelsContainerAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertStoreContainerId** | Pointer to **int32** |  | [optional] 
**NewContainerName** | Pointer to **NullableString** |  | [optional] 
**NewContainerType** | Pointer to **int32** |  | [optional] 
**KeystoreIds** | **[]string** |  | 

## Methods

### NewCSSCMSDataModelModelsContainerAssignment

`func NewCSSCMSDataModelModelsContainerAssignment(keystoreIds []string, ) *CSSCMSDataModelModelsContainerAssignment`

NewCSSCMSDataModelModelsContainerAssignment instantiates a new CSSCMSDataModelModelsContainerAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCSSCMSDataModelModelsContainerAssignmentWithDefaults

`func NewCSSCMSDataModelModelsContainerAssignmentWithDefaults() *CSSCMSDataModelModelsContainerAssignment`

NewCSSCMSDataModelModelsContainerAssignmentWithDefaults instantiates a new CSSCMSDataModelModelsContainerAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertStoreContainerId

`func (o *CSSCMSDataModelModelsContainerAssignment) GetCertStoreContainerId() int32`

GetCertStoreContainerId returns the CertStoreContainerId field if non-nil, zero value otherwise.

### GetCertStoreContainerIdOk

`func (o *CSSCMSDataModelModelsContainerAssignment) GetCertStoreContainerIdOk() (*int32, bool)`

GetCertStoreContainerIdOk returns a tuple with the CertStoreContainerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertStoreContainerId

`func (o *CSSCMSDataModelModelsContainerAssignment) SetCertStoreContainerId(v int32)`

SetCertStoreContainerId sets CertStoreContainerId field to given value.

### HasCertStoreContainerId

`func (o *CSSCMSDataModelModelsContainerAssignment) HasCertStoreContainerId() bool`

HasCertStoreContainerId returns a boolean if a field has been set.

### GetNewContainerName

`func (o *CSSCMSDataModelModelsContainerAssignment) GetNewContainerName() string`

GetNewContainerName returns the NewContainerName field if non-nil, zero value otherwise.

### GetNewContainerNameOk

`func (o *CSSCMSDataModelModelsContainerAssignment) GetNewContainerNameOk() (*string, bool)`

GetNewContainerNameOk returns a tuple with the NewContainerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewContainerName

`func (o *CSSCMSDataModelModelsContainerAssignment) SetNewContainerName(v string)`

SetNewContainerName sets NewContainerName field to given value.

### HasNewContainerName

`func (o *CSSCMSDataModelModelsContainerAssignment) HasNewContainerName() bool`

HasNewContainerName returns a boolean if a field has been set.

### SetNewContainerNameNil

`func (o *CSSCMSDataModelModelsContainerAssignment) SetNewContainerNameNil(b bool)`

 SetNewContainerNameNil sets the value for NewContainerName to be an explicit nil

### UnsetNewContainerName
`func (o *CSSCMSDataModelModelsContainerAssignment) UnsetNewContainerName()`

UnsetNewContainerName ensures that no value is present for NewContainerName, not even an explicit nil
### GetNewContainerType

`func (o *CSSCMSDataModelModelsContainerAssignment) GetNewContainerType() int32`

GetNewContainerType returns the NewContainerType field if non-nil, zero value otherwise.

### GetNewContainerTypeOk

`func (o *CSSCMSDataModelModelsContainerAssignment) GetNewContainerTypeOk() (*int32, bool)`

GetNewContainerTypeOk returns a tuple with the NewContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewContainerType

`func (o *CSSCMSDataModelModelsContainerAssignment) SetNewContainerType(v int32)`

SetNewContainerType sets NewContainerType field to given value.

### HasNewContainerType

`func (o *CSSCMSDataModelModelsContainerAssignment) HasNewContainerType() bool`

HasNewContainerType returns a boolean if a field has been set.

### GetKeystoreIds

`func (o *CSSCMSDataModelModelsContainerAssignment) GetKeystoreIds() []string`

GetKeystoreIds returns the KeystoreIds field if non-nil, zero value otherwise.

### GetKeystoreIdsOk

`func (o *CSSCMSDataModelModelsContainerAssignment) GetKeystoreIdsOk() (*[]string, bool)`

GetKeystoreIdsOk returns a tuple with the KeystoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreIds

`func (o *CSSCMSDataModelModelsContainerAssignment) SetKeystoreIds(v []string)`

SetKeystoreIds sets KeystoreIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


