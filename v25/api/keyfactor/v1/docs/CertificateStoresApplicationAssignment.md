# CertificateStoresApplicationAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** |  | [optional] 
**NewApplicationName** | Pointer to **NullableString** |  | [optional] 
**NewApplicationType** | Pointer to **int32** |  | [optional] 
**KeystoreIds** | **[]string** |  | 

## Methods

### NewCertificateStoresApplicationAssignment

`func NewCertificateStoresApplicationAssignment(keystoreIds []string, ) *CertificateStoresApplicationAssignment`

NewCertificateStoresApplicationAssignment instantiates a new CertificateStoresApplicationAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateStoresApplicationAssignmentWithDefaults

`func NewCertificateStoresApplicationAssignmentWithDefaults() *CertificateStoresApplicationAssignment`

NewCertificateStoresApplicationAssignmentWithDefaults instantiates a new CertificateStoresApplicationAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *CertificateStoresApplicationAssignment) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *CertificateStoresApplicationAssignment) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *CertificateStoresApplicationAssignment) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *CertificateStoresApplicationAssignment) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetNewApplicationName

`func (o *CertificateStoresApplicationAssignment) GetNewApplicationName() string`

GetNewApplicationName returns the NewApplicationName field if non-nil, zero value otherwise.

### GetNewApplicationNameOk

`func (o *CertificateStoresApplicationAssignment) GetNewApplicationNameOk() (*string, bool)`

GetNewApplicationNameOk returns a tuple with the NewApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplicationName

`func (o *CertificateStoresApplicationAssignment) SetNewApplicationName(v string)`

SetNewApplicationName sets NewApplicationName field to given value.

### HasNewApplicationName

`func (o *CertificateStoresApplicationAssignment) HasNewApplicationName() bool`

HasNewApplicationName returns a boolean if a field has been set.

### SetNewApplicationNameNil

`func (o *CertificateStoresApplicationAssignment) SetNewApplicationNameNil(b bool)`

 SetNewApplicationNameNil sets the value for NewApplicationName to be an explicit nil

### UnsetNewApplicationName
`func (o *CertificateStoresApplicationAssignment) UnsetNewApplicationName()`

UnsetNewApplicationName ensures that no value is present for NewApplicationName, not even an explicit nil
### GetNewApplicationType

`func (o *CertificateStoresApplicationAssignment) GetNewApplicationType() int32`

GetNewApplicationType returns the NewApplicationType field if non-nil, zero value otherwise.

### GetNewApplicationTypeOk

`func (o *CertificateStoresApplicationAssignment) GetNewApplicationTypeOk() (*int32, bool)`

GetNewApplicationTypeOk returns a tuple with the NewApplicationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewApplicationType

`func (o *CertificateStoresApplicationAssignment) SetNewApplicationType(v int32)`

SetNewApplicationType sets NewApplicationType field to given value.

### HasNewApplicationType

`func (o *CertificateStoresApplicationAssignment) HasNewApplicationType() bool`

HasNewApplicationType returns a boolean if a field has been set.

### GetKeystoreIds

`func (o *CertificateStoresApplicationAssignment) GetKeystoreIds() []string`

GetKeystoreIds returns the KeystoreIds field if non-nil, zero value otherwise.

### GetKeystoreIdsOk

`func (o *CertificateStoresApplicationAssignment) GetKeystoreIdsOk() (*[]string, bool)`

GetKeystoreIdsOk returns a tuple with the KeystoreIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeystoreIds

`func (o *CertificateStoresApplicationAssignment) SetKeystoreIds(v []string)`

SetKeystoreIds sets KeystoreIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


