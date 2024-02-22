# KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to [**KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedCustomer**](KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedCustomer.md) |  | [optional] 
**IssuedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 
**LicensedProducts** | Pointer to [**[]KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedProduct**](KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedProduct.md) |  | [optional] 

## Methods

### NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense

`func NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense() *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense`

NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense instantiates a new KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicenseWithDefaults

`func NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicenseWithDefaults() *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense`

NewKeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicenseWithDefaults instantiates a new KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseId

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### SetLicenseIdNil

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetLicenseIdNil(b bool)`

 SetLicenseIdNil sets the value for LicenseId to be an explicit nil

### UnsetLicenseId
`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) UnsetLicenseId()`

UnsetLicenseId ensures that no value is present for LicenseId, not even an explicit nil
### GetCustomer

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetCustomer() KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetCustomerOk() (*KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetCustomer(v KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetIssuedDate

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetIssuedDate() time.Time`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetIssuedDateOk() (*time.Time, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetIssuedDate(v time.Time)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetLicensedProducts

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetLicensedProducts() []KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedProduct`

GetLicensedProducts returns the LicensedProducts field if non-nil, zero value otherwise.

### GetLicensedProductsOk

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) GetLicensedProductsOk() (*[]KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedProduct, bool)`

GetLicensedProductsOk returns a tuple with the LicensedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedProducts

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetLicensedProducts(v []KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicensedProduct)`

SetLicensedProducts sets LicensedProducts field to given value.

### HasLicensedProducts

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) HasLicensedProducts() bool`

HasLicensedProducts returns a boolean if a field has been set.

### SetLicensedProductsNil

`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) SetLicensedProductsNil(b bool)`

 SetLicensedProductsNil sets the value for LicensedProducts to be an explicit nil

### UnsetLicensedProducts
`func (o *KeyfactorWebKeyfactorApiModelsLicenseLicenseResponseLicense) UnsetLicensedProducts()`

UnsetLicensedProducts ensures that no value is present for LicensedProducts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


