# LicenseLicenseResponseLicensedProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProductId** | Pointer to **NullableString** |  | [optional] 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**MajorRev** | Pointer to **NullableString** |  | [optional] 
**MinorRev** | Pointer to **NullableString** |  | [optional] 
**LicensedFeatures** | Pointer to [**[]LicenseLicenseResponseLicensedFeature**](LicenseLicenseResponseLicensedFeature.md) |  | [optional] 

## Methods

### NewLicenseLicenseResponseLicensedProduct

`func NewLicenseLicenseResponseLicensedProduct() *LicenseLicenseResponseLicensedProduct`

NewLicenseLicenseResponseLicensedProduct instantiates a new LicenseLicenseResponseLicensedProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseResponseLicensedProductWithDefaults

`func NewLicenseLicenseResponseLicensedProductWithDefaults() *LicenseLicenseResponseLicensedProduct`

NewLicenseLicenseResponseLicensedProductWithDefaults instantiates a new LicenseLicenseResponseLicensedProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProductId

`func (o *LicenseLicenseResponseLicensedProduct) GetProductId() string`

GetProductId returns the ProductId field if non-nil, zero value otherwise.

### GetProductIdOk

`func (o *LicenseLicenseResponseLicensedProduct) GetProductIdOk() (*string, bool)`

GetProductIdOk returns a tuple with the ProductId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductId

`func (o *LicenseLicenseResponseLicensedProduct) SetProductId(v string)`

SetProductId sets ProductId field to given value.

### HasProductId

`func (o *LicenseLicenseResponseLicensedProduct) HasProductId() bool`

HasProductId returns a boolean if a field has been set.

### SetProductIdNil

`func (o *LicenseLicenseResponseLicensedProduct) SetProductIdNil(b bool)`

 SetProductIdNil sets the value for ProductId to be an explicit nil

### UnsetProductId
`func (o *LicenseLicenseResponseLicensedProduct) UnsetProductId()`

UnsetProductId ensures that no value is present for ProductId, not even an explicit nil
### GetDisplayName

`func (o *LicenseLicenseResponseLicensedProduct) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *LicenseLicenseResponseLicensedProduct) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *LicenseLicenseResponseLicensedProduct) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *LicenseLicenseResponseLicensedProduct) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *LicenseLicenseResponseLicensedProduct) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *LicenseLicenseResponseLicensedProduct) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetMajorRev

`func (o *LicenseLicenseResponseLicensedProduct) GetMajorRev() string`

GetMajorRev returns the MajorRev field if non-nil, zero value otherwise.

### GetMajorRevOk

`func (o *LicenseLicenseResponseLicensedProduct) GetMajorRevOk() (*string, bool)`

GetMajorRevOk returns a tuple with the MajorRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMajorRev

`func (o *LicenseLicenseResponseLicensedProduct) SetMajorRev(v string)`

SetMajorRev sets MajorRev field to given value.

### HasMajorRev

`func (o *LicenseLicenseResponseLicensedProduct) HasMajorRev() bool`

HasMajorRev returns a boolean if a field has been set.

### SetMajorRevNil

`func (o *LicenseLicenseResponseLicensedProduct) SetMajorRevNil(b bool)`

 SetMajorRevNil sets the value for MajorRev to be an explicit nil

### UnsetMajorRev
`func (o *LicenseLicenseResponseLicensedProduct) UnsetMajorRev()`

UnsetMajorRev ensures that no value is present for MajorRev, not even an explicit nil
### GetMinorRev

`func (o *LicenseLicenseResponseLicensedProduct) GetMinorRev() string`

GetMinorRev returns the MinorRev field if non-nil, zero value otherwise.

### GetMinorRevOk

`func (o *LicenseLicenseResponseLicensedProduct) GetMinorRevOk() (*string, bool)`

GetMinorRevOk returns a tuple with the MinorRev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinorRev

`func (o *LicenseLicenseResponseLicensedProduct) SetMinorRev(v string)`

SetMinorRev sets MinorRev field to given value.

### HasMinorRev

`func (o *LicenseLicenseResponseLicensedProduct) HasMinorRev() bool`

HasMinorRev returns a boolean if a field has been set.

### SetMinorRevNil

`func (o *LicenseLicenseResponseLicensedProduct) SetMinorRevNil(b bool)`

 SetMinorRevNil sets the value for MinorRev to be an explicit nil

### UnsetMinorRev
`func (o *LicenseLicenseResponseLicensedProduct) UnsetMinorRev()`

UnsetMinorRev ensures that no value is present for MinorRev, not even an explicit nil
### GetLicensedFeatures

`func (o *LicenseLicenseResponseLicensedProduct) GetLicensedFeatures() []LicenseLicenseResponseLicensedFeature`

GetLicensedFeatures returns the LicensedFeatures field if non-nil, zero value otherwise.

### GetLicensedFeaturesOk

`func (o *LicenseLicenseResponseLicensedProduct) GetLicensedFeaturesOk() (*[]LicenseLicenseResponseLicensedFeature, bool)`

GetLicensedFeaturesOk returns a tuple with the LicensedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicensedFeatures

`func (o *LicenseLicenseResponseLicensedProduct) SetLicensedFeatures(v []LicenseLicenseResponseLicensedFeature)`

SetLicensedFeatures sets LicensedFeatures field to given value.

### HasLicensedFeatures

`func (o *LicenseLicenseResponseLicensedProduct) HasLicensedFeatures() bool`

HasLicensedFeatures returns a boolean if a field has been set.

### SetLicensedFeaturesNil

`func (o *LicenseLicenseResponseLicensedProduct) SetLicensedFeaturesNil(b bool)`

 SetLicensedFeaturesNil sets the value for LicensedFeatures to be an explicit nil

### UnsetLicensedFeatures
`func (o *LicenseLicenseResponseLicensedProduct) UnsetLicensedFeatures()`

UnsetLicensedFeatures ensures that no value is present for LicensedFeatures, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


