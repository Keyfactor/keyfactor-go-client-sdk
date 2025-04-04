# LicenseLicenseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyfactorVersion** | Pointer to **NullableString** |  | [optional] [readonly] 
**LicenseData** | Pointer to [**LicenseLicenseResponseLicense**](LicenseLicenseResponseLicense.md) |  | [optional] 

## Methods

### NewLicenseLicenseResponse

`func NewLicenseLicenseResponse() *LicenseLicenseResponse`

NewLicenseLicenseResponse instantiates a new LicenseLicenseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseLicenseResponseWithDefaults

`func NewLicenseLicenseResponseWithDefaults() *LicenseLicenseResponse`

NewLicenseLicenseResponseWithDefaults instantiates a new LicenseLicenseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyfactorVersion

`func (o *LicenseLicenseResponse) GetKeyfactorVersion() string`

GetKeyfactorVersion returns the KeyfactorVersion field if non-nil, zero value otherwise.

### GetKeyfactorVersionOk

`func (o *LicenseLicenseResponse) GetKeyfactorVersionOk() (*string, bool)`

GetKeyfactorVersionOk returns a tuple with the KeyfactorVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyfactorVersion

`func (o *LicenseLicenseResponse) SetKeyfactorVersion(v string)`

SetKeyfactorVersion sets KeyfactorVersion field to given value.

### HasKeyfactorVersion

`func (o *LicenseLicenseResponse) HasKeyfactorVersion() bool`

HasKeyfactorVersion returns a boolean if a field has been set.

### SetKeyfactorVersionNil

`func (o *LicenseLicenseResponse) SetKeyfactorVersionNil(b bool)`

 SetKeyfactorVersionNil sets the value for KeyfactorVersion to be an explicit nil

### UnsetKeyfactorVersion
`func (o *LicenseLicenseResponse) UnsetKeyfactorVersion()`

UnsetKeyfactorVersion ensures that no value is present for KeyfactorVersion, not even an explicit nil
### GetLicenseData

`func (o *LicenseLicenseResponse) GetLicenseData() LicenseLicenseResponseLicense`

GetLicenseData returns the LicenseData field if non-nil, zero value otherwise.

### GetLicenseDataOk

`func (o *LicenseLicenseResponse) GetLicenseDataOk() (*LicenseLicenseResponseLicense, bool)`

GetLicenseDataOk returns a tuple with the LicenseData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseData

`func (o *LicenseLicenseResponse) SetLicenseData(v LicenseLicenseResponseLicense)`

SetLicenseData sets LicenseData field to given value.

### HasLicenseData

`func (o *LicenseLicenseResponse) HasLicenseData() bool`

HasLicenseData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


