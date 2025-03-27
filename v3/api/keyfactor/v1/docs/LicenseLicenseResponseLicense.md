# LicenseLicenseResponseLicense

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LicenseId** | Pointer to **NullableString** |  | [optional] 
**Customer** | Pointer to [**LicenseLicenseResponseLicensedCustomer**](LicenseLicenseResponseLicensedCustomer.md) |  | [optional] 
**IssuedDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **NullableTime** |  | [optional] 
**LicensedProducts** | Pointer to [**[]LicenseLicenseResponseLicensedProduct**](LicenseLicenseResponseLicensedProduct.md) |  | [optional] 

## Methods

### NewLicenseLicenseResponseLicense

`func NewLicenseLicenseResponseLicense() *LicenseLicenseResponseLicense`

NewLicenseLicenseResponseLicense instantiates a new LicenseLicenseResponseLicense object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseResponseLicenseWithDefaults

`func NewLicenseLicenseResponseLicenseWithDefaults() *LicenseLicenseResponseLicense`

NewLicenseLicenseResponseLicenseWithDefaults instantiates a new LicenseLicenseResponseLicense object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLicenseId

`func (o *LicenseLicenseResponseLicense) GetLicenseId() string`

GetLicenseId returns the LicenseId field if non-nil, zero value otherwise.

### GetLicenseIdOk

`func (o *LicenseLicenseResponseLicense) GetLicenseIdOk() (*string, bool)`

GetLicenseIdOk returns a tuple with the LicenseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseId

`func (o *LicenseLicenseResponseLicense) SetLicenseId(v string)`

SetLicenseId sets LicenseId field to given value.

### HasLicenseId

`func (o *LicenseLicenseResponseLicense) HasLicenseId() bool`

HasLicenseId returns a boolean if a field has been set.

### SetLicenseIdNil

`func (o *LicenseLicenseResponseLicense) SetLicenseIdNil(b bool)`

 SetLicenseIdNil sets the value for LicenseId to be an explicit nil

### UnsetLicenseId
`func (o *LicenseLicenseResponseLicense) UnsetLicenseId()`

UnsetLicenseId ensures that no value is present for LicenseId, not even an explicit nil
### GetCustomer

`func (o *LicenseLicenseResponseLicense) GetCustomer() LicenseLicenseResponseLicensedCustomer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *LicenseLicenseResponseLicense) GetCustomerOk() (*LicenseLicenseResponseLicensedCustomer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *LicenseLicenseResponseLicense) SetCustomer(v LicenseLicenseResponseLicensedCustomer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *LicenseLicenseResponseLicense) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetIssuedDate

`func (o *LicenseLicenseResponseLicense) GetIssuedDate() time.Time`

GetIssuedDate returns the IssuedDate field if non-nil, zero value otherwise.

### GetIssuedDateOk

`func (o *LicenseLicenseResponseLicense) GetIssuedDateOk() (*time.Time, bool)`

GetIssuedDateOk returns a tuple with the IssuedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedDate

`func (o *LicenseLicenseResponseLicense) SetIssuedDate(v time.Time)`

SetIssuedDate sets IssuedDate field to given value.

### HasIssuedDate

`func (o *LicenseLicenseResponseLicense) HasIssuedDate() bool`

HasIssuedDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *LicenseLicenseResponseLicense) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *LicenseLicenseResponseLicense) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *LicenseLicenseResponseLicense) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *LicenseLicenseResponseLicense) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### SetExpirationDateNil

`func (o *LicenseLicenseResponseLicense) SetExpirationDateNil(b bool)`

 SetExpirationDateNil sets the value for ExpirationDate to be an explicit nil

### UnsetExpirationDate
`func (o *LicenseLicenseResponseLicense) UnsetExpirationDate()`

UnsetExpirationDate ensures that no value is present for ExpirationDate, not even an explicit nil
### GetLicensedProducts

`func (o *LicenseLicenseResponseLicense) GetLicensedProducts() []LicenseLicenseResponseLicensedProduct`

GetLicensedProducts returns the LicensedProducts field if non-nil, zero value otherwise.

### GetLicensedProductsOk

`func (o *LicenseLicenseResponseLicense) GetLicensedProductsOk() (*[]LicenseLicenseResponseLicensedProduct, bool)`

GetLicensedProductsOk returns a tuple with the LicensedProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedProducts

`func (o *LicenseLicenseResponseLicense) SetLicensedProducts(v []LicenseLicenseResponseLicensedProduct)`

SetLicensedProducts sets LicensedProducts field to given value.

### HasLicensedProducts

`func (o *LicenseLicenseResponseLicense) HasLicensedProducts() bool`

HasLicensedProducts returns a boolean if a field has been set.

### SetLicensedProductsNil

`func (o *LicenseLicenseResponseLicense) SetLicensedProductsNil(b bool)`

 SetLicensedProductsNil sets the value for LicensedProducts to be an explicit nil

### UnsetLicensedProducts
`func (o *LicenseLicenseResponseLicense) UnsetLicensedProducts()`

UnsetLicensedProducts ensures that no value is present for LicensedProducts, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


