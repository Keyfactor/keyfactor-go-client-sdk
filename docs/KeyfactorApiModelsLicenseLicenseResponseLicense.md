# KeyfactorApiModelsLicenseLicenseResponseLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseId** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**KeyfactorApiModelsLicenseLicenseResponseLicensedCustomer**](KeyfactorApiModelsLicenseLicenseResponseLicensedCustomer.md) |  | [optional] 
**IssuedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**LicensedProducts** | Pointer to [**[]KeyfactorApiModelsLicenseLicenseResponseLicensedProduct**](KeyfactorApiModelsLicenseLicenseResponseLicensedProduct.md) |  | [optional] 

## Methods

### NewKeyfactorApiModelsLicenseLicenseResponseLicense

`func NewKeyfactorApiModelsLicenseLicenseResponseLicense() *KeyfactorApiModelsLicenseLicenseResponseLicense`

NewKeyfactorApiModelsLicenseLicenseResponseLicense instantiates a new KeyfactorApiModelsLicenseLicenseResponseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyfactorApiModelsLicenseLicenseResponseLicenseWithDefaults

`func NewKeyfactorApiModelsLicenseLicenseResponseLicenseWithDefaults() *KeyfactorApiModelsLicenseLicenseResponseLicense`

NewKeyfactorApiModelsLicenseLicenseResponseLicenseWithDefaults instantiates a new KeyfactorApiModelsLicenseLicenseResponseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseId

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### GetCustomer

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetCustomer() KeyfactorApiModelsLicenseLicenseResponseLicensedCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetCustomerOk() (*KeyfactorApiModelsLicenseLicenseResponseLicensedCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) SetCustomer(v KeyfactorApiModelsLicenseLicenseResponseLicensedCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetIssuedDate

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetIssuedDate() time.Time`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetIssuedDateOk() (*time.Time, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) SetIssuedDate(v time.Time)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetLicensedProducts

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetLicensedProducts() []KeyfactorApiModelsLicenseLicenseResponseLicensedProduct`

GetLicensedProducts returns the LicensedProducts field if non-nil, zero value otherwise.

### GetLicensedProductsOk

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) GetLicensedProductsOk() (*[]KeyfactorApiModelsLicenseLicenseResponseLicensedProduct, bool)`

GetLicensedProductsOk returns a tuple with the LicensedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedProducts

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) SetLicensedProducts(v []KeyfactorApiModelsLicenseLicenseResponseLicensedProduct)`

SetLicensedProducts sets LicensedProducts field to given value.

### HasLicensedProducts

`func (o *KeyfactorApiModelsLicenseLicenseResponseLicense) HasLicensedProducts() bool`

HasLicensedProducts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


